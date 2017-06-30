package main

import (
	"bufio"
	"fmt"
	"github.com/ntsvetko/gobucks/models"
	"gopkg.in/mgo.v2"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var userDone map[string](chan bool)

var wg sync.WaitGroup
var coll *mgo.Collection

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./GobucksConcurrent <file>")
		return
	}
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("problem opening the file")
		return
	}
	defer file.Close()

	session := models.ConnectToMongo("mongodb://localhost")
	coll = models.GetColl(session, "gobucks", "users")
	defer session.Close()
	scanner := bufio.NewScanner(file)
	userDone = make(map[string](chan bool))
	for scanner.Scan() {
		wg.Add(1)
		text := scanner.Text()
		fmt.Println("command: " + text)
		go parse(text)
	}
	wg.Wait()
}

/*
* parses something of the format "<user> gamble <number>"
 */
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
		_, val := userDone[user]
		if !val {
			userDone[user] = make(chan bool, 10)
			userDone[user] <- true
		}
		concurrentGamble(user, numGamble)
	default:
		errorMessage()
	}
}

/*
* gamble method for concurrency
* takes in a user and an amount to gamble
 */
func concurrentGamble(user string, amount int) {
	select {
	case <-userDone[user]:
		time.Sleep(time.Second * 5) // wait 5 seconds because you can't have instant gratification...
		win, newAmt, err := Gamble(user, amount, coll)
		if err != nil {
			fmt.Println("something went wrong when gambling")
			return
		}
		if win {
			fmt.Println(user + " has won " + strconv.Itoa(amount) + " and now has " + strconv.Itoa(newAmt) + "!")
		} else {
			fmt.Println(user + " has lost " + strconv.Itoa(amount) + " and now has " + strconv.Itoa(newAmt) + "! :(")
		}
		userDone[user] <- true
	}
	defer wg.Done()
}

func errorMessage() {
	fmt.Println("ERROR")
}
