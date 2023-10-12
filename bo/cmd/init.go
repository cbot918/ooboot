/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bo/cmd/pkg"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// initCmd
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init oboot project",
	Long:  `init oboot project long`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := initProject(); err != nil {
			log.Fatal(err)
		}

	},
}

func initProject() error {
	// init config.bo
	_, err := os.Create("config.bo")
	if err != nil {
		return err
	}

	// init Makefile
	fd, err := os.Create("Makefile")
	if err != nil {
		return err
	}
	_, err = fd.Write([]byte(pkg.Makefile()))
	if err != nil {
		return err
	}
	defer fd.Close()

	// go mod init
	if err = goModInit(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func goModInit() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`[^/]+$`)
	dirName := re.FindString(path)

	str := fmt.Sprintf("go mod init %s", dirName)
	if err := exec.Command("/bin/sh", "-c", str).Run(); err != nil {
		return err
	}
	return nil
}

// fiberThreeCmd

// fiberAPICmd

func init() {
	rootCmd.AddCommand(initCmd)
}
