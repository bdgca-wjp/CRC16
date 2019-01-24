package main

import (
	"fmt"

	"github.com/bkzy-wangjp/CRC16"
)

func main() {
	//fmt.Printf("lanuch")

	m_data := "gpxwrsyssuqvrrDGvuqttuGsqq"

	checksum := crc16.DataAndCrcSum([]byte(m_data))
	strsum := crc16.StringAndCrcSum(m_data)

	fmt.Printf("check sum:%X \n", checksum)
	fmt.Printf("check sum:%s \n", strsum)

	str, ok := crc16.StringCheckCRC("mpxwrsyssuqvrrDGvuqttuGsqqBD3C")
	fmt.Printf("原始字符串:%s,校验结果:%t", str, ok)
}
