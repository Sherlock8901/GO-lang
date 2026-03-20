package main
import "fmt"

func main(){
	// Range for slice

	nums:= []int {6,7,8}
	sum:= 0

	for i, num := range nums {
		fmt.Println(num,i)
		sum = sum+num
	}
	fmt.Println(sum)

	// Range for map

	m := map[string]string {"fname": "Dip", "lname": "Patra"}
    for k, v := range m{
		fmt.Println(k,v)
	}

	// You can use it for string also

    name:= "Dip"
	for i, char := range name{
		fmt.Println(i,char) //it print unicode code, we learn it after few classes. Check it in the unicode table
		fmt.Println(i,string(char)) // string() coverted unicode into its string vale
	}
}