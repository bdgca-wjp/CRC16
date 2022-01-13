package crc16

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestCrcSum(t *testing.T) {
	tests := []struct {
		hex_datas string
		rst       string
		hlbyte    bool
	}{
		{"5b226b657931222c22e4bda0e5a5bde4b896e7958c225d", //["key1","你好世界"]
			"96D2", false},
		{"5b226b657931222c22e4bda0e5a5bde4b896e7958c225d", //["key1","你好世界"]
			"D296", true},
		{"5b226b657931225d", //["key1"]
			"D39F", false},
		{"5b226b657931225d", //["key1"]
			"9FD3", true},
	}
	for _, tt := range tests {
		bt, err := hex.DecodeString(tt.hex_datas)
		if err != nil {
			t.Errorf("hex字符串转换为字节出错:%s", err.Error())
		} else {
			rst := CrcSum(bt, tt.hlbyte)
			if err != nil {
				t.Errorf("decode错误:%s", err.Error())
			} else {
				if fmt.Sprintf("%X", rst) != tt.rst {
					t.Errorf("错误,期望值是:%s,得到的值是:%X", tt.rst, rst)
				}
			}
		}
	}
}

func TestCrc32(t *testing.T) {
	m_data := "bdkal;dkgujadgkjkeugo"

	bytesum := DataAndCrcSum([]byte(m_data))
	strsum := StringAndCrcSum(m_data)

	t.Logf("check sum:%X \n", bytesum)
	t.Logf("check sum:%s \n", strsum)

	str, ok := StringCheckCRC("gpxwrsyssuqvrrDGvuqttuGsqqBD3C")
	t.Logf("原始字符串:%s,校验结果:%t", str, ok)

	bts, ok := BytesCheckCRC(bytesum)
	t.Logf("原始字符串:%s,校验结果:%t", bts, ok)
}
