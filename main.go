package main

import "fmt"

func main() {
	result, err := Run("slava.com", "com.whois-servers.net")
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(result)
	}
}
