package main

import (
	"fmt"

	"github.com/BOPR/db"
	"github.com/BOPR/types"
)

func main() {
	// tx := types.NewTx(1, 1, 1, 1, "esds")
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	acc := types.NewAccountLeaf(12, 12, 12, 12)
	db.InsertAccount(acc)

	accs, _ := db.GetAllAccounts()
	fmt.Println("acs", accs)
}
