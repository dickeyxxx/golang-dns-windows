package main

import (
	"fmt"
	"net"

	"github.com/franela/goreq"
)

func main() {
	fmt.Println("starting")

	fmt.Println("google.com:80")
	_, err := net.Dial("tcp", "google.com:80")
	printerr(err)

	fmt.Println("github.com:443")
	_, err = net.Dial("tcp", "github.com:443")
	printerr(err)

	fmt.Println("http://github.com")
	_, err = goreq.Request{Uri: "http://www.github.com"}.Do()
	printerr(err)

	fmt.Println("https://github.com")
	_, err = goreq.Request{Uri: "https://www.github.com"}.Do()
	printerr(err)

	fmt.Println("http://google.com")
	_, err = goreq.Request{Uri: "http://www.google.com"}.Do()
	printerr(err)

	fmt.Println("https://google.com")
	_, err = goreq.Request{Uri: "https://www.google.com"}.Do()
	printerr(err)

	fmt.Println("done")
}

func printerr(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
