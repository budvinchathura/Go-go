//counts duplicate lines in the input

package main

import(
	"fmt"
	"bufio"
	"os"
)

func main() {
	//make Allocates memory for referenced data types (slice, map, chan), plus initializes their underlying data structures
	fmt.Print("Counts duplicate inputs.\nEnter your inputs lines by line.\n")
	counts := make(map[string]int);					//keys are string and values are int
	input := bufio.NewScanner(os.Stdin);
	for input.Scan(){

		if err := input.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		if input.Text()==""{
			break;
		}else{
			counts[input.Text()]++
		}
		
	}
	for line,n := range counts{
		if n>1{
			fmt.Printf("%s \t %d times\n",line,n)
		}
	}
}