package stop

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

// ServerConfig is a set of properties which specify an instance.
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

// Stop stops a GCE instance.
func Stop(w http.ResponseWriter, r *http.Request) {
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

	res, err := computeService.Instances.Stop(s.Project, s.Zone, s.Instance).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Failed to start instance: %v", err)
	}

	fmt.Printf("%#v\n", res)
}
