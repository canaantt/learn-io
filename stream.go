package main

import (
	"bytes"
	"fmt"
	"io"
)

var rawData = []byte("It is a sunny day. It is a sunny day. It is a sunny day. It is a sunny day. It is a sunny day.")
var expectedData = []byte("It is a Sunny day. It is a Sunny day. It is a Sunny day. It is a Sunny day. It is a Sunny day.")


func main() {
	var output bytes.Buffer
	var find = []byte("sunny")
	var repl = []byte("Sunny")

	fmt.Println("=======================================\nRunning Algorithm One")
	output.Reset()
	method1(rawData, find, repl, &output)
	matched := bytes.Compare(expectedData, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, rawData, expectedData, output.Bytes())

	output.Reset()
	method2(rawData, find, repl, &output)
	matched = bytes.Compare(expectedData, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, rawData, expectedData, output.Bytes())

	output.Reset()
	method3(rawData, find, repl, &output)
	matched = bytes.Compare(expectedData, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, rawData, expectedData, output.Bytes())

}

func method1(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	input := bytes.NewBuffer(data)

	size := len(find)
	end := size - 1
	buf := make([]byte, size)

	//fmt.Println("0 -- ", string(buf[:end]))

	if n, err := io.ReadFull(input, buf[:end]); err != nil {
		output.Write(buf[:n])
		return
	}
	//fmt.Println("1 -- ", string(buf))
	for{
		if _, err := io.ReadFull(input, buf[end:]); err != nil {
			output.Write(buf[:end])
			return
		}
		//fmt.Println("2 -- ", string(buf))
		if bytes.Compare(buf, find) == 0 {
			output.Write(repl)
			//fmt.Println("3 -- ", string(buf))
			if n, err := io.ReadFull(input, buf[:end]); err != nil {
				output.Write(buf[:n])
				//fmt.Println("6 --", string(buf))
				break
			}
			//fmt.Println("4 -- ", string(buf))
			continue
		}

		output.WriteByte(buf[0])
		//fmt.Println("5 -- ", string(buf))
		copy(buf, buf[1:])
	}
}

func method2(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	input := bytes.NewReader(data)
	// input := bytes.NewBuffer(data) seems to be the same as above
	size := len(find)
	index := 0

	for{
		b, err := input.ReadByte(); if err != nil {
			break;
		}
		if b == find[index] {
			index++
			if index == size {
				output.Write(repl)
				index = 0
			}
			continue
		}
		if index != 0 {
			output.Write(find[:index])
			//output.WriteByte(b) // faster?
			input.UnreadByte()
			index = 0
			continue
		}
		output.WriteByte(b)
		index = 0
	}

}

func method3(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	input := bytes.NewBuffer(data)
	// input := bytes.NewReader(data) seems to be the same as above
	size := len(find)
	index := 0

	for{
		b, err := input.ReadByte(); if err != nil {
			break;
		}
		if b == find[index] {
			index++
			if index == size {
				output.Write(repl)
				index = 0
			}
			continue
		}
		if index != 0 {
			output.Write(find[:index])
			output.WriteByte(b) // faster?
			//input.UnreadByte()
			index = 0
			continue
		}
		output.WriteByte(b)
		index = 0
	}

}