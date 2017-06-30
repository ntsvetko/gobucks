package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/ntsvetko/gobucks/models"
)

func TestTrueOrFalse(test *testing.T) {

	var trueFalseCount [2]int
	for i := 0; i < 100; i++ {
		if trueOrFalse() == true {
			trueFalseCount[1]++
		} else {
			trueFalseCount[0]++
		}
	}

	fmt.Println("Number of false results: " + strconv.Itoa(trueFalseCount[0]))
	fmt.Println("Number of true results: " + strconv.Itoa(trueFalseCount[1]))
}

func TestGamble(test *testing.T) {
	dbConf, session := models.InitResetDB()
	names, _ := models.SeedUsers(dbConf, "gambletestcollection", session)

	_, currBalance, _ := Gamble(names[0], 1000, session, "test", "gambletestcollection")

	if currBalance != 100 {
		test.Fatalf("@Gamble")
	}

	_, currBalance, _ = Gamble(names[0], -1, session, "test", "gambletestcollection")

	if currBalance != 100 {
		test.Fatalf("@Gamble")
	}

	outcome, currBalance, _ := Gamble(names[0], 85, session, "test", "gambletestcollection")

	if currBalance != 15 && currBalance != 185 {
		test.Fatalf("@Gamble")
	}

	if currBalance == 15 {
		if outcome != false {
			test.Fatalf("@Gamble")
		}
	}

	if currBalance == 185 {
		if outcome != true {
			test.Fatalf("@Gamble")
		}
	}

	Gamble(names[0], 10, session, "test", "gambletestcollection")
}

func TestBalance(test *testing.T) {
	dbConf, session := models.InitResetDB()
	names, userColl := models.SeedUsers(dbConf, "balancetestcollection", session)

	models.AddTransaction(names[0], 10, false, userColl) // first user should now have 90

	currBalance := Balance(names[0], session, "test", "balancetestcollection")

	if currBalance != 90 {
		test.Fatalf("currBalance should be 90 but was %v", currBalance)
	}
}
