package main

import "testing"

func TestCheckEnoughCoin(t *testing.T) {

	params := CheckEnoughCoinParam{
		coin10: 0,
		coin5:  0,
		coin1:  0,
		money:  0,
		cost:   30,
	}

	_, enoughCoin, enoughMoney := CheckEnoughCoin(params)
	if enoughCoin {
		t.Errorf(`expected enoughCoin to be 'false' but got %+v`, enoughCoin)
	}
	if enoughMoney {
		t.Errorf(`expected enoughMoney to be 'false' but got %+v`, enoughMoney)
	}

	params = CheckEnoughCoinParam{
		coin10: 1,
		coin5:  0,
		coin1:  0,
		money:  10,
		cost:   5,
	}

	_, enoughCoin, enoughMoney = CheckEnoughCoin(params)
	if enoughCoin {
		t.Errorf(`expected enoughCoin to be 'false' but got %+v`, enoughCoin)
	}
	if !enoughMoney {
		t.Errorf(`expected enoughMoney to be 'true' but got %+v`, enoughMoney)
	}
}

func TestCheckAllProductOutOfStock(t *testing.T) {
	products := map[string]Product{
		"tewli": {Name: "tewli", Stock: 0, Price: 5},
		"pocky": {Name: "pocky", Stock: 0, Price: 5},
		"oreo":  {Name: "oreo", Stock: 0, Price: 5},
	}

	allOut := CheckAllProductOutOfStock(products)
	if !allOut {
		t.Errorf(`expected allOut to be 'true' but got %+v`, allOut)
	}
}
