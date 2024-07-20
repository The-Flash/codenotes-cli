package cli

import (
	"fmt"

	"github.com/The-Flash/codenotes-cli/internal/notes"
	"github.com/spf13/cobra"
)

var message string
var path string
var lineNo *int
var isSticky *bool

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addNoteCmd)
	addNoteCmd.Flags().StringVarP(&message, "message", "m", "", "message to attach")
	addNoteCmd.MarkFlagRequired("message")
	addNoteCmd.Flags().StringVarP(&path, "path", "p", "", "path to file/directory")
	addNoteCmd.MarkFlagRequired("path")
	lineNo = addNoteCmd.Flags().IntP("line-no", "n", -1, "")
	isSticky = addNoteCmd.Flags().Bool("sticky", false, "Is line note")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a note",
}

var addNoteCmd = &cobra.Command{
	Use:   "note",
	Short: "Add a note",
	Run: func(cmd *cobra.Command, args []string) {
		nc, err := notes.NewNoteConfig(
			message,
			path,
			notes.WithIsSticky(*isSticky),
			notes.WithLineNo(*lineNo),
		)
		if err != nil {
			panic(err)
		}
		if id, err := notes.AddNote(nc); err != nil {
			panic(err)
		} else {
			fmt.Println(id)
		}
	},
}
