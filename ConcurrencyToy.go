package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var userMap map[string]int
var userDone map[string](chan bool)

//var done chan bool
var wg sync.WaitGroup

func main() {
	userMap = make(map[string]int)
	userDone = make(map[string](chan bool))
	for i := 0; i < 3; i = i + 1 {
		wg.Add(4)
		go parse("gamble 10 natalie")
		go parse("gamble 20 vlad")
		go parse("gamble 15 natalie")
		go parse("gamble 50 justin")
	}
	wg.Wait()
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
		_, val := userMap[user]
		if !val {
			addUser(user)
		}
		gamble(user, numGamble)
	default:
		errorMessage()
	}
}

func errorMessage() {
	fmt.Println("error")
}

func gamble(user string, amount int) {
	select {
	case <-userDone[user]:
		oldAmount := userMap[user]
		fmt.Println(user + " gambling " + strconv.Itoa(amount) + ".")
		time.Sleep(time.Second * 5)
		newAmount := oldAmount + amount // in this world, you always win
		userMap[user] = newAmount
		fmt.Println(user + " wins and now has " + strconv.Itoa(newAmount) + ".")
		userDone[user] <- true
	}
	defer wg.Done()
}

func addUser(user string) {
	userMap[user] = 100
	userDone[user] = make(chan bool, 10)
	userDone[user] <- true
}
