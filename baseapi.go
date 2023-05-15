package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A constant Base URL
const BaseUrl = "https://google.com/api/"

func main() {
	// A slice of endpoints for the BaseUrl to query
	var endpoints = []string{"gmail", "maps", "drive"}

	// Print  Endpoints in formatted string
	for _, endpoint := range endpoints {
		fmt.Printf("Endpoint to be queried: %s%s\n", BaseUrl, endpoint)
	}

	// The returned JSON will be in a different types of  format, make slice of map with len(endpoints)
	responses := make([]map[string](interface{}), len(endpoints))

	// Print out the slice of empty maps
	fmt.Println(responses)

	// Make the HTTP requests
	for i, endpoint := range endpoints {
		//    GET request
		resp, err := http.Get(fmt.Sprintf("%s%s", BaseUrl, endpoint))

		if err != nil {
			fmt.Println("Error with endpoint:", endpoint, ", the error was: ", err)
		} else {
			// Read in the response body
			body, respErr := ioutil.ReadAll(resp.Body)

			if respErr != nil {
				fmt.Println("Error with a response body")
				return
			}

			// Unmarshalling the body into our map
			err1 := json.Unmarshal(body, &responses[i])
			if err1 != nil {
				fmt.Println("Error unmarshalling data")
			}
		}
	}

	//   Print out all the responses
	for _, response := range responses {
		fmt.Println("")
		fmt.Println(response)
	}

	// Print Thanks you msg
	fmt.Println("\n ***Thanks you Hava a nice Day!***")
}
