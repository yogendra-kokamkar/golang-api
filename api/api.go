package api

import(
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct{
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct{
	// success code, usually 200
	code int
	//Account Balance
	Balance int64
}

// Error Response
type Error struct{
	// Error Code
	Code int
	// Error Message
	Message string
}