package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	file, err := os.Open(`C:\FakePath\file.txt`) 
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cmd, err := exec.Command("cmd.exe", "/C", "ping", "-n", "1", line).Output()
		if err != nil {
			log.Println(err)
		}
			if err := ioutil.WriteFile("IPs", cmd, 0644); err != nil {
			log.Println(err)
		}

	}

}
