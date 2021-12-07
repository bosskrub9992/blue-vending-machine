package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

type Balance struct {
	Coin1  int
	Coin5  int
	Coin10 int
}

type Product struct {
	Name  string
	Stock int
	Price int
}

// init balance and product in vending machine
var (
	initBalance  = Balance{Coin1: 5, Coin5: 5, Coin10: 5}
	initProducts = map[string]Product{
		"tewli": {Name: "tewli", Stock: 0, Price: 5},
		"pocky": {Name: "pocky", Stock: 2, Price: 5},
		"oreo":  {Name: "oreo", Stock: 3, Price: 5},
	}
)

func main() {
	balance := initBalance
	products := initProducts

	for {
		inputCoin1 := 0
		inputCoin5 := 0
		inputCoin10 := 0

		productNameSorted := []string{}
		for k := range products {
			productNameSorted = append(productNameSorted, k)
		}
		sort.Strings(productNameSorted)

		fmt.Println("\tprice \tstock")
		for _, name := range productNameSorted {
			product := products[name]
			if product.Stock == 0 {
				fmt.Printf("%s\t %d THB\t %s\n", name, product.Price, "Out of stock")
			} else {
				fmt.Printf("%s\t %d THB\t %d\n", name, product.Price, product.Stock)
			}
		}

		fmt.Println("\ncoin for change")
		fmt.Printf("coin1  remain: %d\n", balance.Coin1)
		fmt.Printf("coin5  remain: %d\n", balance.Coin5)
		fmt.Printf("coin10 remain: %d\n\n", balance.Coin10)

		var (
			product Product
			found   bool
		)
		for {
			fmt.Println(">> Please select product: ")
			productRaw := ""
			fmt.Scanln(&productRaw)
			product, found = products[strings.TrimSpace(productRaw)]
			if found {
				if product.Stock == 0 {
					fmt.Printf("%s is out of stock\n", product.Name)
					continue
				}
				break
			} else {
				fmt.Printf("not found product. Please select one of %s\n", strings.Join(productNameSorted, ", "))
				continue
			}
		}

		fmt.Printf(">> Please insert %d THB (only 1,5,10 THB coin): \n", product.Price)
		var custMoney = 0
		for {
			inputWrong := false
			coin := 0
			fmt.Scanln(&coin)
			switch coin {
			case 1:
				inputCoin1++
				balance.Coin1++
				custMoney++
			case 5:
				inputCoin5++
				balance.Coin5++
				custMoney += 5
			case 10:
				inputCoin10++
				balance.Coin10++
				custMoney += 10
			default:
				inputWrong = true
			}

			fmt.Printf(">> sum money: %d THB\n", custMoney)

			if inputWrong {
				fmt.Println(">> Please insert only 1, 5, 10 THB coin!")
				continue
			}

			params := CheckEnoughCoinParam{
				coin10: balance.Coin10,
				coin5:  balance.Coin5,
				coin1:  balance.Coin1,
				money:  custMoney,
				cost:   product.Price,
			}

			remain, enoughCoin, enoughMoney := CheckEnoughCoin(params)
			if !enoughMoney {
				fmt.Printf("still need more %d THB\n", product.Price-custMoney)
				continue
			}
			if !enoughCoin {
				fmt.Println("vending machine has not enough coin to change")
				fmt.Printf("here is your money: %d\n", custMoney)
				balance.Coin10 = balance.Coin10 - inputCoin10
				balance.Coin5 = balance.Coin5 - inputCoin5
				balance.Coin1 = balance.Coin1 - inputCoin1
				break
			}

			// success

			// update balance
			balance.Coin10 = remain.coin10
			balance.Coin5 = remain.coin5
			balance.Coin1 = remain.coin1

			// update stock
			product.Stock--
			products[product.Name] = product

			fmt.Printf("here is your %s\n", product.Name)
			change := custMoney - product.Price
			if change > 0 {
				fmt.Printf("And your change is %d\n", change)
			}

			fmt.Printf("Thank you!\n\n")
			time.Sleep(5 * time.Second)
			break
		}

		allOutOfStock := CheckAllProductOutOfStock(products)
		if allOutOfStock {
			fmt.Println("all product are out of stock")
			break
		}
	}
}

type CheckEnoughCoinParam struct {
	coin10, coin5, coin1, money, cost int
}

type CheckEnoughCoinResult struct {
	coin1  int
	coin5  int
	coin10 int
}

func CheckEnoughCoin(params CheckEnoughCoinParam) (*CheckEnoughCoinResult, bool, bool) {

	if params.cost > params.money {
		return nil, false, false
	}

	if (params.coin10 == 0) && (params.coin5 == 0) && (params.coin1 == 0) {
		return nil, false, true
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
	if coin1Need >= params.coin1 {
		return nil, false, true
	}
	coin1Remain = coin1Remain - coin1Need

	return &CheckEnoughCoinResult{
		coin10: coin10Remain,
		coin5:  coin5Remain,
		coin1:  coin1Remain,
	}, true, true
}

func CheckAllProductOutOfStock(products map[string]Product) bool {
	allOut := true
	for _, product := range products {
		if product.Stock > 0 {
			allOut = false
			break
		}
	}
	return allOut
}
