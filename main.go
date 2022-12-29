package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	sl := []string{}
	f1, err := os.Open("poem.txt")
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		stroka := scanner.Text()
		for i := range stroka {
			sl = append(sl, string(stroka[i]))
			fmt.Println(string(stroka[i]))
		}
	}
}
Footer
© 2022 GitHub, Inc.
Footer navigation
Terms
Privacy
Security
Status
