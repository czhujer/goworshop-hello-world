package main

import (
	"fmt"
	"rsc.io/quote"
	"time"
)

//go mod https://github.com/golang/go/wiki/Modules#example

func printCustomDate() string {
	dt := time.Now()
	dtf := dt.Format("01-02-2006 15:04:05")
	//dtf := dt.String()

	fmt.Println("Dnes je: ", dtf)

	return dtf
}

func on_finish() {
	fmt.Println("Finished")
}

func main() {
	defer on_finish()

	fmt.Println(quote.Hello())
	fmt.Println("ahoj svete")

	printCustomDate()

	//https://github.com/davecgh/go-spew
}
