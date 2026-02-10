package docker

import (
	"context"
	"log"

	"github.com/moby/moby/client"
)

func GetContainerEnvs(id string) ([]string) {
    ctx := context.Background()
    apiClient, err := client.New(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer apiClient.Close()

	inspect, err := apiClient.ContainerInspect(ctx, id, client.ContainerInspectOptions{})
	if err != nil {
		log.Fatal(err)
	}
    return inspect.Container.Config.Env
}
