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

	params = CheckEnoughCoinParam{
		coin10: 10,
		coin5:  10,
		coin1:  10,
		money:  12,
		cost:   5,
	}
	remain, enoughCoin, enoughMoney := CheckEnoughCoin(params)
	if !enoughCoin {
		t.Errorf(`expected enoughCoin to be 'true' but got %+v`, enoughCoin)
	}
	if !enoughMoney {
		t.Errorf(`expected enoughMoney to be 'true' but got %+v`, enoughMoney)
	}
	if remain.coin10 != 10 {
		t.Errorf(`expected remain.coin10 to be 10 but got %+v`, remain.coin10)
	}
	if remain.coin5 != 9 {
		t.Errorf(`expected remain.coin10 to be 9 but got %+v`, remain.coin10)
	}
	if remain.coin1 != 8 {
		t.Errorf(`expected remain.coin10 to be 8 but got %+v`, remain.coin10)
	}

	params = CheckEnoughCoinParam{
		coin10: 0,
		coin5:  5,
		coin1:  5,
		money:  20,
		cost:   5,
	}
	remain, enoughCoin, enoughMoney = CheckEnoughCoin(params)
	if !enoughCoin {
		t.Errorf(`expected enoughCoin to be 'true' but got %+v`, enoughCoin)
	}
	if !enoughMoney {
		t.Errorf(`expected enoughMoney to be 'true' but got %+v`, enoughMoney)
	}
	if remain.coin10 != 0 {
		t.Errorf(`expected remain.coin10 to be 0 but got %+v`, remain.coin10)
	}
	if remain.coin5 != 2 {
		t.Errorf(`expected remain.coin10 to be 2 but got %+v`, remain.coin10)
	}
	if remain.coin1 != 5 {
		t.Errorf(`expected remain.coin10 to be 5 but got %+v`, remain.coin10)
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

	products = map[string]Product{
		"tewli": {Name: "tewli", Stock: 1, Price: 5},
		"pocky": {Name: "pocky", Stock: 1, Price: 5},
		"oreo":  {Name: "oreo", Stock: 1, Price: 5},
	}
	allOut = CheckAllProductOutOfStock(products)
	if allOut {
		t.Errorf(`expected allOut to be 'false' but got %+v`, allOut)
	}
}
