package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
  Number int `json:"-"`
  Balance int `json:"balance"`
}

func main() {
  account := Account{Number: 1, Balance: 100}
  res, err := json.Marshal(account)
  if err != nil {
    panic(err)
  }

  fmt.Println(string(res))

  err = json.NewEncoder(os.Stdout).Encode(account)
  if err != nil {
    panic(err)
  }

  jsonString := []byte(`{"n":1,"b":100}`)
  var accountX Account

  err = json.Unmarshal(jsonString, &accountX)
  if err != nil {
    panic(err)
  }

  fmt.Println(accountX.Balance)
}
