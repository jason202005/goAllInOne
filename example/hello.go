package main
// ref : https://www.youtube.com/watch?v=C8LgvuEBraI&ab_channel=JakeWright
import (
	"fmt"
	"errors"
	"math"
)

func main(){
	fmt.Println("Hello World~")
	// var x int
	// x = 5
	x := 5 // go can convert the type by using :=
	y := 4
	// var sum int = x + y
	sum := x + y

	if sum > 6 {
		fmt.Println("More than 6")
	}

	var a [5]int // declare an array
	a[2] = 7 // accessing the array

	// if you want to drop the var, we can do the same above
	b := [5]int {5,4,3,2,1}

	//slice can be declared without a fixed size of an array
	c := [] int {1,2,3,4,5} 
	c = append (c,13) // slice is useful when we need to append a new element to the array

	// map in GO
	// it is useful, they're like dictionaries in python or associative array in php 
	// map[THE_TYPE_OF_THE_KEY]VALUE_IN_THE_MAP

	// to create the map, we use the build in make function
	vertices := make(map[string]int)
	// it is similar to indexing array
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	// to remove something from the map, we can use the following
	delete(vertices, "square")
	fmt.Println(vertices)
	// to get particular value in a map, we can use the following code segments to get the value by key
	fmt.Println(vertices["triangle"])

	//for loop  is the only type of loop in GO
	for i := 0; i<5; i++{
		fmt.Println(i)
	}
	// while loop in GO
	i:=0

	for i<5 {
		fmt.Println(i)
		i++
	}

	// we can ilterate over the element in an array or a slice by using range 
	arr := []string{"a","b","c"}
	//range arr to get the length of arr
	for index, value := range arr {
		fmt.Println("index", index, "value:", value)
	}
	//we can do the similar range things for map
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	fmt.Println(c) 
	fmt.Println(b) 
	fmt.Println(a) // all the elements will be initized with 0 value
	fmt.Println(sum)

	result := sumfunction(2,3)
	fmt.Println(result)

	// it is useful since GO doesnt have excepotions if you are used to working in other languages like java.
	res, err := sqrt(-16) 
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// to initisalize a person
	p := person{name: "Jason", age:22}
	fmt.Println(p)
	fmt.Println(p.name) // to get a specific field we can use dot notation

	ref := 7
	fmt.Println(&ref) // & return the memory address of the variable // so this gives a pointer to ref
	inc(&ref)
	fmt.Println(ref)
}
// passing pointer in GO
func inc (x *int){
	*x++; //use * to dereferening the pointer 
}

// function in GO
func sumfunction(x int, y int) int { // int is the return type
	return x + y
}

// return a float64 value function
func sqrt(x float64) (float64,error) { // error is the built-in type. we can use it by importing "errors"
	if x < 0{
		return 0, errors.New("Undefined for negative number") // use New to create an error
		 // the fuction will only accept input which is > 0, if < 0, return an error
	}
	return math.Sqrt(x), nil // since this function will return two value, we return nil if input > 0
	//  "Math" package is used 
}

// struct is a collection of fields so you can group things together to create a more logical type
type person struct {
	name string
	age int
}

