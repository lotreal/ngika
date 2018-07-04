package main

import (
	"io"
	"os"

	"fmt"
	// "log"
	// "time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	// "github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
)

const (
	NgPort = "80/tcp"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	fmt.Println("This is a info")

	config := &container.Config{
		Image: "nginx",
		Volumes: map[string]struct{}{
			"/go/src/github.com/hg2c/example": {},
		},
		ExposedPorts: nat.PortSet{
			NgPort: struct{}{},
		},
	}

	hostConfig := &container.HostConfig{
		Binds: []string{
			"/Volumes/Kayle/w/hg2c/golang:/go/src/github.com/hg2c/example:rw",
		},
		PortBindings: nat.PortMap{
			NgPort: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8008",
				},
			},
		},
		// PortBindings: nat.PortMap
		// PublishAllPorts: true,
		// Mounts: []mount.Mount{
		// 	mount.Mount{
		// 		Type:     "bind",
		// 		Source:   "/Volumes/Kayle/w/hg2c/golang",
		// 		Target:   "/go/src/github.com/hg2c/example",
		// 		ReadOnly: false,
		// 		BindOptions: &mount.BindOptions{
		// 			Propagation: mount.PropagationRPrivate,
		// 		},
		// 	},
		// },
	}

	containerName := ""

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, containerName)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	fmt.Println("This is a info")

	// ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// _, errC := cli.ContainerWait(ctx2, resp.ID)
	// if err := <-errC; err != nil {
	// 	log.Fatal(err)
	// }
	if _, err = cli.ContainerWait(ctx, resp.ID); err != nil {
		panic(err)
	}
	fmt.Println("This is a info")

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}
