package main

import (
	"log"
	"wordChange/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v\n",err)
	}
}
