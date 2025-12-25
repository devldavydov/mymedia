package main

import (
	"log"

	"github.com/devldavydov/mymedia/internal/mymediacli/exifr"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mymediacli",
	Short: "MyMedia Cli",
	Long:  "Command line tool for media (photo, video) operations",
}

func init() {
	rootCmd.AddCommand(exifr.Cmd)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
