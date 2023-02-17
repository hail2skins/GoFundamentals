package main

import (
	"GoFundamentals/menu"
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "GoFundamentals/menu"
)

var in = bufio.NewReader(os.Stdin)

func main() {
loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')
		//choice = strings.TrimSpace(choice)

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			err := menu.Add()
			if err != nil {
				fmt.Println(fmt.Errorf("invalid input: %w", err))
			}
		case "q":
			break loop
		default:
			fmt.Println("Unkown option")
		}
	}
}
