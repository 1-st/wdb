package stardict

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	embed2 "wdb/lib/embed"
)

// Info contains dictionary options
type Info struct {
	Version string
	Is64    bool
	Options map[string]string
}

func decodeOption(str string) (key string, value string, err error) {
	a := strings.Split(str, "=")

	if len(a) < 2 {
		return "", "", errors.New("Invalid file format: " + str)
	}

	return a[0], a[1], nil
}

// ReadInfo reads ifo file and collects dictionary options
func ReadInfo(filename string, embed bool) (info *Info, err error) {
	var r *bufio.Reader

	if !embed {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
		}
		r = bufio.NewReader(file)
		defer file.Close()
	} else {
		fs, err := embed2.Assets().Open(filename)
		if err != nil {
			fmt.Println(err)
		}
		r = bufio.NewReader(fs)
	}

	_, err = r.ReadString('\n')

	if err != nil {
		return
	}

	version, err := r.ReadString('\n')

	if err != nil {
		return
	}

	kn, kv, err := decodeOption(version[:len(version)-1])

	if err != nil {
		return
	}

	if kn != "version" {
		err = errors.New("Version missing (should be on second line)")
		return
	}

	if kv != "2.4.2" && kv != "3.0.0" && kv != "2.4.2\r" && kv != "3.0.0\r" {
		err = errors.New("Stardict version should be either 2.4.2 or 3.0.0")
		return
	}

	info = new(Info)

	info.Version = kv

	info.Options = make(map[string]string)

	for {
		option, err := r.ReadString('\n')

		if err != nil && err != io.EOF {
			return info, err
		}

		if err == io.EOF && len(option) == 0 {
			break
		}

		kn, kv, err = decodeOption(option[:len(option)-1])

		if err != nil {
			return info, err
		}

		info.Options[kn] = kv

		if err == io.EOF {
			break
		}
	}

	if val, ok := info.Options["idxoffsetbits"]; ok {
		if val == "64" {
			info.Is64 = true
		}
	} else {
		info.Is64 = false
	}

	return
}
