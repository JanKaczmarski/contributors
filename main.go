package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Contributor struct {
	Login         string `json:"login"`
	Contributions int    `json:"contributions"`
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Print("Token not found. You must set it in your environemt like")
		log.Print("export GITHUB_TOKEN=<token>")
		os.Exit(1)
	}

	url := "https://api.github.com/repos/golang/go/contributors"
	// using Pointer semantics
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "token"+token)

	// Create http.Client and make the request
	cl := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := cl.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println("API responded with", res.Status)
		// Copy res.Body to Stderr on local machine
		io.Copy(os.Stderr, res.Body)
		os.Exit(1)
	}

	var cons []Contributor
	if err := json.NewDecoder(res.Body).Decode(&cons); err != nil {
		log.Fatal(err)
	}

	for i, con := range cons {
		fmt.Println(i, con.Login, con.Contributions)
	}

}
