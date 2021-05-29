package auth

import (
	"fmt"
	config "homefill/backend/configs"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

func GetUserInfo(state string, code string) ([]byte, error) {
	if state != config.State {
		return nil, fmt.Errorf("invalid auth state")
	}

	token, err := config.GOOGLEAuthConfig.Exchange(oauth2.NoContext, code)

	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}