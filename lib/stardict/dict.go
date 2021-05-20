package stardict

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	embed2 "wdb/lib/embed"
)

// Dict implements in-memory dictionary
type Dict struct {
	buffer []byte
}

// GetSequence returns data at the given offset
func (d Dict) GetSequence(offset uint64, size uint64) []byte {
	return d.buffer[offset:(offset + size)]
}

// ReadDict reads dictionary into memory
func ReadDict(filename string, info *Info, embed bool) (dict *Dict, err error) {

	var r io.Reader

	if !embed {
		reader, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
		}
		defer reader.Close()
		if strings.HasSuffix(filename, ".dz") { // if file is compressed then read it from archive
			r, err = gzip.NewReader(reader)
		} else {
			r = reader
		}
	} else {
		reader, err := embed2.Assets().Open(filename)
		if err != nil {
			fmt.Println(err)
		}
		defer reader.Close()
		if strings.HasSuffix(filename, ".dz") { // if file is compressed then read it from archive
			r, err = gzip.NewReader(reader)
		} else {
			r = reader
		}
	}

	buffer, err := ioutil.ReadAll(r)

	if err != nil {
		return
	}

	dict = new(Dict)
	dict.buffer = buffer

	return
}
