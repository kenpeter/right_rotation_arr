package main
import "fmt"

func main() {
	input := []int{1, 3, 5, 7, 9}
	k := 2

	// real roate k
	k = k % len(input)

	// start 1, because 1 rotation
	// <= k, because k times  
	for i:=1; i<=k; i++ {
		input = rotateOnce(input)
	}

	fmt.Println(input)
}


// rotate 1 right
func rotateOnce(input []int) []int {
	// remember last 
	last := input[len(input)-1]

	// start at last 
	for i:=len(input)-1; i>0; i-- {
		// last = last - 1 
		input[i] = input[i-1]		
	}

	// last becomes 1st
	input[0] = last
		
	return input
}
