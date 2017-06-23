package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: './gobucks <user>'")
		return
	}
	username := os.Args[1]
	session := connectToMongo("mongodb://localhost")
	defer session.Close()
	coll := getColl(session, "gobucks", "users")
	user := findOrCreateUser(username, coll)
<<<<<<< HEAD
	fmt.Println(user)
	repl(user.Name)
=======
	repl(user.name)
>>>>>>> 809017ca7269e4728fb90369395f5a53f9f2921b
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
			return
		}
		numGamble, err := strconv.Atoi(arr[1])
		if err != nil {
			errorMessage()
			return
		}
		fmt.Println(numGamble)
	default:
		errorMessage()
	}
}

func errorMessage() {
	fmt.Println("please enter 'gamble <number>'")
}
