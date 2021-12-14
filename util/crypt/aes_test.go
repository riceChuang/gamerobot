package crypt

import (
	"robot/internal/util"
	"testing"
)

func TestAES(t *testing.T) {
	src := "account=dsgdev_ssp_624762318&gameId=10&ip=127.0.0.1&password=123qwe&s=0"
	AESKey := "xyz123!QAZ2wsxGE"
	result := "Qylp4ZZuc1Olg1TOiduCZjf/LUHLcGd0Z85KptzSvy1SYPoL0BOEn0fHDcZQW5xm4O3AJXitGPlF2ioOD7lb4TnhAmDjGUNQ72eMPDfQXtg="
	encrypt := AesEncrypt(src, AESKey)

	util.AssertEqual(t, encrypt, result, "encrypt doesn't match the result")
}
