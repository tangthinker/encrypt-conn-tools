# encrypt-conn-tools (中文文档)

简单易用的 Go 语言加密库，用于常见的加密操作。它提供了 ECDH 密钥交换、ECDSA 签名/验签和 AES-256-GCM 加密的高级封装。

[English Documentation](README.md)

## 特性

- **ECDH (椭圆曲线 Diffie-Hellman)**: 使用 P-256 曲线安全地生成共享密钥。
- **ECDSA (椭圆曲线数字签名算法)**: 使用 P-256 曲线密钥对数据进行签名和验证。
- **AES-256-GCM**: 使用 256 位密钥进行认证加密和解密。
- **十六进制编码**: 所有密钥、签名和加密数据都作为十六进制编码的字符串处理，便于存储和传输。

## 安装

```bash
go get github.com/tangthinker/encrypt-conn-tools
```

## 构建共享库 (C-Shared Library)

为了构建 C 共享库（供 C、Python、Node.js 等调用）并应用代码混淆，我们推荐使用 [garble](https://github.com/burrowers/garble)。

```shell
# 安装混淆工具
go install mvdan.cc/garble@latest

# 使用 garble 构建
# -tiny: 优化二进制体积
# -literals: 混淆字符串字面量
# -seed=random: 随机化构建种子
garble -tiny -literals -seed=random build -buildmode=c-shared -o libencrypt.dylib main.go
```

### 交叉编译

支持使用标准的 `GOOS` 和 `GOARCH` 环境变量进行交叉编译。您可以使用 `Makefile` 快速构建所有目标。

**注意**：由于使用了 `c-shared` 模式，您需要安装并指定目标平台的 C 交叉编译器（CC）并启用 CGO。

**使用 Makefile:**

项目提供了一个 `Makefile` 来简化构建流程。

```bash
# 构建当前平台的库（或所有已配置编译器的平台）
make

# 构建特定平台
# 请确保环境变量中 CC 指向了正确的交叉编译器，或者直接在命令中指定
CC=x86_64-linux-gnu-gcc make linux-amd64
CC=x86_64-w64-mingw32-gcc make windows-amd64
```

**手动构建示例:**

```shell
# 示例：在 macOS 上编译 Linux (x86_64) 版本
# 需安装 x86_64-linux-musl-gcc 或类似工具
GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-linux-musl-gcc \
garble -tiny -literals -seed=random build -buildmode=c-shared -o libencrypt.so main.go

# 示例：编译 Windows (x86_64) 版本
# 需安装 x86_64-w64-mingw32-gcc
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
garble -tiny -literals -seed=random build -buildmode=c-shared -o libencrypt.dll main.go
```

### 常见平台编译参数参考

| 目标平台 (Target Platform) | 架构 (Arch) | GOOS | GOARCH | C 编译器示例 (CC) | 输出文件 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Linux (x86_64)** | AMD64 | `linux` | `amd64` | `x86_64-linux-musl-gcc` | `libencrypt.so` |
| **Linux (ARM64)** | ARM64 | `linux` | `arm64` | `aarch64-linux-musl-gcc` | `libencrypt.so` |
| **Windows (x86_64)** | AMD64 | `windows` | `amd64` | `x86_64-w64-mingw32-gcc` | `libencrypt.dll` |
| **Windows (ARM64)** | ARM64 | `windows` | `arm64` | `aarch64-w64-mingw32-gcc` | `libencrypt.dll` |
| **macOS (Intel)** | AMD64 | `darwin` | `amd64` | `o64-clang` / `clang` | `libencrypt.dylib` |
| **macOS (Apple Silicon)** | ARM64 | `darwin` | `arm64` | `oa64-clang` / `clang` | `libencrypt.dylib` |

*注：交叉编译时 `CC` 需要设置为对应的交叉编译器路径或命令。*


## 使用方法

在您的 Go 项目中导入该包：

```go
import "github.com/tangthinker/encrypt-conn-tools/pkg"
```

### ECDH 密钥交换

生成密钥对并计算共享密钥。

```go
// 1. 为 Alice 和 Bob 生成密钥对
alicePub, alicePriv := pkg.GenerateKeyPairECDH()
bobPub, bobPriv := pkg.GenerateKeyPairECDH()

// 2. 计算共享密钥 (Alice 使用她的私钥和 Bob 的公钥)
sharedKeyAlice, err := pkg.GenerateSharedKey(alicePriv, bobPub)
if err != nil {
    panic(err)
}

// 3. 计算共享密钥 (Bob 使用他的私钥和 Alice 的公钥)
sharedKeyBob, err := pkg.GenerateSharedKey(bobPriv, alicePub)
if err != nil {
    panic(err)
}

// 两个共享密钥应该完全相同
if sharedKeyAlice == sharedKeyBob {
    println("共享密钥建立成功！")
}
```

### ECDSA 签名与验证

使用私钥签名数据，并使用公钥验证签名。

```go
// 1. 生成 ECDSA 密钥对
pubKey, privKey := pkg.GenerateKeyPairECDSA()

data := "重要的交易数据"

// 2. 对数据签名
signature := pkg.SignECDSA(data, privKey)

// 3. 验证签名
valid := pkg.VerifyECDSA(data, pubKey, signature)
if valid {
    println("签名有效！")
}
```

### AES-256 加密

使用 32 字节的十六进制编码密钥加密和解密数据。

```go
// 派生一个 32 字节的密钥 (例如，从 ECDH 共享密钥派生)
// DeriveKey 函数通过哈希输入因子来创建一个合适的密钥
key := pkg.DeriveKey("some-shared-secret", "salt") 

plaintext := "你好，世界！"

// 1. 加密
ciphertext := pkg.Encrypt(plaintext, key)
println("加密后:", ciphertext)

// 2. 解密
decrypted := pkg.Decrypt(ciphertext, key)
println("解密后:", decrypted)
```

## 许可证

MIT

