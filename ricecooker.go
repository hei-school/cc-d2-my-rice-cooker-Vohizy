package main

import (
	"fmt"
	"strconv"
	"time"
)

// RiceCooker represents a rice cooker.
type RiceCooker struct {
	power bool
}

// checkPower checks the power status of the rice cooker.
func (rc *RiceCooker) checkPower() bool {
	return rc.power
}

// turnOn turns on the rice cooker.
func (rc *RiceCooker) turnOn() {
	rc.power = true
	fmt.Println("Hello!! You are on the rice cooker. The rice cooker is turned on.")
}

// turnOff turns off the rice cooker.
func (rc *RiceCooker) turnOff() {
	rc.power = false
	fmt.Println("The rice cooker is turned off.")
}

// chooseFunction handles user input to perform actions on the rice cooker.
func (rc *RiceCooker) chooseFunction(userInput string) {
	switch userInput {
	case "1":
		rc.turnOn()
		action, err := rc.getValidNumber("What do you want to do? (1 to cook / 2 to boil): ", []int{1, 2})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		switch action {
		case 1:
			duration, err := rc.getValidNumber("How many minutes do you want to cook? ", nil)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			temperature, err := rc.getValidNumber("At what temperature do you want to cook? ", nil)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			rc.cook(duration, temperature)
		case 2:
			rc.boil()
		default:
			fmt.Println("Unrecognized action.")
		}
	case "2":
		rc.turnOff()
	default:
		fmt.Println("Unrecognized command.")
	}
}

// getValidNumber prompts the user for a valid number and validates the input.
func (rc *RiceCooker) getValidNumber(promptMessage string, validChoices []int) (int, error) {
	var userInput string
	for {
		fmt.Print(promptMessage)
		fmt.Scanln(&userInput)
		number, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Error: Please enter a valid number.")
			continue
		}
		if len(validChoices) > 0 {
			valid := false
			for _, choice := range validChoices {
				if number == choice {
					valid = true
					break
				}
			}
			if !valid {
				fmt.Println("Error: Please enter a valid number among the specified choices.")
				continue
			}
		}
		return number, nil
	}
}

// cook simulates cooking for a specified duration and temperature.
func (rc *RiceCooker) cook(duration, temperature int) {
	fmt.Printf("Cooking in progress for %d minute(s) at %d°C...\n", duration, temperature)
	rc.sleep(duration)
	fmt.Println("Cooking completed!")
}

// boil simulates boiling water.
func (rc *RiceCooker) boil() {
	fmt.Println("Boiling water in progress...")
	boilingTime := 10 // Simulation duration (10 minutes for example)
	currentTemperature := 0

	for i := 1; i <= boilingTime; i++ {
		currentTemperature += 10
		rc.sleep(1)
		fmt.Printf("Water temperature: %d°C - Remaining time: %d minutes\n", currentTemperature, boilingTime-i)
	}

	fmt.Println("Water has reached 100°C. Automatic shutdown.")
}

// sleep pauses execution for the specified number of minutes.
func (rc *RiceCooker) sleep(minutes int) {
	milliseconds := int64(minutes) * 60 * 1000
	startTime := currentTimeMillis()
	for currentTimeMillis() < startTime+milliseconds {
	}
}

// currentTimeMillis returns the current time in milliseconds.
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func main() {
	myRiceCooker := &RiceCooker{}

	fmt.Print("Hello! You are on the rice cooker. Do you want to turn it on? (1 for yes / 2 for no): ")
	userDecision, err := myRiceCooker.getValidNumber("", []int{1, 2})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	myRiceCooker.chooseFunction(strconv.Itoa(userDecision))
}
