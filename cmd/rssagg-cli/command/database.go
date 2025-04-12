package command

func MigrateUp(args []string) error {
	return runCmd("go", "tool", "goose", "up")
}

func MigrateDown(args []string) error {
	return runCmd("go", "tool", "goose", "down")
}

func MigrateCreate(args []string) error {
	args = append([]string{"tool", "goose", "create"}, args...)

	return runCmd("go", args...)
}

func Generate(args []string) error {
	return runCmd("go", "tool", "sqlc", "generate")
}
