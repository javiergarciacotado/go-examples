package greeting

import "fmt"

func init() { //called when a package is declared
	fmt.Println("greeting/day.go ==> ", init())
}

var morning = "Good morning"
var Morning = "Hey," + morning
