package gmmu

import (
	"gitlab.com/akita/akita/v3/sim"
	"gitlab.com/akita/mem/v3/vm"
)

// A Builder can build GMMU component
type Builder struct {
	engine                   sim.Engine
	freq                     sim.Freq
	log2PageSize             uint64
	pageTable                vm.PageTable
	migrationServiceProvider sim.Port
	maxNumReqInFlight        int
	pageWalkingLatency       int
}

// MakeBuilder creates a new builder
func MakeBuilder() Builder {
	return Builder{
		freq:              1 * sim.GHz,
		log2PageSize:      12,
		maxNumReqInFlight: 16,
	}
}

// WithEngine sets the engine to be used with the GMMU
func (b Builder) WithEngine(engine sim.Engine) Builder {
	b.engine = engine
	return b
}

// WithFreq sets the frequency that the MMU to work at
func (b Builder) WithFreq(freq sim.Freq) Builder {
	b.freq = freq
	return b
}

// WithLog2PageSize sets the page size that the mmu support.
func (b Builder) WithLog2PageSize(log2PageSize uint64) Builder {
	b.log2PageSize = log2PageSize
	return b
}

// WithPageTable sets the page table that the MMU uses.
func (b Builder) WithPageTable(pageTable vm.PageTable) Builder {
	b.pageTable = pageTable
	return b
}

// WithMigrationServiceProvider sets the destination port that can perform
// page migration.
func (b Builder) WithMigrationServiceProvider(p sim.Port) Builder {
	b.migrationServiceProvider = p
	return b
}

// WithMaxNumReqInFlight sets the number of requests can be concurrently
// processed by the MMU.
func (b Builder) WithMaxNumReqInFlight(n int) Builder {
	b.maxNumReqInFlight = n
	return b
}

// WithPageWalkingLatency sets the number of cycles required for walking a page
// table.
func (b Builder) WithPageWalkingLatency(n int) Builder {
	b.pageWalkingLatency = n
	return b
}

// Build returns a newly created GMMU component
func (b Builder) Build(name string) *GMMU {
	gmmu := new(GMMU)
	gmmu.TickingComponent = *sim.NewTickingComponent(
		name, b.engine, b.freq, gmmu)

	b.createPorts(name, gmmu)
	b.createPageTable(gmmu)
	b.configureInternalStates(gmmu)

	return gmmu
}

func (b Builder) configureInternalStates(gmmu *GMMU) {
	gmmu.MigrationServiceProvider = b.migrationServiceProvider
	gmmu.migrationQueueSize = 4096
	gmmu.maxRequestsInFlight = b.maxNumReqInFlight
	gmmu.latency = b.pageWalkingLatency
	gmmu.PageAccessedByDeviceID = make(map[uint64][]uint64)
}

func (b Builder) createPageTable(gmmu *GMMU) {
	if b.pageTable != nil {
		gmmu.pageTable = b.pageTable
	} else {
		gmmu.pageTable = vm.NewPageTable(b.log2PageSize)
	}
}

func (b Builder) createPorts(name string, gmmu *GMMU) {
	gmmu.topPort = sim.NewLimitNumMsgPort(gmmu, 4096, name+".ToTop")
	gmmu.AddPort("Top", gmmu.topPort)
	gmmu.migrationPort = sim.NewLimitNumMsgPort(gmmu, 1, name+".MigrationPort")
	gmmu.AddPort("Migration", gmmu.migrationPort)

	gmmu.topSender = sim.NewBufferedSender(
		gmmu.topPort, sim.NewBuffer(name+".TopSenderBuffer", 4096))
}
