package models

type Transaction struct {
	amountBefore  int
	amountGambled int
	result        bool
	amountAfter   int
}

// Doens't check if transaction is legal
func addTransaction(name string, prevAmt int, betAmt int, outcome bool) {

}
