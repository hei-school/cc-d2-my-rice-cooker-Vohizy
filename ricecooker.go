package main

import (
	"fmt"
	"strconv"
)

type RiceCooker struct {
	power bool
}

func (rc *RiceCooker) checkPower() bool {
	return rc.power
}

func (rc *RiceCooker) turnOn() {
	rc.power = true
	fmt.Println("Salut !! Vous êtes sur le rice cooker. Le rice cooker est allumé.")
}

func (rc *RiceCooker) turnOff() {
	rc.power = false
	fmt.Println("Le rice cooker est éteint.")
}

func (rc *RiceCooker) chooseFunction(userInput string) {
	switch userInput {
	case "1":
		rc.turnOn()
		action, err := rc.getValidNumber("Que voulez-vous faire ? (1 pour cuire / 2 pour bouillir): ", []int{1, 2})
		if err != nil {
			fmt.Println("Erreur:", err)
			return
		}
		switch action {
		case 1:
			duration, err := rc.getValidNumber("Combien de minutes souhaitez-vous cuire ? ", nil)
			if err != nil {
				fmt.Println("Erreur:", err)
				return
			}
			temperature, err := rc.getValidNumber("À quelle température souhaitez-vous cuire ? ", nil)
			if err != nil {
				fmt.Println("Erreur:", err)
				return
			}
			rc.cook(duration, temperature)
		case 2:
			rc.boil()
		default:
			fmt.Println("Action non reconnue.")
		}
	case "2":
		rc.turnOff()
	default:
		fmt.Println("Commande non reconnue.")
	}
}

func (rc *RiceCooker) getValidNumber(promptMessage string, validChoices []int) (int, error) {
	var userInput string
	for {
		fmt.Print(promptMessage)
		fmt.Scanln(&userInput)
		number, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Erreur: Veuillez saisir un nombre valide.")
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
				fmt.Println("Erreur: Veuillez saisir un nombre valide parmi les choix spécifiés.")
				continue
			}
		}
		return number, nil
	}
}

func (rc *RiceCooker) cook(duration, temperature int) {
	fmt.Printf("Cuisson en cours pendant %d minute(s) à %d°C...\n", duration, temperature)
	rc.sleep(duration)
	fmt.Println("Cuisson terminée!")
}

func (rc *RiceCooker) boil() {
	fmt.Println("Bouillir de l'eau en cours...")
	boilingTime := 10 // Durée de la simulation (10 minutes par exemple)
	currentTemperature := 0

	for i := 1; i <= boilingTime; i++ {
		currentTemperature += 10
		rc.sleep(1)
		fmt.Printf("Température de l'eau: %d°C - Durée restante: %d minutes\n", currentTemperature, boilingTime-i)
	}

	fmt.Println("L'eau a atteint 100°C. Arrêt automatique.")
}

func (rc *RiceCooker) sleep(minutes int) {
	milliseconds := int64(minutes) * 60 * 1000
	startTime := currentTimeMillis()
	for currentTimeMillis() < startTime+milliseconds {
	}
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func main() {
	myRiceCooker := &RiceCooker{}

	fmt.Print("Salut !! Vous êtes sur le rice cooker. Voulez-vous l'allumer ? (1 pour oui / 2 pour non): ")
	userDecision, err := myRiceCooker.getValidNumber("", []int{1, 2})
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	myRiceCooker.chooseFunction(strconv.Itoa(userDecision))
}
