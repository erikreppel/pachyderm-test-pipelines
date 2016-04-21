package main

import (
	"bufio"
	"fmt"
	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pachyderm/pachyderm/src/client/pfs"
	"log"
	"os"
	"path"
)

// CreateFileOfNumbers writes a file and returns the path
func CreateFileOfNumbers(n int, filePath string) (outPath string, e error) {

	f, err := os.Create(filePath)
	if err != nil {
		return "An error occured", err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, i)
	}
	return filePath, nil
}

func main() {

	apiClient, err := client.NewFromAddress("10.0.0.215:650")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Established connection to Pachyderm")

	repoName := "numbers"

	err = pfs.CreateRepo(apiClient, repoName)
	if err != nil {
		log.Fatal(err)
	}

	commit, err := pfs.StartCommit(apiClient, repoName, "", "")
	if err != nil {
		log.Fatal(err)
	}

	defer pfs.FinishCommit(apiClient, repoName, commit.ID)

	filePath := path.Join("/pfs", repoName, commit.ID, "numbers")

	_, err = CreateFileOfNumbers(10000, filePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created a file of numbers in pachyderm")

}
