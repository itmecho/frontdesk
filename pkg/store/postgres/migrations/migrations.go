// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// 1563449552_create_user_table.down.sql
// 1563449552_create_user_table.up.sql
package migrations

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

var __1563449552_create_user_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1563449552_create_user_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1563449552_create_user_tableDownSql,
		"1563449552_create_user_table.down.sql",
	)
}

func _1563449552_create_user_tableDownSql() (*asset, error) {
	bytes, err := _1563449552_create_user_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1563449552_create_user_table.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1563449552, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1563449552_create_user_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xca\xb1\x0a\xc2\x30\x10\x87\xf1\xbd\x4f\xf1\x1f\x5b\xf0\x0d\x9c\xaa\x64\x10\x15\xa4\x74\xe9\x24\x47\xef\xc4\x40\x92\xc6\x24\x67\x5f\x5f\x30\xb8\xe8\xe2\xfa\x7d\xbf\xfd\x60\xfa\xd1\x60\xec\x77\x27\x03\xcd\x92\x32\xda\x06\x00\x2c\x43\xd5\x32\x2e\xc3\xe1\xdc\x0f\x13\x8e\x66\xda\xbc\x47\x20\x2f\x78\x52\x9a\xef\x94\x10\x96\x82\xa0\xce\xd5\x25\x9e\xac\xfb\x79\xd0\x60\x1f\x2a\x95\x44\xca\x79\x5d\x12\x7f\x54\xad\x73\x12\x2a\xc2\x57\x2a\x28\xd6\x4b\x2e\xe4\x23\x58\x6e\xa4\xae\x20\x2c\x6b\xdb\x55\xa7\x91\xff\x72\x2c\x4e\xbe\x5d\xd3\x6d\x5f\x01\x00\x00\xff\xff\x3c\x7c\x21\xb4\xed\x00\x00\x00")

func _1563449552_create_user_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1563449552_create_user_tableUpSql,
		"1563449552_create_user_table.up.sql",
	)
}

func _1563449552_create_user_tableUpSql() (*asset, error) {
	bytes, err := _1563449552_create_user_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1563449552_create_user_table.up.sql", size: 237, mode: os.FileMode(420), modTime: time.Unix(1563713317, 0)}
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
	"1563449552_create_user_table.down.sql": _1563449552_create_user_tableDownSql,
	"1563449552_create_user_table.up.sql":   _1563449552_create_user_tableUpSql,
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
	"1563449552_create_user_table.down.sql": &bintree{_1563449552_create_user_tableDownSql, map[string]*bintree{}},
	"1563449552_create_user_table.up.sql":   &bintree{_1563449552_create_user_tableUpSql, map[string]*bintree{}},
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
