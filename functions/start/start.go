package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

type ServerConfig struct {
	Project  string
	Zone     string
	Instance string
}

func readServerConfig() *ServerConfig {
	p := os.Getenv("SERVER_PROJECT")
	z := os.Getenv("SERVER_ZONE")
	i := os.Getenv("SERVER_INSTANCE")

	return &ServerConfig{p, z, i}
}

func main() {
	s := readServerConfig()

	ctx := context.Background()
	gcpClient, err := google.DefaultClient(ctx, compute.CloudPlatformScope)

	if err != nil {
		log.Fatalf("Failed to create GCP Client: %v", err)
	}

	computeService, err := compute.New(gcpClient)

	if err != nil {
		log.Fatalf("Failed to create GCE service: %v", err)
	}

	res, err := computeService.Instances.Start(s.Project, s.Zone, s.Instance).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Failed to start instance: %v", err)
	}

	fmt.Printf("%#v\n", res)
}
