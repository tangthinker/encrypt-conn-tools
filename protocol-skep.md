# 安全密钥交换协议 Secure Key Exchange Protocol SKEP

SKEP在客户端、服务端内置ECDSA签名密钥

流程：

1. 客户端开始，发送客户端公钥
```text
client:start:<client_pub>:<client_nonce>
```
2. 服务端，发送服务端公钥匙 + 签名，sign = sign_func(server_pub, client_pub, client_nonce)
```text
server:start:<server_pub>:<sign>
```
3. 客户端验证后响应ok
```text
client:ok
or
client:err
```
4. 服务端收到client ok后响应ok，正式建立连接
```text
server:ok
or
server:err
```

服务端、客户端对无法解密的数据进行丢弃

协议中：
- 使用ECDSA防止中间人攻击
- 使用客户端随机数nonce防止重放攻击