package main

import "fmt"

func main() {
	result, err := Run("slava.com")
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(result.Registrant.Email)
	}
}
