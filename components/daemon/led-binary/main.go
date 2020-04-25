package main

import (
	"log"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

const (
	YELLOW = 18
	GREEN  = 23
)

func main() {
	input := os.Args[1]
	log.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	defer rpio.Close()
	// reset board start
	yellow_pin := rpio.Pin(YELLOW)
	green_pin := rpio.Pin(GREEN)
	yellow_pin.Output()
	green_pin.Output()
	yellow_pin.Write(rpio.Low)
	green_pin.Write(rpio.Low)

	var active_pin rpio.Pin
	log.Println("activating pin:", input)
	if input == "YELLOW" {
		active_pin = rpio.Pin(YELLOW)
	} else if input == "GREEN" {
		active_pin = rpio.Pin(GREEN)
	}
	defer active_pin.Write(rpio.Low)
	for {
		active_pin.Toggle()
		time.Sleep(time.Second)
	}

}
