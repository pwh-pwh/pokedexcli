package cli

import "os"

func exitCommand() error {
	os.Exit(0)
	return nil
}
