// Diet arranged by dietitian
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Your menu now is ... ")
	t:= time.Now() 
	switch {
	case t.Hour() < 11:
		fmt.Println("Breakfast:")
		fmt.Println("Meal A: A cup of unsweetened soy milk, 2 eggs & an apple")
		fmt.Println("Meal B: A cup of unsweetened soy milk, 2 eggs & wholegrain toast")

	case t.Hour() < 16:
		fmt.Println("Lunch:")
		fmt.Println("Brown rice/ sweet potato, good protein & vegetable")

	default:
		fmt.Println("Dinner:")
		fmt.Println("Chicken breast/ salmon, fried vegetable & sweet potato/ green bean")

	}
	fmt.Println("Enjoy your healthy meal! :)")
}
