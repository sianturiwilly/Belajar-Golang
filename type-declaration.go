package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpWilli NoKTP = "3275091901930013"
	var marriedStatus Married = false
	fmt.Println(noKtpWilli)
	fmt.Println(marriedStatus)
}
