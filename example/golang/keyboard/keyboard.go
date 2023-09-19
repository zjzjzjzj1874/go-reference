package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("Listening for keyboard input. Press 'Q' to quit.")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC || char == 'q' || char == 'Q' {
			break
		}

		fmt.Printf("You pressed: %c\r\n", char)
	}
}
