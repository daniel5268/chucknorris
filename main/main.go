package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JokesResponse struct {
	Joke string `json:"value"`
}

func main() {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random/")
	if err != nil {
		fmt.Println("ups:", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var jr JokesResponse
	json.Unmarshal(body, &jr)
	resp.Body.Close()
	fmt.Println(jr.Joke)
}
