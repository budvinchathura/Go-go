package main

import (
	"fmt"
	"os"
	// "io/ioutil"
	"io"
	"net/http"
	"strings"
)

func main(){
	for _, url := range os.Args[1:]{
		if !strings.HasPrefix(url,"http://"){
			url = "http://"+url
		}
		resp, err := http.Get(url)
		if err!= nil{
			fmt.Fprint(os.Stderr,"fetch: %v\n",err)
			os.Exit(1);
		}
		io.Copy(os.Stdout,resp.Body);			//without buffering
		// b,err := ioutil.ReadAll(resp.Body)	//with buffer
		header := resp.Header
		code := resp.StatusCode
		resp.Body.Close();
		// if err!= nil{
		// 	fmt.Fprintf(os.Stderr,"error in reading %s : %v",url,err);
		// 	os.Exit(1);
		// }
		fmt.Printf("%d\n",code)
		fmt.Printf("%s\n",header)
		// fmt.Printf("%s",b)
	}
}