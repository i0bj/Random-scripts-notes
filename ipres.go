package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	file, _ := os.Open(`txt file with IPs`) 
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		addr, _ := net.LookupAddr(line)
		fmt.Println(addr)
	}
}
