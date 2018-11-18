package main
import "fmt"

func main() {
	input := []int{1, 3, 5, 7, 9}
	k := 2

	// real roate
	k = k % len(input)

	// start 1
	// roate k
	for i:=1; i<=k; i++ {
		input = rotateOnce(input)
	}

	fmt.Println(input)
}


// rotate 1 right
func rotateOnce(input []int) []int {
	// get last element
	last := input[len(input)-1]

	// start at last index, until i still has
	for i:=len(input)-1; i>0; i-- {
		// last assigned by 2nd last
		input[i] = input[i-1]		
	}

	input[0] = last
		
	return input
}
