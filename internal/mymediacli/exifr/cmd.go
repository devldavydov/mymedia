package exifr

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "exifr",
	Short: "EXIF rename",
	Long:  "Rename photo/video to their EXIF timestamp in format",
	Run:   run,
}

var flags struct {
	dir    string
	format string
}

func init() {
	curWD, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	Cmd.PersistentFlags().StringVarP(
		&flags.dir,
		"dir",
		"d",
		curWD,
		"Path to directory (current dir if empty)",
	)

	Cmd.PersistentFlags().StringVarP(
		&flags.format,
		"format",
		"f",
		"20060102_150405",
		"Timestamp Go format",
	)
}

func run(cmd *cobra.Command, args []string) {
	fmt.Printf("Do Exif rename in dir '%s' with format '%s'\n", flags.dir, flags.format)
}
