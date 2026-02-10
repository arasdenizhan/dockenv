package docker

import (
	"bytes"
	"context"
	"log"

	"github.com/moby/moby/client"
)

func GetImageEnvs(name string) ([]string) {
    ctx := context.Background()
    
    apiClient, err := client.New(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	defer apiClient.Close()

    buffer := bytes.Buffer{}
    image, err := apiClient.ImageInspect(ctx, name, client.ImageInspectWithRawResponse(&buffer))
    if err != nil {
        log.Fatal(err)
    }
    return image.Config.Env
}
