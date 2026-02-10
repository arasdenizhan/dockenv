package docker

import (
	"context"
	"fmt"
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
    envs := inspect.Container.Config.Env

    fmt.Println("ENV VARIABLES:")
    for _, e := range envs {
        fmt.Println(e)
    }
    return []string{}
}
