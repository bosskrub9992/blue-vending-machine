package cmd

import (
	"errors"
	"fmt"
	"math"

	"github.com/bosskrub9992/blue-vending-machine/db"
	"github.com/bosskrub9992/blue-vending-machine/db/model"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start vending machine",
	RunE: func(cmd *cobra.Command, args []string) error {

		DB, err := db.GetDB()
		if err != nil {
			return err
		}

		for {
			products := []model.Product{}
			err = DB.Find(&products).Error
			if err != nil {
				return err
			}

			balance := model.Balance{}
			err := DB.First(&balance).Error
			if err != nil {
				return err
			}

			fmt.Println("\t\t price\t stock")
			for _, product := range products {
				fmt.Printf("%s\t %d THB\t %d left\n", product.Name, product.Price, product.Stock)
			}
			fmt.Printf("\ncoin  1 remain: %d\n", balance.Coin1)
			fmt.Printf("coin  5 remain: %d\n", balance.Coin5)
			fmt.Printf("coin 10 remain: %d\n", balance.Coin10)

			var custMoney = 0
			for {
				prompt := promptui.Select{
					Label: "Insert Coin",
					Items: []string{"1 THB", "5 THB", "10 THB", "Select product"},
				}
				_, coin, err := prompt.Run()
				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return err
				}

				updateBalance := model.Balance{}
				err = DB.First(&updateBalance).Error
				if err != nil {
					return err
				}

				if coin != "Select product" {
					switch coin {
					case "1 THB":
						custMoney++
						updateBalance.Coin1++
					case "5 THB":
						custMoney += 5
						updateBalance.Coin5++
					case "10 THB":
						custMoney += 10
						updateBalance.Coin10++
					}
					fmt.Printf("You insert %q\n", coin)

				} else {
					if custMoney == 0 {
						fmt.Println("Please insert 1, 5, 10 THB coin")
					} else {
						break
					}
				}

				fmt.Println("sum money: ", custMoney)

				err = DB.Where("balance_id = ?", updateBalance.BalanceId).Save(&updateBalance).Error
				if err != nil {
					return err
				}
			}

			products = []model.Product{}
			err = DB.Where("price <= ? AND stock != 0 ORDER BY price DESC", custMoney).Find(&products).Error
			if err != nil {
				return err
			}

			balance = model.Balance{}
			err = DB.First(&balance).Error
			if err != nil {
				return err
			}

			productNames := []string{}
			for _, product := range products {
				productNames = append(productNames, product.Name)
			}

			if len(productNames) == 0 {
				fmt.Println("not enough stock")
				break
			}

			prompt := promptui.Select{
				Label: "Select Product",
				Items: productNames,
			}
			_, productName, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return err
			}
			product := model.Product{}
			err = DB.Where("name = ?", productName).First(&product).Error
			if err != nil {
				return err
			}

			params := CheckEnoughCoinParam{
				coin10: balance.Coin10,
				coin5:  balance.Coin5,
				coin1:  balance.Coin1,
				money:  custMoney,
				cost:   int(product.Price),
			}
			remainCoin, enough, err := CheckEnoughCoin(params)
			if err != nil {
				return err
			}
			if !enough {
				fmt.Println("not enough coin for change")
				fmt.Println("here is your money: ", custMoney)
				break
			} else {
				updateBalance := model.Balance{}
				err = DB.First(&updateBalance).Error
				if err != nil {
					return err
				}

				updateBalance.Coin1 = remainCoin.coin1
				updateBalance.Coin5 = remainCoin.coin5
				updateBalance.Coin10 = remainCoin.coin10

				err = DB.Where("balance_id = ?", updateBalance.BalanceId).Save(&updateBalance).Error
				if err != nil {
					return err
				}
			}

			fmt.Printf("Here is your %q and your change: %d THB\n\n", productName, custMoney-int(product.Price))

			// update stock
			if product.Stock > 0 {
				product.Stock = product.Stock - 1
				err = DB.Where("product_id = ?", product.ProductID).Save(&product).Error
				if err != nil {
					return err
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

type CheckEnoughCoinParam struct {
	coin10, coin5, coin1, money, cost int
}

type CheckEnoughCoinResult struct {
	coin1  int
	coin5  int
	coin10 int
}

func CheckEnoughCoin(params CheckEnoughCoinParam) (*CheckEnoughCoinResult, bool, error) {

	if params.money < params.cost {
		return nil, false, errors.New("not enough money")
	}

	var (
		coin10Remain = params.coin10
		coin5Remain  = params.coin5
		coin1Remain  = params.coin1

		change = params.money - params.cost
	)

	coin10Need := int(math.Floor(float64(change) / 10.0))
	if coin10Need > params.coin10 {
		coin10Need = params.coin10
	}
	coin10Remain = coin10Remain - coin10Need

	change = change - coin10Need*10

	coin5Need := int(math.Floor(float64(change) / 5.0))
	if coin5Need > params.coin5 {
		coin5Need = params.coin5
	}
	coin5Remain = coin5Remain - coin5Need

	change = change - coin5Need*5

	coin1Need := int(math.Floor(float64(change) / 1.0))
	if coin1Need > params.coin1 {
		return nil, false, nil
	}
	coin1Remain = coin1Remain - coin1Need

	return &CheckEnoughCoinResult{
		coin10: coin10Remain,
		coin5:  coin5Remain,
		coin1:  coin1Remain,
	}, true, nil
}
