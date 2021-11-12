// Code generated for package templates by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../templates/index.tmpl
package templates

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x58\xff\x8f\x1b\x47\x15\xff\x99\xfb\x2b\x1e\x9b\x36\xb6\x13\xef\xae\xcf\x49\x44\xe4\x78\x1d\x20\xa9\x4a\x25\x68\x4e\xc9\xf1\x03\x42\x28\x1a\xef\x3e\xdb\x73\x37\x9e\x99\xce\xcc\xda\x31\xd1\x49\x0d\xa8\x04\x50\x53\x04\xad\x38\x09\x45\x14\x21\x15\xaa\x0a\x55\x48\xb4\x51\x69\x89\xfa\xcf\xc4\xbe\xe6\x27\xfe\x85\x6a\xf6\x9b\xd7\x3e\xef\xe5\x92\x48\xa7\x5d\xef\xbc\xf7\x3e\x9f\x37\x6f\xde\xcc\xbc\x77\x5b\xdd\xef\x5e\xbf\x71\x6d\xf7\x67\x3b\xaf\xc1\xc8\x8c\x59\x6f\xab\x6b\x5f\xc0\x08\x1f\x06\xce\x2f\x47\x6e\xc8\x1d\x3b\x86\x24\xea\x6d\x01\x74\x0d\x35\x0c\x7b\xaf\x8b\x9d\xf6\x0e\xcc\xdf\x79\xf0\xcd\x67\xff\x9d\xff\xfe\x6f\x4f\xdf\xfe\xf0\x9b\x7f\xfe\xe6\xe9\xdf\xff\xb8\xf8\xfc\xde\xfc\x93\x3f\x3d\x79\xfc\xfe\xff\xff\xf7\xae\xd5\xb8\x76\xfd\x4d\xd8\x28\xef\xfa\x29\x90\x85\x1c\xa3\x21\x10\x8e\x88\xd2\x68\x02\x27\x36\x03\xf7\xb2\x03\xfe\x52\xc4\xc9\x18\x03\x67\x42\x71\x2a\x85\x32\x0e\x84\x82\x1b\xe4\x26\x70\xa6\x34\x32\xa3\x20\xc2\x09\x0d\xd1\x4d\x3e\x36\xd8\x45\xa8\x43\x45\xa5\xa1\x82\x97\x4c\x6f\xa1\x89\x25\xcc\x44\xac\x60\x42\x23\x14\x20\x38\x58\x87\x39\x9a\xa9\x50\xfb\x30\x10\x0a\x06\x0a\xd1\xdb\x80\xb8\x8f\xb3\xa9\x50\x91\x2e\xc1\xc9\xb6\x6c\xc2\x90\x89\x3e\x36\x73\x88\x26\x28\x24\xcc\xd0\x31\x36\x81\xd1\x09\x36\x61\x8a\x7d\x65\xc2\x66\x6e\x05\x11\xda\x71\x35\x5b\x5a\x14\x12\xaa\x8d\xa2\xfd\xd8\x3a\x9d\x4b\x73\x47\x18\xe5\xfb\x30\x52\x38\x08\x9c\x91\x31\x52\x77\x7c\x7f\x20\xb8\xd1\xde\x50\x88\x21\x43\x22\xa9\xf6\x42\x31\xf6\x43\xad\xaf\x0e\xc8\x98\xb2\x59\x70\x43\x22\x3f\x7f\x8b\x70\xed\x80\x42\x16\x38\xda\xcc\x18\xea\x11\xa2\x59\x01\x4d\x65\x23\xa1\x4c\x18\x1b\xa0\xa1\x0d\x58\x4a\x24\xe3\x3e\xa3\xa1\x3f\x20\x13\x3b\xea\x49\x3e\x74\xd6\xcd\x4a\x90\x2b\x36\xa1\xd6\x7e\x12\x18\x2f\xd4\x3a\xa7\x4b\xb4\xed\x2f\x00\x2f\x22\x86\xb8\x03\x26\xa6\x70\x37\x19\x00\x08\x05\x13\xaa\x03\x67\x06\x83\xc1\x95\x6c\xa8\x2f\x54\x84\xaa\x03\xdb\xf2\x0e\x68\xc1\x68\x04\x23\xcd\x48\xbd\xd5\x84\xd6\xab\x4d\xd8\x6e\xd9\xa7\x77\xb1\x91\xab\x8f\x89\x1a\x52\xee\xf6\x85\x31\x62\xdc\x81\xcb\xf2\x4e\x2e\x31\x78\xc7\xb8\x84\xd1\x21\xef\x40\x88\xdc\xa0\xca\x25\x49\xfe\x74\xe0\xc2\xa5\xd6\x52\x7b\x84\x74\x38\x32\x9d\xf6\x85\xd2\x98\x24\x51\x44\xf9\xd0\x35\x42\x76\x60\xbb\x10\x1c\xac\xcf\x46\xde\x5d\x71\xa6\xb3\xae\x69\x84\x74\x35\x86\xc9\x12\x7b\x4c\x0c\xc5\xdd\x15\x47\xb6\xbf\x57\xe2\xec\x93\x70\x7f\xa8\x44\xcc\xa3\x0e\xc4\x8a\xd5\xb3\xd8\x5a\x2b\xbb\x18\x0d\xe0\xc2\x55\x28\x91\x98\x25\x45\xd7\xcf\xa2\xdc\xf5\xd3\xcd\xdb\xed\x8b\x68\x06\xc9\x60\xe0\x94\x11\xfb\x8c\x84\xfb\x57\xd2\x15\x8d\xe8\x04\x42\x46\xb4\x0e\x1c\x9b\x8c\x84\x72\x54\xa0\xa5\x42\x12\x39\xb9\x6d\x44\xb5\x64\x64\xd6\x01\x2e\x38\xa6\x76\x6b\x96\x84\x4f\x88\x76\x8f\x01\xf4\xba\x7e\x44\x27\xc7\xf5\x4b\xa1\xc8\xd0\x00\xba\x24\x97\xda\x59\xe6\x49\x75\xa6\x90\xaf\x22\x58\x1d\xd7\xee\x74\x18\xd1\x08\x4b\x4a\x00\x3f\xa6\x13\x4c\xb6\xf6\x9b\xe9\x3e\x5a\xda\x2f\x9d\xb1\x1f\x24\xf3\x6b\xb3\x8b\x76\x2e\x4a\x30\xed\x12\x85\x64\xe9\x64\x49\xa3\x58\xf9\x95\x69\x6e\xd0\xb1\x9b\xfe\x58\x2c\xfb\x4c\x84\xfb\x57\xd2\xa5\x4f\x53\xb0\x62\xa6\x05\x84\x1b\x8a\x98\x1b\xa7\xe7\xae\xb0\x55\x29\xdb\xbc\x5f\x09\xcb\xfc\xd3\xbf\x2e\x0e\x1f\xcd\x1f\x7e\x7c\xf4\xe5\xd7\x55\x21\xa9\x98\x45\x1e\x8b\x0a\x0f\x8b\x50\x69\x26\xa6\xa8\x1c\x48\x4e\xfa\xc0\xb9\x89\x51\x1c\x22\x28\x61\x48\x92\xf4\x5a\x22\x46\x2b\x3e\x75\x69\x8e\x41\x73\xe3\x5e\xd7\xa7\xbd\xcd\xfe\x55\xb0\x26\xf8\x08\x24\x34\x49\x9c\x33\xf2\xd7\x38\xe9\x33\xf4\x23\xaa\xed\xbb\x70\xa2\x92\x3e\x45\x79\x01\xfa\x01\xd1\xa6\x34\xe9\x37\x78\xa8\x90\xe8\x53\x4f\x3b\x33\x7f\x81\x69\xa3\xb6\x67\x6f\x11\x6b\xfb\x55\x39\xb9\x54\xfa\xdc\x1c\x94\x0f\x44\x25\x68\x2a\xac\xc6\xac\x4c\xa7\x09\x2a\xbd\xba\x14\x93\x96\x77\xd1\xdb\xbe\x78\xdc\xb0\x62\x73\x26\x69\x1e\xda\x1d\x54\xbd\x33\x13\xb1\x6b\xaf\xe3\x12\x91\xeb\x76\x92\x3f\xf8\xe9\xee\xb5\x13\xd8\x96\x3f\x4a\xa8\xda\xae\xa7\x36\x34\xcc\xf7\x41\x97\xd1\x5e\x57\x4b\xc2\x8f\x6b\xb8\x76\xd8\xe9\x2d\x1e\xbe\xbd\xf8\xf3\xbf\x8f\x7e\xf7\xdb\xc5\xc3\x7f\x75\xec\x01\x2d\x09\xcf\x4c\x68\x14\x38\x7d\x34\xc4\xe9\x4d\x60\xdb\x6b\x79\xad\x5c\xea\x33\xda\x83\xef\x9c\x92\xe0\xe8\x3f\x5f\x1d\x7d\xf5\xe1\xe2\xf0\xd1\xd3\xc3\xcf\x36\x10\x58\xfd\xdb\x2a\xe6\x69\x14\xdc\x32\xc5\x29\x09\x7e\xb4\xbb\xbb\x03\x8b\xcf\xef\x3d\xbd\xff\x87\x2a\xfc\x30\xe2\x36\xe0\x2f\x86\x6f\xcf\xe9\x67\xe0\xcb\xb6\x7c\x71\xfc\xc4\xff\xf9\x17\xff\x98\x7f\xfa\xf8\x04\xff\xfb\xd3\x97\xf0\xfe\x64\x74\xd9\x96\x2f\x81\xfe\xe4\xf1\xd7\x47\x1f\x7c\x7c\xf4\xde\xfd\x2a\xf8\x9d\xf6\xce\x4d\x7b\xcc\xbc\x18\x43\x9a\x3e\xf3\xf7\x1f\xcc\xbf\xfc\xa0\x32\x7d\x50\x4b\xc1\x75\x29\x7f\x60\xac\xd7\x48\xc8\x5a\x7d\x3a\xa4\x66\x14\xf7\x93\xaa\x74\x4c\x23\xb1\xaf\xfd\xa1\x70\x65\xdb\x5e\xfa\x6a\x92\x1c\x96\x44\x0d\x6d\xed\x7f\xbb\xcf\x08\xdf\x2f\x2e\xc7\xb4\x10\x54\x18\x39\x09\xbf\xc2\x88\x2a\x0c\xed\x9d\xd7\xcb\x7a\x8f\xc5\x7b\x1f\xcd\xdf\xf9\x68\xf1\xee\xfd\xf9\x83\x43\xe8\xba\xf6\x16\xcf\x7d\xd9\xb8\x69\x29\xa7\x86\x12\xe6\x32\x41\xa2\xa2\x2e\x29\xd7\x0c\x5d\x3a\x1e\x16\xca\x63\x32\x44\x77\xaa\x88\x94\xd6\x49\xff\xf8\xc1\xb3\xaa\x51\x55\xde\x68\x49\x39\xb7\xf2\x0d\x67\x93\x9d\xcf\xf6\x09\x15\x83\x95\xb7\x9f\x21\xbf\xf0\x0c\xf9\xc5\x67\xc8\x2f\xad\x7a\x7e\xfc\xd4\x4b\x1b\xa8\x54\x3c\x21\x0a\x6c\x22\x4c\x35\x04\xf0\xf3\x5f\x5c\x29\x06\x43\xc1\xc3\x58\x29\xdb\xbe\x04\xd0\x5a\x8e\x2b\xc2\xa3\x7d\x9c\x2d\x07\xf6\x6d\xd7\x60\xd0\xde\x73\xb7\xb5\x21\xca\xea\xf3\x98\xb1\xcd\x1a\xc8\xa3\x93\xe4\x9b\x65\x37\x38\x9b\x41\x00\x46\xc5\xb8\x94\x8c\x63\x06\x01\x6c\x97\x5c\x9b\xea\x1f\x44\x91\x85\x70\x9c\x74\x90\x0e\xa0\xce\x44\x98\x5c\xd4\x9e\x54\xc2\x88\x50\x30\x08\x02\xc8\x52\xd9\x69\xe4\x45\x7a\xc9\x76\xaa\x6d\x8e\x3b\xe7\x0b\xc3\x91\xd0\xe6\xbc\xe3\x1b\x45\x42\xbc\x4a\xa3\x40\xb6\x65\xea\x6c\xc6\x72\x00\xc8\x34\x16\xed\x4e\x19\xe9\x39\x81\x92\x67\x28\xb8\x16\x0c\x6d\x0f\x51\x4f\xb1\x92\x3e\xa8\xeb\x67\xcb\x76\xf6\xad\x58\x24\x7d\x41\xb6\x8e\xa0\x55\x58\xb4\x67\x7b\x79\x77\xb6\xa7\xaf\x4e\x82\xcb\x36\x13\x4e\x34\xcb\xb7\x74\x18\x71\x6f\x4f\x27\x1d\xac\xf2\x38\x1a\x9f\xcb\xb1\xbf\xf7\x56\x8c\x6a\xf6\x7d\x46\xac\x97\x27\x40\xa5\x99\x34\x88\x79\xda\xfd\xc4\xf6\x40\xc7\xeb\xc4\x90\x7a\xa3\x88\x0a\x1d\xd4\xd7\xd2\xc0\x3d\x96\x38\xbd\x0b\xad\x56\xab\x05\x67\xcf\x82\x99\x49\x14\x2b\x16\x8d\x20\x70\x72\x86\xd2\xb2\x95\x73\xa7\x5e\xf4\x8b\x07\x5b\xd9\x8f\x3c\xbb\xd7\x13\xe3\xb9\x52\x03\x56\x21\xf2\x90\x1d\x5f\xd8\x21\x9a\x5b\x56\xb3\x20\x59\xcb\x8c\x0d\x40\xa7\xc4\xc9\xde\xaf\x78\x64\x8f\xdc\xa9\x2f\xf1\x62\xc5\x3a\x4b\xd0\x66\x31\x6e\x6f\xd4\xdd\x99\xc4\x0e\xd4\xf6\xb4\xe0\xb5\xa5\xc4\xd6\x0a\x22\x36\x1d\xb0\xb1\x5e\x0e\xeb\x38\x0c\x51\xeb\xce\x72\x15\xeb\x16\xa3\x51\xf2\x1d\xe0\x95\x7a\xed\x4c\xb9\x26\xa8\x35\xbc\x91\x19\xb3\xfa\x4f\x88\x19\x79\x49\xf3\x99\x18\x79\x76\x62\x70\x2e\xd9\xa0\xe7\x6c\x33\x0f\xbe\x7d\xe6\xaf\x46\xfa\x86\xf3\x50\x83\xd7\x6b\xc5\x9a\xad\x10\x64\x45\x41\x15\x81\x6c\xbf\x24\x7e\x52\x14\x54\xa1\x93\xc9\xd0\xd6\x15\x3f\x24\x3c\x4a\x1a\x38\x38\x07\x97\x4f\x49\xd7\x97\xba\x7a\x46\x27\x32\xee\xb4\x37\x13\x9e\x86\x2f\xa5\xda\xc4\x9a\xd7\x0f\x1b\x88\x37\x33\xfb\xb0\x31\x04\x0d\xeb\xc7\x2a\xf7\xab\x15\x13\xcd\xea\xd1\x9c\x31\x81\x4b\xeb\x82\x9b\xa9\xa4\x71\x65\xab\x64\xb7\xbc\x66\x82\x74\x69\x11\x95\x5e\x03\xf6\x8a\xbe\x17\xbc\xb5\x7e\x79\x85\x26\xb1\x2d\x79\x75\xb0\xcc\xef\x50\x8c\x25\x43\x83\xe5\x04\x5f\x4d\x6e\x8d\x66\x37\xdd\x1b\xf5\x2a\x15\x58\x39\xda\x4a\x82\x83\x26\xb4\x6d\x78\x4a\xd4\xf9\xce\xcd\xc6\x0e\xd2\x57\xd9\x3e\x0b\x83\xef\x2f\xfe\xf2\xab\xc5\xe1\xa3\x27\x5f\x3c\x78\x7a\xf8\xc9\xd1\xaf\xef\x65\xa3\xd6\xa1\x37\xb8\x41\x35\x21\xac\xf0\xa8\x9e\x9d\x4a\xbe\x5f\x04\x3d\xaf\xa2\x6a\x0d\x2f\xd4\xba\x5e\x4b\x8a\xac\x5a\xb3\x5a\xd4\x08\x82\x9a\x1a\xf6\xeb\xed\x4b\x97\x9a\x90\x3f\x1a\x35\xb8\x0a\x4e\x52\x9a\x75\xc0\x99\x8e\xa8\x41\xa7\x91\x53\x1d\x34\x2f\xb4\x5a\x8d\xf2\x1d\xb4\xd5\xf5\xfb\x22\x9a\xf5\xb6\xb6\xba\x7e\xf2\x5f\xe6\x6f\x03\x00\x00\xff\xff\xe4\x0d\xef\x47\x76\x16\x00\x00"

func indexTmplBytes() ([]byte, error) {
	return bindataRead(
		_indexTmpl,
		"index.tmpl",
	)
}

func indexTmpl() (*asset, error) {
	bytes, err := indexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"index.tmpl": indexTmpl,
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
	"index.tmpl": &bintree{indexTmpl, map[string]*bintree{}},
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
