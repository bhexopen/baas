import requests
import json
import time
from binascii import hexlify, unhexlify
import ed25519

domain     = "https://baas.bluehelix.com"
private_key = "ce19ff3824c46d589c7ccad54028f1e010645c27732bcb369e7b19b4962863d36510e490e5fbf93d839b374e3139fe5eed036c5b9c58d56ca8993a68153adb69"
chain      = "BAAS-TEST"
api_key     = "0bea9a7c38d944a2a0c8af4058665153"

def get_address_count():
    timestamp = str(int(time.time() * 1000))
    path = "/api/v1/address/unused/count?chain="+chain

    sign_msg = create_sign_msg("GET", path, timestamp, {})

    signing_key = ed25519.SigningKey(private_key, "", "hex")
    signature = signing_key.sign(sign_msg)
    print("signature = ", hexlify(signature))

    headers  = {
        "BWAAS-API-KEY": api_key,
        "BWAAS-API-TIMESTAMP": timestamp,
        "BWAAS-API-SIGNATURE": hexlify(signature)
    }

    res = requests.get(
        url=domain+path,
        headers=headers)

    print(res.text)

def add_address():
    timestamp = str(int(time.time() * 1000))

    data = {
        "chain": chain,
        "addr_list": [
            "BAAS-TEST-address-123456",
            "BAAS-TEST-address-654321"
        ]
    }
    sign_msg = create_sign_msg(
        "POST", "/api/v1/address/add", timestamp, data)

    signing_key = ed25519.SigningKey(private_key, "", "hex")
    signature = signing_key.sign(sign_msg)
    print("signature = ", hexlify(signature))

    headers = {
        "BWAAS-API-KEY": api_key,
        "BWAAS-API-TIMESTAMP": timestamp,
        "BWAAS-API-SIGNATURE": hexlify(signature),
        "Content-Type": "application/json"
    }

    res = requests.post(
        url=domain+"/api/v1/address/add",
        data=json.dumps(data), 
        headers=headers)

    print(res.text)

def create_key():
    signing_key, verifying_key = ed25519.create_keypair()
    print("the private key is", signing_key.to_ascii(encoding="hex"))
    print("the public key is", verifying_key.to_ascii(encoding="hex"))


def create_sign_msg(method,url, timestamp, body):
    params_list = [method, url, timestamp]
   
    if method == "POST":
        sorted_body = sorted(body.items(),  key=lambda d: d[0], reverse=False)
        print("sorted_body= ", sorted_body)

        data_list = []
        for data in sorted_body:
            if isinstance(data[1],list):
                value = "["+" ".join(data[1])+"]"
                key = data[0]
                data_list.append(key+"="+value)
            else:
                data_list.append("=".join(data))

        body_params = "&".join(data_list)
        params_list.append(body_params)

    params_str = "|".join(params_list)
    print("params_str= ", params_str)
    return params_str
        

create_key()
get_address_count()
add_address()
