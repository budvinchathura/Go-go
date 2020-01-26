package main

import(
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int);
	fileNames := os.Args[1:];
	if len(fileNames) == 0 {
		fmt.Print("Counts duplicate inputs.\nEnter your inputs lines by line.\nEnter ^Z to stop the program.\n")
		countLines(os.Stdin,counts);
	}else{
		for _, arg := range fileNames{
			f, err := os.Open(arg);				//open file
			if err != nil{
				fmt.Fprintf(os.Stderr," %v\n",err)			//%v for any in a natural format
				continue;
			}
			countLines(f,counts);
			f.Close();
		}
		
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("\"%s\" \t %d times\n",line,n)
		}
	}
}

func countLines(f *os.File,cnts map[string]int){
	input := bufio.NewScanner(f);
	for input.Scan(){
		cnts[input.Text()]++
	}
}