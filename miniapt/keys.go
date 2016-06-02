package miniapt

import (
	"errors"
	"os"
	"path/filepath"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

// SaveKey saves an OpenPGP EntityList.
func SaveKey(keysPath string, entityList openpgp.EntityList) error {
	if len(entityList) == 0 {
		return errors.New("no keys provided to save")
	}
	keyID := entityList[0].PrimaryKey.KeyIdString()
	f, err := os.Create(filepath.Join(keysPath, keyID+".asc"))
	if err != nil {
		return err
	}
	defer f.Close()
	for _, e := range entityList {
		encoder, err := armor.Encode(f, openpgp.PublicKeyType, nil)
		if err != nil {
			return err
		}
		err = e.Serialize(encoder)
		if err != nil {
			encoder.Close()
			return err
		}
		encoder.Close()
	}
	return nil
}
