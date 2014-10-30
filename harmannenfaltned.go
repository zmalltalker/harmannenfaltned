package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"errors"
	"os"
	"flag"
)

type Mannen struct {
	FaltNed bool `json:"falt_ned"`
}

func getResponse() (error, bool) {
	resp, err := http.Get("http://www.vondess.com/mannen/api")
	defer resp.Body.Close()
	if err == nil {
		body, ioerr := ioutil.ReadAll(resp.Body)
		if ioerr != nil {
			return errors.New("IO error"), false
		}
		var mannen Mannen
		jsonerr := json.Unmarshal(body, &mannen)
		if jsonerr != nil {
			return errors.New("OMG"), false
		}
		return nil, mannen.FaltNed
	} else {
		return errors.New("API error"), false
	}
}
func main(){
	runQuiet := flag.Bool("q", false, "Run quietly")
	flag.Parse()
	err, v := getResponse()
	if err == nil {
		if v {
			if !*runQuiet {
				fmt.Println("Ja")
			}
			os.Exit(1)
		} else {
			if !*runQuiet {
				fmt.Println("Nei")
			}
			os.Exit(0)
		}
	} else {
		if !*runQuiet {
			fmt.Printf("Error: %v\n", err)
		}
		os.Exit(2)
	}
}
