// Code generated by go-bindata.
// sources:
// ../api/schema.json
// DO NOT EDIT!

package parse

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x93\x41\x73\x83\x20\x10\x85\xef\xfc\x0a\x87\xf6\xd8\xd4\xa4\xd3\x5c\xf2\x2b\x7a\xef\xf4\x40\xe2\xc6\x6c\x46\xd0\x2e\x98\xa9\xd3\xc9\x7f\xaf\x40\x4b\xb4\x22\x4e\x6e\xf2\x58\xde\xfb\x7c\xe2\x37\xcb\x32\xfe\xa8\x0f\x27\x90\x82\xef\x32\x7e\x32\xa6\xd9\xe5\xf9\x59\xd7\x6a\xe5\xd5\xe7\x9a\xca\xbc\x20\x71\x34\xab\xf5\x6b\xee\xb5\x07\xfe\x64\xcf\x61\x61\x8f\xe4\x7e\x61\xba\x06\xec\xb2\xde\x9f\xe1\x60\xbc\x26\x8a\x02\x0d\xd6\x4a\x54\x6f\x54\x37\x40\x06\x41\xf7\x33\x47\x51\x69\x70\x03\xcd\x50\xb6\x2c\xbd\x56\xe1\x05\xc2\x2a\x6e\xec\xf4\x25\x73\x37\x14\x09\xf0\x21\x70\x81\x6a\x24\x0d\x92\x54\x2b\xf7\x40\x21\xc9\xed\x49\xf1\x85\xb2\x95\xfd\xf6\x66\xac\xa3\xfa\xd5\xd7\x41\xbe\xde\x26\xb8\xee\xb4\x01\xa9\xe7\xa2\x04\x91\xe8\xc6\x49\x18\x99\x9f\xaf\xe1\x8e\x32\xd2\x95\x24\x8a\x59\xae\x27\x5d\x52\xa2\xaa\x49\x61\xb7\xf7\x21\xd0\x31\xc2\x01\x08\x2a\x03\xe5\x12\xc9\xcb\x76\x9b\x64\xd9\x2c\xb2\xb4\xad\xbb\xe8\x09\x10\x6d\x08\x55\xc9\xff\x3b\xb1\x84\x2f\x27\xf8\x6c\x91\xc0\x3a\xbf\xc7\x3f\xc2\x6c\x29\x71\xc2\x91\xf8\xc1\x62\x1c\x7f\x4f\x81\x25\x4e\x31\xc9\x0f\xb7\x98\x0d\xdd\xad\x9b\x73\x9a\xba\xf8\x9f\x98\xd9\xc9\x2b\xfb\x09\x00\x00\xff\xff\xac\x6a\xb4\x24\x64\x04\x00\x00")

func schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaJson,
		"schema.json",
	)
}

func schemaJson() (*asset, error) {
	bytes, err := schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.json", size: 1124, mode: os.FileMode(420), modTime: time.Unix(1438180577, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"schema.json": schemaJson,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"schema.json": &bintree{schemaJson, map[string]*bintree{
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

