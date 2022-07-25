package main

import (
	f "fmt"

	"github.com/jinnyl22/go-prac/banking"
	"github.com/jinnyl22/go-prac/something"
)

func main() {
	something.Say()
	account := banking.Account{Owner: "hi", Balance: 20000}
	var a int = 1
	b := &a
	*b = 2
	f.Println(account)
	f.Println(&a, *b, b, a)
}
