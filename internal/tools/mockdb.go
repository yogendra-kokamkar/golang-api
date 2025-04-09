package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Yogendra": {
		AuthToken: "123ABC",
		Username:  "Yogendra",
	},
	"Bittu": {
		AuthToken: "456DEF",
		Username:  "Bittu",
	},
	"Kokamkar": {
		AuthToken: "789GHI",
		Username:  "Kokamkar",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Yogendra": {
		Coins:    100,
		Username: "Yogendra",
	},
	"Bittu": {
		Coins:    200,
		Username: "Bittu",
	},
	"Kokamkar": {
		Coins:    300,
		Username: "Kokamkar",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
