package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string


	//i:=1 is equivalent to
	// var i int
	// var i = 0
	// var i int  = 0

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " | "
		fmt.Println(os.Args[i])
	}
	fmt.Println(s)

	fmt.Println();
	fmt.Println();



	var s2, sep2 string;
	// _ is a blank identifier 
	//used whene ver syntax requires a variable name but program logic does not
	//Go does not permit declaring unused variables unless it is a blank identifier

	for _, arg := range os.Args[1:]{
		s2 += sep2 + arg;
		sep2  = " / ";
	}
	fmt.Println(s2);


	fmt.Println();
	fmt.Println();
	
	for index,value := range(os.Args){
		fmt.Println("Arg "+strconv.Itoa(index)+" = "+value);
	}
}
