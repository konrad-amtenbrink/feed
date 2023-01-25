package cmd

import (
	"fmt"
	"os"

	"github.com/konrad-amtenbrink/feed/cmd/server"
	"github.com/spf13/cobra"
)

func Execute() {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "-",
	}

	cmd.AddCommand(server.NewCmd())

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%v'", err)
		os.Exit(1)
	}
}
