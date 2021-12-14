package connect

import (
	"io"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	URL := "wss://qiangzhuangniuniu-dev.zzishare.com:5008"
	ws := NewConn(URL)
	ws.Connect()
	defer ws.Close()

	for {
		bytes, err := ws.Read()
		if err != nil {
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			break
		}
		log.Println("recv: ", bytes)
	} //
}
