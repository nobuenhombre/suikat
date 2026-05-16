# secreto

Пакет `secreto` предоставляет функции шифрования и дешифрования данных с использованием AES-GCM и scrypt.

## Обзор

Пакет содержит два набора функций:
- **Стандартные** (`Encrypt`/`Decrypt`) — с автоматическим генерированием соли и производством ключа из пароля через scrypt
- **Упрощённые** (`EncryptSimple`/`DecryptSimple`) — с использованием прямого ключа

## Стандартное шифрование

### Encrypt

Шифрует данные, генерируя ключ из пароля:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/secreto"
)

password := []byte("my-secret-password")
data := []byte("Hello, secret world!")

ciphertext, err := secreto.Encrypt(password, data)
if err != nil {
    panic(err)
}
fmt.Printf("Зашифровано: %x\n", ciphertext)
```

### Decrypt

Дешифрует данные, используя тот же пароль:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/secreto"
)

password := []byte("my-secret-password")
ciphertext := []byte{...} // результат Encrypt

plaintext, err := secreto.Decrypt(password, ciphertext)
if err != nil {
    panic(err)
}
fmt.Println(string(plaintext)) // Hello, secret world!
```

### DeriveKey

Производит ключ из пароля и соли через scrypt:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/secreto"
)

password := []byte("my-password")

// Соль генерируется автоматически
key, salt, err := secreto.DeriveKey(password, nil)
if err != nil {
    panic(err)
}
fmt.Printf("Ключ: %x\n", key)
fmt.Printf("Соль: %x\n", salt)

// Соль можно указать вручную
customSalt := make([]byte, secreto.SaltLength)
key, _, err = secreto.DeriveKey(password, customSalt)
```

## Упрощённое шифрование

### EncryptSimple

Шифрует данные напрямую с заданным ключом:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/secreto"
)

// Ключ должен быть 32 байта (AES-256)
key := []byte("0123456789abcdef0123456789abcdef")
data := []byte("Hello, simple encryption!")

ciphertext, err := secreto.EncryptSimple(key, data)
if err != nil {
    panic(err)
}
fmt.Printf("Зашифровано: %x\n", ciphertext)
```

### DecryptSimple

Дешифрует данные напрямую с заданным ключом:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/secreto"
)

key := []byte("0123456789abcdef0123456789abcdef")
ciphertext := []byte{...} // результат EncryptSimple

plaintext, err := secreto.DecryptSimple(key, ciphertext)
if err != nil {
    panic(err)
}
fmt.Println(string(plaintext)) // Hello, simple encryption!
```

### GenerateSimpleKey

Генерирует случайный ключ для упрощённого шифрования:

```go
import (
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/secreto"
)

key, err := secreto.GenerateSimpleKey()
if err != nil {
    panic(err)
}
fmt.Printf("Сгенерированный ключ: %x\n", key)
// Длина ключа: 32 байта
```

## Пример: сохранение и восстановление секрета

```go
import (
    "encoding/base64"
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/secreto"
)

// Сохранение
func saveSecret(password []byte, secret string) (string, error) {
    ciphertext, err := secreto.Encrypt(password, []byte(secret))
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Восстановление
func loadSecret(password []byte, encoded string) (string, error) {
    ciphertext, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        return "", err
    }
    plaintext, err := secreto.Decrypt(password, ciphertext)
    if err != nil {
        return "", err
    }
    return string(plaintext), nil
}

password := []byte("master-password-123")
secret := "my-private-data"

encoded, err := saveSecret(password, secret)
if err != nil {
    panic(err)
}
fmt.Printf("Зашифровано: %s\n", encoded)

decoded, err := loadSecret(password, encoded)
if err != nil {
    panic(err)
}
fmt.Printf("Расшифровано: %s\n", decoded) // my-private-data
```

## Пример: работа с простыми ключами

```go
import (
    "encoding/base64"
    "fmt"
    "github.com/nobuenhombre/suikat/pkg/secreto"
)

func encryptData(data string) (string, error) {
    key, err := secreto.GenerateSimpleKey()
    if err != nil {
        return "", err
    }

    ciphertext, err := secreto.EncryptSimple(key, []byte(data))
    if err != nil {
        return "", err
    }

    // В реальном приложении ключ нужно сохранить отдельно!
    return base64.StdEncoding.EncodeToString(ciphertext) + "|" + base64.StdEncoding.EncodeToString(key), nil
}

func decryptData(encoded string) (string, error) {
    parts := []byte(encoded)
    split := 0
    for i, b := range parts {
        if b == '|' {
            split = i
            break
        }
    }

    ciphertext, _ := base64.StdEncoding.DecodeString(string(parts[:split]))
    key, _ := base64.StdEncoding.DecodeString(string(parts[split+1:]))

    plaintext, err := secreto.DecryptSimple(key, ciphertext)
    if err != nil {
        return "", err
    }
    return string(plaintext), nil
}

data := "sensitive information"
encoded, err := encryptData(data)
if err != nil {
    panic(err)
}
fmt.Printf("Зашифровано: %s\n", encoded)

decoded, err := decryptData(encoded)
if err != nil {
    panic(err)
}
fmt.Printf("Расшифровано: %s\n", decoded) // sensitive information
```

## Константы

```go
import "github.com/nobuenhombre/suikat/pkg/secreto"

// Длина соли (32 байта)
_ = secreto.SaltLength // 32

// Длина ключа (32 байта = 256 бит)
_ = secreto.KeyLength // 32
```

## Примечания

- Стандартные функции (`Encrypt`/`Decrypt`) используют scrypt с параметрами: `N=1048576, r=8, p=1`
- Соль автоматически генерируется и хранится в конце зашифрованных данных
- Для упрощённого шифрования ключ должен быть ровно 32 байта
- AES-GCM обеспечивает как шифрование, так и аутентификацию данных
