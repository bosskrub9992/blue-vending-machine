package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bosskrub9992/blue-vending-machine/db"
	"github.com/bosskrub9992/blue-vending-machine/db/model"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate-db",
	Short: "Start the API server",
	RunE: func(cmd *cobra.Command, args []string) error {

		DB, err := db.GetDB()
		if err != nil {
			return err
		}

		migrator := DB.Migrator()

		forceMigrate, _ := cmd.Flags().GetBool("force-migrate")
		if forceMigrate {
			for _, table := range model.Migrations {
				if migrator.HasTable(table) {
					err = migrator.DropTable(table)
					if err != nil {
						fmt.Printf("DropTable error: %+v\n", err)
						return err
					}
				}
			}
		}

		for _, mod := range model.Migrations {
			if migrator.HasTable(mod) {
				err := migrator.AutoMigrate()
				if err != nil {
					return err
				}
			} else {
				err := migrator.CreateTable(mod)
				if err != nil {
					return err
				}
				fmt.Println("add table", mod.TableName())
			}
		}

		if forceMigrate {
			fmt.Println("force migrate-db success")
		} else {
			fmt.Println("migrate-db success")
		}

		// ping
		err = DB.Select("1").Error
		if err != nil {
			return err
		}
		fmt.Println("ping success")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().Bool("force-migrate", false, "Drop all db, then create new ones")
}
