package blobaccess

import (
	"bytes"
	"io"
	"sync"

	"github.com/mandelsoft/goutils/errors"
	"github.com/mandelsoft/vfs/pkg/vfs"

	"github.com/open-component-model/ocm/pkg/blobaccess/bpi"
	"github.com/open-component-model/ocm/pkg/iotools"
	"github.com/open-component-model/ocm/pkg/utils"
)

func Cast[I interface{}](acc BlobAccess) I {
	return bpi.Cast[I](acc)
}

////////////////////////////////////////////////////////////////////////////////

// TempFile holds a temporary file that should be kept open.
// Close should never be called directly.
// It can be passed to another responsibility realm by calling Release
// For example to be transformed into a TemporaryBlobAccess.
// Close will close and remove an unreleased file and does
// nothing if it has been released.
// If it has been releases the new realm is responsible.
// to close and remove it.
type TempFile struct {
	lock       sync.Mutex
	temp       vfs.File
	filesystem vfs.FileSystem
}

func NewTempFile(dir string, pattern string, fss ...vfs.FileSystem) (*TempFile, error) {
	fs := utils.FileSystem(fss...)
	temp, err := vfs.TempFile(fs, dir, pattern)
	if err != nil {
		return nil, err
	}
	return &TempFile{
		temp:       temp,
		filesystem: fs,
	}, nil
}

func (t *TempFile) Name() string {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.temp.Name()
}

func (t *TempFile) FileSystem() vfs.FileSystem {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.filesystem
}

func (t *TempFile) Release() vfs.File {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.temp != nil {
		t.temp.Sync()
	}
	tmp := t.temp
	t.temp = nil
	return tmp
}

func (t *TempFile) Writer() io.Writer {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.temp
}

func (t *TempFile) Sync() error {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.temp.Sync()
}

func (t *TempFile) AsBlob(mime string) BlobAccess {
	return ForTemporaryFile(mime, t.Release(), t.filesystem)
}

func (t *TempFile) Close() error {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.temp != nil {
		name := t.temp.Name()
		t.temp.Close()
		t.temp = nil
		return t.filesystem.Remove(name)
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// BlobData can be applied directly to a function result
// providing a BlobAccess to get the data for the provided blob.
// If the blob access providing function provides an error
// result it is passed to the caller.
func BlobData(blob DataGetter, err ...error) ([]byte, error) {
	if len(err) > 0 && err[0] != nil {
		return nil, err[0]
	}
	return blob.Get()
}

// BlobReader can be applied directly to a function result
// providing a BlobAccess to get a reader for the provided blob.
// If the blob access providing function provides an error
// result it is passed to the caller.
func BlobReader(blob DataReader, err ...error) (io.ReadCloser, error) {
	if len(err) > 0 && err[0] != nil {
		return nil, err[0]
	}
	return blob.Reader()
}

// DataFromProvider extracts the data for a given BlobAccess provider.
func DataFromProvider(s BlobAccessProvider) ([]byte, error) {
	blob, err := s.BlobAccess()
	if err != nil {
		return nil, err
	}
	defer blob.Close()
	return blob.Get()
}

// ReaderFromProvider gets a reader for a BlobAccess provided by
// a BlobAccesssProvider. Closing the Reader also closes the BlobAccess.
func ReaderFromProvider(s BlobAccessProvider) (io.ReadCloser, error) {
	blob, err := s.BlobAccess()
	if err != nil {
		return nil, err
	}
	r, err := blob.Reader()
	if err != nil {
		blob.Close()
		return nil, err
	}
	return iotools.AddReaderCloser(r, blob), nil
}

// MimeReaderFromProvider gets a reader for a BlobAccess provided by
// a BlobAccesssProvider. Closing the Reader also closes the BlobAccess.
// Additionally, the mime type of the blob is returned.
func MimeReaderFromProvider(s BlobAccessProvider) (io.ReadCloser, string, error) {
	blob, err := s.BlobAccess()
	if err != nil {
		return nil, "", err
	}
	mime := blob.MimeType()
	r, err := blob.Reader()
	if err != nil {
		blob.Close()
		return nil, "", err
	}
	return iotools.AddReaderCloser(r, blob), mime, nil
}

type GenericData interface{}

type GenericDataGetter interface {
	Get() (interface{}, error)
}

const KIND_DATASOURCE = "data source"

// GetData provides data as byte sequence from some generic
// data sources like byte arrays, strings, DataReader and
// DataGetters. This means we can pass all BlobAccess or DataAccess
// objects.
// If no an unknown data source is passes an ErrInvalid(KIND_DATASOURCE)
// is returned.
func GetData(src GenericData) ([]byte, error) {
	switch t := src.(type) {
	case []byte:
		return t, nil
	case string:
		return []byte(t), nil
	case DataGetter:
		return t.Get()
	case DataReader:
		var buf bytes.Buffer
		r, err := t.Reader()
		if err != nil {
			return nil, err
		}
		defer r.Close()
		_, err = io.Copy(&buf, r)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
	return nil, errors.ErrInvalidType(KIND_DATASOURCE, src)
}

// GetGenericData evaluates some input provided by well-known
// types or interfaces and provides some data output
// by mapping the input to either a byte sequence or
// some specialized object.
// If the input type is not known an ErrInvalid(KIND_DATASOURCE)
// // is returned.
// In extension to GetData, it additionally evaluates the interface
// GenericDataGetter to map the input to some evaluated object.
func GetGenericData(src GenericData) (interface{}, error) {
	switch t := src.(type) {
	case GenericDataGetter:
		return t.Get()
	default:
		return GetData(src)
	}
}
