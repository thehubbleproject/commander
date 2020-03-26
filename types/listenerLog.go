package types

import (
	"math/big"
)

type ListenerLog struct {
	DBModel
	LastRecordedBlock string `json:"lastRecordedBlock"`
}

func (l *ListenerLog) BigInt() *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(l.LastRecordedBlock, 10)
	return n
}
