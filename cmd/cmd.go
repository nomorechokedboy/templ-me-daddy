package cmd

import (
	"log"
	"templ-me-daddy/cmd/serve"

	"github.com/spf13/cobra"
)

func new() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.AddCommand(serve.New())

	return cmd
}

func Exec() {
	cmd := new()
	if err := cmd.Execute(); err != nil {
		log.Fatal("Exec err: ", err)
	}
}
