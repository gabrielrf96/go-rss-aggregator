package tree

import (
	"fmt"
	"iter"
	"os"
	"text/tabwriter"
)

type CommandTree struct {
	Key            string
	commands       map[string]*Command
	sections       map[string]*CommandTree
	commandOrder   []string
	sectionOrder   []string
	defaultCommand *Command
}

func (tree *CommandTree) Commands() iter.Seq[*Command] {
	return func(yield func(*Command) bool) {
		for _, commandKey := range tree.commandOrder {
			if !yield(tree.commands[commandKey]) {
				return
			}
		}
	}
}

func (tree *CommandTree) Sections() iter.Seq[*CommandTree] {
	return func(yield func(*CommandTree) bool) {
		for _, sectionKey := range tree.sectionOrder {
			if !yield(tree.sections[sectionKey]) {
				return
			}
		}
	}
}

func (tree *CommandTree) AddCommand(command *Command) {
	tree.commandOrder = append(tree.commandOrder, command.Key)
	tree.commands[command.Key] = command
}

func (tree *CommandTree) SetDefaultCommand(command *Command) {
	tree.defaultCommand = command
}

func (tree *CommandTree) NewCommand(key string, handler CommandHandler, help string) *Command {
	command := &Command{
		Key:     key,
		Help:    help,
		handler: handler,
	}

	tree.AddCommand(command)

	return command
}

func (tree *CommandTree) NewDefaultCommand(key string, handler CommandHandler, help string) *Command {
	command := tree.NewCommand(key, handler, help)

	tree.SetDefaultCommand(command)

	return command
}

func (tree *CommandTree) AddSection(section *CommandTree) {
	tree.sectionOrder = append(tree.sectionOrder, section.Key)
	tree.sections[section.Key] = section
}

func (tree *CommandTree) NewSection(
	key string,
	fn func(section *CommandTree),
) *CommandTree {
	section := NewCommandTree()
	section.Key = key

	fn(section)

	tree.AddSection(section)

	return section
}

func (tree *CommandTree) Run(args ...string) error {
	if len(args) == 0 {
		if tree.defaultCommand != nil {
			tree.defaultCommand.Run([]string{})

			return nil
		}

		return fmt.Errorf("invalid option: %s", tree.Key)
	}

	key := args[0]
	args = args[1:]

	section, ok := tree.sections[key]
	if ok {
		return section.Run(args...)
	}

	command, ok := tree.commands[key]
	if ok {
		return command.Run(args)
	}

	return fmt.Errorf("invalid option: %s", key)
}

func NewCommandTree() *CommandTree {
	return &CommandTree{
		commandOrder: make([]string, 0, 10),
		commands:     make(map[string]*Command, 10),
		sectionOrder: make([]string, 0, 10),
		sections:     make(map[string]*CommandTree, 10),
	}
}

type CommandHandler func(args []string) error

type Command struct {
	Key     string
	Help    string
	handler CommandHandler
}

func (command *Command) Run(args []string) error {
	return command.handler(args)
}

func Print(t *CommandTree) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 5, ' ', 0)

	printTree(t, w, t.Key)

	w.Flush()
}

func printTree(t *CommandTree, w *tabwriter.Writer, prefix string) {
	if t.Key != "" {
		fmt.Fprintf(w, "\n %s:\n", t.Key)
	}

	for command := range t.Commands() {
		fmt.Fprintf(w, " %s%s\t%s\n", prefix, command.Key, command.Help)
	}

	for section := range t.Sections() {
		sectionPrefix := fmt.Sprintf("%s %s ", prefix, section.Key)
		printTree(section, w, sectionPrefix)
	}
}
