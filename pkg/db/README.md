# db

Пакет `db` предоставляет коннекторы для подключения к SQL-базам данных (MariaDB, PostgreSQL) и общие типы для работы с ними.

## Структура

- `types/` — общие типы (интерфейсы логирования, конфигурации)
- `connectors/maria-db/` — коннектор для MariaDB (через `database/sql`)
- `connectors/postgres-pgx-db/` — коннектор для PostgreSQL (через `pgx` + `pgxpool`)

## MariaDB

### Подключение

```go
import (
    "github.com/nobuenhombre/suikat/pkg/db/connectors/maria-db"
    "github.com/nobuenhombre/suikat/pkg/ge"
)

cfg := &mariadb.Config{
    Protocol: mariadb.ProtocolTCP, // или mariadb.ProtocolUNIXSocket
    Address:  "localhost:3306",
    Schema:   "mydb",
    User:     "root",
    Password: "password",
    Charset:  "utf8mb4",
}

conn, err := mariadb.New(cfg, func(sql string, du time.Duration, params ...any) {
    log.Printf("SQL: %s | %v | %+v", sql, du, params)
})
if err != nil {
    log.Fatal(ge.Pin(err))
}
defer conn.Close()
```

### Выполнение запросов

```go
// Exec
result, err := conn.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Alice", "alice@example.com")

// Query
rows, err := conn.Query("SELECT id, name FROM users WHERE email LIKE ?", "%@example.com")
defer rows.Close()
for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
    fmt.Println(id, name)
}

// QueryRow
var name string
err := conn.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&name)

// Транзакции
tx, err := conn.Begin()
result, err := tx.Exec("INSERT INTO orders (user_id) VALUES (?)", 1)
err = tx.Commit()
```

## PostgreSQL (pgx)

### Подключение

```go
import (
    "github.com/nobuenhombre/suikat/pkg/db/connectors/postgres-pgx-db"
    "context"
)

cfg := &postgrespgxdb.Config{
    Host:     "localhost",
    Port:     "5432",
    Name:     "mydb",
    User:     "postgres",
    Password: "password",
    SSLMode:  "disable",
}

conn, err := postgrespgxdb.New(cfg, func(sql string, du time.Duration, params ...any) {
    log.Printf("SQL: %s | %v | %+v", sql, du, params)
})
if err != nil {
    log.Fatal(err)
}
defer conn.Close()
```

### Подключение через PgBouncer

```go
cfg := &postgrespgxdb.Config{
    Host:                "localhost",
    Port:                "6432",
    Name:                "mydb",
    User:                "postgres",
    Password:            "password",
    UseConnectionPooler: true,
    StatementCacheMode:  postgrespgxdb.SCMDescribe, // для PgBouncer
}
```

### Выполнение запросов

```go
// Exec
tag, err := conn.Exec(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2)", "Alice", "alice@example.com")
fmt.Println(tag.RowsAffected())

// Query
rows, err := conn.Query(context.Background(), "SELECT id, name FROM users WHERE email LIKE $1", "%@example.com")
defer rows.Close()
for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
    fmt.Println(id, name)
}

// QueryRow
var name string
err := conn.QueryRow(context.Background(), "SELECT name FROM users WHERE id = $1", 1).Scan(&name)

// Batch
batch := &pgx.Batch{}
batch.Queue("INSERT INTO users (name) VALUES ($1)", "Alice")
batch.Queue("INSERT INTO users (name) VALUES ($1)", "Bob")
results := conn.SendBatch(context.Background(), batch)
for i := 0; i < batch.Len(); i++ {
    tag, err := results.Exec()
    fmt.Printf("Row %d: %s\n", i, tag.String())
}
results.Close()

// Транзакции
tx, err := conn.Begin(context.Background())
tag, err := tx.Exec(context.Background(), "INSERT INTO orders (user_id) VALUES ($1)", 1)
err = tx.Commit(context.Background())
```

### Проверка ошибок

```go
import "github.com/nobuenhombre/suikat/pkg/db/connectors/postgres-pgx-db"

// Проверка на дубликат по уникальному ограничению
if postgrespgxdb.IsConstraintError(err, "users_email_key") {
    log.Println("Email уже занят")
}

// Проверка на недостаточное количество затронутых строк
if _, ok := err.(*postgrespgxdb.AffectedLessThenNecessaryError); ok {
    log.Println("Затронуто меньше строк, чем ожидалось")
}
```
