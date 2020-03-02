package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"golang.org/x/crypto/ed25519"
)

const (
	domain     = "https://baas.bluehelix.com"
	privateKey = "ce19ff3824c46d589c7ccad54028f1e010645c27732bcb369e7b19b4962863d36510e490e5fbf93d839b374e3139fe5eed036c5b9c58d56ca8993a68153adb69"
	chain      = "BAAS-TEST"
	apiKey     = "0bea9a7c38d944a2a0c8af4058665153"
)

func main() {
	testCreateKey()
	testGetAddressCount()
	testAddAddress()
}

func testGetAddressCount() {
	timestamp := fmt.Sprintf("%v", time.Now().UnixNano()/1e6)

	path := "/api/v1/address/unused/count?chain=" + chain
	msg := createSignMsg(path, "GET", timestamp, nil)
	fmt.Printf(" msg:%v\n", msg)

	priBytes, _ := hex.DecodeString(privateKey)
	pri := ed25519.PrivateKey(priBytes)

	signMsg := ed25519.Sign(pri, []byte(msg))
	signature := hex.EncodeToString(signMsg)

	fmt.Printf("signature:%v\n", signature)

	client := &http.Client{}
	request, err := http.NewRequest("GET", domain+path, nil)

	request.Header.Add("BWAAS-API-KEY", apiKey)
	request.Header.Add("BWAAS-API-SIGNATURE", signature)
	request.Header.Add("BWAAS-API-TIMESTAMP", timestamp)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("http get err:%v\n", err)
		return
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp code:%v, body:%v\n", resp.StatusCode, string(result))
}

func testAddAddress() {
	timestamp := fmt.Sprintf("%v", time.Now().UnixNano()/1e6)

	postData := map[string]interface{}{
		"chain": chain,
		"addr_list": []string{
			"BAAS-TEST-address-12345",
			"BAAS-TEST-address-54321",
		},
	}

	path := "/api/v1/address/add"
	msg := createSignMsg(path, "POST", timestamp, postData)
	fmt.Printf(" msg:%v\n", msg)

	priBytes, _ := hex.DecodeString(privateKey)
	pri := ed25519.PrivateKey(priBytes)

	signMsg := ed25519.Sign(pri, []byte(msg))
	signature := hex.EncodeToString(signMsg)

	fmt.Printf("signature:%v\n", signature)

	postBody, _ := json.Marshal(postData)

	client := &http.Client{}
	request, err := http.NewRequest("POST", domain+path, bytes.NewReader(postBody))

	request.Header.Add("BWAAS-API-KEY", apiKey)
	request.Header.Add("BWAAS-API-SIGNATURE", signature)
	request.Header.Add("BWAAS-API-TIMESTAMP", timestamp)
	request.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("http get err:%v\n", err)
		return
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp code:%v, body:%v\n", resp.StatusCode, string(result))
}

func testCreateKey() (string, string) {
	pub, pri, _ := ed25519.GenerateKey(rand.Reader)
	pubStr, priStr := hex.EncodeToString(pub), hex.EncodeToString(pri)
	fmt.Printf("pub:%v\n, pri:%v\n", pubStr, priStr)

	return priStr, pubStr
}

func createSignMsg(url, method, timestamp string, mapBody map[string]interface{}) string {
	result := strings.Join([]string{method, url, timestamp}, "|")

	if mapBody != nil {
		var keys []string

		for key := range mapBody {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		bodyParam := "|"
		for index, key := range keys {
			if index == len(keys)-1 {
				bodyParam = bodyParam + fmt.Sprintf("%s=%v", key, mapBody[key])
			} else {
				bodyParam = bodyParam + fmt.Sprintf("%s=%v&", key, mapBody[key])
			}
		}

		result = result + bodyParam
		fmt.Printf("makeParams result:%v\n", result)
		return result
	}

	fmt.Printf("makeParams result:%v\n", result)
	return result
}
