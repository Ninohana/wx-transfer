package weixin

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetAccessToken(appid string, secret string) (string, int, error) {
	resp, _ := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	type ResponseBody struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	var responseBody ResponseBody
	json.Unmarshal(body, &responseBody)

	log.Printf("获取到的accessToken将于%ds后过期\n", responseBody.ExpiresIn)

	return responseBody.AccessToken, responseBody.ExpiresIn, nil
}

func Post(url string, header http.Header, data []byte) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body, nil
}

func InvokeCloudFunc(accessToken string, envId string, funcName string, data []byte) ([]byte, error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	url := "https://api.weixin.qq.com/tcb/invokecloudfunction?access_token=" + accessToken + "&env=" + envId + "&name=" + funcName
	println(string(data))
	body, err := Post(url, headers, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}
