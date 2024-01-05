package helpers

func Commands(prev []string, cmds ...string) []string {
	return append(prev, cmds...)
}
