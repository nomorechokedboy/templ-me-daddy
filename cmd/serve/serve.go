package serve

import (
	"log"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ArbitraryArgs,
		Use:   "serve",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Serving dinner...")
			app := NewApp()
			app.RunHttpServer()
		},
	}

	return cmd
}
