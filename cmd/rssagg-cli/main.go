package main

import (
	"fmt"
	"os"

	cmd "github.com/gabrielrf96/go-practice-rss-aggregator/cmd/rssagg-cli/command"
	"github.com/gabrielrf96/go-practice-rss-aggregator/cmd/rssagg-cli/tree"
)

func main() {
	t := getCommandTree()

	err := t.Run(os.Args[1:]...)
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}
}

func getCommandTree() *tree.CommandTree {
	var t = tree.NewCommandTree()

	t.NewDefaultCommand(
		"help",
		func(_ []string) error { return help(t) },
		"Show the list of available commands",
	)

	t.NewCommand("run", cmd.RunServer, "Start the RSS aggregator server")

	t.NewSection("db", func(section *tree.CommandTree) {
		section.NewCommand("up", cmd.MigrateUp, "Run pending migrations")
		section.NewCommand("down", cmd.MigrateDown, "Roll back last migration")
		section.NewCommand("new", cmd.MigrateCreate, "{{NAME}} [{sql|go}] Create a new migration")
		section.NewCommand("gen", cmd.Generate, "Generate query-related code using sqlc")
	})

	return t
}

func help(t *tree.CommandTree) error {
	cmd.Help(t)

	return nil
}
