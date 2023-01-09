package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v2/pkg/bindings"
	"github.com/containers/podman/v2/pkg/bindings/containers"
	"github.com/containers/podman/v2/pkg/specgen"
)

func connectPodman() context.Context {
	sock_dir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sock_dir + "/podman/podman.sock"
	connection, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		fmt.Println("Podman Socket Connection error: ", err)
	}
	return connection
}

func createContainer(connection context.Context) {
	rawImage := "ubuntu"
	s := specgen.NewSpecGenerator(rawImage, false)
	s.Terminal = true
	r, err := containers.CreateWithSpec(connection, s)
	if err != nil {
		fmt.Println(err)
	}
	err = containers.Start(connection, r.ID, nil)
	if err != nil {
		fmt.Println(err)
	}
}
