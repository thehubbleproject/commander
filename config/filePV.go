package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/types"
)

//-------------------------------------------------------------------------------

// FilePVKey stores the immutable part of PrivValidator.
type FilePVKey struct {
	Address types.Address  `json:"address"`
	PubKey  crypto.PubKey  `json:"pub_key"`
	PrivKey crypto.PrivKey `json:"priv_key"`

	filePath string
}

// Save persists the FilePVKey to its filePath.
func (pvKey FilePVKey) Save() {
	outFile := pvKey.filePath
	if outFile == "" {
		panic("cannot save PrivValidator key: filePath not set")
	}

	jsonBytes, err := json.MarshalIndent(pvKey, "", "  ")
	if err != nil {
		panic(err)
	}
	err = cmn.WriteFileAtomic(outFile, jsonBytes, 0600)
	if err != nil {
		panic(err)
	}

}

//-------------------------------------------------------------------------------

// FilePV implements PrivValidator using data persisted to disk
// to prevent double signing.
// NOTE: the directories containing pv.Key.filePath and pv.LastSignState.filePath must already exist.
// It includes the LastSignature and LastSignBytes so we don't lose the signature
// if the process crashes after signing but before the resulting consensus message is processed.
type FilePV struct {
	Key FilePVKey
}

// GenFilePV generates a new validator with randomly generated private key
// and sets the filePaths, but does not call Save().
func GenFilePV(keyFilePath, stateFilePath string) *FilePV {
	// privKey := ed25519.GenPrivKey()
	privKey := secp256k1.GenPrivKey()
	return &FilePV{
		Key: FilePVKey{
			Address:  privKey.PubKey().Address(),
			PubKey:   privKey.PubKey(),
			PrivKey:  privKey,
			filePath: keyFilePath,
		},
	}
}

// LoadFilePV loads a FilePV from the filePaths.  The FilePV handles double
// signing prevention by persisting data to the stateFilePath.  If either file path
// does not exist, the program will exit.
func LoadFilePV(keyFilePath, stateFilePath string) *FilePV {
	return loadFilePV(keyFilePath, stateFilePath, true)
}

// LoadFilePVEmptyState loads a FilePV from the given keyFilePath, with an empty LastSignState.
// If the keyFilePath does not exist, the program will exit.
func LoadFilePVEmptyState(keyFilePath, stateFilePath string) *FilePV {
	return loadFilePV(keyFilePath, stateFilePath, false)
}

// If loadState is true, we load from the stateFilePath. Otherwise, we use an empty LastSignState.
func loadFilePV(keyFilePath, stateFilePath string, loadState bool) *FilePV {
	keyJSONBytes, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		cmn.Exit(err.Error())
	}
	pvKey := FilePVKey{}
	err = json.Unmarshal(keyJSONBytes, &pvKey)
	if err != nil {
		cmn.Exit(fmt.Sprintf("Error reading PrivValidator key from %v: %v\n", keyFilePath, err))
	}

	// overwrite pubkey and address for convenience
	pvKey.PubKey = pvKey.PrivKey.PubKey()
	pvKey.Address = pvKey.PubKey.Address()
	pvKey.filePath = keyFilePath

	return &FilePV{
		Key: pvKey,
	}
}

// LoadOrGenFilePV loads a FilePV from the given filePaths
// or else generates a new one and saves it to the filePaths.
func LoadOrGenFilePV(keyFilePath, stateFilePath string) *FilePV {
	var pv *FilePV
	if cmn.FileExists(keyFilePath) {
		pv = LoadFilePV(keyFilePath, stateFilePath)
	} else {
		pv = GenFilePV(keyFilePath, stateFilePath)
		pv.Save()
	}
	return pv
}

// GetAddress returns the address of the validator.
// Implements PrivValidator.
func (pv *FilePV) GetAddress() types.Address {
	return pv.Key.Address
}

// GetPubKey returns the public key of the validator.
// Implements PrivValidator.
func (pv *FilePV) GetPubKey() crypto.PubKey {
	return pv.Key.PubKey
}

// Save persists the FilePV to disk.
func (pv *FilePV) Save() {
	pv.Key.Save()
}

func (pv *FilePV) signBatch(chainID string, vote *types.Vote) error {
	// height, round, step := vote.Height, vote.Round, voteToStep(vote)

	// lss := pv.LastSignState

	// sameHRS, err := lss.CheckHRS(height, round, step)
	// if err != nil {
	// 	return err
	// }

	// signBytes := vote.SignBytes(chainID)

	// // We might crash before writing to the wal,
	// // causing us to try to re-sign for the same HRS.
	// // If signbytes are the same, use the last signature.
	// // If they only differ by timestamp, use last timestamp and signature
	// // Otherwise, return error
	// if sameHRS {
	// 	if bytes.Equal(signBytes, lss.SignBytes) {
	// 		vote.Signature = lss.Signature
	// 	} else if timestamp, ok := checkVotesOnlyDifferByTimestamp(lss.SignBytes, signBytes); ok {
	// 		vote.Timestamp = timestamp
	// 		vote.Signature = lss.Signature
	// 	} else {
	// 		err = fmt.Errorf("conflicting data")
	// 	}
	// 	return err
	// }

	// // It passed the checks. Sign the vote
	// sig, err := pv.Key.PrivKey.Sign(signBytes)
	// if err != nil {
	// 	return err
	// }
	// pv.saveSigned(height, round, step, signBytes, sig)
	// vote.Signature = sig
	return nil
}
