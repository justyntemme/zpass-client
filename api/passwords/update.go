package passwords

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/stormentt/zpass-client/api"
	"github.com/stormentt/zpass-client/keyvault"
	"net/http"
)

// Update changes the password refered to by selector to newPassword
func Update(selector string, newPassword string) error {
	log.Info("Updating password")
	encrypted, _ := keyvault.PassCrypter.Encrypt([]byte(newPassword))
	req := api.NewRequest()
	req.
		Dest("passwords/"+selector, "PATCH").
		Set("password", encrypted)

	response, err := req.Send()
	if err != nil {
		log.Error(err)
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Unable to update, code: %v", response.StatusCode)
	}

	return nil
}
