package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codenotes",
	Short: "CLI for Codenotes",
	Long: `Codenotes is a tool for attaching notes/messages to files/folder.
	You can attach notes to lines in a file	
	`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
