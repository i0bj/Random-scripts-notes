package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Usage: Press Enter for new line or task you did.")
	fmt.Println("Usage: Press Ctrl + c when you are done documenting tasks.")
	var user string
	fmt.Println("Enter user name, Case Sensitive ex... 'useR' : ")
	fmt.Scanln(&user)
	dt := time.Now()
	f, _ := os.Create(fmt.Sprintf(`\\server\path\path` + user + `\` + user + dt.Format("-Monday-") + dt.Format("01022006") + ".txt"))
	defer f.Close()

	fmt.Println("Enter what you did today >> ")
	for {
		usr := bufio.NewReader(os.Stdin)
		usr.WriteTo(f)
	}
}
