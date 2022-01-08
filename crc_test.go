package crc16

import (
	"testing"
)

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
