package gmmu

import (
	"log"
	"reflect"
	"gitlab.com/akita/akita/v3/sim"
	"gitlab.com/akita/akita/v3/tracing"
	"gitlab.com/akita/mem/v3/vm"
)

type transaction struct {
	req       *vm.TranslationReq
	page      vm.Page
	cycleLeft int
	migration *vm.PageMigrationReqToDriver
}

// MMU is the default mmu implementation. It is also an akita Component.
type GMMU struct {
	sim.TickingComponent

	topPort       sim.Port
	migrationPort sim.Port

	MigrationServiceProvider sim.Port

	topSender sim.BufferedSender

	pageTable           vm.PageTable  //PageTable
	latency             int
	maxRequestsInFlight int

	walkingTranslations      []transaction
	migrationQueue           []transaction //QUEUE
	migrationQueueSize       int
	currentOnDemandMigration transaction
	isDoingMigration         bool

	toRemoveFromPTW        []int
	PageAccessedByDeviceID map[uint64][]uint64
}

// Tick defines how the MMU update state each cycle
func (gmmu *GMMU) Tick(now sim.VTimeInSec) bool {
	madeProgress := false

	madeProgress = gmmu.topSender.Tick(now) || madeProgress
	madeProgress = gmmu.sendMigrationToDriver(now) || madeProgress
	madeProgress = gmmu.walkPageTable(now) || madeProgress
	madeProgress = gmmu.processMigrationReturn(now) || madeProgress
	madeProgress = gmmu.parseFromTop(now) || madeProgress

	return madeProgress
}

func (gmmu *GMMU) trace(now sim.VTimeInSec, what string) {
	ctx := sim.HookCtx{
		Domain: gmmu,
		Item:   what,
	}

	gmmu.InvokeHook(ctx)
}

func (gmmu *GMMU) walkPageTable(now sim.VTimeInSec) bool {
	madeProgress := false
	for i := 0; i < len(gmmu.walkingTranslations); i++ {
		if gmmu.walkingTranslations[i].cycleLeft > 0 {
			gmmu.walkingTranslations[i].cycleLeft--
			madeProgress = true
			continue
		}

		madeProgress = gmmu.finalizePageWalk(now, i) || madeProgress
	}

	tmp := gmmu.walkingTranslations[:0]
	for i := 0; i < len(gmmu.walkingTranslations); i++ {
		if !gmmu.toRemove(i) {
			tmp = append(tmp, gmmu.walkingTranslations[i])
		}
	}
	gmmu.walkingTranslations = tmp
	gmmu.toRemoveFromPTW = nil

	return madeProgress
}

func (gmmu *GMMU) finalizePageWalk(
	now sim.VTimeInSec,
	walkingIndex int,
) bool {
	req := gmmu.walkingTranslations[walkingIndex].req
	page, found := gmmu.pageTable.Find(req.PID, req.VAddr)

	if !found {
		panic("page not found")
	}

	gmmu.walkingTranslations[walkingIndex].page = page

	if page.IsMigrating {
		return gmmu.addTransactionToMigrationQueue(walkingIndex)
	}

	if gmmu.pageNeedMigrate(gmmu.walkingTranslations[walkingIndex]) {
		return gmmu.addTransactionToMigrationQueue(walkingIndex)
	}

	return gmmu.doPageWalkHit(now, walkingIndex)
}

func (gmmu *GMMU) addTransactionToMigrationQueue(walkingIndex int) bool {
	if len(gmmu.migrationQueue) >= gmmu.migrationQueueSize {
		return false
	}

	gmmu.toRemoveFromPTW = append(gmmu.toRemoveFromPTW, walkingIndex)
	gmmu.migrationQueue = append(gmmu.migrationQueue,
		gmmu.walkingTranslations[walkingIndex])

	page := gmmu.walkingTranslations[walkingIndex].page
	page.IsMigrating = true
	gmmu.pageTable.Update(page)

	return true
}

func (gmmu *GMMU) pageNeedMigrate(walking transaction) bool {
	if walking.req.DeviceID == walking.page.DeviceID {
		return false
	}

	if !walking.page.Unified {
		return false
	}

	if walking.page.IsPinned {
		return false
	}

	return true
}

func (gmmu *GMMU) doPageWalkHit(
	now sim.VTimeInSec,
	walkingIndex int,
) bool {
	if !gmmu.topSender.CanSend(1) {
		return false
	}
	walking := gmmu.walkingTranslations[walkingIndex]

	rsp := vm.TranslationRspBuilder{}.
		WithSendTime(now).
		WithSrc(gmmu.topPort).
		WithDst(walking.req.Src).
		WithRspTo(walking.req.ID).
		WithPage(walking.page).
		Build()

	gmmu.topSender.Send(rsp)
	gmmu.toRemoveFromPTW = append(gmmu.toRemoveFromPTW, walkingIndex)

	tracing.TraceReqComplete(walking.req, gmmu)

	return true
}

func (gmmu *GMMU) sendMigrationToDriver(
	now sim.VTimeInSec,
) (madeProgress bool) {
	if len(gmmu.migrationQueue) == 0 {
		return false
	}
	trans := gmmu.migrationQueue[0]
	req := trans.req
	page, found := gmmu.pageTable.Find(req.PID, req.VAddr)
	if !found {
		panic("page not found")
	}
	trans.page = page

	if req.DeviceID == page.DeviceID || page.IsPinned {
		gmmu.sendTranlationRsp(now, trans)
		gmmu.migrationQueue = gmmu.migrationQueue[1:]
		gmmu.markPageAsNotMigratingIfNotInTheMigrationQueue(page)

		return true
	}

	if gmmu.isDoingMigration {
		return false
	}

	migrationInfo := new(vm.PageMigrationInfo)
	migrationInfo.GPUReqToVAddrMap = make(map[uint64][]uint64)
	migrationInfo.GPUReqToVAddrMap[trans.req.DeviceID] =
		append(migrationInfo.GPUReqToVAddrMap[trans.req.DeviceID],
			trans.req.VAddr)

	gmmu.PageAccessedByDeviceID[page.VAddr] =
		append(gmmu.PageAccessedByDeviceID[page.VAddr], page.DeviceID)

	migrationReq := vm.NewPageMigrationReqToDriver(
		now, gmmu.migrationPort, gmmu.MigrationServiceProvider)
	migrationReq.PID = page.PID
	migrationReq.PageSize = page.PageSize
	migrationReq.CurrPageHostGPU = page.DeviceID
	migrationReq.MigrationInfo = migrationInfo
	migrationReq.CurrAccessingGPUs = unique(gmmu.PageAccessedByDeviceID[page.VAddr])
	migrationReq.RespondToTop = true

	err := gmmu.migrationPort.Send(migrationReq)
	if err != nil {
		return false
	}

	trans.page.IsMigrating = true
	gmmu.pageTable.Update(trans.page)
	trans.migration = migrationReq
	gmmu.isDoingMigration = true
	gmmu.currentOnDemandMigration = trans
	gmmu.migrationQueue = gmmu.migrationQueue[1:]

	return true
}

func (gmmu *GMMU) markPageAsNotMigratingIfNotInTheMigrationQueue(
	page vm.Page,
) vm.Page {
	inQueue := false
	for _, t := range gmmu.migrationQueue {
		if page.PAddr == t.page.PAddr {
			inQueue = true
			break
		}
	}

	if !inQueue {
		page.IsMigrating = false
		gmmu.pageTable.Update(page)
		return page
	}

	return page
}

func (gmmu *GMMU) sendTranlationRsp(
	now sim.VTimeInSec,
	trans transaction,
) (madeProgress bool) {
	req := trans.req
	page := trans.page

	rsp := vm.TranslationRspBuilder{}.
		WithSendTime(now).
		WithSrc(gmmu.topPort).
		WithDst(req.Src).
		WithRspTo(req.ID).
		WithPage(page).
		Build()
	gmmu.topSender.Send(rsp)

	return true
}

func (gmmu *GMMU) processMigrationReturn(now sim.VTimeInSec) bool {
	item := gmmu.migrationPort.Peek()
	if item == nil {
		return false
	}

	if !gmmu.topSender.CanSend(1) {
		return false
	}

	req := gmmu.currentOnDemandMigration.req
	page, found := gmmu.pageTable.Find(req.PID, req.VAddr)
	if !found {
		panic("page not found")
	}

	rsp := vm.TranslationRspBuilder{}.
		WithSendTime(now).
		WithSrc(gmmu.topPort).
		WithDst(req.Src).
		WithRspTo(req.ID).
		WithPage(page).
		Build()
	gmmu.topSender.Send(rsp)

	gmmu.isDoingMigration = false

	page = gmmu.markPageAsNotMigratingIfNotInTheMigrationQueue(page)
	page.IsPinned = true
	gmmu.pageTable.Update(page)

	gmmu.migrationPort.Retrieve(now)

	return true
}

func (gmmu *GMMU) parseFromTop(now sim.VTimeInSec) bool {
	if len(gmmu.walkingTranslations) >= gmmu.maxRequestsInFlight {
		return false
	}

	req := gmmu.topPort.Retrieve(now)
	if req == nil {
		return false
	}

	tracing.TraceReqReceive(req, gmmu)

	switch req := req.(type) {
	case *vm.TranslationReq:
		gmmu.startWalking(req)
	default:
		log.Panicf("MMU canot handle request of type %s", reflect.TypeOf(req))
	}

	return true
}

func (gmmu *GMMU) startWalking(req *vm.TranslationReq) {
	translationInPipeline := transaction{
		req:       req,
		cycleLeft: gmmu.latency,
	}

	gmmu.walkingTranslations = append(gmmu.walkingTranslations, translationInPipeline)
}

func (gmmu *GMMU) toRemove(index int) bool {
	for i := 0; i < len(gmmu.toRemoveFromPTW); i++ {
		remove := gmmu.toRemoveFromPTW[i]
		if remove == index {
			return true
		}
	}
	return false
}

func unique(intSlice []uint64) []uint64 {
	keys := make(map[int]bool)
	list := []uint64{}
	for _, entry := range intSlice {
		if _, value := keys[int(entry)]; !value {
			keys[int(entry)] = true
			list = append(list, entry)
		}
	}
	return list
}
