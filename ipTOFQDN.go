package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	file, _ := os.Open(`file.txt`) // txt file with IP addresses
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		addr, _ := net.LookupAddr(line)
		fmt.Println(addr)
	}

}
