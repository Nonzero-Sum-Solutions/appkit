package cmd

import (
	"fmt"
	"os"

	gmodel "github.com/mattmunz/appkit/model/gen/appkit"
	"github.com/spf13/cobra"
)

type commandRunner func(cmd *cobra.Command, args []string)

func WrapRunner(run func(_ gmodel.CLI, _ *cobra.Command, _ []string) error, cli gmodel.CLI) commandRunner {
	return func(command *cobra.Command, args []string) {
		if err := run(cli, command, args); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
			os.Exit(1)
		}
	}
}
