package main

import (
	"fmt"
	"log"
)

func main() {
	var inputone, inputtwo float64
	var operator string

	for {
		fmt.Print("Input Operator (-, +, *, /): ")
		fmt.Scanln(&operator)

		if operator == "+" || operator == "-" || operator == "*" || operator == "/" {
			break
		} else {
			fmt.Println("Komputer tidak mengerti input user, tolong masukkan operator yang benar!")
		}
	}

	for {
		fmt.Print("Enter First Number: ")
		_, err := fmt.Scanln(&inputone)
		if err != nil {
			log.Println("tolong masukkan angka")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter Second Number: ")
		_, err := fmt.Scanln(&inputtwo)
		if err != nil {
			log.Println("tolong masukkan angka")
			continue
		}
		break
	}

	switch operator {

	case "+":
		fmt.Printf("Result :  %.2f\n", inputone+inputtwo)

	case "-":
		fmt.Printf("Result : %.2f\n", inputone-inputtwo)

	case "*":
		fmt.Printf("Result : %.2f\n", inputone*inputtwo)

	case "/":
		if inputtwo != 0 {
			fmt.Printf("Result : %.2f\n", inputone/inputtwo)
		} else {
			fmt.Println("Kalkulator tidak mengerti input user!")
		}
	}

}
