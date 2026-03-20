package main
import "fmt"
import "maps"

func main(){
	m := make(map[string]string)
	m["name"]= "Golang"
	m["area"]= "Backend"

	fmt.Println(m["name"], m["area"])
	fmt.Println(m["num"], m["area"]) // If something is doesn't match then it return zero value
	fmt.Println(m)

	// delete function

	delete(m, "area")
	fmt.Println(m) // print:  map[name: Golang]

	// Check if the element is in the map or not

	// new way to writing map
	ma:= map[string]int {"price":40, "phone":3}
	v, ok:= ma["phone"] //it return two values in "ok" it return present or not, in "v" it return the value of phone
	fmt.Println(v)
    if ok{
		fmt.Println("all ok")
	} else{
		fmt.Println("all ok")
	}

	// check two maps is equal or not
    m1 := map[string]int{"a": 3, "b": 2}
	m2 := map[string]int{"b": 2, "a": 1}
	fmt.Println(maps.Equal(m1, m2))

}