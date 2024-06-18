package cli

import (
	"fmt"

	"github.com/The-Flash/codenotes-cli/internal/notes"
	"github.com/spf13/cobra"
)

var directory string

func init() {
	initCmd.PersistentFlags().StringVar(&directory, "directory", ".", "Directory to initialize codenotes project")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a directory/repo as a codenotes project",
	Run: func(cmd *cobra.Command, args []string) {
		if p, err := notes.InitializeProject(directory); err != nil {
			panic(err)
		} else {
			fmt.Println("Initialized codenotes project " + p)
		}
	},
}
