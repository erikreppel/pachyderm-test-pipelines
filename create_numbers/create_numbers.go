package main

import (
	"bufio"
	"fmt"
	"github.com/pachyderm/pachyderm/src/client"
	"log"
)

func main() {
	c, err := client.NewFromAddress("localhost:30650")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Established connection to Pachyderm")

	repoName := "test_data"
	log.Println("Creating a new repo named", repoName)
	err = c.CreateRepo(repoName)
	log.Println(err)

	log.Println("Successfully created the repo", repoName)

	repoInfo, err := c.InspectRepo(repoName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(repoInfo)

	log.Println("Starting a commit")
	branch := ""
	var parentCommitID string
	commit, err := c.StartCommit(repoName, parentCommitID, branch)
	if err != nil {
		log.Fatal(err)
	}

	commitID := commit.ID
	defer c.FinishCommit(repoName, commitID)

	for i := 0; i < 100; i++ {

		fileName := fmt.Sprintf("test/this/numbers%d", i)
		file, err := c.PutFileWriter(repoName, commitID, fileName, "")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		w := bufio.NewWriter(file)
		for j := i * 10; j < i*10+10; j++ {
			fmt.Fprintln(w, j)
		}
		err = w.Flush()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("/pfs/%s/%s/%s", repoName, commitID, fileName)
	}
}
