package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)

func main() {
	startTime := time.Now();
	channel := make(chan string);

	for _, url := range os.Args[1:]{
		if !strings.HasPrefix(url,"http://"){
			url = "http://"+url
		}
		go fetch(url,channel);
	}

	for range os.Args[1:]{
		fmt.Println(<-channel);			//print all responses
	}
	fmt.Printf("%s elapsed\n",time.Since(startTime))
}

func fetch(url string, ch chan<- string){
	start := time.Now();
	resp, err := http.Get(url)
	if err!=nil{
		ch <- fmt.Sprint(err)		//send response to channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard,resp.Body);	//discarding just to count the bytes
	resp.Body.Close();
	if err!=nil{
		ch <- fmt.Sprintf("%v",err);
		return
	}
	t := time.Since(start);
	ch <- fmt.Sprintf("%s \t %d    %s",t,nbytes,url)
}