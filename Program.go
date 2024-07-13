package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func BypassV3(getURL string, postURL string, bg string) (string, error) {
	client := &http.Client{}
	resp, err := client.Get(getURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	recaptchaTokenRe := regexp.MustCompile(`id="recaptcha-token" value="(.*?)"`)
	recaptchaTokenMatch := recaptchaTokenRe.FindStringSubmatch(string(body))
	if len(recaptchaTokenMatch) < 2 {
		return "", fmt.Errorf("recaptcha token not found")
	}
	recaptchaToken := recaptchaTokenMatch[1]
	vRe := regexp.MustCompile(`v=(.*?)&`)
	vMatch := vRe.FindStringSubmatch(getURL)
	if len(vMatch) < 2 {
		return "", fmt.Errorf("v parameter not found")
	}
	v := vMatch[1]
	kRe := regexp.MustCompile(`&k=(.*?)&`)
	kMatch := kRe.FindStringSubmatch(getURL)
	if len(kMatch) < 2 {
		return "", fmt.Errorf("k parameter not found")
	}
	k := kMatch[1]
	coRe := regexp.MustCompile(`&co=(.*?)&`)
	coMatch := coRe.FindStringSubmatch(getURL)
	if len(coMatch) < 2 {
		return "", fmt.Errorf("co parameter not found")
	}
	co := coMatch[1]
	data := url.Values{}
	data.Set("v", v)
	data.Set("reason", "q")
	data.Set("c", recaptchaToken)
	data.Set("k", k)
	data.Set("co", co)
	data.Set("hl", "en")
	data.Set("size", "invisible")
	data.Set("chr", "%5B89%2C64%2C27%5D")
	data.Set("vh", "13599012192")
	data.Set("bg", bg)
	req, err := http.NewRequest("POST", postURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	rrespRe := regexp.MustCompile(`"rresp","(.*?)"`)
	rrespMatch := rrespRe.FindStringSubmatch(string(body))
	if len(rrespMatch) < 2 {
		return "", fmt.Errorf("rresp not found")
	}
	return rrespMatch[1], nil
}

func postRequest(solution string) (string, error) {
	url := "https://antcpt.com/score_detector/verify.php"
	payload := map[string]string{
		"g-recaptcha-response": solution,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshaling payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://antcpt.com")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Referer", "https://antcpt.com/score_detector/")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Gpc", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 status code received: %v", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	resultBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling response: %v", err)
	}

	return string(resultBytes), nil
}
