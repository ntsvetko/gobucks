package main

import (
	"bufio"
	"fmt"
	"strconv"
	//"gopkg.in/mgo.v2"
	"os"
	"strings"
)

func main() {
	repl("natalie")
}

func repl(user string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(user + "> ")
	for scanner.Scan() {
		input := scanner.Text()
		parse(input)
		fmt.Print(user + "> ")
	}
}

func parse(input string) {
	arr := strings.Split(input, " ")
	switch arr[0] {
	case "gamble":
		if len(arr) < 2 {
			errorMessage()
		}
		numGamble, err := strconv.Atoi(arr[1])
		if err != nil {
			errorMessage()
		}
		fmt.Println(numGamble)
	default:
		errorMessage()
	}
}

func errorMessage() {
	fmt.Println("please enter 'gamble <number>'")
}
