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
	project  string
	zone     string
	instance string
}

func readCredential(file string) Credentials {
	c := Credentials{}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.close()

	data, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatal(err)
	}

	c, err = yaml.Marshal(&data)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func main() {
	c := readCredential("../credentials.yaml")

	ctx := context.Background()
	gcpClient, err := google.DefaultClient(ctx, compute.CloudPlatformScope)

	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(gcpClient)

	if err != nil {
		log.Fatal(err)
	}

	res, err := computeService.Instances.Start(c.project, c.zone, c.instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", res)
}
