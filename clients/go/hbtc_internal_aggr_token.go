/*
 * *******************************************************************
 * @项目名称: go
 * @文件名称: hbtc_internal_aggr_token.go
 * @Date: 2020/08/20
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
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/ed25519"
)

const (
	domain     = "https://baas.bluehelix.com"
	privateKey = "7b1b2a1e0026723a93afd9271e3520199aee4b867c1ff329c0469bfe40440020910aa5a57367fc43b437b19abb67ec01b747675eba726bcdfd4990708607f4c3"
	pubKey     = "910aa5a57367fc43b437b19abb67ec01b747675eba726bcdfd4990708607f4c3"
	chain      = "HBTC-AGGR-CHAIN"
	apiKey     = "40221a1881ae4c49b6662986ae58f667"
)

func main() {
	// getAddressCount()
	// addAddress()
	// deposit("SXP8")
	// deposit("BAND8")
	// deposit("MANA8")
	// deposit("WAVES8")
	// deposit("LEND8")
	// deposit("WNXM8")
	// deposit("ANT8")
	// deposit("MKR8")
	// deposit("KSM8")
	// deposit("OCEAN8")
	// deposit("DOCK8")
	// deposit("SUSHI8")
	// deposit("IDEX8")
	// deposit("SNX8")
	// deposit("MTA8")
	// deposit("RING8")
	// deposit("YAMV28")
	// deposit("YFV8")
	// deposit("CVP8")
	// deposit("CRV8")
	// deposit("LEND8")
	// deposit("UMA8")
	// deposit("REN8")
	// deposit("BNT8")
	// deposit("DIA8")
	// deposit("ANKR8")
	// deposit("TRADE8")
	// deposit("YFI8")
	// deposit("YFII8")
	// deposit("PNK8")
	// deposit("LRC8")
	// deposit("KNC8")
	// deposit("TRB8")
	// deposit("BZRX8")
	// deposit("RING8")
	// deposit("STORJ8")
	// deposit("MANA8")
	// deposit("PEARL8")
	// deposit("BTS8")
	// deposit("NBS8")
	// deposit("SUN8")
	// deposit("GOF8")
	// deposit("UNI8")
	// deposit("AVAX8")
}

func getAddressCount() {
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

func addAddress() {
	timestamp := fmt.Sprintf("%v", time.Now().UnixNano()/1e6)

	postData := map[string]interface{}{
		"chain": chain,
		"addr_list": []string{
			"HBTC-AGGR-CHAIN-ADDRESS",
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

func deposit(tokenID string) {
	timestamp := fmt.Sprintf("%v", time.Now().UnixNano()/1e6)
	txHash, _ := uuid.NewV4()
	blockTime := strconv.FormatInt(time.Now().Unix(), 10)

	postData := map[string]interface{}{
		"token_id":     tokenID,
		"from":         "HBTC-AGGR-CHAIN-POOL",
		"to":           "HBTC-AGGR-CHAIN-ADDRESS",
		"amount":       "100000000",
		"tx_hash":      txHash.String(),
		"index":        "0",
		"block_height": timestamp,
		"block_time":   blockTime,
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

func ceateKey() (string, string) {
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
