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

func main() {
	apiClient, err := client.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Established connection to Pachyderm")

	repoName := "data"
	log.Println("Creating a new repo named", repoName)

	pfs.CreateRepo(apiClient, repoName)
	log.Println("Successfully created the repo", repoName)

	repoInfo, err := pfs.InspectRepo(apiClient, repoName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(repoInfo)

	log.Println("Starting a commit")
	branch := ""
	var parentCommitID string
	commit, err := pfs.StartCommit(apiClient, repoName, parentCommitID, branch)
	if err != nil {
		log.Fatal(err)
	}

	commitID := commit.ID
	defer pfs.FinishCommit(apiClient, repoName, commitID)

	for i := 0; i < 10; i++ {

		fileName := fmt.Sprintf("numbers%d", i)
		filePath := path.Join("/pfs", repoName, commitID, fileName)

		file, err := os.Create(filePath)
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
