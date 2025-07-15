package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func generateKey() []byte {
    key := make([]byte, 32) // AES-256
    _, err := rand.Read(key)
    if err != nil {
        panic(err)
    }
    ioutil.WriteFile("secret.key", []byte(base64.StdEncoding.EncodeToString(key)), 0644)
    return key
}

func loadKey() []byte {
    keyData, err := ioutil.ReadFile("secret.key")
    if err != nil {
        panic(err)
    }
    key, _ := base64.StdEncoding.DecodeString(string(keyData))
    return key
}

func encryptFile(filename string) {
    key := loadKey()
    plaintext, _ := ioutil.ReadFile(filename)

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

    ioutil.WriteFile(filename+".enc", ciphertext, 0644)
}

func decryptFile(filename string) {
    key := loadKey()
    ciphertext, _ := ioutil.ReadFile(filename)

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)

    ioutil.WriteFile(filename+".dec", ciphertext, 0644)
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run encryptor.go [generate|encrypt|decrypt] filename")
        return
    }

    action := os.Args[1]
    file := os.Args[2]

    switch action {
    case "generate":
        generateKey()
    case "encrypt":
        encryptFile(file)
    case "decrypt":
        decryptFile(file)
    default:
        fmt.Println("Unknown action")
    }
}
