package mail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mailgo/model"
	"net/http"
	"os"
	"strings"
	"sync"
)

func UpdateTokenAcces(wg *sync.WaitGroup, acces *model.Token) {
	defer wg.Done()

	url := "https://oauth2.googleapis.com/token"
	method := "POST"
	cid := os.Getenv("CID")
	csecret := os.Getenv("CSECRET")
	rtoken := os.Getenv("RTOKEN")

	//payloadformat := fmt.Sprintf("client_id=%s&client_secret=%s&refresh_token%s=&grant_type=refresh_token", cid, csecret, rtoken)

	payload := strings.NewReader("client_id=" + cid + "&client_secret=" + csecret + "&refresh_token=" + rtoken + "&grant_type=refresh_token")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(body, acces)

}
