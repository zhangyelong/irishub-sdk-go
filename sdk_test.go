package main

import (
	"fmt"
	"irishub-sdk-go/app"
	"testing"
)

func TestGetAccount(t *testing.T) {
	sdk := app.NewSdk().WithHttpTimeout(10).WithLcdAddr("http://localhost:1317")
	acc, err := sdk.Bank.GetAccount("faa1nl2dxgelxu9ektxypyul8cdjp0x3ksfqcgxhg7")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(acc.GetCoins())
}