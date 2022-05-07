package helper

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Base64Conv(text string) string {
	encodedText := base64.StdEncoding.EncodeToString([]byte(text))
	return encodedText
}

var Client = &http.Client{}

func Authentication(baseURL, clientId, secretKey string) error {
	basic := fmt.Sprintf("%s:%s", clientId, secretKey)
	basicToken := Base64Conv(basic)
	url := fmt.Sprintf("%s/passport/oauth/token", baseURL)
	method := "POST"
	// payload := strings.NewReader(`{
	// 	"grant_type": ""
	// }`)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basicToken))
	res, err := Client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(string(body))
	return nil
}

// curl https://sandbox.interswitchng.com/passport/oauth/token \
// -H "Authorization: Basic Base64(CLIENT_ID:SECRET_KEY)" \
// -H "Content-Type: application/x-www-form-urlencoded" \
// -D '{
//     "grant_type": "client_credentials"
// }'
// -X POST
