package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Print("Token not found. You must set it in your environemt like")
		log.Print("export GITHUB_TOKEN=<token>")
		os.Exit(1)
	}

	c, err := github.NewClient(token)
	if err != nil {
		log.Fatal(err)
	}

	if err := process("ardanlabs/gotraining", cl); err != nil {
		log.Fatal(err)
	}
}

type contributorLister interface {
	ContributorList(string) ([]github.Contributor, error)
}

func process(repo string, c contributorLister) error {
	cons, err := c.ContributorList(repo)
	if err != nil {
		return err
	}

	for i, con := range cons {
		fmt.Println(i, con.Login, con.Contributions)
	}

	return nil
}
