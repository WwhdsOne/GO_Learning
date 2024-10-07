package main

import (
	"fmt"
	"github.com/wwh/bankcore"
	"log"
	"net/http"
	"strconv"
)

var accounts = map[float64]*bankcore.Account{}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}
	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}
func main() {
	accounts[1001] = &bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}
	accounts[1002] = &bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "Mary",
			Address: "New York, New York",
			Phone:   "(212) 555 0174",
		},
		Number: 1002,
	}
	http.HandleFunc("/statement", statement)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
