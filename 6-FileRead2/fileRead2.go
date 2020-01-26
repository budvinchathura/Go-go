package main

import(
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	// "strconv"
)

func main() {
	counts := make(map[string]int);
	if len(os.Args) == 1{
		fmt.Println("Please provide file names as arguments");
		os.Exit(0);
	}
	for _, fileName := range os.Args[1:]{
		data, err := ioutil.ReadFile(fileName);				//reads whole file
		if err!=nil{
			fmt.Println(err);
			continue;
		}
		for _, line := range strings.Split(string(data),"\n"){
			counts[line]++
		}
	}

	for line,n := range counts{
		if n>1{
			// fmt.Printf(line+"\t"+strconv.Itoa(n)+" times\n");		//will not print correctly becuase \r is a special character
			fmt.Printf("%q \t %d times\n",line,n);
		}
	}

}