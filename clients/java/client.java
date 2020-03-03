package baas;

import com.alibaba.fastjson.JSONObject;
import okhttp3.*;
import okio.ByteString;
import org.bouncycastle.crypto.AsymmetricCipherKeyPair;
import org.bouncycastle.crypto.Signer;
import org.bouncycastle.crypto.generators.Ed25519KeyPairGenerator;
import org.bouncycastle.crypto.params.Ed25519KeyGenerationParameters;
import org.bouncycastle.crypto.params.Ed25519PrivateKeyParameters;
import org.bouncycastle.crypto.params.Ed25519PublicKeyParameters;
import org.bouncycastle.crypto.signers.Ed25519Signer;

import java.security.SecureRandom;
import java.util.ArrayList;
import java.util.List;
import java.util.TreeMap;

public class BaasDemo {

    private static String domain = "https://baas.bluehelix.com";
    private static String privateKey = "ce19ff3824c46d589c7ccad54028f1e010645c27732bcb369e7b19b4962863d36510e490e5fbf93d839b374e3139fe5eed036c5b9c58d56ca8993a68153adb69";
    private static String testChain = "BAAS-TEST";
    private static String apiKey = "0bea9a7c38d944a2a0c8af4058665153";


    private static OkHttpClient HTTP_CLIENT = new OkHttpClient();

    private static byte[] hex2bytes(String s) {
        return ByteString.decodeHex(s).toByteArray();
    }

    private static String bytes2Hex(byte[] b) {
        return ByteString.of(b).hex();
    }


    private static String generateSignature(String content, String key) throws Exception {
        byte[] contentByte = content.getBytes();

        Ed25519PrivateKeyParameters privateKeyRebuild = new Ed25519PrivateKeyParameters(hex2bytes(key), 0);
        System.out.println(bytes2Hex(privateKeyRebuild.generatePublicKey().getEncoded()));

        // create the signature
        Signer signer = new Ed25519Signer();
        signer.init(true, privateKeyRebuild);
        signer.update(contentByte, 0, contentByte.length);
        byte[] signature = signer.generateSignature();

        // verify the signature
       /* Signer verifier = new Ed25519Signer();
        verifier.init(false, privateKeyRebuild.generatePublicKey());
        verifier.update(contentByte, 0, contentByte.length);
        boolean shouldVerify = verifier.verifySignature(signature);*/
        return bytes2Hex(signature);

    }

    private static String composeParams(TreeMap<String, Object> params) {
        StringBuffer sb = new StringBuffer();
        params.forEach((s, param) -> {
            if (param instanceof List) {
                String values = "[";
                for (String address : (List<String>) param
                ) {
                    values = values + address + " ";
                }
                String encodeAddressed = values.substring(0, values.length() - 1) + "]";
                sb.append(s).append('=').append(encodeAddressed).append("&");
            } else {
                sb.append(s).append("=").append(param).append("&");
            }
        });
        if (sb.length() > 0) {
            sb.deleteCharAt(sb.length() - 1);
        }
        return sb.toString();
    }

    private static String request(String method, String path, TreeMap<String, Object> params, String apiKey, String apiSecret, String host) throws Exception {
        method = method.toUpperCase();
        String timestamp = String.valueOf(System.currentTimeMillis());
        String content = method + "|" + path + "|" + timestamp;
        if (params != null && params.size() > 0) {
            String paramString = composeParams(params);
            content = content + "|" + paramString;
        }
        System.out.println(content);

        String signature = generateSignature(content, apiSecret);
        System.out.println(signature);

        Request.Builder builder = new Request.Builder()
                .addHeader("BWAAS-API-KEY", apiKey)
                .addHeader("BWAAS-API-TIMESTAMP", timestamp)
                .addHeader("BWAAS-API-SIGNATURE", signature);

        Request request;
        if ("GET".equalsIgnoreCase(method)) {
            request = builder
                    .url(host + path )
                    .build();
        } else if ("POST".equalsIgnoreCase(method)) {

            JSONObject paramJson = new JSONObject(params);

            MediaType JSON = MediaType.parse("application/json;charset=utf-8");

            String jasonData = paramJson.toJSONString();
            RequestBody requestBody = RequestBody.create(JSON, jasonData);

            request = builder
                    .url(host + path)
                    .post(requestBody)
                    .build();
        } else {
            throw new RuntimeException("not supported http method");
        }
        try (Response response = HTTP_CLIENT.newCall(request).execute()) {
            String body = response.body().string();
            System.out.println(body);
            return body;
        }
    }


    public static void main(String... args) throws Exception {
        testUnusedCountApi();
        testAddressAddApi();
        //testCreateKey();
    }

    public static void testUnusedCountApi() throws Exception {
        String key = apiKey;
        String secret = privateKey;
        String host = domain;
        String url = "/api/v1/address/unused/count?chain=" + testChain;
        String res = request("GET", url, null, key, secret, host);
        System.out.println(res);
    }

    public static void testAddressAddApi() throws Exception {
        TreeMap<String, Object> params = new TreeMap<>();
        List<String> addressList = new ArrayList<>();
        addressList.add("testaddress11111");
        addressList.add("testaddress22222");
        params.put("chain", testChain);
        params.put("addr_list", addressList);
        String key = apiKey;
        String secret = privateKey;
        String host = domain;
        String res = request("POST", "/api/v1/address/add", params, key, secret, host);
        System.out.println(res);
    }

    public static void testCreateKey() {
        SecureRandom randomData = new SecureRandom();
        Ed25519KeyPairGenerator keyPairGenerator = new Ed25519KeyPairGenerator();
        keyPairGenerator.init(new Ed25519KeyGenerationParameters(randomData));
        AsymmetricCipherKeyPair asymmetricCipherKeyPair = keyPairGenerator.generateKeyPair();
        Ed25519PrivateKeyParameters privateKey = (Ed25519PrivateKeyParameters) asymmetricCipherKeyPair.getPrivate();
        Ed25519PublicKeyParameters publicKey = (Ed25519PublicKeyParameters) asymmetricCipherKeyPair.getPublic();

        System.out.println(bytes2Hex(privateKey.getEncoded()));
        System.out.println(bytes2Hex(publicKey.getEncoded()));
        System.out.println(bytes2Hex(privateKey.generatePublicKey().getEncoded()));
    }

}
