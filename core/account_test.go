package core

import (
	"fmt"
	"testing"

	"github.com/BOPR/common"
)

func TestPubkeyHashCreation(t *testing.T) {
	bz, err := ABIEncodePubkey("90718dcbc2477c86294742fb72bf098ba85ff671b88c8d79b2e09ce19bdbd88fd87047aaebc775b168372752aa8bc4e5be1ca5d39284fed00722f341927888c3")
	if err != nil {
		panic(err)
	}

	hash := common.Keccak256(bz)

	fmt.Println(hash.String())
}
