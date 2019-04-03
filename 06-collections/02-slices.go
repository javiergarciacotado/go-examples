package main

import "fmt"

func secondMain() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var sliceOfPrimes []int = primes[1:4] //dynamically-sized, flexible view of the elements of an array

	fmt.Println(sliceOfPrimes)

	names := [4]string{
		"Gael",
		"Miriam",
		"Javi",
		"gallego",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[2:4]
	fmt.Println(a, b)

	b[0] = "anotherOne" //changing the elements of a slice, modifies the elements of its underlying array
	fmt.Println(names)
	fmt.Println(a, b)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, false},
		{5, true},
		{7, true},
	}
	fmt.Println(s)

	var nilSlice []int //non-value slice is nil!
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nil!")
	}

	sliceByMake1 := make([]int, 5) //len(sliceByMake) = 5
	printSlice("sliceByMake1", sliceByMake1)

	sliceByMake2 := make([]int, 0, 5)
	printSlice("sliceByMake2", sliceByMake2)

	sliceByMake3 := sliceByMake2[:2]
	printSlice("sliceByMake3", sliceByMake3)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}

	powFrom := make([]int, 10)
	for index := range powFrom { //omitting value
		powFrom[index] = 1 << uint(index)
	}

	for _, value := range powFrom {
		fmt.Printf("%v\n", value)
	}

}

func printSlice(sliceIdentifier string, slice []int) {
	fmt.Printf("%s len=%d, cap=%d %v\n", sliceIdentifier, len(slice), cap(slice), slice)
}
