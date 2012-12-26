// map
package main

import (
	"fmt"
)

func main() {
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, // <--- must have a colon in the end
	}

	fmt.Println(monthdays["Jan"])
	
	monthdays["February"] = 29   // can add an element
	
	days, present := monthdays["Aug"]                   // check the presence if element
	fmt.Println("Presence of \"Aug\" is ", present, " it has ", days, " days")
	
	monthdays["Oct"] = 0, false       // delete an element
	
	
}
