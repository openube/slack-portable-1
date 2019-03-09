// Code generated by go-bindata. DO NOT EDIT.
// sources:
// res/Slack.lnk (1.798kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _slackLnk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x94\x4f\x68\x2b\x55\x14\xc6\x7f\x69\xa3\xb1\xa2\x90\x6a\x10\x5b\x2a\xcd\x26\x22\xca\x8c\xc9\x34\x6d\x32\x41\xa1\x25\x7f\x48\xe9\xd0\x84\xc4\x40\x91\x11\xbc\x4e\xa7\x34\x34\x21\x63\x12\xa1\x75\x21\xb8\x2a\x96\xba\xd1\x45\x2b\x28\xad\x22\xb4\x82\x06\xeb\xb6\x52\x5c\x5a\x17\xd2\x45\xa5\x6a\x41\xad\xe8\x42\x41\x0b\xba\x12\x5c\xc8\xdc\x24\x36\xed\x2b\x3c\x78\xbb\xf7\x78\x1f\xcc\x9c\xb9\xe7\xcf\x77\xbe\x3b\x97\x7b\x0c\xc0\x13\xe8\xc3\xc5\x17\xf2\x4d\xe6\x64\x12\x82\xc0\xf9\xa7\x7d\x83\x87\xad\x63\xcf\xbd\x27\x6d\x2b\x76\xf6\x7f\xf9\xb2\x75\xec\xf1\x2f\x78\x65\xa2\x87\xcb\x38\x25\xc0\x68\xfe\xa7\xdc\x51\xf0\xf7\x44\xd9\xff\xc1\xb7\xf7\xf1\x54\x38\xfc\xde\x10\x4f\x27\x13\x26\x37\x22\x4f\x44\x5a\x6d\xd6\xf1\xf9\x69\x54\x84\xb5\xc4\x33\x0c\xe0\xe5\xfc\x40\x9b\xad\xf9\x5c\xbf\x0a\x6c\x8e\xb9\xf2\xee\xb9\x52\x9d\x78\xe0\x4f\x1a\x54\x10\x58\x2c\x01\x01\x4c\x34\x5c\x69\xda\x6c\xd6\x17\x6c\xf3\xa9\xf6\xb2\x4d\xea\x12\x67\x4d\x72\xbe\x2f\x39\xbd\x57\x38\xb7\xde\xcd\xf5\x70\xaa\xd8\x2c\x63\x03\x8f\x02\x53\xc0\x48\x67\xd3\xae\x55\x3a\x35\x93\xc0\x20\xd0\x0f\x24\xf7\x47\x4c\xbf\xeb\x4c\x26\x4c\xd9\xdf\xbc\x50\xc1\x83\x14\xff\xa7\x0e\x92\xc2\xa6\xc1\x12\x4d\x6a\x38\x3c\x44\x92\x04\x66\x4f\x73\x13\x81\x83\x83\xc2\x18\xaa\x7c\x62\x3c\xd9\xc9\x2a\xd1\xc0\xa6\x4e\x03\x13\x8b\x3a\x82\x57\x59\xc1\x64\x4a\x56\xa4\x10\x34\x11\x98\x18\xd4\xb0\x10\x54\xae\x65\x56\x29\x63\x51\x23\xd0\x0f\x3e\xd8\x0e\x95\x8a\xe9\x42\xbe\x90\xcb\x4c\x1b\xe9\x90\x39\xe5\x38\x29\xd1\x14\xa6\x51\xb3\x44\xa5\xb3\x19\xe1\x38\x6a\xd9\xaa\x5d\x73\x94\x77\x0a\x42\x94\x28\x92\xa6\x40\x9e\x02\x39\x32\x4c\x63\x90\x26\x74\xcb\x7f\xf7\x2e\x6e\x47\xbc\xd8\xbe\xd0\xdb\x73\x9d\xb5\x63\x29\x8b\x62\xde\x6e\x74\xe3\xf3\xad\x56\x73\xe5\xaf\x8f\xd3\x6f\xbd\xfe\xeb\x47\xdf\xef\xae\x7a\xbf\x7b\x7e\xe3\x70\x6f\xf8\xb7\xc1\x77\x8c\xe7\x3e\xeb\xbf\x7f\xcd\xb8\x59\x7c\xdd\x03\x03\xb0\xfd\x06\x10\x29\xe6\x8b\x3f\xaf\xcd\x65\x3e\x37\xe2\xc9\xfd\x7f\x1f\x7e\xfb\xf1\xcd\xea\xd7\xd5\xee\x70\x1a\x05\xdc\x79\x55\x44\x21\x82\xc2\x38\x0a\x9a\xfc\x1a\x23\x86\x4e\x94\x28\x3a\x3a\x13\xc4\x99\x90\x31\x9d\x71\xc2\x44\xe5\x5a\x23\x2e\x33\xc7\x89\x12\x26\x4c\x84\x98\xac\x98\x90\x6c\x6d\x4f\x1b\x07\x1d\x21\xa5\x27\x8c\xad\x95\x2d\x7d\x66\xf7\xe8\xec\x58\x39\xfb\xe6\xef\x17\x80\xe1\xae\x90\x10\xa0\x23\xd0\xd1\x98\x47\x10\x27\x86\x22\x05\x2c\xa0\x49\x71\x11\x04\x51\x14\x04\x1a\x2f\x49\xb1\x31\x22\xc4\x65\xd3\x28\x36\x0b\x58\x44\xa4\x1c\x17\xe9\xee\x58\x77\xc9\x87\x40\x5e\x98\x2a\x2a\x0d\x5e\xe6\x15\xca\xd4\xa9\x63\x53\x91\x9e\x8b\xa9\xdc\x3b\xf5\xdb\xd0\x3b\xea\xf7\x1e\xa9\xa6\x3e\x79\xd3\xc9\xee\x64\x27\xd5\x0f\x9f\x5d\x5e\x7f\x0c\x58\x74\x13\xb2\xc0\x8f\x3f\x0c\xff\xb1\xfa\xcf\x79\x76\xa3\x75\x3a\xe3\xfb\xea\xb5\x70\xef\x79\xff\x17\x00\x00\xff\xff\x71\x6c\x57\x8d\x06\x07\x00\x00")

func slackLnkBytes() ([]byte, error) {
	return bindataRead(
		_slackLnk,
		"Slack.lnk",
	)
}

func slackLnk() (*asset, error) {
	bytes, err := slackLnkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Slack.lnk", size: 1798, mode: os.FileMode(438), modTime: time.Unix(1547773182, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc5, 0xed, 0x4a, 0x31, 0xef, 0x53, 0x4a, 0x6c, 0x45, 0x4c, 0x6e, 0xed, 0x5, 0x25, 0x87, 0xdc, 0xe3, 0xc, 0xda, 0xed, 0x9, 0x34, 0xcf, 0x33, 0xcd, 0x51, 0xb2, 0xc9, 0xb0, 0x97, 0x9b, 0xef}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"Slack.lnk": slackLnk,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"Slack.lnk": &bintree{slackLnk, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}