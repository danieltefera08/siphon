package main

import (
	"fmt"
	"os"

	"siphon/cli/dtui/internal/config"
	"siphon/cli/dtui/internal/docker"
	"siphon/cli/dtui/plugins/compose"
	"siphon/cli/dtui/plugins/containers"
	"siphon/cli/dtui/plugins/deploy"
	"siphon/cli/dtui/plugins/images"
	"siphon/cli/dtui/plugins/settings"
	"siphon/cli/dtui/plugins/home"
	"siphon/cli/uix"
)

func main() {
	configPath := config.DefaultPath()
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	cfg := config.Load(configPath)

	client, err := docker.NewClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Docker client error: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	if err := client.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "Docker daemon not running: %v\n", err)
		os.Exit(1)
	}

	app := uix.NewApp("dtui > ")

	containerPlugin := containers.New(client)
	imagePlugin := images.New(client)
	composePlugin := compose.New(client, cfg)
	deployPlugin := deploy.New(client, cfg)
	settingsPlugin := settings.New(cfg)

	app.Register(containerPlugin)
	app.Register(imagePlugin)
	app.Register(composePlugin)
	app.Register(deployPlugin)
	app.Register(settingsPlugin)

	welcomeScreen := home.NewScreen()
	app.SetHome(func() string {
		welcomeScreen.SetSize(app.SafeWidth(), app.BodyHeight())
		return welcomeScreen.View()
	})

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
