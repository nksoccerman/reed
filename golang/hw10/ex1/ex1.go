package main

import(
	"fmt"	
	"strconv"
)

func convtobits(n int) string {
	var bits string = ""
	for n != 0  {
		var bit = n & 1
		n = n >> 1 
		bits = strconv.Itoa(bit) + bits
	}
	return bits
}

func revstring(s string) string {
	var new_s string = ""
	var i = 0
	for i = len(s); i > 0; i-- {
		new_s = new_s + string(s[i-1])
	}
	return new_s
}

func main() {
        var s string = ""
        var x int = 0

        fmt.Println("Enter an integer: ")
        fmt.Scanln(&s)
        x, _ = strconv.Atoi(s)	
	for x > 0 {
		var fbits = convtobits(x)
		var rbits = revstring(fbits)
		if fbits == rbits {
			fmt.Printf("%d - %s\n", x, fbits)
		}
		x--
	}
}	
