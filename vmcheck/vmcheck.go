// Package main contains the code to read cpuid information and prints them out

package main

import (
	"encoding/binary"
	"fmt"
)

func cpuid_low(arg1, arg2 uint32) (eax, ebx, ecx, edx uint32) // implemented in cpuidlow_amd64.s
func main() {
	var HV uint32
	HV = uint32(1 << 31)
	_, _, c, _ := cpuid_low(0x1, 0)
	fmt.Printf("reg=0x%x\n", c)
	fmt.Printf("reg=%d\n", c)
	fmt.Println("(c & HV) == HV :", (c&HV) == HV)

	var leaf uint32
	leaf = 0x40000000
	fmt.Printf("leaf=0x%08x\n", leaf)
	_, b, c, d := cpuid_low(leaf, 0)
	buf := make([]byte, 12)
	binary.LittleEndian.PutUint32(buf, b)
	binary.LittleEndian.PutUint32(buf[4:], c)
	binary.LittleEndian.PutUint32(buf[8:], d)
	fmt.Println(string(buf))
}
