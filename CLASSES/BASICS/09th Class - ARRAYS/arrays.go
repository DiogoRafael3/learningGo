package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array [5]int   //to declare an array use var <VARNAME> [<ARRSIZE>]<TYPE>
	fmt.Println(array) //empty declaration fills an array with 5 0s

	array[0] = 1
	array[2] = 3
	fmt.Println(array)

	array2 := [3]int{1, 2, 3} //using type inference we still need to declare the type of the array
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4} //... allows the creation of an array with the amount of elements given, this still creates a static array
	fmt.Println(array3)

	//to create a dynamic array, we create a slice in go

	slice := []int{12, 13, 4, 6} //can also declare a slice from the start by using var <VARNAME> []<TYPE>
	fmt.Println(slice)
	slice = append(slice, 2)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3)) //as we can see through these prints, slices and arrays are totally different things

	slice2 := array2[1:2] //we can also create slices from pieces of arrays, in this case index 1 is inclusive but index 2 is exclusive, which means this will create a slice with one element only
	fmt.Println(slice2)

	array2[1] = 100 //slices work like pointers, so if the original array is changed and you have created a slice from it, then the position that was changed in the array will also change in the slice
	fmt.Println(array2)
	fmt.Println(slice2)

	//INTERNAL ARRAYS
	slice3 := make([]int32, 2, 3) //this creates a slice of strings, with the initial size 2 and max size 10, to be specific the make function creates allocates a specific space to whatever you put in the args
	fmt.Println(slice3)           //the function creates an array, then it returns a slice that references the first 10 positions of the array
	fmt.Println(len(slice3))      //length
	fmt.Println(cap(slice3))      //capacity

	slice3 = append(slice3, 3)
	slice3 = append(slice3, 4)
	fmt.Println(slice3)
	fmt.Println(cap(slice3)) //when go analyses that your slice created via make is about to overflow, then it creates another array of the same size to add to the slice
	//this means when i added two elements to my size 2 capacity 3 slice, it became a size 4 capacity 6 slice
	//go will never let your run out of space in an array
	//when you don't define a capacity for your slice, then the length of it will be considered the capacity
	slice4 := make([]int32, 2)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
