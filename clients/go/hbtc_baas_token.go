/*
 * *******************************************************************
 * @项目名称: go
 * @文件名称: hbtc_baas_token.go
 * @Date: 2020/10/11
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

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

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/ed25519"
)

const (
	domain     = "https://baas.bluehelix.com"
	privateKey = "6180ead1719869c41da5479f3ac0c3dc86a140ffc92f19c8e5734532288d2b104e66c61578550e2c30f12e9d7946173f23066e3f76896aafb2b4df4bdb3952da"
	pubkey     = "4e66c61578550e2c30f12e9d7946173f23066e3f76896aafb2b4df4bdb3952da"
	chain      = "HBTC-AGGR-CHAIN"
	apiKey     = "40221a1881ae4c49b6662986ae58f667"
)

func main() {
	// testCreateKey()
	// testGetAddressCount()
	// testAddAddress()
	deposit("TEST1")
}

func deposit(tokenID string) {
	timestamp := fmt.Sprintf("%v", time.Now().UnixNano()/1e6)

	txHash, _ := uuid.NewV4()
	postData := map[string]interface{}{
		"token_id":     tokenID,
		"from":         "HBTC-AGGR-CHAIN-POOL",
		"to":           "HBTC-AGGR-CHAIN-ADDRESS",
		"amount":       "100000000",
		"tx_hash":      txHash.String(),
		"index":        "0",
		"block_height": timestamp,
		"block_time":   timestamp,
	}

	path := "/api/v1/notify/deposit"
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
