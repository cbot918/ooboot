package cmd

import (
	"bo/cmd/pkg"
	"log"
	"os"

	"os/exec"

	"github.com/spf13/cobra"
)

var fiberCmd = &cobra.Command{
	Use:   "fiber",
	Short: "generate fiber spa code",
	Long:  `this is bo`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := fiber(); err != nil {
			log.Fatal(err)
		}
	},
}

func fiber() error {
	fd, err := os.Create("main.go")
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = fd.Write([]byte(pkg.FiberSpa()))
	if err != nil {
		return err
	}
	return nil
}

var fiber3Cmd = &cobra.Command{
	Use:   "fiber3",
	Short: "generate fiber three layer code",
	Long:  `this is bo`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := fiber3(); err != nil {
			log.Fatal(err)
		}
	},
}

func fiber3() error {
	str := "curl -o a.tar https://getsub.fiveplanet.online/?url=https://github.com/cbot918/template/tree/main/go-three-layer-poc && tar -xf a.tar && rm a.tar && rm go-three-layer-poc/go.mod && mv go-three-layer-poc/* . && rm -rf go-three-layer-poc"
	if err := exec.Command("/bin/sh", "-c", str).Run(); err != nil {
		log.Fatal(err)
	}

	// init Makefile
	fd, err := os.Create(".env")
	if err != nil {
		return err
	}
	_, err = fd.Write([]byte(pkg.Env()))
	if err != nil {
		return err
	}
	defer fd.Close()

	return nil
}

var fiber3cCmd = &cobra.Command{
	Use:   "fiber3c",
	Short: "generate fiber three layer clean repository structure",
	Long:  `this is bo`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := fiber3c(); err != nil {
			log.Fatal(err)
		}
	},
}

func fiber3c() error {
	str := "curl -o a.tar https://getsub.fiveplanet.online/?url=https://github.com/cbot918/template/tree/main/go-three-clean && tar -xf a.tar && rm a.tar && rm go-three-clean/go.mod && mv go-three-clean/* . && rm -rf go-three-clean"
	if err := exec.Command("/bin/sh", "-c", str).Run(); err != nil {
		log.Fatal(err)
	}
	// init Makefile
	fd, err := os.Create(".env")
	if err != nil {
		return err
	}
	_, err = fd.Write([]byte(pkg.Env()))
	if err != nil {
		return err
	}
	defer fd.Close()

	return nil
}

func init() {
	rootCmd.AddCommand(fiberCmd)
	rootCmd.AddCommand(fiber3Cmd)
	rootCmd.AddCommand(fiber3cCmd)
}
