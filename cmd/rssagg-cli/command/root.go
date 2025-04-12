package command

import (
	"fmt"

	"github.com/gabrielrf96/go-practice-rss-aggregator/cmd/rssagg-cli/tree"
)

func RunServer(args []string) error {
	return runCmd("./bin/rssagg-server", args...)
}

func Help(t *tree.CommandTree) error {
	fmt.Print("\nAvailable commands:\n\n")

	tree.Print(t)

	return nil
}
