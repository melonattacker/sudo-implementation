package main

import (
	"fmt"
	"syscall"
)

func main() {
	ruid := syscall.Getuid()
	euid := syscall.Geteuid()
	fmt.Println(ruid, euid)
}