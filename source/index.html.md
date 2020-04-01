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

# 关于 BLUEHELIX BAAS

## 概述
BLUEHELIX BAAS 提供REST风格的API（HTTPS + JSON)，方便BHOP客户自助接入第三方公链。

在请求API接口之前，需要申请APIKEY, 使用ED25519算法生成公私钥对，用户自己保存私钥，公钥在上币申请时进行提交，得到APIKEY。

<aside class="notice">
客户端：全节点运营方
服务端：BAAS
</aside>

## 申请方式
- 工单系统
- 邮箱 globalbd@bhex.com

## 上币申请材料

参数|值
-----------|------------
币种ID | ABC
所属公链| ABC
代币精度| 8
代币总量| 100亿
IP地址 | 100.100.100.100 （用作IP白名单限制）

## 客户端代码示例
提供3种编程语言（Python, Golang, Java)的用户端代码供用户使用 [https://github.com/bhexopen/baas/clients] (https://github.com/bhexopen/baas)。

# API签名认证
  签名前准备的数据如下： HTTP_METHOD + | + HTTP_REQUEST_PATH + | + TIMESTAMP + | + PARAMS 连接完成后，对数据进行 ED25519 签名，签名后的 bytes 进行 Hex 编码。

## 域名
- 正式环境：https://baas.bluehelix.com

## HTTP方法
GET POST

## TIMESTAMP
访问 API 时的 UNIX EPOCH 时间戳 (精确到毫秒), 过期时间120000ms。

## 完成示例

### POST请求

METHOD    | URL | TIMESTAMP
-----------|-----------------------|----------------------
POST |    https://baas.bluehelix.com/api/v1/test                   | 1580887996488

参数见右：

```
{
  "side": 1,
  "amount": "100.0543",
  "token_id": "ABC",
  "tx_hash":"0x1234567890",
  "block_height": 1000000
}
```
在进行签名之前，需要对请求参数，按照key的首字母进行排序，得到如下数据： `POST|/api/v1/test|1580887996488|amount=100.0543&block_height=1000000&side=1&token_id=ABC&tx_hash=0x1234567890`

使用您本地生成的 private_key（私钥），对数据进行ED25519签名，并对签名后的bytes进行 Hex 编码, 得到最终签名signature。

在HTTP请求时，写入header，即可通过校验:

- BWAAS-API-KEY
- BWAAS-API-SIGNATURE
- BWAAS-API-TIMESTAMP

### GET请求

METHOD    | URL | TIMESTAMP
-----------|-----------------------|----------------------
GET |    https://baas.bluehelix.com/api/v1/test?chain=ABC                   | 1580887996488

在进行签名之前，需要对请求参数，按照key的首字母进行排序，得到如下数据： `GET|/api/v1/test?chain=ABC|1580887996488`

使用您本地生成的 private_key（私钥），对数据进行ED25519签名，并对签名后的bytes进行 Hex 编码, 得到最终签名signature。

在HTTP请求时，写入header，即可通过校验:

- BWAAS-API-KEY
- BWAAS-API-SIGNATURE
- BWAAS-API-TIMESTAMP
  

# 接口列表

## 获取剩余地址数量

> 获取剩余地址数量

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

> 返回结果：

```json
{
    "code": 10000,
    "msg": "success",
    "data": 1000
}
```

HTTP Request：

`GET /api/v1/address/unused/count`


请求参数：

参数 | 类型| 必须| 说明
-----------|-----------|-----------|-----------
chain | string| 是|那个链,使用主网代币

响应结果：

参数 | 类型| 说明
-----------|-----------|-----------
code | int| 详情见返回类型表
msg | string | 返回内容；失败时为错误信息
data | int | 剩余地址数量

<aside class="notice">
检查是否需要重新生成地址时使用，需要定期调用，视新增用户速度决定。建议每小时调用一次。
</aside>

## 添加充值地址

> 添加充值地址

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

> 返回结果：

```json
{
    "code": 10000,
    "msg": "success",
}
```

HTTP Request：

`POST /api/v1/address/add`


请求参数：

参数 | 类型| 必须| 说明
-----------|-----------|-----------|-----------
chain | string| 是|那个链,使用主网代币
addr_list | []string | 是|地址列表

响应结果：

参数 | 类型| 说明
-----------|-----------|-----------
code | int| 详情见返回类型表
msg | string | 返回内容；失败时为错误信息


<aside class="notice">
当检查到剩余地址小于特定值，比如1000时，重新生成一批地址，并调用此接口添加充值地址，建议每次添加的充值地址不超过100个，
如果需要导入大量地址，可以多次调用此接口。

对于使用memo的公链，客户端可只调用此服务添加一个充值地址即可，由服务端给来分配memo供充值使用。
</aside>


## 充值到账通知

> 充值到账通知

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
        "index": "1",
        "block_height": "124",
        "block_time": "1234"
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

> 返回结果：

```json
{
    "code": 10000,
    "msg": "success",
}
```

HTTP Request：

`POST /api/v1/notify/deposit`

请求参数：

参数 | 类型| 必须| 说明
-----------|-----------|-----------|-----------
token_id| string| 是| 币种ID
from | string | 是|从哪个地址转出来
to | string | 是|转给那个地址
memo| string| 可选| memo标识
amount| string| 是| 充值金额
tx_hash| string|是 |交易hash
index| string | 是| 该充值所在交易中的位置
block_height| string| 是| 区块高度
block_time| string| 是| 区块时间（秒）


响应结果：

参数 | 类型| 说明
-----------|-----------|-----------
code | int| 详情见返回类型表
msg | string | 返回内容；失败时为错误信息

<aside class="notice">
当有用户充值时，调用此接口，为保证充值可靠性，充值需要逐笔执行，超过1条的话可以多次调用此接口。

客户端必须保证充币的真实可靠，因客户端通知错误充值带来的损失，由客户端执行者承担。
</aside>



## 获取待处理提现请求

> 获取待处理提现请求

```shell
curl
  -X GET
  -H "BWAAS-API-KEY: 123"
  -H "BWAAS-API-TIMESTAMP: 1580887996488"
  -H "BWAAS-API-SIGNATURE: f321da3"
  https://baas.bluehelix.com/api/v1/withdrawal/orders
```

```golang
```

```python
```

```javascript
```

> 返回结果：

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

请求参数：

参数 | 类型| 必须| 说明
-----------|-----------|-----------|-----------
chain | string| 是|那个链,使用主网代币

响应结果：

参数 | 类型| 说明
-----------|-----------|-----------
code | int| 详情见返回类型表
msg | string | 返回内容；失败时为错误信息
data | []order | 待处理提现订单列表

order 信息：

参数 | 类型| 说明
-----------|-----------|-----------
order_id| string | 订单id
token_id| string| 提现币种
to | string | 提现给那个地址
memo | string | memo标记
amount | string | 提现金额


<aside class="notice">
应定期轮询是否有用户提现需求 ，轮询周期建议不大于出块间隔的十分之一。

如出块时间为15s，建议每1s轮询一次，出块时间为15min的话，建议15s轮询一次。

为对账方便，服务端返回提币金额为净提币金额，链处理所需的手续费由客户端管理。

该接口每次最多返回50个未处理订单。
</aside>

## 提现处理完成通知

> 提现处理完成通知

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

> 返回结果：

```json
{
    "code": 10000,
    "msg": "success",
}
```

HTTP Request：

`POST /api/v1/notify/withdrawal`

请求参数：

参数 | 类型| 必须| 说明
-----------|-----------|-----------|-----------
order_id| string | 是|订单id
token_id| string| 是|提现币种
to | string | 是|提现给那个地址
memo | string | 是|memo标记
amount | string | 是|提现金额
tx_hash| string | 是|交易hash
block_height|int |是|区块高度
block_time|int |是|区块时间（秒）

响应结果：

参数 | 类型| 说明
-----------|-----------|-----------
code | int| 详情见返回类型表
msg | string | 返回内容；失败时为错误信息



<aside class="notice">
提币成功（提币交易已被区块链打包并确认执行成功）后调用此接口，为保证提币结果反馈可靠性，每次仅通知一笔提现处理结果。
客服端负责保证提币执行后才调用此接口
</aside>

## 提现处理失败通知

> 提现处理失败通知

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
    "reason": "invalid address"
  }
  '
  https://baas.bluehelix.com/api/v1/notify/failed
```

```golang
```

```python
```

```javascript
```

> 返回结果：

```json
{
    "code": 10000,
    "msg": "success",
}
```

HTTP Request：

`POST /api/v1/notify/failed`

请求参数：

参数 | 类型| 必须| 说明
-----------|-----------|-----------|-----------
order_id| string | 是|订单id
token_id| string| 是|提现币种
reason| string | 是|失败原因


响应结果：

参数 | 类型| 说明
-----------|-----------|-----------
code | int| 详情见返回类型表
msg | string | 返回内容；失败时为错误信息



<aside class="notice">
提币订单无法处理或者处理失败时，调用此接口，该订单会变为无效订单，订单资产会退回给用户。
</aside>



## 定期对账

> 定期对账

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

> 返回结果：

```json
{
    "code": 10000,
    "msg": "success",
}
```

HTTP Request：

`POST /api/v1/asset/verify`

请求参数：

参数 | 类型| 必须| 说明
-----------|-----------|-----------|-----------
token_id| string| 是|提现币种
total_deposit_amount | string | 是|总充值金额
total_withdrawal_amount | string | 是|总提现金额
last_block_height|string |是|对账最高区块高度

响应结果：

参数 | 类型| 说明
-----------|-----------|-----------
code | int| 详情见返回类型表
msg | string | 返回内容；失败时为错误信息



<aside class="notice">
定期进行资产对账，客服端定期（每小时或者每天进行对账，客户端根据链的出块时间，交易数量合理确定）向服务端反馈特定链上资产的充提情况（截止到asset_info中指定的区块高度）。

如果有未处理完成的提现订单，建议处理完成后进行对账。当服务端发现客户端反馈的资产信息与服务端不一致，将返回错误，并暂停该币种的充值提现。
</aside>

# 返回值列表

返回值 | 类型| 说明
-----------|-----------|-----------
10000 | SUCCESS| 成功
10001 | INVALID_SIGN| 无效签名
10002 | INVALID_APIKEY | 无效的api_key
10004 | INVALID_CHAIN | 无效的chain
10005 | INVALID_TOKEN_ID | 无效的token_id
10006 | INVALID_PARAMS | 无效的参数
10007 | INVALID_TO_ADDRESS | 无效的充币地址
10008 | INVALID_ORDER_ID | 无效的订单id
10009 | INVALID_AMOUNT | 无效的amount值
10010 | INVALID_DECIMALS | 无效的精度
10011 | INVALID_BLOCK_HEIGHT | 无效的区块高度
10012 | INVALID_BLOCK_TIME | 无效的区块时间
10013 | INVALID_TXHASH | 无效的tx_hash
10014 | INVALID_INDEX | 无效的交易index
10015 | NETWORK_ERROR | 网络错误
10016 | REPEAT_DEPOSIT | 重复充值
10017 | ASSET_VERIFY_FAILED| 资产校验失败
10018 | DEPOSIT_SUSPENDED| 充值暂停
10019 | WITHDRAWAL_SUSPENDED| 提现暂停
10020 | TIMESTAMP_EXPIRED| 时间戳过期
10021 | MEMO_REQUIRED| 需要memo
10022 | NEED_WAIT | 通知提现失败时，需要等待队列释放