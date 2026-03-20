package main
import "fmt"
import "time"

func main(){
	// Normal Switch

	// i:= 5
	// switch i{
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Other")
	// }

	// Multicondition Switch exclusive for GO

	switch time.Now(). Weekday(){ //time.Now() → gets current date & time
		//.Weekday() → extracts the day (like Monday)
	case time.Saturday, time.Sunday: 
	fmt.Println("it's Weekend")

	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Working day")
	}
}