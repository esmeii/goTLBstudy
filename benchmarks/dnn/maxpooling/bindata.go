// Code generated by go-bindata.
// sources:
// maxpooling.hsaco
// DO NOT EDIT!

package maxpooling

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _maxpoolingHsaco = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdf\x6f\xdb\xd6\x15\x3e\xf7\x92\xa2\x28\x4a\xa6\xe5\x1f\x8d\x3d\x37\xc8\x18\x2f\xa8\x07\xa3\x56\x65\x5a\x11\xb4\xa0\xc0\xe2\xd8\xb3\x53\xd4\x4e\x9c\x7a\x4b\x96\x15\x85\xc1\x98\xb4\x24\x47\x22\x05\x89\x4e\xec\x62\x52\xe4\x6c\x89\x6d\xa0\xc0\xba\xcc\x28\x82\x22\xa8\x5f\x37\xa0\x7f\xc0\x9e\x6c\x19\x7d\xdb\x53\xe2\xd7\xf5\xa1\xc0\xd0\xff\x60\x1b\xf6\x64\x0d\x97\xbc\xd7\x12\xb5\xc8\xfb\x95\xa1\xa8\xc6\x0f\xd0\x3d\xba\xdf\xb9\xe7\xdc\xef\xdc\x7b\x95\x84\x21\xf9\xe0\x47\x73\x33\x18\xa1\xcb\x40\xc1\xc1\x57\x80\xc8\x17\xc5\xed\x33\x47\x46\x75\xed\xa8\xc3\xa5\x40\x84\xcb\x10\x01\x09\x04\x00\xe0\x9b\xc6\xb5\xda\x1a\xf2\x5a\x91\xf2\x88\xc6\xb5\x43\xb4\xcf\x6b\x21\xda\x88\x0b\x34\xe9\x6b\xb5\x9f\xf3\x5e\xdb\x1c\x47\xb4\x42\x9c\xf2\x2d\xb6\x00\x5e\xcb\xe2\xf0\xbf\x19\xc7\xea\x7b\xef\x6b\x5b\xe7\xff\x85\xb8\x66\x7d\x04\x37\xbe\xb6\x75\xe1\x94\x75\x69\x07\x9e\x7e\xd6\xb1\xdb\x6f\xb5\x4f\xa3\x5e\x7b\xda\xda\x33\x3d\x43\xa4\xf2\x96\x75\x45\x4d\x1f\xa1\x91\x4a\x21\xfd\xf7\xbf\x7b\xf8\x09\xa6\x6b\x86\x5a\x13\xce\x6b\xeb\x0b\x96\x95\x9b\xb1\x8a\xf7\xb5\xa2\xee\x04\x89\x74\xdc\xe4\xfc\xf4\x49\x0c\xe1\x5f\x77\xce\xa1\xcb\xf3\x10\x3c\x59\x53\xc6\x4d\xce\x4f\xcf\x2e\xfc\xc4\x1d\x7b\x1e\x00\xc2\x94\xd7\xf2\x7a\x7a\xd9\x1c\xd3\xf2\x3a\xf9\x64\x4a\x1a\x31\xe9\xec\x87\xcb\xb9\xb1\xf4\xca\x7a\x2a\x3e\x41\xf3\xcb\x51\x00\x89\xc6\x8c\x8d\x8d\x49\x37\x8d\x62\x29\x6b\x99\x97\x14\x86\xf7\x95\xf1\x37\x95\xb8\xf2\x81\xf4\xae\x51\x34\x8d\x5c\xa9\xe1\x91\x14\x65\x4c\xb9\xa6\xe5\x8d\x06\xa5\x28\x8a\xb7\x36\x89\x50\x8b\x1b\xf9\x3b\x56\xae\x69\xe4\x88\x77\xd0\xe5\xbb\xfa\x88\x33\x70\x4e\x33\xd3\x6b\x5a\xba\x91\xf0\x7a\xc1\x30\xa7\xe6\x94\x29\x8f\xf7\x44\xa2\x23\x4d\x55\x3e\x70\xbc\x93\xc5\x74\xc9\x23\x44\x72\xcd\x4b\x24\x9a\x76\xa6\x68\x68\x7a\x49\x62\xc4\x8f\x37\x0a\x86\x67\x54\xd6\xb4\x4f\x9c\x8b\xd9\x0f\xbd\xe1\x89\x13\xd7\x64\x2e\x9b\x36\x2f\xbd\xd4\x75\x53\xcb\xad\x19\xef\x66\x4d\x9d\xb9\xaf\x6c\x38\x94\x77\x00\x99\x98\x0d\x78\x67\x42\x6d\x24\x5e\x5e\xbe\xb1\xa6\xe5\x1a\xa9\xa7\x8d\x15\x6d\x2d\x67\xb7\xaf\xe9\x8e\x65\xdb\x56\x7e\x49\xd7\x6c\xad\x7d\x59\x23\x2b\x39\x4b\xb3\x47\x47\xda\xd7\x96\x6a\x5f\x5b\xaa\x7d\x6d\xb3\x39\xeb\x8e\x96\xbb\xb2\xb6\xb2\x62\x14\xdb\x17\x38\xd3\x5c\xa0\xae\x17\x17\x0b\xda\xb2\xc1\xca\x74\x53\xfc\x17\x0b\x60\xae\xe5\x3b\x69\x3f\x97\x33\x9a\x49\x7e\x6e\x9d\x54\x53\xc6\xc8\xa6\x33\x76\x27\x55\x74\x3f\xab\xdb\x99\x4e\x2a\xa8\x60\x59\x39\x43\x5f\xea\xbc\x9d\xa2\x85\x75\xdc\x86\xdd\x75\xfe\x56\x5e\xea\xc4\x9a\xee\x77\x52\x4d\x25\xbb\x98\xd5\x8d\xce\xda\x27\x5a\x53\x47\xed\x53\x41\xd3\x3b\x6b\x93\x48\x41\x1d\xb5\x43\xb6\x55\xf8\xff\xf9\x67\xee\x2b\x52\x7e\x35\xab\xeb\x86\xe9\x4e\x7e\x7d\x65\xa5\x64\xd8\x3f\x3d\x65\x7b\x92\x89\xff\xfd\xfc\xb7\xbf\xe1\xf9\x7f\xf6\xcf\xe7\x9f\xb2\x74\x63\xa1\x68\x15\x4e\x2e\x2b\x69\x08\xb9\x10\xd6\x8a\xe9\x45\x23\x9d\x37\x4c\xdb\x55\x38\x1e\x67\x9a\x67\x8b\xd6\x5a\x81\xfa\x66\xb2\xeb\x86\xee\x0e\x88\x53\xf7\x42\x31\x7b\x4f\xb3\x8d\xf6\x03\xbc\xd9\x69\x91\xac\xb4\x5b\xda\x3d\x63\xa5\x68\xb1\x59\x15\xe5\x64\xa9\xae\xad\xe5\x17\x67\x17\xde\x6b\x5c\x02\x4f\x34\x79\x6e\x7a\x3c\xe3\xec\x7c\xce\x6b\xeb\x33\x39\xcd\xbe\x65\x15\xef\xba\xaa\x9d\xa4\xea\xc5\xa4\x14\x8b\xc5\xa4\xff\xe0\x3f\x5e\x7c\xf8\xf0\xe1\xc3\x87\x0f\x1f\x3e\x7c\xf8\xf0\xe1\xe3\x5b\x82\xc6\xfd\x5c\xd1\xbd\xc3\xfa\x0f\x37\x6e\xbd\xc0\xe8\x73\x78\xee\xdc\x73\xf5\x5e\x2c\x65\x9a\xbe\x0f\x43\xc4\xe3\xe3\x79\x5e\xa8\xd7\xeb\xf5\x57\xa9\xfb\x55\x01\x03\xae\xf1\x00\xf0\x0b\xc0\x35\xd2\x9f\x02\x5c\xbb\x0a\x00\x0f\xe0\xa3\x03\xa8\xc3\x63\xa2\x5a\x04\xf1\x89\x08\xa0\x22\x00\x15\xc3\x76\x79\x55\xa9\xee\xff\x11\xb6\x0f\xbe\x00\x5c\xbb\x40\x56\x0d\xe3\xda\xf7\x49\x2c\xc6\xb5\x37\x49\x2e\x8c\x6b\x71\x00\xf8\x02\xe3\x5a\x82\xf8\x39\x5c\x4b\x11\x3f\x87\x6b\x6f\x13\x3f\x27\xd4\x2e\x3b\x6b\x89\x6b\x11\x67\xd5\x85\x1a\x59\xfd\x29\x84\x6b\x83\xa4\xcf\xe3\xda\x59\x32\x0e\xe1\x9a\x42\xb5\x74\x1d\xf6\x6e\x22\x2c\x54\x00\xf3\x15\x6e\x0f\x3e\xe6\x00\x6d\x22\x38\xb3\x2d\xec\xc1\xc7\x02\xa0\xcd\x68\x20\xf8\xa4\xda\xf5\xeb\xcd\xde\xd0\x93\xcd\x28\xff\x9b\xfd\x2a\xda\xdd\x47\x30\xb0\x7d\x26\x82\x2b\x7b\x20\x0e\x0f\x44\xe4\x0a\x9a\xc5\x95\x3a\xc6\x12\x40\xf5\x3a\xea\xc6\x95\x00\x3c\x3a\x42\x63\x00\x02\x3c\x76\x6c\x55\x8a\x26\x00\xfe\xf0\xbc\x1a\xc1\x10\x00\x78\x11\xe8\x21\xf6\xf1\x51\x80\xc3\xc0\xa3\x88\x2a\x84\xc4\xd1\x00\x17\x51\x03\x1c\x4e\x22\x80\x17\xa8\x0b\x03\x22\xb1\x21\x0c\xc1\x59\xa9\x22\xd0\x7c\x9b\x58\x56\x0f\x71\x54\xad\x4b\x92\x33\x17\xdf\xf5\xac\x1c\x02\x78\x51\x7d\xf8\x3b\x24\x84\xc4\xe4\x19\xf1\x93\x32\x4f\xfb\x7c\x9f\xf8\x46\x55\xdc\x2d\x07\xba\xa5\x0a\x07\xe2\x68\x90\xc3\x50\xed\xdd\x2d\xa3\x1e\x0c\x7b\xa2\x38\x8c\x42\x78\x14\x85\x22\x09\x04\x8f\x8e\x02\x6f\x01\xf0\x44\xcf\x5b\x00\x7b\x11\x79\xb8\x8a\xa9\x5e\x91\xe8\x80\x17\x24\xc6\xd1\x13\xc6\x10\xec\x12\x55\xbe\x5b\x1c\x45\x61\x59\x45\x61\xaa\xb7\xbb\xa1\xd7\xa9\x9d\xe4\x0b\x3f\x2b\x07\xa9\x96\x40\x48\x4c\x0e\x34\x6b\xeb\x16\xdf\xd8\xc4\x92\xea\xe8\xe3\x44\x38\xc4\x58\xad\xca\xbb\x65\x9e\xc3\x20\x44\xc4\xd1\x86\x3e\x9c\xe0\x49\xbe\x2e\x80\x20\x3c\x3a\x0a\x62\x0c\x75\x2c\x55\xea\xf5\x07\x75\xbe\x4b\x4c\x12\x1f\xff\x1a\xc0\x9e\x1c\x1d\x0e\xc2\xaf\x8e\x1e\xca\x18\x22\xa2\x98\x14\x45\xac\xe2\x6e\x5e\x0d\x60\x3c\x58\x15\xc5\x21\xae\x47\x48\xf1\xdc\xc3\xf2\xaa\xf2\x74\xff\x22\x6c\x1f\x10\x8d\x42\x90\xec\xcf\xa3\xa3\xa8\x88\xa1\x8a\x3e\xdd\x0f\x20\x29\x49\xf8\xc0\x00\x40\x97\x24\x27\x43\x72\x54\xad\xca\xf2\x50\x34\x1a\x1d\x04\x38\x7c\x1e\xec\xc1\xd0\x23\xc9\x6a\xaf\x24\x25\x08\x2f\x74\x47\xd4\xe0\x2d\xb9\xd2\x27\x49\x43\xc1\x70\x34\xc1\x74\x81\xf2\x6c\xbf\x1f\xb6\x0f\xc8\x3a\x86\x88\xa6\x08\x06\xdc\xd7\xab\x72\xaf\xf5\xa7\xaa\xe8\xb3\x7d\x91\x93\x2b\x00\x0b\x5f\x86\x00\xc2\x87\xb2\xac\xfe\xb2\xb7\x57\xad\xf6\xf7\xa7\x48\xfd\xab\xe7\x3e\xdb\x2a\xc0\x47\x07\xe1\xf0\xf6\xcf\x03\x03\x12\x54\xce\x1d\xef\xfc\xb9\xbe\x73\x50\x39\x77\xbc\x55\x39\x7b\xbc\xb5\x29\x8a\x2a\xcf\xed\x96\x57\x87\x3e\xdd\x8a\x46\x22\x6a\x65\xe8\x78\xe7\x4f\xc4\x3f\x74\xbc\x55\x19\x3c\xde\xda\x03\x3c\x2c\x90\x39\x01\x43\x37\xc6\x15\x39\x12\x51\x51\xb7\x9c\xe2\x9d\xdf\xd7\xe3\xf2\x6a\xff\xee\x16\x40\xe1\x4b\x21\x00\x50\xe9\x3f\xde\xb9\x56\xdf\x39\x00\xd8\x3c\xf8\xa6\xff\xac\xf0\xe1\xc3\x87\x0f\x1f\x3e\x7c\xf8\xf0\xe1\xe3\xdb\x8e\x93\x67\xcd\xe9\x73\xe2\x61\xda\x1d\xa4\x36\x40\x6d\x86\xfa\xd9\x55\x3f\x7d\x8c\x1c\xfe\x72\x5c\xb7\x88\x8d\x53\x3f\x7b\xae\x7c\x01\xbf\x7c\xbe\xb9\xac\x79\xd7\x28\x5e\x52\xe6\xe6\xa6\x95\x44\x2c\x0e\xcb\x39\xcd\x4c\x2b\xf7\xdc\xc7\x9b\x09\xa3\x9c\xae\x17\x51\xd5\x6f\x9f\xf7\xf2\x41\xca\x17\x2e\x78\x79\x99\xf2\xbf\x6d\xe1\x5f\xa3\xfc\x57\x2d\x79\x5e\xa7\xfc\xb3\xef\x79\xf9\xf3\x94\xff\x5b\x0b\x3f\x42\x79\x68\xc9\x3f\x46\xf9\xb3\x2d\xfc\x04\xe5\x17\x5a\xf8\xab\xa4\xc1\xc1\xc6\x7b\x02\x14\x3f\x68\xf3\x1c\x3e\xc4\xf4\x0d\xb3\xb4\x91\x87\x58\xda\x5c\x8b\x65\xb4\x52\x06\x68\x4b\x78\xbb\x08\x31\xd3\xb2\x0d\x88\xd9\xc6\xba\xed\x70\x5a\x3e\xbb\x0c\xb1\x65\x2b\x9f\x37\x4c\x1b\x62\xa5\x8d\xbc\xad\xdd\x81\x58\x29\x53\xb2\x8b\xee\x37\xd7\xc2\x95\x2b\xf1\xa5\x71\xb7\x8d\xbb\xc6\xed\xa9\x4e\x3b\xe1\xb4\x09\xa7\xbd\xe8\xb4\x49\xa7\x4d\xb5\x3e\xe0\xbf\x34\x7d\xfb\xda\xe4\xfc\x3b\x53\xf0\x6a\x80\xe8\xd9\x3c\x39\x56\xd8\x6b\xd9\xb2\xf1\x74\x2c\x7b\x67\x80\x9d\xe3\x10\x00\xfc\xb5\x5e\xb7\x58\x18\x3b\xaf\xcc\x2a\x4d\xf3\x40\x53\x3c\x43\x1f\xfd\x2d\xb0\x78\x76\xbe\x99\x1d\x6c\x89\xe7\x5b\xec\x77\xe8\xfb\x0b\x2c\x9e\xfd\x9e\x98\x8d\xbe\xa4\xde\x66\x9c\xa7\x67\x9c\xc5\xb7\x7b\xbf\x03\x5a\xe6\x65\x18\xf1\xbe\xb3\xd1\xf6\xbd\x99\x76\x02\xc6\x68\x2c\xc7\x88\x36\xef\xb3\xb0\x34\x6c\xfd\x98\xac\x24\x4d\xd9\x72\xbc\xa1\x40\x89\x0b\x6d\xa6\x67\xf6\x87\xcd\x7b\xdf\x84\xa7\x34\x5e\xa1\x03\xc3\xf4\xd3\xba\xff\xb3\xcd\xda\x9b\xf0\xfb\x71\xd7\xde\x6e\x33\x3f\xc3\x8d\x36\xf1\x3d\x6a\xc3\x7f\x5a\xfc\xdf\x03\x00\x00\xff\xff\x65\xe2\xad\x3b\x68\x35\x00\x00")

func maxpoolingHsacoBytes() ([]byte, error) {
	return bindataRead(
		_maxpoolingHsaco,
		"maxpooling.hsaco",
	)
}

func maxpoolingHsaco() (*asset, error) {
	bytes, err := maxpoolingHsacoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "maxpooling.hsaco", size: 13672, mode: os.FileMode(493), modTime: time.Unix(1544059272, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"maxpooling.hsaco": maxpoolingHsaco,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"maxpooling.hsaco": &bintree{maxpoolingHsaco, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

