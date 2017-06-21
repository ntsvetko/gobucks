package main

import (
	"testing"
)

func TestConnectToMongo(test *testing.T) {
	connectToMongo("localhost:27017")
}
