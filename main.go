package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Equanox/gotron"
	"github.com/pkg/errors"

	proto "github.com/talos-systems/talos/api/os"
	"github.com/talos-systems/talos/cmd/osctl/pkg/client"
	"github.com/talos-systems/talos/pkg/constants"
)

type UIEvent struct {
	*gotron.Event
	Type string `json:"type"`
}

func newClient() (*client.Client, error) {
	t, creds, err := client.NewClientTargetAndCredentialsFromConfig("/Users/tgerla/.talos/config")
	if err != nil {
		return nil, errors.Wrap(err, "error getting client credentials")
	}
	c, err := client.NewClient(creds, t, constants.OsdPort)
	if err != nil {
		return nil, errors.Wrap(err, "error constructing client")
	}

	return c, nil
}

func main() {
	// Create a new browser window instance
	window, err := gotron.New("ui/dist")
	if err != nil {
		panic(err)
	}

	// Alter default window size and window title.
	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980
	window.WindowOptions.Title = "Talos"

	// Start the browser window.
	// This will establish a golang <=> nodejs bridge using websockets,
	// to control ElectronBrowserWindow with our window object.
	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	// Open dev tools must be used after window.Start
	// window.OpenDevTools()

	window.On(&gotron.Event{Event: "version"}, func(bin []byte) {
		c, err := newClient()
		if err != nil {
			log.Printf("%+v", err)
		}
		version, err := c.Version(context.Background())
		if err != nil {
			fmt.Printf("error getting version: %+v", err)
		}
		b, err := json.Marshal(version)
		if err != nil {
			fmt.Printf("error marshalling response: %+v", err)
		}
		window.Send(&UIEvent{Event: &gotron.Event{string(b)}, Type: "version"})
	})

	window.On(&gotron.Event{Event: "services"}, func(bin []byte) {
		c, err := newClient()
		if err != nil {
			log.Printf("%+v", err)
		}
		reply, err := c.ServiceList(context.Background())
		if err != nil {
			fmt.Printf("error getting services: %+v", err)
		}
		b, err := json.Marshal(reply.Services)
		if err != nil {
			fmt.Println("error:", err)
		}
		window.Send(&UIEvent{Event: &gotron.Event{string(b)}, Type: "services"})
	})

	window.On(&gotron.Event{Event: "containers"}, func(bin []byte) {
		c, err := newClient()
		if err != nil {
			log.Printf("%+v", err)
		}
		reply, err := c.Containers(context.Background(), "system", proto.ContainerDriver_CONTAINERD)
		if err != nil {
			fmt.Printf("error getting containers: %+v", err)
		}
		b, err := json.Marshal(reply.Containers)
		if err != nil {
			fmt.Println("error:", err)
		}
		window.Send(&UIEvent{Event: &gotron.Event{string(b)}, Type: "containers"})
	})

	window.On(&gotron.Event{Event: "processes"}, func(bin []byte) {
		c, err := newClient()
		if err != nil {
			log.Printf("%+v", err)
		}
		reply, err := c.Processes(context.Background())
		if err != nil {
			fmt.Printf("error getting processes: %+v", err)
		}
		b, err := json.Marshal(reply.Processes)
		if err != nil {
			fmt.Println("error:", err)
		}
		window.Send(&UIEvent{Event: &gotron.Event{string(b)}, Type: "processes"})
	})

	window.On(&gotron.Event{Event: "routes"}, func(bin []byte) {
		c, err := newClient()
		if err != nil {
			log.Printf("%+v", err)
		}
		reply, err := c.Routes(context.Background())
		if err != nil {
			fmt.Printf("error getting routes: %+v", err)
		}
		b, err := json.Marshal(reply.Routes)
		if err != nil {
			fmt.Println("error:", err)
		}
		window.Send(&UIEvent{Event: &gotron.Event{string(b)}, Type: "routes"})
	})

	window.On(&gotron.Event{Event: "interfaces"}, func(bin []byte) {
		c, err := newClient()
		if err != nil {
			log.Printf("%+v", err)
		}
		reply, err := c.Interfaces(context.Background())
		if err != nil {
			fmt.Printf("error getting interfaces: %+v", err)
		}
		b, err := json.Marshal(reply.Interfaces)
		if err != nil {
			fmt.Println("error:", err)
		}
		window.Send(&UIEvent{Event: &gotron.Event{string(b)}, Type: "interfaces"})
	})

	window.On(&gotron.Event{Event: "mounts"}, func(bin []byte) {
		c, err := newClient()
		if err != nil {
			log.Printf("%+v", err)
		}
		reply, err := c.Mounts(context.Background())
		if err != nil {
			fmt.Printf("error getting mounts: %+v", err)
		}
		b, err := json.Marshal(reply.Stats)
		if err != nil {
			fmt.Println("error:", err)
		}
		window.Send(&UIEvent{Event: &gotron.Event{string(b)}, Type: "mounts"})
	})

	// Wait for the application to close
	<-done
}
