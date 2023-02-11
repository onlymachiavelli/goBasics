package main 

//making request api 
import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main () {


	const targetURL = "https://randomuser.me/api" 
	resp, err := http.Get(targetURL)
	if err != nil {
		fmt.Println(err)
	}

	bddy, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println(e)
	}
	
	
	defer resp.Body.Close()
	fmt.Println(string(bddy))
	




}