package main
import (
	"fmt"
	"time"
	"math/rand"
)
func main(){
	rand.Seed(time.Now().UnixNano())
	secretNumber:= rand.Intn(100) + 1
	 var guess int

	 fmt.Println("Get a number between 1 - 100")
	 for{
		fmt.Println("Enter a guess: ")
		fmt.Scanln(&guess)
		 if guess < secretNumber {
			fmt.Println("The number is too low")
		 }else if guess > secretNumber{
			fmt.Println("The number is too high")
		 }else{
			fmt.Println("The number is correct")
			break
		 }
	 }
	 
}