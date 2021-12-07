package cmd

import (
	"fmt"

	"github.com/bosskrub9992/blue-vending-machine/db"
	"github.com/bosskrub9992/blue-vending-machine/db/model"
	"github.com/spf13/cobra"
)

var initDbCmd = &cobra.Command{
	Use:   "init-db",
	Short: "Start the API server",
	RunE: func(cmd *cobra.Command, args []string) error {

		DB, err := db.GetDB()
		if err != nil {
			return err
		}

		initProducts := []model.Product{
			{
				Name:  "Kitkat Choc",
				Stock: 5,
				Price: 1,
			},
			{
				Name:  "Tewli Twin",
				Stock: 5,
				Price: 5,
			},
			{
				Name:  "Pocky Choc",
				Stock: 5,
				Price: 20,
			},
			{
				Name:  "Pocky Straw",
				Stock: 5,
				Price: 8,
			},
		}

		initCoins := []model.Balance{
			{
				Coin1:  20,
				Coin5:  15,
				Coin10: 10,
			},
		}

		err = DB.Create(&initProducts).Error
		if err != nil {
			return err
		}

		err = DB.Create(&initCoins).Error
		if err != nil {
			return err
		}

		fmt.Println("init-db success")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initDbCmd)
}
