# 1. Project Setup

Langkah awal untuk membuat file `go.mod` ketikan perintah berikut:

```
go mod init <nama_module>
```

### A. Instal Library Yang Dibutuhkan

#### 1. Go Fiber (http server)

```
go get github.com/gofiber/fiber/v2
```

#### 2. Air (hot reload)

```
go install github.com/gofiber/fiber/v2
```

buat file .air.toml kemudian tambahkan kode berikut:

```
root = "."
tmp_dir = "tmp"

[build]
    bin = "main
    cmd = "go build -o {{.Output}} {{.Input}}"
    exclude = ["tmp/*", "client/*"]
    include = ["**/*.go"]
    ignore = ["tmp/*"]
```

### B. Contoh Pembuatan API Sederhana

edit file `main.go` dengan kode berikut:

```
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	log.Fatal(app.Listen(":5000"))
}
```
