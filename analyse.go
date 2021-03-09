package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/hbollon/go-edlib"
)

const (
	packageEndpoint string = "https://replicate.npmjs.com/_all_docs"
)

type packageFeed struct {
	Packages []npmPackage `json:"rows"`
}

type npmPackage struct {
	ID string `json:"id"`
}

func main() {
	document, err := http.Get(packageEndpoint)
	if err != nil {
		panic(err)
	}

	defer document.Body.Close()

	// File
	// f, _ := os.Open("test.json")
	// defer f.Close()

	feed := &packageFeed{}

	start := time.Now()
	packageNames := make([]string, 0, 0)

	err := json.NewDecoder(f).Decode(document.Body)

	if err != nil {
		panic(err)
	}

	for _, pkg := range feed.Packages {
		packageNames = append(packageNames, pkg.ID)
	}

	fmt.Printf("Population finished in %s\n", time.Since(start))
	fmt.Printf("%v\n", len(packageNames))

	matches, err := edlib.FuzzySearchThreshold("this-is-an-unique-package-name", packageNames, 0.1, edlib.DamerauLevenshtein)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Matching finished in %s\n", time.Since(start))
	fmt.Printf("%v\n", matches)
}
