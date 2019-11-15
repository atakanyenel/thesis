package main

import (
	"context"
	"log"
	"os/exec"
	"time"
)


func main() {
	pods := make(map[string]func())
	log.Println("starting Program")
	ctx, cancel := context.WithCancel(context.Background())
	pods["mypod"] = cancel

	go func(command string) {
		cmd := exec.CommandContext(ctx, "sh", "-c", command)
		if err := cmd.Start(); err != nil { //start
			log.Println(err)
		}
		if err := cmd.Wait(); err != nil { //wait
			log.Println("waiting on cmd:", err)
		}
	}("ncat -l 8080")
	<-time.After(10 * time.Second)
	log.Println("closing with context")
	pods["mypod"]()
	delete(pods, "mypod")

}
