package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var x, y uint8

	prints("We have two variables, which have not yet been assigned.\nWould you like to pick a number? Y/N")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter text: ")
	yn, _ := reader.ReadString('\n')

	if yn == "Y" {
		prints("\nOkay, choose the first variable (x): ")
		_, err := fmt.Scanf("%d", &x)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			prints("\nOkay. Choose the next one: ")
			_, err := fmt.Scanf("%d", &y)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

}

func prints(x string) {
	fmt.Print(x)
}

func printb(x bool) {
	fmt.Println(x)
}

func printi(x uint8) {
	fmt.Print(x)
}

func twice(x, y uint8) (uint8, uint8) {
	prints("This function should return 2 uinsigned ints.\n")

	if x > 1 {
		prints("X is bigger than 1\n")
	} else {
		prints("X is smaller than 1.\n")
	}

	if (y * x) > 10 {
		prints("The value of X times Y is: ")
		printi(x * y)
		prints(". This is important because if one number is too big, the uint8 won't be able to handle it")
	}

	return x + 1, y + 1
}

func runRobberyAttempt() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	openedVault := rand.Intn(100)
	leftSafely := rand.Intn(5)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time? You didn't even make it close to the vault. Idiot.")
	}

	if isHeistOn && (openedVault >= 70) {
		fmt.Println("\n\nGrab and go!")
	} else if isHeistOn && !(openedVault >= 70) {
		isHeistOn = false
		fmt.Println("\n\nYou failed miserably to open the fault. Moron.")
	}

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("\n\nFAILED to leave safely!")
		case 1:
			isHeistOn = false
			fmt.Print("\n\nTurns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("\n\nYou buffoon, failure.")
		case 3:
			fmt.Println("\n\nGet to the car! We made it.")
		default:
			fmt.Println("\n\nStart the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)

		fmt.Println("\n\nWe stole a whopping $", amtStolen)
	}

	fmt.Println("Rolls:\nEluded Guards: ", eludedGuards, "%.\nOpened Vault: ", openedVault, "%.\nLeft Safely: ", leftSafely, "%.")
}
