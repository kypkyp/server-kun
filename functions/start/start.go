package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
	"gopkg.in/yaml.v2"
)

type Credentials struct {
	Project  string
	Zone     string
	Instance string
}

func readCredential(file string) Credentials {
	c := Credentials{}

	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open credentials file: %v", err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalf("Failed to read credentials file: %v", err)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("Failed to read credentials file: %v", err)
	}

	return c
}

func main() {
	c := readCredential("../../credentials.yml")

	ctx := context.Background()
	gcpClient, err := google.DefaultClient(ctx, compute.CloudPlatformScope)

	if err != nil {
		log.Fatalf("Failed to create GCP Client: %v", err)
	}

	computeService, err := compute.New(gcpClient)

	if err != nil {
		log.Fatalf("Failed to create GCE service: %v", err)
	}

	res, err := computeService.Instances.Start(c.Project, c.Zone, c.Instance).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Failed to start instance: %v", err)
	}

	fmt.Printf("%#v\n", res)
}
