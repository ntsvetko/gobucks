package main

import (
	"bufio"
	"fmt"
	"github.com/ntsvetko/gobucks/models"
	"os"
	"strconv"
	"strings"
)

var coll *mgo.Collection

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: './gobucks <user>'")
		return
	}
	username := os.Args[1]
	session := models.ConnectToMongo("mongodb://localhost")
	coll = models.GetColl(session, "gobucks", "users")
	defer session.Close()
	user := models.FindOrCreateUser(username, coll)
	fmt.Println(user)
	repl(user.Name)
}

func repl(user string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(user + "> ")
	for scanner.Scan() {
		input := scanner.Text()
		parse(input + " " + user)
		fmt.Print(user + "> ")
	}
}

func parse(input string) {
	arr := strings.Split(input, " ")
	switch arr[0] {
	case "gamble":
		if len(arr) < 3 {
			errorMessage()
			return
		}
		numGamble, err := strconv.Atoi(arr[1])
		if err != nil {
			errorMessage()
			return
		}
		user := arr[2]
		err = Gamble(user, numGamble, coll)
		if err != nil {
			fmt.Println("something went wrong when gambling")
			return
		}
	default:
		errorMessage()
	}
}

func errorMessage() {
	fmt.Println("please enter 'gamble <number>'")
}
