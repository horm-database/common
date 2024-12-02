package compress

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"sync"

	"github.com/horm-database/common/json"
)

var (
	gzipWriterPool = sync.Pool{}
	gzipReaderPool = sync.Pool{}
)

// JsonMarshalAndCompress first json encoding, then compression.
func JsonMarshalAndCompress(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}

	gzipWriter, ok := gzipWriterPool.Get().(*gzip.Writer)
	if !ok {
		gzipWriter = gzip.NewWriter(buf)
	} else {
		gzipWriter.Reset(buf)
	}

	defer gzipWriterPool.Put(gzipWriter)

	if err := json.Api.NewEncoder(gzipWriter).Encode(v); err != nil {
		_ = gzipWriter.Close()
		return nil, err
	}

	if err := gzipWriter.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// DecompressJsonUnmarshal decompression and json decode
func DecompressJsonUnmarshal(in []byte, v interface{}) error {
	gzipReader, ok := gzipReaderPool.Get().(*gzip.Reader)

	defer func() {
		if gzipReader != nil {
			gzipReaderPool.Put(gzipReader)
		}
	}()

	var err error
	if ok {
		err = gzipReader.Reset(bytes.NewReader(in))
	} else {
		gzipReader, err = gzip.NewReader(bytes.NewReader(in))
	}

	var reader io.Reader = gzipReader
	if err != nil {
		reader = bytes.NewReader(in)
	}

	return json.Api.NewDecoder(reader).Decode(v)
}

func Decompress(in []byte) ([]byte, error) {
	gzipReader, ok := gzipReaderPool.Get().(*gzip.Reader)

	defer func() {
		if gzipReader != nil {
			gzipReaderPool.Put(gzipReader)
		}
	}()

	var err error
	if ok {
		err = gzipReader.Reset(bytes.NewReader(in))
	} else {
		gzipReader, err = gzip.NewReader(bytes.NewReader(in))
	}

	var reader io.Reader = gzipReader
	if err != nil {
		reader = bytes.NewReader(in)
	}

	return ioutil.ReadAll(reader)
}
