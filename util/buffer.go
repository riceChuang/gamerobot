package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"runtime"
)

//IntToBytes int to byte
func IntToBytes(n int32, bigEndian bool) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})

	var order binary.ByteOrder
	order = binary.BigEndian
	if !bigEndian {
		order = binary.LittleEndian
	}

	binary.Write(bytesBuffer, order, tmp)
	return bytesBuffer.Bytes()
}

//BytesToInt bytes to int
func BytesToInt(data []byte, bigEndian bool) int32 {
	bytesBuffer := bytes.NewBuffer(data)

	var order binary.ByteOrder
	order = binary.BigEndian
	if !bigEndian {
		order = binary.LittleEndian
	}

	var x int32

	if bytesBuffer == nil {
		pc, file, line, _ := runtime.Caller(6)
		f := runtime.FuncForPC(pc)
		fmt.Printf("BytesToInt 錯誤 %v:%v FuncName: %v", file, line, f.Name())
	}
	if order == nil {
		pc, file, line, _ := runtime.Caller(6)
		f := runtime.FuncForPC(pc)
		fmt.Printf("order 錯誤 %v:%v FuncName: %v", file, line, f.Name())
	}

	binary.Read(bytesBuffer, order, &x)

	return x
}
