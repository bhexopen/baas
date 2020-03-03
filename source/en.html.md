---
title: BLUEHELIX BAAS API documentation

language_tabs: # must be one of https://git.io/vQNgJ
  - shell
  - python
  - javascript
  - golang

toc_footers:
- <a href="en.html">English</a>
- <a href="index.html">中文</a>

search: true
---

# About BLUEHELIX BAAS

## Overview
BLUEHELIX BAAS provides a REST API (HTTPS + JSON) to facilitate Bluehelix Cloud customers' for self-service access to third-party public chains.

Before requesting for an API socket, users need to apply for an APIKEY by using ED25519 algorithm to generate the Public & Private Keys. Users are required to save the Public & Private Keys. Submit the application for coin listing and attach the Public & Private Keys accordingly. The APIKEY will then be issued to the applicant.

## How to apply?
- Submit a Ticket Request
- Email to globalbd@bhex.com

## Information Required For Coin Listing

Parameter|Value
-----------|------------
Digital Asset Symbol | ABC
Explorer| ABC
Decimals| 8
Total Supply| 10 billion
IP address | 100.100.100.100  (used as IP whitelist restriction)

## Client Sample
Developer can code using any one of the following programming languages (Python, Golang, Java) 
https://github.com/bhexopen/baas.

# API Signature Authentication
  The data needs to be signed as the following: HTTP_METHOD + | + HTTP_REQUEST_PATH + | + TIMESTAMP + | + PARAMS

  The API signature should sign data with ED25519 signature after connection and sign the bytes with hex encoding.


## Domain name
- Prod env：https://baas.bluehelix.com

## HTTPMethods
GET POST

## TIMESTAMP
When accessing API, the UNIX EPOCH timestamp(ms), expires in 120000ms.

## Complete Sample

### POST Request

METHOD    | URL | TIMESTAMP
-----------|-----------------------|----------------------
POST |    https://baas.bluehelix.com/api/v1/test                   | 1580887996488

The request data shown on the right:：

```
{
  "side": 1,
  "amount": "100.0543",
  "token_id": "ABC",
  "tx_hash":"0x1234567890",
  "block_height": 1000000
}
```

Before signing in, you need to sort the request parameters according to the first letter of the key to obtain the following data `POST|/api/v1/test/|1580887996488|amount=100.0543&block_height=1000000&side=1&token_id=ABC&tx_hash=0x1234567890`

Use your locally generated private_key to sign off the data with ED25519. Then Hex-encode the binary result to obtain the final signature.


In the HTTP request, enter the header to bypass the verification:

- BWAAS-API-KEY
- BWAAS-API-SIGNATURE
- BWAAS-API-TIMESTAMP

### GET Request

METHOD    | URL | TIMESTAMP
-----------|-----------------------|----------------------
GET |    https://baas.bluehelix.com/api/v1/test?chain=ABC                   | 1580887996488

Before signing in, you need to sort the request parameters according to the first letter of the key to obtain the following data `GET|/api/v1/test?chain=ABC|1580887996488`

Use your locally generated private_key to sign off the data with ED25519. Then Hex-encode the binary result to obtain the final signature.


In the HTTP request, enter the header to bypass the verification:

- BWAAS-API-KEY
- BWAAS-API-SIGNATURE
- BWAAS-API-TIMESTAMP


# API List

## Count the number of unused address

> Count the number of unused address

```shell
curl
  -X GET
  -H "BWAAS-API-KEY: 123"
  -H "BWAAS-API-TIMESTAMP: 1580887996488"
  -H "BWAAS-API-SIGNATURE: f321da3"
  https://baas.bluehelix.com/api/v1/address/unused/count/
?chain=ABC
```

```golang
```

```python
```

```javascript
```

> Reponse：

```json
{
    "code": 10000,
    "msg": "success",
    "data": 1000
}
```

HTTP Request：

`GET /api/v1/address/unused/count?chain=ABC`


Request parameters：

Parameter | Type| Mandatory| Description
-----------|-----------|-----------|-----------
chain | string| yes|The chain,use mainnet token

Response：

Parameter | Type| Description
-----------|-----------|-----------
code | int| Please see the retruen code list
msg | string | Returned content; error message if failed
data | int | Number of unused address

<aside class="notice">
It is recommended that you generate the count of unused address/addresses regularly, preferably on an hourly basis. Any increase in the number of new users will change the final count.
</aside>

## Add deposit address

> Add deposit address

```shell
curl
  -X POST
  -H "BWAAS-API-KEY: 123"
  -H "BWAAS-API-TIMESTAMP: 1580887996488"
  -H "BWAAS-API-SIGNATURE: f321da3"
  -- data '
    {
      "chain":"ABC",
      "addr_list":[
        "addr_111",
        "addr_222"
      ]
    }
  '
  https://baas.bluehelix.com/api/v1/address/add
```

```golang
```

```python
```

```javascript
```

> Reponse：

```json
{
    "code": 10000,
    "msg": "success",
}
```

HTTP Request：

`POST /api/v1/address/add`


Request parameters：

Parameter | Type| Mandatory| Description
-----------|-----------|-----------|-----------
chain | string| yes|The chain,use mainne token
addr_list | []string | yes|list of address

Response:

Parameter | Type| Description
-----------|-----------|-----------
code | int| Please see the retruen code list
msg | string | Returned content; error message if failed


<aside class="notice">
Check that the balance number of addresses is not less than the specified value. If the value is less than 1,000, please generate a new batch of addresses by calling the add deposit address. It is recommended to add no more than 100 deposit addresses per call. Should you require a larger number of addresses, please perform several add deposit requests accordingly.

For Blockchain using Tag, client needs only to request for one (1) deposit address. Server will assign a unique tag to each address.
</aside>


## Deposit Notify

> Deposit Notify

```shell
curl
  -X POST
  -H "BWAAS-API-KEY: 123"
  -H "BWAAS-API-TIMESTAMP: 1580887996488"
  -H "BWAAS-API-SIGNATURE: f321da3"
  -- data '
    {
        "token_id": "ABC",
        "from": "addr1",
        "to": "addr2",
        "memo":"1234",
        "amount": "124.23",
        "tx_hash": "1234",
        "index": 1,
        "block_height": 124,
        "block_time": 1234
    }
  '
  https://baas.bluehelix.com/api/v1/notify/deposit
```

```golang
```

```python
```

```javascript
```

> Reponse：

```json
{
    "code": 10000,
    "msg": "success",
}
```

HTTP Request：

`POST /api/v1/notify/deposit`

Request parameters:

Parameter | Type| Mandatory| Description
-----------|-----------|-----------|-----------
token_id| string| yes|digital asset symbol
from | string | yes|from which address
to | string | yes|to which address
memo| string| optional| memo
amount| string| yes| deposit amount
tx_hash| string|yes |transaction hash
index| int | yes| he location of the deposit made
block_height| int64| yes| block height
block_time| int64| yes| block time (seconds)

Response：

Parameter | Type| Description
-----------|-----------|-----------
code | int| Please see the retruen code list
msg | string | Returned content; error message if failed

<aside class="notice">
A call deposit notice will help to ensure the user receives notification after a deposit was made. Each call caters to one (1) deposit action. 

The client is responsible to provide a proof of authenticity of all deposits made. Users will incur losses caused by any incorrect notification sent by the client.
</aside>



## Generate pending withdrawal orders

> Generate pending withdrawal orders

```shell
curl
  -X GET
  -H "BWAAS-API-KEY: 123"
  -H "BWAAS-API-TIMESTAMP: 1580887996488"
  -H "BWAAS-API-SIGNATURE: f321da3"
  https://baas.bluehelix.com/api/v1/withdrawal/orders?chain=ABC
```

```golang
```

```python
```

```javascript
```

> Reponse：

```json
{
    "code": 10000,
    "msg": "success",
    "data"::[
        {
            "order_id": "1234",
            "token_id":"ABC",
            "to": "bhexaddr1",
            "memo": "bhexmemo",
            "amount": "12.34"
        },
        {
            "order_id": "2345",
            "token_id":"ABC",
            "to": "bhexaddr1",
            "memo": "bhexmemo",
            "amount": "12.34"
        }
    ]
}
```

HTTP Request：

`GET /api/v1/withdrawal/orders?chain=ABC`

Request parameters:

Parameter | Type| Mandatory| Description
-----------|-----------|-----------|-----------
chain | string| yes|The chain,use mainnet token


Response：

Parameter | Type| Description
-----------|-----------|-----------
code | int| Please see the retruen code list
msg | string | Returned content; error message if failed
data | []order | List of pending withdrawal orders

orderinfo：

Parameter | Type| Description
-----------|-----------|-----------
order_id| string | order id
token_id| string| withdrawal currency
to | string | withdrawal address
memo | string | memo
amount | string | withdrawal amount


<aside class="notice">
A request for user's withdrawal is required. It is recommended that the number of requests should not be more than one-tenth of the interval for each block generated. 

If the time taken for a block to generate is 15s, it is recommended that you put in your request every 1s. 

If the time taken for a block to generate is 15mins, it is recommended that you put in your request every 15s.

To reconcile the assets, the net withdrawal amount is the withdrawal amount returned. The processing fee for on-chain transaction is managed by the client. 

Each request generates up to 50 outstanding orders at a time.
</aside>

## Successful withdrawal notify

> Successful withdrawal notify

```shell
curl
  -X POST
  -H "BWAAS-API-KEY: 123"
  -H "BWAAS-API-TIMESTAMP: 1580887996488"
  -H "BWAAS-API-SIGNATURE: f321da3"
  --data '
  {
    "order_id": "1234",
    "token_id": "ABC",
    "to": "bhexaddr1",
    "memo": "bhexmemo",
    "amount": "12.34",
    "tx_hash": "0x5f99810a4154379e5b7951419a77250f020be54b78acb9a8747ff8b0ec75769d",
    "block_height": 6581548,
    "block_time": 1540480255
  }
  '
  https://baas.bluehelix.com/api/v1/notify/withdrawal
```

```golang
```

```python
```

```javascript
```

> Reponse：

```json
{
    "code": 10000,
    "msg": "success",
}
```

HTTP Request：

`POST /api/v1/notify/withdrawal`

Request parameters:

Parameter | Type| Mandatory| Description
-----------|-----------|-----------|-----------
order_id| string | yes|order id
token_id| string| yes| withdrawal currency
to | string | yes|to wihich address
memo | string | yes|memo
amount | string | yes|withdrawal amount
tx_hash| string | yes|transaction hash
block_height|int |yes|block height
block_time|int |yes|block time (seconds)

Response：

Parameter | Type| Description
-----------|-----------|-----------
code | int| Please see the retruen code list
msg | string | Returned content; error message if failed

<aside class="notice">
*After a successful withdrawal (blockchain shows the transactions on the distributed public ledger), calling for a Successful Withdrawal Notice is required. To ensure the reliability of the withdrawal result, only one withdrawal result is given at a time. It is the responsibility of the client to ensure that a request is called only after the withdrawal is executed.
</aside>


## Asset verification

> Asset verification

```shell
curl
  -X POST
  -H "BWAAS-API-KEY: 123"
  -H "BWAAS-API-TIMESTAMP: 1580887996488"
  -H "BWAAS-API-SIGNATURE: f321da3"
  --data '
  {
    "token_id": "ABC",
    "total_deposit_amount": "100000.567",
    "total_withdrawal_amount": "10000",
    "total_fee_amount": "100",
    "last_block_height": "100000"
  }
  '
  https://baas.bluehelix.com/api/v1/asset/verify
```

```golang
```

```python
```

```javascript
```

> Reponse：

```json
{
    "code": 10000,
    "msg": "success",
}
```

HTTP Request：

`POST /api/v1/asset/verify`

Request parameters:

Parameter | Type| Mandatory| Description
-----------|-----------|-----------|-----------
token_id| string| yes|withdrawal currency
total_deposit_amount | string | yes|total deposit amount
total_withdrawal_amount | string | yes|total withdrawal amount
total_fee_amount | string | yes|total on-chain transaction fee
last_block_height|string |yes|verification of highest block height

Response：

Parameter | Type| Description
-----------|-----------|-----------
code | int| Please see the retruen code list
msg | string | Returned content; error message if failed



<aside class="notice">
Asset verification is performed on a regular basis, and the customer service terminal regularly (Reconciliation is performed every hour or every day, and the client determines the transaction amount reasonably based on the block generation time of the chain) reports the deposit and withdrawal of assets on a specific chain to the server (till the block height specified in asset_info). If there are outstanding withdrawal orders, it is recommended to perform reconciliation after processing is completed.

When the server finds that the asset information returned by the client is inconsistent with the server, it will return an error and suspend the deposit and withdrawal of the currency.
</aside>

# Return code list

Code |Type| Description
-----------|-----------|-----------
10000 | SUCCESS| successful
10001 | INVALID_SIGN| invalid signature
10002 | INVALID_APIKEY | invalid api_key
10004 | INVALID_CHAIN | invalid chain
10005 | INVALID_TOKEN_ID | invalid token_id
10006 | INVALID_PARAMS | invalid paramter
10007 | INVALID_TO_ADDRESS | invalid deposit address
10008 | INVALID_ORDER_ID | invalid order id
10009 | INVALID_AMOUNT | invalid amount
10010 | INVALID_DECIMALS | invalid decimals
10011 | INVALID_BLOCK_HEIGHT | invalid block height
10012 | INVALID_BLOCK_TIME | invalid block time
10013 | INVALID_TXHASH | invalid tx hash
10014 | INVALID_INDEX | invalid tx index
10015 | NETWORK_ERROR | network error
10016 | REPEAT_DEPOSIT | repeat deposit
10017 | ASSET_VERIFY_FAILED| asset verification failed
10018 | DEPOSIT_SUSPENDED| deposit suspended
10019 | WITHDRAWAL_SUSPENDED| withdrawal suspended
10020 | TIMESTAMP_EXPIRED| timestamp expired