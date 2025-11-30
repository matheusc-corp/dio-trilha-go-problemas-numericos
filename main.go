package main

import "fmt"

func dividirPor3() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(i, " ")
		}
	}
}

func pinPan() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("Pinpan ")
		} else if i%3 == 0 {
			fmt.Print("Pin ")
		} else if i%5 == 0 {
			fmt.Print("Pan ")
		} else {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	fmt.Println("Numeros diviseis por 3")
	dividirPor3()

	fmt.Println("\nPin pan")
	pinPan()
}
