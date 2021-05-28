package auth

import (
	"fmt"
	config "homefill/backend/configs"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

func GetUserInfo(state string, code string) ([]byte, error) {
	if state != config.State {
		errStr := "Invalid Auth State"
		return nil, errors.New(errStr)
	}

	token, err := config.GOOGLEAuthConfig.Exchange(oauth2.NoContext, code)

	if err != nil {
		errStr := fmt.Sprintf("Code Exchange Failed: %s", err.Error())
		return nil, errors.New(errStr)
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		errStr := fmt.Sprintf("Failed Getting User Info: %s", err.Error())
		return nil, errors.New(errStr)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errStr := fmt.Sprintf("Failed Reading Response Body: %s", err.Error())
		return nil, errors.New(errStr)
	}

	return contents, nil
}
