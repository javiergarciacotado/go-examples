package main

func fourthMain() { //rename to main for its execution

	var i interface{}
	Describe(i)

	i = 42 //may hold values of any type
	Describe(i)

	i = "hello"
	Describe(i)

}

// func describe(i interface{}) {
// fmt.Printf("(%v, %T)\n", i, i)
// }
