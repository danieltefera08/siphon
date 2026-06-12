package main

import (
	"fmt"
	"os"

	"siphon/cli/dtui/internal/config"
	"siphon/cli/dtui/plugins/compose"
	cfgplugin "siphon/cli/dtui/plugins/config"
	"siphon/cli/dtui/plugins/containers"
	"siphon/cli/dtui/plugins/deploy"
	"siphon/cli/dtui/plugins/images"
	"siphon/cli/dtui/plugins/test"
	"siphon/cli/uix"
)

func main() {
	cfg := config.Load(config.DefaultPath())
	app := buildApp(cfg)
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func buildApp(cfg config.Config) *uix.Shell {
	app := uix.NewApp("dtui > ")
	app.RegisterModule(test.New())
	app.RegisterModule(containers.New())
	app.RegisterModule(images.New())
	app.RegisterModule(compose.New(cfg))
	app.RegisterModule(deploy.New(cfg))
	app.RegisterModule(cfgplugin.New(cfg, ""))
	app.AppendMessage(uix.RoleSystem, "DTUI — Docker Terminal UI. Type / to open the command palette.")
	app.AppendMessage(uix.RoleSystem, "Modules: /containers (ctr), /images (img), /compose, /deploy, /config (cfg), /test.")
	app.AppendMessage(uix.RoleSystem, "Docker modules require the Docker daemon; the app starts without it.")
	app.AppendMessage(uix.RoleSystem, "Prompt modes: / commands, @ references, # file picker, ! disabled shell prefix.")
	return app
}
