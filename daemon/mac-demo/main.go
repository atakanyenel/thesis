package main

import "os/exec"

import "os"

import "time"

func main() {
	input := os.Args[1]

	for {
		_, err := exec.Command("say", input).Output()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 2)
	}
}
