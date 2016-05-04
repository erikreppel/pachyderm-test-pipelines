package main

import (
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/pachyderm/pachyderm/src/client"
	ppsclient "github.com/pachyderm/pachyderm/src/client/pps"
	"golang.org/x/net/context"
	"io"
	"log"
	"os"
)

func main() {
	pachAddr := os.Getenv("PACHD_PORT_650_TCP_ADDR")
	apiClient, err := client.NewFromAddress(pachAddr + ":650")
	if err != nil {
		log.Fatalln(err)
	}

	pipelinePath := "pipeline.json"

	pipelineFile, err := os.Open(pipelinePath)
	if err != nil {
		log.Fatalf("Error opening %s: %s\n", pipelinePath, err.Error())
	}
	defer func() {
		if err := pipelineFile.Close(); err != nil {
			log.Fatalf("Error closing%s: %s", pipelinePath, err.Error())
		}
	}()

	var request ppsclient.CreatePipelineRequest
	decoder := json.NewDecoder(pipelineFile)
	for {
		message := json.RawMessage{}
		if err := decoder.Decode(&message); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Error reading from stdin: %s", err.Error())
			}
		}
		if err := jsonpb.UnmarshalString(string(message), &request); err != nil {
			log.Fatalf("Error reading from stdin: %s", err.Error())
		}
		if err := apiClient.CreatePipeline(
			context.Background(),
			&request,
		); err != nil {
			log.Fatalf("Error from CreatePipeline: %s", err.Error())
		}
	}
}
