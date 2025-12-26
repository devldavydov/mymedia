package exifr

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/devldavydov/mymedia/internal/common/exif"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "exifr",
	Short: "EXIF rename",
	Long:  "Rename photo/video to their EXIF timestamp in format",
	Run:   run,
}

var flags struct {
	dir     string
	format  string
	pattern string
	dry     bool
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

	Cmd.PersistentFlags().StringVarP(
		&flags.pattern,
		"pattern",
		"p",
		"",
		"Filename pattern for processing",
	)

	Cmd.PersistentFlags().BoolVar(
		&flags.dry,
		"dry",
		false,
		"Dry run",
	)
}

func run(cmd *cobra.Command, args []string) {
	entries, err := os.ReadDir(flags.dir)
	if err != nil {
		log.Fatalf("Failed to read directory '%s': %v\n", flags.dir, err)
	}

	var totalRenamed int64
	for _, entry := range entries {
		fileName := entry.Name()

		if flags.pattern != "" && !strings.Contains(fileName, flags.pattern) {
			continue
		}

		if !flags.dry {
			if err := exif.RenameFile(flags.dir, fileName); err != nil {
				log.Printf("Failed to rename '%s': %v\n", fileName, err)
				continue
			}
		} else {
			log.Printf("File to be renamed: '%s'\n", fileName)
		}

		totalRenamed += 1
	}

	var strTotal strings.Builder
	strTotal.WriteString("Total ")
	if flags.dry {
		strTotal.WriteString("to be ")
	}
	strTotal.WriteString(fmt.Sprintf("renaimed: %d", totalRenamed))
	log.Print(strTotal.String())
}
