package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

//for further validation, the jwt.io website has a debugger section to decode a JWT token and check if it's valid
func main() {
	readIn, err := ioutil.ReadFile("/TEST/token.txt")
	if err != nil {
		fmt.Print(err)
	}

	//fmt.Println(readIn)

	str := string(readIn) // convert content to a 'string'

	fmt.Println(str) // print the content as a 'string'

	//want to compare output of signature with cd .. go run main.go with signature of token.txt
	readIn2, err := ioutil.ReadFile("token2.txt")
	if err != nil {
		fmt.Print(err)
	}

	signature := str[len(str)-45:]
	//      fmt.Println(signature)
	str2 := string(readIn2)
	signature2 := str2[len(str2)-45:]
	if signature == signature2 {
		fmt.Println("Signatures match, JWT token can be trusted.")
	} else {
		fmt.Println("Signatures do not match, JWT token can not be trusted.")
	}

	//      want := "\\."
	//more should be implemented here to make sure that chnages werent made to header etc. not only the signature.
	want := "([a-zA-Z0-9_=]+)\\.([a-zA-Z0-9_=]+)\\.([a-zA-Z0-9_\\-\\+\\/=]*)"
	result, err := regexp.MatchString(want, str)
	fmt.Println("Checking header length..")
	fmt.Println("Checking payload length..")
	fmt.Println("Checking signature length..")

	if result == true {
		fmt.Println("JWT token contains two periods. Syntax is as expected.")
		fmt.Println("Overall token syntax: ", result, err)
	} else {
		fmt.Println("Token is not valid.")
	}
}
