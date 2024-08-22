package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zfirdavs/test_repo/greeting"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("must provide an argument")
	}

	name := os.Args[1]
	fmt.Println(greeting.Greet(name))
}
