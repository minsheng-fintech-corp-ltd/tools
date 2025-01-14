// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// assets statically implements the virtual filesystem provided to vfsgen.
var assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰CompressedFileInfo{
			name:             "/",
			modTime:          time.Date(2019, 7, 15, 11, 11, 14, 868923683, time.UTC),
			uncompressedSize: 2730,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x51\x8f\xdb\x36\x12\x7e\xb6\x7e\xc5\x9c\x92\x87\xbb\x83\x4f\x4e\xf2\x76\xbe\xcb\x2d\x7c\x9b\xdd\xd6\x68\xb0\x2e\xe2\x4d\x83\xa0\x4d\x81\x31\x35\x92\x26\xa1\x48\x86\xa4\xec\x18\x4d\xfa\xdb\x8b\x21\x65\x5b\x9b\xf6\xb1\x5d\x60\x1f\x2c\x0e\x3f\x7e\xfc\xbe\x6f\x86\x8b\x05\x5c\x5b\x77\xf4\xdc\x76\x11\x9e\x3d\x79\xfa\x6f\x58\x87\xc8\x16\x56\x43\xec\xac\x0f\xc5\x62\x51\x2c\x16\xf0\x92\x15\x99\x40\x35\x0c\xa6\x26\x0f\xb1\x23\x58\x39\x54\x1d\x9d\x56\xe6\xf0\x03\xf9\xc0\xd6\xc0\xb3\xea\x09\xfc\x5d\x0a\xca\x71\xa9\xfc\xc7\x7f\x04\xe2\x68\x07\xe8\xf1\x08\xc6\x46\x18\x02\x41\xec\x38\x40\xc3\x9a\x80\x3e\x29\x72\x11\xd8\x80\xb2\xbd\xd3\x8c\x46\x11\x1c\x38\x76\xe9\x9c\x11\xa5\x12\x8c\xb7\x23\x86\xdd\x45\x64\x03\x08\xca\xba\x23\xd8\x66\x5a\x08\x18\x47\xd2\xf2\xd7\xc5\xe8\x96\x8b\xc5\xe1\x70\xa8\x30\x11\xae\xac\x6f\x17\x3a\x97\x86\xc5\xcb\xf5\xf5\xcd\xdd\xf6\xe6\x5f\xcf\xaa\x27\xe3\xa6\xd7\x46\x53\x08\xe0\xe9\xe3\xc0\x9e\x6a\xd8\x1d\x01\x9d\xd3\xac\x70\xa7\x09\x34\x1e\xc0\x7a\xc0\xd6\x13\xd5\x10\xad\x90\x3e\x78\x8e\x6c\xda\x39\x04\xdb\xc4\x03\x7a\x12\x98\x9a\x43\xf4\xbc\x1b\xe2\x03\xcd\x4e\x14\x39\x3c\x28\xb0\x06\xd0\x40\xb9\xda\xc2\x7a\x5b\xc2\xff\x57\xdb\xf5\x76\x2e\x20\x6f\xd6\xf7\xdf\x6e\x5e\xdf\xc3\x9b\xd5\xab\x57\xab\xbb\xfb\xf5\xcd\x16\x36\xaf\xe0\x7a\x73\xf7\x62\x7d\xbf\xde\xdc\x6d\x61\x73\x0b\xab\xbb\xb7\xf0\xdd\xfa\xee\xc5\x1c\x88\x63\x47\x1e\xe8\x93\xf3\x72\x03\xeb\x81\x45\x4d\xaa\x93\x74\x5b\xa2\x07\x14\x1a\x9b\x29\x05\x47\x8a\x1b\x56\xa0\xd1\xb4\x03\xb6\x04\xad\xdd\x93\x37\x6c\x5a\x70\xe4\x7b\x0e\xe2\x6a\x00\x34\xb5\xc0\x68\xee\x39\x62\x4c\x9f\x7e\x77\xaf\xaa\x90\x92\xfb\xb3\xb1\x35\x05\xe5\x79\x47\x21\x15\xe1\x1e\x59\x27\x15\x1b\x26\x5d\x87\x33\x85\x96\x8c\x75\xc8\xa0\xac\x69\xb8\x1d\x7c\x82\x4f\x08\x89\xf9\xc6\xc9\x6f\xd4\xa7\x6d\xe8\x09\xd8\xd4\xac\x50\xb4\x4b\x39\x41\xf8\x38\x50\x48\xdb\x7a\xf4\x1f\x32\x8f\xde\xd6\x83\x26\x50\xd6\x7b\x0a\xce\x9a\x3a\x88\x61\x72\xe0\x37\x56\xd4\xb9\x7e\x7d\x73\xaa\xb1\x0d\x20\x78\x72\x36\x70\xb4\xfe\x38\x87\x43\xc7\xaa\x13\x9b\x62\x37\xda\xe9\x49\xc9\x12\x38\x4f\x0d\x7f\x92\x04\xa7\x6c\x38\x54\x1f\xb0\xa5\x90\x78\xb0\x49\xf0\x23\xa8\xa4\xa3\x77\xd6\x47\x20\x54\x1d\x58\x71\x27\x5d\xe8\xd6\x8a\x4b\xd8\x3b\x4d\xf3\x24\x42\xcb\xb1\x1b\x76\x95\xb2\xfd\x82\xe5\x16\x0b\x74\x3c\x9f\x42\x71\x80\x32\xad\x54\x79\xb1\xac\x8a\xbc\xb2\x04\x49\x91\x69\x8b\xc2\x3a\x32\xe8\xf8\x6a\x09\xbf\x14\xb3\xc5\x02\x02\xe9\xe6\xda\x1a\xe9\x13\xaa\x4f\x2e\x0b\xcf\x8e\x52\x4c\x50\x6b\xf0\xd4\x90\x27\xa3\x28\x40\xe8\xec\xa0\x6b\xd8\x89\xb4\x4a\x0f\xf5\xa8\x2c\x9b\x04\x26\x54\x36\x8e\xcc\xea\xfb\x35\xd8\x21\xba\x21\xce\xc1\x93\x1a\x7c\xe0\x3d\xe9\x63\x55\xcc\x1e\x1c\xb7\x84\x9d\xb5\x1a\x3e\xc3\x3f\x1b\xd4\x81\x8a\x84\x71\x2b\xee\xdd\xb2\x8e\xe4\xa1\xa6\x86\x0d\x85\x24\x7a\x3b\x68\x3c\xa7\x56\x1c\x14\x33\xf4\xd9\xed\x68\xc1\xf6\x1c\xa1\xf1\xb6\x4f\x6e\x08\x56\xe6\x50\xc1\x3a\x8a\x34\xd6\xe8\xa3\x6c\xb1\x87\xdc\x93\x4d\x3e\xe4\x04\xd0\x61\x04\xac\x6b\xf9\xe7\x9c\xa4\x84\xa1\xac\x09\xd1\x23\x9b\x18\xaa\x4c\x6e\xac\x3d\x65\x0b\x76\x18\x58\x41\x3c\x3a\x0a\xa0\xd0\xc8\xe0\xda\x11\x78\xea\xed\x9e\xea\xf1\xf0\x04\x85\x06\xc8\x7b\xeb\x93\x99\x61\x50\xdd\x84\xfc\x2e\xcd\xb7\xac\xe8\xee\x78\x9e\x7a\x51\xb2\x70\x96\x25\x87\xfa\xe3\x80\x5a\x4c\x1a\x0b\x89\x3d\x6c\x76\xef\x49\xc5\xc4\xa1\x4a\xb9\x61\x13\xa2\xcc\xc7\xf9\x59\x8b\x9e\x0d\xf7\x43\x9f\x8f\x3c\xcd\xc2\xa0\x3a\xea\x11\x6c\xde\xce\x61\x82\x8d\x01\xb6\x69\x75\x31\xee\xac\x8a\x59\x73\xf1\xe6\xea\x1c\xa9\x2f\x63\x37\x53\xb2\x23\x90\x4a\xfd\x75\x49\x52\xa0\x28\x73\x2f\xf7\x71\x4b\x86\xa4\x71\x4d\x2b\x6a\x60\xdb\x7a\x6a\x45\xc3\x53\x6a\xa4\x9b\x8b\x94\x24\x8c\x22\xbd\x04\x25\x64\xe0\x4c\x35\x47\xa2\x4e\x9f\xc6\x4b\x5c\x9a\x8e\x4c\xf4\x4c\xa1\x2a\x50\xeb\x73\xc2\x2d\x3a\xbe\x65\x4d\x06\x7b\x1a\x5b\x35\x1d\x93\x7e\x4f\x66\x8b\xf0\xa2\xfa\xab\xfc\x56\xc5\x6c\xba\x7f\x09\xcf\x7f\x2d\xab\xf7\xc1\x9a\xc7\x65\x8e\x6b\xe4\x28\x83\x61\x4f\xde\x73\x3d\x4e\xb0\xc7\x79\xa0\xa5\x79\x04\x18\x82\x55\x3c\x9d\x41\x67\xbe\x55\x31\x4b\xdb\x2f\xdd\x39\xdb\xe7\xd7\x51\xce\x79\x54\xfe\xbc\xff\xa9\xae\xca\x47\xd2\x20\xe5\xfe\x29\x3a\xdd\xe1\xd3\x72\xd4\xfb\x04\x22\x0a\xb3\x74\x48\x8f\x4e\x04\x99\x7c\x97\xde\xd3\x18\x79\x4f\xa7\x81\xe6\xad\x8d\x79\x94\xe4\xc1\x15\xad\x40\x39\x6f\x95\xbc\x04\xce\xdb\xdc\x15\xa2\xe0\x14\xff\xbf\x2f\xd8\xff\x6f\x09\x3f\x56\x55\x95\x25\xed\x6d\x7d\x99\xad\x61\x04\x9b\x6c\xcf\x2f\x5e\x4a\xb3\xbc\x7c\x17\x71\xcf\xa2\xa6\x67\x57\x3c\x4c\x7f\xa7\x5a\xf9\x30\x81\x19\xcb\x1c\x79\xd1\x1f\xce\x38\x80\x10\xc8\xa1\xff\x3a\x38\xe9\x66\x69\x88\x5e\x40\x8a\x99\x90\xbd\x5a\x42\x89\x5a\x97\xf0\x19\xca\x11\x6e\xf4\x2f\x55\xde\x26\xd2\x3d\x9a\x01\xb5\x3e\x9e\xb3\x5b\x27\xd1\x34\x87\x28\xca\xfe\xf1\xfd\xaa\x62\x76\x81\xb8\xca\x22\x65\x37\xdf\x15\x7f\x42\xfc\xd2\xa3\x87\x8e\x13\xd4\x81\x45\x1f\x56\x1f\x20\xf5\x01\x0e\x3a\x42\x86\x6d\x60\x30\x67\xd6\x5f\x65\xf6\xea\xaf\x09\xed\xd5\x12\xfe\xf6\xbc\x1c\x01\xc7\xd8\x4e\xba\x5e\x20\x4f\x5f\x65\x52\x9b\x07\x4e\x55\x69\x5e\x9c\xee\x30\xa9\xcd\x73\x32\x4c\x2c\x18\x1f\xcb\xc9\xcb\x8a\xb1\xab\xce\x9d\x72\xf5\xa0\x55\x8a\x2f\xef\x8a\xdf\x02\x00\x00\xff\xff\xc2\xc7\xf6\xab\xaa\x0a\x00\x00"),
		},
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
