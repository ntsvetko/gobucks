package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ntsvetko/gobucks/models"
	//"github.com/pkg/profile"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	//"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
	"time"
)

// map of boolean channels, saying if the user is done
var userDone map[string](chan bool) = make(map[string](chan bool))

// fun concurrency things
var wg sync.WaitGroup
var rwm sync.RWMutex

// fun mongo things
var session *mgo.Session

var username string // to make things a bit easier in repl mode

const databaseString = "gobucks"
const collectionString = "users"

var cmode bool // whether or not you're in file/concurrency mode

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: './gobucks -c <filepath>' or './gobucks <user>'")
		return
	}

	cpuprofile := "gobucks.prof"
	f, err := os.Create(cpuprofile)
	if err != nil {
		log.Fatal(err)
	}
	//runtime.SetCPUProfileRate(100000000)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	go func() { // run an http server, needed for profiling
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	concurrent := flag.Bool("c", false, "take in a file with gamble commands")
	flag.Parse()

	cmode = *concurrent
	session = models.ConnectToMongo("mongodb://localhost")
	defer session.Close()

	if cmode { // take in a file and read through it for commands
		filepath := os.Args[2]
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println("ERROR: problem opening the file at " + filepath + ".")
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			wg.Add(1)
			text := scanner.Text()
			fmt.Println("read command: " + text)
			go parse(text)
		}
		wg.Wait()
	} else { // take in a user and give them a repl to gamble in
		repl(os.Args[1])
	}
}

/* repl for the single-user mode), takes in username*/
func repl(user string) {
	_, val := userDone[user]
	if !val {
		userDone[user] = make(chan bool, 10)
		userDone[user] <- true
	}
	username = user
	scanner := bufio.NewScanner(os.Stdin)
	printBalance(user)
	fmt.Print(user + "> ")
	for scanner.Scan() {
		input := scanner.Text()
		wg.Add(1)
		go parse(input + " " + user)
		wg.Wait()
		fmt.Print(user + "> ")
	}
}

/*
* parses something of the format "gamble <user> <number>"
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
		rwm.Lock()
		_, val := userDone[user]
		if !val {
			userDone[user] = make(chan bool, 10)
			userDone[user] <- true
		}
		rwm.Unlock()
		concurrentGamble(user, numGamble)
	case "balance":
		printBalance(username)
		defer wg.Done()
	default:
		errorMessage()
	}
}

/*
* gamble method for concurrency
* takes in a user and an amount to gamble
 */
func concurrentGamble(user string, amount int) {
	rwm.Lock()
	channel := userDone[user]
	rwm.Unlock()
	select {
	case <-channel:
		defer func() {
			channel <- true
		}()
		defer wg.Done()
		time.Sleep(time.Second * 2) // wait a second because you can't have instant gratification...
		win, newAmt, err := Gamble(user, amount, session, databaseString, collectionString)
		if err != nil {
			fmt.Print("ERROR: ")
			log.Println(err)
			return
		}
		if win {
			fmt.Println(user + " has won " + strconv.Itoa(amount) + " and now has " + strconv.Itoa(newAmt) + "! :D")
		} else {
			fmt.Println(user + " has lost " + strconv.Itoa(amount) + " and now has " + strconv.Itoa(newAmt) + "! :(")
		}
	}
}

/* tells users when they input nonsense */
func errorMessage() {
	if cmode {
		fmt.Println("ERROR: please input something of form 'gamble <num> <user>' in the file you pass in.")
	} else {
		fmt.Println("ERROR: please input something of form 'gamble <num>'")
	}
	wg.Done()
}

func printBalance(user string) {
	balance := Balance(user, session, databaseString, collectionString)
	fmt.Println(user + "'s balance: " + strconv.Itoa(balance))
}
