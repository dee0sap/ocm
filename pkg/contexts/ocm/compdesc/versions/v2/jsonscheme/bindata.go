// Code generated for package jsonscheme by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../../../../../resources/component-descriptor-v2-schema.yaml
package jsonscheme

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

var _ResourcesComponentDescriptorV2SchemaYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5b\x6f\xdb\xb6\x17\x7f\xf7\xa7\x38\x68\x02\xd0\x69\x22\x3b\xf1\xff\xdf\x01\xf5\x4b\x90\xb5\xd8\x50\x6c\x6b\x86\xb4\xdb\xc3\x12\xaf\xa0\xa5\x63\x9b\x99\x44\x7a\x24\xed\x44\xbd\x7c\xf7\x81\xa4\xa8\x8b\x2d\xcb\x72\x9c\xb6\x1b\xd0\xbc\x44\xa4\x0e\xcf\x8d\xbf\x73\x21\xe5\x43\x16\x0d\x81\xcc\xb4\x9e\xab\x61\xbf\x3f\xa5\x32\x42\x8e\xb2\x17\xc6\x62\x11\xf5\x55\x38\xc3\x84\xaa\x7e\x28\x92\xb9\xe0\xc8\x75\x10\xa1\x0a\x25\x9b\x6b\x21\x83\xe5\x80\x74\x0e\x1d\x45\x89\xc3\xad\x12\x3c\x70\xb3\x3d\x21\xa7\xfd\x48\xd2\x89\xee\x0f\x4e\x07\xa7\xc1\xd9\x20\x63\x48\x3a\x9e\x0d\x13\x7c\x08\xe4\xc7\x4c\x2a\xbc\xf0\x72\xe0\x65\x2e\x07\x96\x03\x28\x96\x4d\x18\x67\x66\x95\x1a\x76\x00\x12\xd4\xd4\xfc\x07\xd0\xe9\x1c\x87\x40\xc4\xf8\x16\x43\x4d\xec\x54\x55\x44\x6e\x01\x14\x16\xd8\xf5\x11\xd5\xd4\x2d\x90\xf8\xf7\x82\x49\x8c\x1c\x47\x80\x00\x88\x93\xfb\x3b\x4a\xc5\x04\x77\x54\x73\x29\xe6\x28\x35\x43\xe5\xe9\x2a\x44\x7e\x32\x57\x49\x69\xc9\xf8\x94\x74\x3a\x00\x31\x1d\x63\xbc\x51\xdf\x1a\xf1\x9c\x26\x48\x8a\xe1\x92\xc6\x0b\xb4\x9c\x72\x6b\x5e\xd3\x04\x2b\x1c\xbd\x38\x33\x95\xd0\xfb\x9f\x91\x4f\xf5\x6c\x08\x83\x67\xcf\x9c\xf6\x54\x6b\x94\xc6\x21\x7f\x5e\xd3\xe0\xfd\x69\xf0\xbc\x77\x13\x8c\x8e\xaf\x7b\x23\x33\x1c\x7d\x18\x9c\xfc\xff\x53\xff\x3a\x70\xaf\xfa\xef\x7a\xa3\xa7\x87\x56\x20\x8b\x90\x6b\xa6\xd3\x0b\xad\x25\x1b\x2f\x34\xfe\x84\xa9\x93\x9b\x30\x9e\x0b\xd9\x20\x62\xd4\xbd\x0e\xde\x1d\x67\xcf\x4f\xfd\xe4\xd1\xb9\x63\x2d\x31\xa6\xf7\x18\xbd\xc1\x64\x89\xd2\xf1\x3c\x00\x4d\xff\x42\x0e\x13\x29\x12\x50\xf6\x85\xc1\x12\x50\x1e\x01\x8d\x6e\x17\x4a\x63\x04\x5a\x00\x8d\x63\x71\x07\x94\x83\xb0\xdb\x4c\x63\x88\x91\x46\x8c\x4f\x81\x2c\xc9\x09\x24\xf4\x56\xc8\x40\xf0\x38\x3d\xb1\x4b\xed\xb8\x97\x30\x9e\xcd\x7a\x59\x33\xa6\x20\x41\xca\x15\xe8\x19\xc2\x44\x18\xae\x86\x89\xf3\xa5\x02\x2a\xd1\x88\x82\x25\x8d\x59\x54\xd5\x57\x79\x85\xcf\x7a\x83\xde\xff\xca\xcf\xc1\x44\x88\xe3\x31\x95\xd9\xdc\xb2\x4c\xb0\xac\xa3\x38\xeb\x0d\xfc\x53\x4e\x56\xa2\xcf\x1f\x2b\xcb\xca\xce\x5e\x8e\xce\xbb\xa7\x1f\xaf\xcf\x82\xe7\xa3\x9b\xe8\xe9\x51\xf7\x7c\x78\xd3\x2b\x4f\x1c\x9d\xd7\x4f\x05\xdd\xee\xf9\xb0\x98\xfc\x78\x13\xd9\x3d\xba\x08\xfe\x08\x46\xd7\xa7\xc1\x73\xff\xec\x59\xb6\x24\x3e\xf2\x12\x8f\xbb\xe5\x17\xc7\x96\x49\x65\xc6\x52\x1e\x92\x3a\x18\xd7\x41\x6f\x63\x04\x65\xa1\x99\x9a\xa0\x50\x43\xf8\x00\x87\x12\x27\x43\x20\x07\xfd\x52\xde\xe8\xd7\x41\x99\xc0\x27\x07\xc5\xb9\x50\x4c\x0b\x99\xbe\x10\x5c\xe3\xbd\xde\x25\x58\x0d\xd5\xa6\x14\x61\x39\x34\x64\x06\x11\xb2\xab\x7a\xd9\x34\x8e\x2f\x27\x85\x94\x5a\x8b\xd6\xd4\x2e\x72\xc6\xaa\x9e\x56\xd3\x31\x55\xf8\x9b\x8c\x49\x3e\xb7\xae\xb0\xf9\xcb\xc8\xca\x53\xb5\x69\xc6\xfd\x55\x52\xd2\x2f\x74\x3e\x67\x7c\xda\x72\x29\x00\xf2\x45\x32\x84\x6b\xb2\x90\xf1\xaf\x54\xcf\xc8\x09\x10\x35\xa3\x83\x67\xdf\x05\x11\x9b\xa2\xd2\x64\xd4\x59\xe1\xb3\x2b\x67\xeb\xe1\x29\x53\x5a\xa6\x64\x64\x5c\x4e\xc3\x10\x95\x6a\x59\x3d\x8c\x2b\x2c\x15\x4c\x84\xcc\x96\xa2\x82\xae\x19\xe1\xbd\x46\x6e\x52\xbf\x3a\xda\x82\x8d\x0e\xc0\x94\xe9\xd9\x62\x7c\xd1\x2c\xbb\x11\x5c\x76\x68\x76\xbc\xb4\x83\x76\x66\xf2\x20\xf0\xad\xfa\xc9\x29\x98\xfb\x3b\x13\xb4\x65\xb9\x01\x65\x33\x45\x28\x92\x84\xe9\xa6\x10\xe0\x82\xe3\x3e\x7e\xd9\xd3\xee\xd7\x82\xa3\x03\x86\x12\x0b\x19\xe2\xcb\x3c\xbe\x76\x50\xc7\xd4\xeb\x7c\xb0\x74\x0d\x41\x3e\x36\x1c\xf2\x81\x83\xd0\x06\xc5\x79\x5e\xd4\x1b\x14\x6f\x9f\xdb\xb2\x25\x78\xaf\x25\x7d\x95\x11\x0c\x77\xe4\xe3\x99\x2c\x57\xbb\x9c\x0d\x09\xa9\x54\x22\xc9\x03\x61\x98\x63\xd0\xb6\x4d\x6a\x6d\x29\x95\x92\xa6\xc5\x4a\xa6\x31\xa9\xa4\xaf\x5a\xcd\x2c\x2f\xbf\xa8\x9c\x02\xec\x98\xa7\x97\x93\x32\x8b\x0d\xf9\xd6\xad\x23\xdb\x09\xcb\xd1\xde\x82\xdc\xb4\xd0\x9e\xd8\x00\x51\x86\x57\x3e\xb0\xb6\x66\x28\x6a\x82\x10\x25\xf2\x10\x6d\x67\x04\xdd\xa2\x67\x8f\x45\x48\xe3\xa3\x0c\xd8\x9b\xa2\xc5\x6f\xf9\x1b\x8c\x31\xd4\x42\x3e\x14\x21\x9f\x61\xb7\xca\x0d\xef\x95\xb7\xf2\xa1\x7e\xc9\x39\xb5\xed\xba\x2b\x85\xad\xdc\x8d\x37\x9f\x0a\x6a\x5a\xf4\x8d\x76\xd6\x8a\x68\xca\x02\x70\x00\x34\xd4\x0b\x1a\xc7\xe9\xb0\x90\x14\xd8\x0a\x75\xd7\x07\x35\xc7\x90\xd1\x18\x24\x1a\xfa\xd0\x0a\xf9\xef\x26\x8e\xcf\x83\x28\x89\x2e\x1e\xde\xe6\x69\x69\xc7\x2e\xc0\x33\x50\xad\x8f\x6f\x19\x60\xe0\xc0\xae\xb7\x51\x59\x70\x39\xc9\xce\x21\x0b\xa5\x21\xa1\x3a\x9c\x95\x90\xaa\xd6\x8a\xc9\x7a\x43\x10\x53\x9d\xa3\xd1\x4e\x95\xb3\xd4\xb7\x1a\x93\x5b\xe5\xb2\xea\x23\xc1\xc9\x31\x2b\xda\x20\xb7\x09\xad\xab\x9c\x85\x80\x69\x76\x4d\x0f\x29\x39\x8d\xbf\x7a\xcd\x6b\x59\xf1\x36\x90\x89\x90\x7d\x1f\x8b\xb5\x82\xb7\x81\xda\x5a\xff\x03\x8b\x51\xa5\x4a\x63\xb2\xeb\xca\xcb\xaa\x30\x77\x8a\x7a\x95\xd0\xe9\x5e\x6d\xa4\x1d\x32\xc3\x25\xaf\x34\x8f\xd2\x5f\x56\xcf\x1f\xd9\xf6\x55\xc4\x6c\x39\x1f\x16\xb6\xb6\x34\xac\x62\x56\x00\x24\xa6\xa9\x0f\x94\xfd\x6c\x01\x92\xa9\x43\x60\x54\x77\x00\xa8\x26\xcd\x0b\xa3\x7c\xb5\x08\xeb\x19\x42\x42\x39\x9b\x98\x93\x5d\xb3\xd0\x04\x23\x46\xdf\x56\x94\xab\xb2\x7f\x6b\x78\x19\x22\x97\x98\xc5\xc4\x72\x77\x5e\x71\x19\xd5\x61\xd7\x69\xa0\x40\x8b\x2d\x12\xdd\x81\xb3\x49\x9c\xa3\xf0\xa2\x34\x95\x53\xd4\x18\x41\x68\xce\xde\x7c\x9b\x41\x8a\xbd\x6f\xb4\xc5\xbc\x07\xc6\x61\x9c\x6a\x54\x5e\xc6\xd8\x38\x7b\x95\x2f\x5f\x24\x63\xb3\xa1\x1d\x80\x8d\x91\xb4\x47\x0c\x4c\x58\x8c\x45\x01\xdb\x17\x31\x35\x1a\x16\xe8\xf1\xa2\x36\xf9\xc5\xbf\x2f\xbb\x03\xf4\x8c\x6a\x60\xca\xda\x6e\xdc\xcf\xb8\x7d\xf7\xc4\xbc\x54\x4f\x20\x62\xd2\xb6\xb1\xe9\xc6\xfd\xf0\x7e\xbb\x7c\x40\x6c\x7d\x21\x87\x5d\xae\xc6\x59\x33\x38\xab\xc0\xb4\xf1\x0e\x77\x4c\xcf\x32\xd7\x84\x0b\x29\x91\x6b\xa8\xbb\x0a\x6f\xf2\x92\x4f\xab\x57\x59\xab\xb2\xcf\x0d\x76\xb9\x67\xae\x73\xe2\xb7\xa6\x65\x7b\x1d\xb1\x9b\xf1\xe5\x3b\x85\x4d\x15\xbf\x54\x72\x2d\x5c\x8a\x53\xec\x1e\xc1\xb4\xf0\x97\x5b\x7b\x96\x5d\xa3\x4c\xee\xaa\x45\xc3\x45\x56\x07\x60\x8a\x1c\x25\x0b\xbf\xe2\x25\x54\xa6\x81\xbb\x87\xca\x06\xdf\xa2\xee\x5f\x10\x75\xc5\xc6\xb8\xf9\xaf\x1b\x74\x15\xa0\x56\xaf\x47\x5a\xdf\x8a\xec\x7c\x0d\xb2\x0e\xa2\xb5\xcf\x0e\xaa\xf4\x72\x2e\xc5\x92\x45\x85\xbb\x03\x20\x95\xe3\x72\xf5\x6a\x25\x6f\x82\x55\x85\x7f\x65\xc5\x36\x60\xb6\xbf\x59\xd9\x03\x35\xeb\x36\xef\x0c\x82\xb5\x5b\xc6\xa6\xe3\xd4\xda\x57\x21\x02\x07\xbe\x90\xc7\xe9\x09\xdc\x21\x08\x1e\xa7\xd9\x97\x50\xdb\xef\x0a\xee\xaf\x99\xfd\x1e\x6c\x81\xf9\x67\x03\x73\xb6\x7d\x8f\x74\xd4\x5e\xb9\x96\x2f\x7d\x59\x58\xc5\xd0\xe3\x08\x5c\x67\x5c\x80\xe0\xa1\x96\xb5\xdf\xfb\xf2\xf5\x14\x69\x09\x96\x4a\x97\xd6\x6a\xd1\x4a\x8d\xb1\xb9\xa4\xde\xa5\xf0\xe1\x53\xa7\xd3\x59\x49\x2c\xe5\xac\x11\x00\x49\xd0\xfd\x94\xa2\x1c\xd9\xa4\x53\x8d\xdb\xe2\x27\x1b\xb5\x0a\x79\x16\x2b\x09\xad\x79\x83\x48\xe7\x9f\x00\x00\x00\xff\xff\x6a\xe0\xcf\x9c\xc5\x22\x00\x00")

func ResourcesComponentDescriptorV2SchemaYamlBytes() ([]byte, error) {
	return bindataRead(
		_ResourcesComponentDescriptorV2SchemaYaml,
		"../../../../../../../resources/component-descriptor-v2-schema.yaml",
	)
}

func ResourcesComponentDescriptorV2SchemaYaml() (*asset, error) {
	bytes, err := ResourcesComponentDescriptorV2SchemaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../../../../../resources/component-descriptor-v2-schema.yaml", size: 8901, mode: os.FileMode(436), modTime: time.Unix(1644329280, 0)}
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
	"../../../../../../../resources/component-descriptor-v2-schema.yaml": ResourcesComponentDescriptorV2SchemaYaml,
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"..": &bintree{nil, map[string]*bintree{
				"..": &bintree{nil, map[string]*bintree{
					"..": &bintree{nil, map[string]*bintree{
						"..": &bintree{nil, map[string]*bintree{
							"..": &bintree{nil, map[string]*bintree{
								"resources": &bintree{nil, map[string]*bintree{
									"component-descriptor-v2-schema.yaml": &bintree{ResourcesComponentDescriptorV2SchemaYaml, map[string]*bintree{}},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}},
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