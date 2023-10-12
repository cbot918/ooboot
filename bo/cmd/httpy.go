package cmd

import (
	"bo/cmd/pkg"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var httpyCmd = &cobra.Command{
	Use:   "httpy",
	Short: "generate httpy template code",
	Long:  `this is bo`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := httpy(); err != nil {
			log.Fatal(err)
		}
	},
}

func httpy() error {
	fd, err := os.Create("main.go")
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = fd.Write([]byte(pkg.GetHttpy()))
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(httpyCmd)
}
