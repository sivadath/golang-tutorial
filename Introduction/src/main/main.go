package main

import (
	"fmt"
    "prime"
)



func main() {
	var status2 bool
	status2 = prime.IsPrime(1)
	fmt.Println("1 IsPrime :", status2)
	fmt.Println("2 IsPrime :", prime.IsPrime(2))
	fmt.Println("3 IsPrime :", prime.IsPrime(3))
	fmt.Println("4 IsPrime :", prime.IsPrime(4))
	fmt.Println("5 IsPrime :", prime.IsPrime(5))
	//Declaration and instantiation with :=
	status := true
	fmt.Println("Boolean variable declared and instatiated:", status)

	//Arrays:
	var arr [3]int //Array declaration and instantiation.
	fmt.Println("Array right after defining:",arr)
	arr[1] = 1
	fmt.Println("Array with some value changed:",arr)

	arr2 := [3]int{1,2,3} //Array declared and assigned value on the same line.
	fmt.Println("Array declared and assigend value in the same line:", arr2)

	//arr = append(arr,1) //Not allowed for arrays.

	//Slices:
	var slc []int //Slice declaration.
	fmt.Println("Slice just delcared and not instantiated:", slc)
	slc = make([]int,0) //Slice instantiation. Second argument is the size of slice.
	fmt.Println("Slice after instantiated:", slc)

	
	slc2 := []int{} //Slice declaration and instantiation in same line.
	//slc2 := make([]int,0) //Can also be done this way.
	fmt.Println("Slice declared and instatiated in same line:", slc2)
	slc2 = append(slc2,1) //Appending values to slice.
	fmt.Println("Values added to slice:", slc2)

	slc3 := []int{1,2,3} //Slice declaration, instantiation, values assigning in same line.
	fmt.Println("Slice declared and instatiated in same line:", slc3)
	slc3 = append(slc3,4)
	fmt.Println("Slice after appending another value:", slc3)

	fmt.Println("Lenght of the slice:", len(slc3)) //Shows element count.

	//Maps:
	var mp map[string]int //Map declaration.
	//mp["one"] = 1 will fail as not instantiated
	fmt.Println("Map declared and not instantiated:",mp)
	mp = make(map[string]int) //Map instantiation.
	fmt.Println("Map after instantiation:",mp)
	mp["one"] = 1
	mp["two"] = 2
	fmt.Println("Map with some keys inserted:",mp)

	mp2 := map[string]int{"Okati":1,"Rendu":2} //Map declaration and instantiation in same line.
	fmt.Println("Map declares, instantiated with few keys in same line:",mp2)

	fmt.Println("Length of a map:", len(mp)) //Shows keys count.


	fmt.Println("Variable imported from other package:",prime.I)
	//fmt.Println(prime.i) //Not possible as the variable is not title cased.
}
