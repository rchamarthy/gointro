package main

import (
	"fmt"
	"os"
	"sync"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/rchamarthy/gointro/ggrep"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "ggrep",
		Short: "grep with a golang twist",
		Long:  "grep multiple files in parallel to increase performance",
		RunE:  grepCmd,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func grepCmd(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("need a pattern and at least one file")
	}
	pattern := args[0]
	files := args[1:]

	c := make(chan ggrep.Line, 256)
	wgScanners := &sync.WaitGroup{}

	for _, file := range files {
		wgScanners.Add(1)
		s := ggrep.New(file, pattern, false)
		go func() {
			defer wgScanners.Done()
			s.Scan(c)
		}()
	}

	wgPrinter := &sync.WaitGroup{}
	wgPrinter.Add(1)
	var err error
	go func() {
		defer wgPrinter.Done()
		for l := range c {
			fmt.Println(l.WithFileAndNum())
			if l.Error != nil {
				err = multierror.Append(err, l.Error)
			}
		}
	}()

	wgScanners.Wait()
	close(c)
	wgPrinter.Wait()

	return err
}
