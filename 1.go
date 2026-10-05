package main
import "fmt"

func main() {
	var value int

	fmt.Print("Enter a value (1-5): ")
	fmt.Scan(&value)

	switch value {
	case 1:
		fmt.Println("One: This is the first value.")
	case 2:
		fmt.Println("Two: This is the second value")
	case 3:
		fmt.Println("Three: This is the third value.")
	case 4:
		fmt.Println("Four: This is the fourth value.")
	case 5:
		fmt.Println("Five: This is the fifth value.")
	default:
		fmt.Println("Invalid value. Please enter a number from 1 to 5.")
	}
}
