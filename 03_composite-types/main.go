package main

import (
	"fmt"
)

func main() {
	// Arrays
	// Arrays are rarely used directly in Go. All of the elements in the array must be of the type that is specified (this doesn't mean they are always of the same type).
	var anArray [3]int
	fmt.Println("an array with zero values", anArray)

	var anotherArray = [3]int{10, 20, 30}
	fmt.Println("an array with values", anotherArray)
 
	// leaving off the number
	var withoutNumberOfElementsArray = [...]int{10, 20, 30}
	fmt.Println("an array without number of elements", withoutNumberOfElementsArray)

	// sparse array: an array where most elements are set to their zero value
	var sparseArray = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println("an array where most of elements are zero", sparseArray)

	// comparing arrays
	var firstArrayToBeCompared = [...]int{1, 2, 3}
	var secondArrayToBeCompared = [3]int{1, 2, 3}
	fmt.Println("arrays are equal?", firstArrayToBeCompared == secondArrayToBeCompared)

	// one-dimensional arrays
	var oneDimensionalArray [2][3]int //no matrix support!
	fmt.Println("an array one dimension", oneDimensionalArray)

	// length of the array
	fmt.Println("length of the array", len(oneDimensionalArray))

	// How would be an array without declaring its size?
	// |
	// |
	// V
	// slice: length is not part of the type of a slice

	var aSlice = []int{10, 20, 30} //Using [...] makes an array, using [] makes a slice
	fmt.Println("a slice ", aSlice)

	// nil slice 
	var nilSlice[] int
	fmt.Println("slice is nil?", nilSlice == nil)

	// appending values
	var appendSlice[] int
	appendSlice = append(appendSlice, 1)
	fmt.Println("slice values", appendSlice)
	// more than one value at a time
	appendSlice = append(appendSlice, 2, 3, 4)
	fmt.Println("slice values", appendSlice)

	// Capacity: number of consecutive memory locations reserved
	var sliceCapacity[] int
	fmt.Println("slice capacity inital", sliceCapacity, len(sliceCapacity), cap(sliceCapacity))
	sliceCapacity = append(sliceCapacity, 10)
	fmt.Println("slice capacity after append 10", sliceCapacity, len(sliceCapacity), cap(sliceCapacity))
	sliceCapacity = append(sliceCapacity, 20)
	fmt.Println("slice capacity after append 20", sliceCapacity, len(sliceCapacity), cap(sliceCapacity))
	sliceCapacity = append(sliceCapacity, 30)
	fmt.Println("slice capacity after append 30", sliceCapacity, len(sliceCapacity), cap(sliceCapacity))
	sliceCapacity = append(sliceCapacity, 40)
	fmt.Println("slice capacity after append 40", sliceCapacity, len(sliceCapacity), cap(sliceCapacity))
	sliceCapacity = append(sliceCapacity, 50)
	fmt.Println("slice capacity after append 50", sliceCapacity, len(sliceCapacity), cap(sliceCapacity))
	// While it’s nice that slices grow automatically, it’s far more efficient to size them once. If you know how many things you plan to put into a slice, create the slice with the correct initial capacity.
	// |
	// |
	// V
	// make: to create an empty slice that already has a length or capacity specified
	var makeInitial= make([] int, 5)
	fmt.Println("make initial len/cap", makeInitial, len(makeInitial), cap(makeInitial))

	var makeAnotherWithZeroLength = make([] int, 0, 10)
	fmt.Println("make another initial len/cap", makeAnotherWithZeroLength, len(makeAnotherWithZeroLength), cap(makeAnotherWithZeroLength))
	makeAnotherWithZeroLength = append(makeAnotherWithZeroLength, 5, 6, 7, 8)
	fmt.Println("make another initial len/cap", makeAnotherWithZeroLength, len(makeAnotherWithZeroLength), cap(makeAnotherWithZeroLength))

	// Tip: never specify a capacity that's less than length (mostly sure your program will panic at runtime

	// wip

	// Slicing slices

	// converting arrays to slices

	// Maps
}
