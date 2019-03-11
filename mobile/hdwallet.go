package geth

import (
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

//CreateRandomMnemonic will return mnemonic created randomly
func CreateRandomMnemonic() (string, error) {

	entropy, _ := bip39.NewEntropy(128)
	mnemonic, err := bip39.NewMnemonic(entropy)

	return mnemonic, err
}

//GetPrivateKeyFromMnemonic will convert mnemonic to privatekey
func GetPrivateKeyFromMnemonic(mnemonic string) ([]byte, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("Account address: %s\n", account.Address.Hex())

	privateKey, err := wallet.PrivateKeyBytes(account)
	if err != nil {
		return nil, err
	}

	return privateKey, nil

}
