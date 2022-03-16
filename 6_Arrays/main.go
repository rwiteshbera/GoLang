package main

import "fmt"

func main() {
	var name = [6]string{"Nicole", "Nana", "Peter"}
	fmt.Println(name)

	var id [5]int
	id[0] = 5
	id[1] = 10
	id[2] = 20

	fmt.Println(id[0])
	fmt.Println(id[1])
	fmt.Println(id)

	// Length of the array
	fmt.Println("Length of id array:", len(id))

	// What if we don't know the size of the array
	// Can we make a list that is dynamic in size

	// Slice is an abstraction of an array
	// Slices are more flexible and powerful
	// Slices are also indexed-based and have a size, but is resized when needed

	//  Internally, slice and an array are connected with each other, a slice is a reference to an underlying array. It is allowed to store duplicate elements in the slice.

	mySlice := name[0:3]

	fmt.Println(mySlice)

	// Length of the slice
	fmt.Println("Length of the slice:", len(mySlice))

	// Capacity of the slice
	fmt.Println("Capacity of the slice:", cap(mySlice))

}
