package main

import (
	"fmt"

	"github.com/BOPR/db"
	"github.com/BOPR/types"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// tx1 := types.NewTx(1, 1, 1, 1, "esds")
	// tx2 := types.NewTx(1, 1, 1, 1, "esds")
	// tx3 := types.NewTx(1, 1, 1, 1, "esds")
	// fmt.Println(db.InsertTx(&tx1))
	// fmt.Println(db.InsertTx(&tx2))
	// fmt.Println(db.InsertTx(&tx3))

	tx4 := types.NewTx(1, 1, 1, 1, "esds")

	fmt.Println(db.InsertTx(&tx4))

	tx5 := types.NewTx(1, 1, 1, 1, "esds")
	// tx5.Status = "processing"
	fmt.Println(db.InsertTx(&tx5))

	fmt.Println("popping error", db.PopTxs())

}
