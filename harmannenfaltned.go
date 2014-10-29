package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type Mannen struct {
	FaltNed bool `json:"falt_ned"`
}

func main(){
	resp, err := http.Get("http://www.vondess.com/mannen/api")
	defer resp.Body.Close()
	if err == nil {
		body, ioerr := ioutil.ReadAll(resp.Body)
		if ioerr != nil {
			panic("IO error")
		}
		var mannen Mannen
		jsonerr := json.Unmarshal(body, &mannen)
		if jsonerr != nil {
			panic("OMG")
		}
		if mannen.FaltNed {
			fmt.Println("Ja")
		} else {
			fmt.Println("Nei")
		}
	}
}
