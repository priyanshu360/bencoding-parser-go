package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

var unmarshalTestData, err = os.ReadFile("C:/Users/priya/Downloads/BB78F5D3E3061DF46BC61D95DB6E39EF02C93DBB.torrent")

func main() {
	r := bytes.NewReader(unmarshalTestData)
	read := bufio.NewReader(r)
	_, err = read.ReadByte()

	for err == nil {
		read.UnreadByte()
		fmt.Println(Decode(read))
		_, err = read.ReadByte()
	}
}

func Decode(r *bufio.Reader) (interface{}, error) {
	ch, err := r.ReadByte()
	if err != nil {
		return nil, err
	}

	switch ch {
	case 'i':
		return ParseInt(r)
	case 'l':
		return ParseList(r)
	case 'd':
		return ParseDict(r)
	default:
		return ParseString(r)
	}

}

func ParseInt(r *bufio.Reader) (interface{}, error) {
	data, err := r.ReadBytes('e')
	if err != nil {
		return nil, err
	}

	data = data[:len(data)-1]
	intData, err := strconv.ParseInt(string(data), 10, 64)
	return intData, nil
}

func ParseList(r *bufio.Reader) (interface{}, error) {
	list := []interface{}{}
	for {
		next, err := r.ReadByte()
		if err != nil {
			return nil, err
		}

		switch next {
		case 'e':
			return list, nil
		default:
			r.UnreadByte()
			data, err := Decode(r)
			if err != nil {
				return nil, err
			}

			list = append(list, data)
		}
	}
}

func ParseDict(r *bufio.Reader) (interface{}, error) {
	dict := map[string]interface{}{}
	for {
		next, err := r.ReadByte()
		if err != nil {
			return nil, err
		}

		switch next {
		case 'e':
			return dict, nil
		default:
			r.UnreadByte()
			key, err := Decode(r)
			if err != nil {
				return nil, err
			}

			value, err := Decode(r)
			if err != nil {
				return nil, err
			}

			dict[key.(string)] = value
		}
	}
}

func ParseString(r *bufio.Reader) (interface{}, error) {
	r.UnreadByte()
	strLen, err := r.ReadBytes(':')
	if err != nil {
		return nil, err
	}

	strLen = strLen[:len(strLen)-1]
	strLenInt, err := strconv.ParseInt(string(strLen), 10, 64)
	buf := make([]byte, strLenInt)
	_, err = readAtLeast(r, buf, int(strLenInt))
	if err != nil {
		return nil, err
	}

	return string(buf), nil
}

func readAtLeast(r *bufio.Reader, buf []byte, min int) (n int, err error) {
	if len(buf) < min {
		return 0, io.ErrShortBuffer
	}
	for n < min && err == nil {
		var nn int
		nn, err = r.Read(buf[n:])
		n += nn
	}

	if n >= min {
		err = nil
	} else if n > 0 && err == io.EOF {
		err = io.ErrUnexpectedEOF
	}

	return
}
