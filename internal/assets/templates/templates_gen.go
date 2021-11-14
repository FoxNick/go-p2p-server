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

var _indexTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x58\x6d\x8b\x1b\xc9\xf1\x7f\xfd\xdf\x4f\x51\xff\xf1\xad\x25\x61\x8d\x46\xbb\xf6\x26\x46\xd6\xc8\x18\xdb\x1c\x07\x89\x6d\x62\xdf\x8b\x10\xc2\xd2\x9a\x29\x8d\xda\xea\xe9\x9e\xeb\xee\x91\xac\x98\x85\xbb\x04\x5f\x1e\x38\x07\x92\x3b\xb2\x21\x98\x38\x04\x8e\x1c\x47\x12\x02\x71\x4c\x72\x17\x73\x5f\xc6\xd2\xda\xaf\xf2\x15\x42\xcf\x93\x66\xb4\x1a\xaf\x77\x03\x8b\x66\xa6\x1e\x7e\x55\x5d\x55\xdd\x55\xbd\x5b\xfd\xff\xbf\x71\xfb\xfa\xbd\xef\xdf\xb9\x09\x63\x1d\xb2\xc1\x56\xdf\x3c\x80\x11\x1e\xb8\xd6\x8f\xc6\xb6\xc7\x2d\x43\x43\xe2\x0f\xb6\x00\xfa\x9a\x6a\x86\x83\x77\xc5\x9d\xdd\x3b\xb0\x78\xf4\xf8\xd5\xb3\x7f\x2d\x7e\xf1\x87\xd7\x1f\x3e\x7d\xf5\xa7\x8f\x5f\xff\xf1\x57\xcb\x7f\x7c\xb4\xf8\xf2\xd7\x2f\x5f\x7c\xfa\x9f\x7f\x7f\x62\x24\xae\xdf\xb8\x05\x1b\xf9\x7d\x27\x05\x32\x90\x21\x6a\x02\xde\x98\x48\x85\xda\xb5\x62\x3d\xb2\x2f\x5b\xe0\xac\x58\x9c\x84\xe8\x5a\x53\x8a\xb3\x48\x48\x6d\x81\x27\xb8\x46\xae\x5d\x6b\x46\x7d\x3d\x76\x7d\x9c\x52\x0f\xed\xe4\x63\x83\x9e\x8f\xca\x93\x34\xd2\x54\xf0\x92\xea\x5d\xd4\x71\x04\x73\x11\x4b\x98\x52\x1f\x05\x08\x0e\xc6\x61\x8e\x7a\x26\xe4\x04\x46\x42\xc2\x48\x22\x76\x36\x20\x4e\x70\x3e\x13\xd2\x57\x25\xb8\x68\x37\x6a\x43\xc0\xc4\x10\xdb\x39\x44\x1b\x24\x12\xa6\x69\x88\x6d\x60\x74\x8a\x6d\x98\xe1\x50\x6a\xaf\x9d\x6b\x81\x8f\x86\x2e\xe7\x2b\x8d\x82\x43\x95\x96\x74\x18\x1b\xa7\x73\x6e\xee\x08\xa3\x7c\x02\x63\x89\x23\xd7\x1a\x6b\x1d\xa9\x9e\xe3\x8c\x04\xd7\xaa\x13\x08\x11\x30\x24\x11\x55\x1d\x4f\x84\x8e\xa7\xd4\xd5\x11\x09\x29\x9b\xbb\xb7\x23\xe4\x17\xee\x12\xae\x2c\x90\xc8\x5c\x4b\xe9\x39\x43\x35\x46\xd4\x15\xd0\x94\x37\x16\x52\x7b\xb1\x06\xea\x99\x80\xa5\x86\xa2\x78\xc8\xa8\xe7\x8c\xc8\xd4\x50\x3b\x11\x0f\xac\x75\xb5\x12\x64\x45\xc7\x53\xca\x49\x02\xd3\x31\x0e\x4d\xdd\x87\x0f\xa1\x33\x45\xa9\xcc\xca\x0e\x0e\x72\xfb\x89\xba\x79\x03\xe8\xf8\x44\x13\x7b\xc4\xc4\x0c\x1e\x26\x04\x00\x4f\x30\x21\x7b\x70\x6e\x34\x1a\x5d\xc9\x48\x43\x21\x7d\x94\x3d\xd8\x89\x1e\x80\x12\x8c\xfa\x30\x56\x8c\x34\xbb\x6d\xe8\x6e\xb7\x61\xa7\x6b\x7e\x3b\x97\x5a\xb9\x78\x48\x64\x40\xb9\x3d\x14\x5a\x8b\xb0\x07\x97\xa3\x07\x39\x47\xe3\x03\x6d\x13\x46\x03\xde\x03\x0f\xb9\x46\x99\x73\x92\x82\xea\xc1\xc5\xbd\xee\x4a\x7a\x8c\x34\x18\xeb\xde\xee\xc5\x12\x2d\x22\xbe\x4f\x79\x60\x6b\x11\xf5\x60\xa7\x60\x1c\xac\xaf\x26\x7a\x58\x71\xa6\xb7\x2e\xa9\x45\x64\x2b\xf4\x92\x9c\x77\x98\x08\xc4\xc3\x8a\x23\x3b\xdf\x2e\xd9\x1c\x12\x6f\x12\x48\x11\x73\xbf\x07\xb1\x64\xcd\x2c\xd8\x46\xcb\x64\xa7\x05\x5c\xd8\x12\x23\x24\x7a\x65\xa2\xef\x64\x51\xee\x3b\xe9\x6e\xee\x0f\x85\x3f\x87\x84\xe8\x5a\x65\xc4\x21\x23\xde\xe4\x4a\x9a\x62\x9f\x4e\xc1\x63\x44\x29\xd7\x32\xd5\x49\x28\x47\x09\x2a\x92\x48\x7c\x2b\xd7\xf5\xa9\x8a\x18\x99\xf7\x80\x0b\x8e\xa9\xde\x9a\x26\xe1\x53\xa2\xec\x63\x00\x83\xbe\xe3\xd3\xe9\x71\xf9\x52\x28\x32\x34\x80\x3e\xc9\xb9\x66\x95\x79\x95\x9d\x2b\xf8\x55\x04\x23\x63\x9b\xad\x0f\x63\xea\x63\x49\x08\xe0\x3b\x74\x8a\xc9\x5e\xbf\x95\x6e\xac\x95\xfe\xca\x19\xf3\x41\x32\xbf\x36\xbb\x68\xd6\x22\x05\x53\x36\x91\x48\x56\x4e\x96\x24\x8a\xcc\x57\x96\xb9\x41\xc6\x9c\x02\xc7\x62\x39\x64\xc2\x9b\x5c\x49\x53\x9f\x96\x60\xcd\x4a\x0b\x08\xdb\x13\x31\xd7\xd6\xc0\xae\x58\xab\x13\x36\x75\x5f\x09\xcb\xe2\xaf\xbf\x5f\x1e\x3e\x5f\x3c\xf9\xe2\xe8\xab\x6f\xea\x42\x52\xb3\x8a\x3c\x16\x35\x1e\x16\xa1\x52\x4c\xcc\x50\x5a\x90\x1c\xfd\xae\xf5\x3d\xf4\x63\x0f\x41\x0a\x4d\x92\xa2\x57\x11\xa2\x5f\xf1\xa9\x4f\x73\x0c\x9a\x2b\x0f\xfa\x0e\x1d\x6c\xf6\xaf\xc6\x6a\x82\x8f\x40\x3c\x9d\xc4\x39\x33\x7e\x93\x93\x21\x43\xc7\xa7\xca\x3c\x0b\x27\x6a\xcd\xa7\x28\x67\x30\x3f\x22\x4a\x97\x16\xfd\x1e\xf7\x24\x12\xf5\xd6\xcb\xce\xd4\xcf\xb0\x6c\x54\xe6\x30\x2e\x62\x6d\xbe\x6a\x17\x97\x72\x4f\x6d\x83\xf2\x91\xa8\x05\x4d\x99\xf5\x98\xb5\xe5\x94\xb5\x87\x12\xf0\xb4\xdb\xb9\xd4\xd9\xb9\x74\x5c\xb1\x66\x73\x26\x65\xee\x99\x1d\x54\xbf\x33\x13\xb6\x6d\xfa\x73\xc9\x90\x6d\xf7\x92\x3f\x78\xff\xde\xf5\x37\x58\x5b\xbd\x94\x50\x95\xc9\xa7\xd2\xd4\xcb\xf7\x41\x9f\xd1\x41\x5f\x45\x84\x1f\x97\xb0\x0d\xd9\x1a\x2c\x9f\x7c\xb8\xfc\xcd\xdf\x8e\x7e\xfe\xb3\xe5\x93\x3f\xf7\xcc\x01\x1d\x11\x9e\xa9\x50\xdf\xb5\x86\xa8\x89\x35\x98\x56\x5b\x66\x2e\xe5\x30\x3a\x80\xff\x7b\x4b\x43\x47\x7f\xff\xfa\xe8\xeb\xa7\xcb\xc3\xe7\xaf\x0f\x9f\x6d\x30\x64\xe4\xf7\x65\xcc\xd3\x68\xd8\x65\x13\xa7\x32\xf0\xea\xd9\xd3\x57\x2f\x5e\xd4\x19\x50\x73\xc5\x84\x39\xf5\xcf\x62\xe0\xfa\x9d\xf7\x73\xf4\x1a\x78\x2f\x8a\xcf\x0e\xbf\xf8\xf8\xd1\xe2\x2f\xbf\x7d\xf9\xe2\x9b\xa3\xcf\xbe\xa8\xb5\x10\x62\x78\x76\x0b\xef\x8a\x5b\x71\x08\x3d\xa8\x03\x0f\x84\x14\xb1\xa6\x1c\xf7\x79\x1c\xfe\x2f\x49\x58\x7c\xfa\x78\xf1\xd5\x67\xb5\x59\x46\x15\x09\xae\x4a\x69\x86\x50\xad\x19\x21\x6b\xf3\x65\x40\xf5\x38\x1e\x26\x53\x65\x48\x7d\x31\x51\x4e\x20\xec\x68\xd7\xf4\x68\x39\x45\x79\x6e\xfb\xe6\xb7\xb6\x2f\xdf\xdc\xbe\xb6\xb7\x7d\x73\x6f\xfb\xf2\xde\xf6\xb5\x3d\x0b\x34\x91\x81\x19\xe7\xf7\x87\x8c\xf0\x49\xd1\xde\xd2\x51\x4e\xa2\x6f\x25\x2e\x49\xf4\xa9\x44\xcf\x74\xad\x01\x2c\x7f\xf9\xf9\xe2\xd1\xe7\xcb\x4f\x7e\xba\x78\x7c\x08\x7d\xdb\x74\xe0\xdc\xb1\x8d\x1b\x8e\x72\xaa\x29\x61\xb6\xc9\x49\x31\x53\x94\xfb\x7d\x9f\x86\x41\x21\x1c\x92\x00\xed\x99\x24\x51\x64\x4e\x63\xe7\xf8\xa1\x51\x95\xa8\x1b\x4d\x54\x44\x39\x37\xfc\x0d\xe7\x8a\x59\xc9\xce\x1b\xba\xbd\xe1\xef\x9e\xc0\xbf\x78\x02\xff\xd2\x09\xfc\xbd\xaa\xe7\xc7\x4f\xac\xf4\x36\x94\xb2\xa7\x44\x82\xa9\x8a\x99\x02\x17\x7e\xf0\xc3\x2b\x05\xd1\x13\xdc\x8b\xa5\x34\x77\x11\x17\xba\x2b\xba\x24\xdc\x9f\xe0\x7c\x45\x98\x98\x2b\x80\x46\xd3\xa3\xf6\x95\x26\xd2\xc8\xf3\x98\xb1\xcd\x12\xc8\xfd\x37\xf1\x37\xf3\x6e\x73\x36\x07\x17\xb4\x8c\x71\xc5\x09\x63\x06\x2e\xec\x94\x5c\x9b\xa9\x6b\xbe\x6f\x20\x2c\x2b\x25\xd2\x11\x34\x99\xf0\x92\x26\xdb\x89\xa4\xd0\xc2\x13\x0c\x5c\x17\xb2\xba\xb6\x5a\xf9\x80\x5d\xd2\x9d\x29\x53\xf0\xd6\x85\x42\x71\x2c\x94\xbe\x60\x39\x33\x75\x95\xfa\x6e\xb4\x1b\xa5\x9e\x66\x26\x0e\x00\x99\xc2\xe2\x9e\x52\x86\x39\x0d\x4a\xf2\xeb\x09\xae\x04\x43\x33\xf9\x37\x53\xa0\xe4\xf6\xd2\x77\xb2\x84\x9d\xff\x20\x16\xc9\x34\x9f\x65\x10\x94\xf4\x8a\x5b\xd6\xfd\xfc\x92\x75\x7f\xc3\x1d\x6b\xf0\xd6\x18\xf7\x3f\x88\x51\xce\x4f\x09\x92\x96\xd2\x28\xe6\xe9\xd5\x25\x8e\x7c\xa2\xf1\x06\xd1\xa4\xd9\x2a\x22\x43\x47\xcd\xb5\x3a\xb0\x8f\x55\xce\xe0\x62\xb7\xdb\xed\xc2\xf9\xf3\xa0\xe7\x11\x8a\x8a\x46\xcb\x75\xad\xdc\x42\x29\x6f\xe5\xe2\x69\x16\x97\xbd\x83\xad\xec\x25\x2f\xef\xf5\xca\x38\x55\x6d\x40\x15\x22\x3f\x13\x8f\x27\x37\x40\x7d\xd7\x48\x16\x46\xd6\xaa\x63\x03\xd0\x5b\xe2\x64\xcf\x77\x3a\xe4\x3e\x79\xd0\x5c\xe1\xc5\x92\xf5\x56\xa0\xed\x82\x6e\xe6\x9b\x7b\xf3\x08\x7b\xd0\xb8\xaf\x04\x6f\xac\x38\xa6\xc1\x8b\x58\xf7\xc0\xc4\x7a\x45\x56\xb1\xe7\xa1\x52\xbd\x55\x16\x9b\x06\xa3\x55\xf2\x1d\xe0\x9d\x66\xe3\x5c\xb9\x8f\x37\x5a\x9d\xb1\x0e\x59\xf3\xbb\x44\x8f\x3b\xc9\xcd\x31\x51\xea\x18\xde\x3e\x99\x06\xfb\x11\xca\x16\x5c\x80\x06\x6c\x37\x8a\xdc\x54\x80\xb2\x8e\x5d\x07\xe4\x45\xf1\xc9\x18\x59\x4f\xae\xc3\x08\x31\x3c\x19\xa3\xd2\x7a\x73\xa4\x44\xbd\xc2\xd9\xac\x9c\x8d\x4d\x15\xb5\xb4\x2f\xe6\xac\xd6\x95\xad\x92\xe2\xea\x68\x75\x13\xd9\x08\x51\xaa\x35\xe4\x4e\x71\x4f\x83\xce\xda\xfd\xae\x62\x27\xd1\x2d\xb9\x75\xb0\x4a\xa9\x27\xc2\x88\xa1\xc6\x72\x4e\xab\xf9\x54\xa8\xef\xa5\xe5\xd0\xac\x13\x81\xca\x6e\x2e\x31\x0e\xda\xb0\xb3\xd7\xed\x96\x4d\xe7\xc5\x9a\xd1\x0e\xd2\x47\x59\x3f\x0b\x83\xe3\x2c\x7f\xf7\xe3\xe5\xe1\xf3\x97\xff\x7c\xfc\xfa\xf0\xcb\xa3\x9f\x7c\x94\x51\x8d\x43\xef\x71\x8d\x72\x4a\x58\xe1\x51\x33\xdb\x88\x8e\x53\x44\x3d\x9f\x19\x1a\xad\x8e\xa7\x54\xb3\x91\x8c\x14\x8d\x76\x3d\xab\xe5\xba\x0d\x19\x0c\x9b\xbb\x7b\x7b\x6d\xc8\x7f\x5a\x0d\xb8\x0a\x56\x32\x88\xf4\xc0\x9a\x8d\xa9\x46\xab\x95\x9b\x3a\x68\x5f\xec\x76\x5b\xe5\xd3\x77\xab\xef\x0c\x85\x3f\x1f\x6c\x6d\xf5\x9d\xe4\xdf\xa4\xff\x0d\x00\x00\xff\xff\x7f\x56\xcf\x94\x37\x15\x00\x00"

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
