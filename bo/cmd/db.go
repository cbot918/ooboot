package cmd

import (
	"bo/cmd/pkg/db"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "db command",
	Long:  "db command long",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("#\navalable commands: createdb / init / mig / sqlx\n#")
	},
}

var createdbCmd = &cobra.Command{
	Use:   "createdb",
	Short: "create database container",
	Long:  "create database container long",
	Run: func(cmd *cobra.Command, args []string) {
		if err := createdb(); err != nil {
			log.Fatal(err)
		}
	},
}

func createdb() error {
	c := " docker run --name psqltemp -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=psqltemp -d postgres"
	if err := exec.Command("/bin/sh", "-c", c).Run(); err != nil {
		return err
	}
	return nil
}

var migCmd = &cobra.Command{
	Use:   "mig",
	Short: "db migration",
	Long:  "db migration long",
	Run: func(cmd *cobra.Command, args []string) {
		if err := mig(); err != nil {
			log.Fatal(err)
		}
	},
}

func mig() error {

	fmt.Println("run migration!!")

	// dbSource := "DB_SOURCE=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	// migrationUrl := "MIGRATION_URL=file://db/migration"

	// if err := db.RunDBMigration(migrationUrl, dbSource); err != nil {
	// 	return err
	// }
	return nil
}

var sqlxCmd = &cobra.Command{
	Use:   "sqlx",
	Short: "sqlx command",
	Long:  "sqlx command long",
	Run: func(cmd *cobra.Command, args []string) {
		if err := sqlx(); err != nil {
			log.Fatal(err)
		}
	},
}

func sqlx() error {
	s := db.NewSqlx()

	fd, err := os.Create("sqlx.go")
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = fd.Write([]byte(s.PkgAndImport()))
	if err != nil {
		return err
	}

	_, err = fd.Write([]byte(s.Schema()))
	if err != nil {
		return err
	}
	_, err = fd.Write([]byte(s.Types()))
	if err != nil {
		return err
	}
	_, err = fd.Write([]byte(s.Connect()))
	if err != nil {
		return err
	}
	_, err = fd.Write([]byte(s.Init()))
	if err != nil {
		return err
	}
	_, err = fd.Write([]byte(s.CRUD()))
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(dbCmd)

	dbCmd.AddCommand(createdbCmd)
	dbCmd.AddCommand(migCmd)
	dbCmd.AddCommand(sqlxCmd)
}
