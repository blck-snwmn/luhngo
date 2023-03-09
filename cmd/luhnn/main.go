package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/blck-snwmn/luhngo"
)

var (
	input  string
	verify bool
)

func main() {
	flag.StringVar(&input, "i", "", "input")
	flag.BoolVar(&verify, "v", false, "verify")
	flag.Parse()

	if input == "" {
		log.Fatalf("input is empty")
	}
	l, _ := luhngo.NewLuhnn("abcdef")
	if verify {
		err := l.Verify(input)
		if err != nil {
			fmt.Println("NG")
		} else {
			fmt.Println("OK")
		}
	} else {
		d, err := l.Generate(input)
		if err != nil {
			log.Fatalf("invalid input: %+v", err)
		}
		fmt.Println(d)
	}
}
