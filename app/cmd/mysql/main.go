package main

import (
	"fmt"
	"time"

	"github.com/urataku/go-docker/internal/mysql"
)

func main() {
	db := mysql.MysqlConnection()
	var user User
	db.First(&user)
	fmt.Println(user)
}

type User struct {
	ID        int       `json:id`
	Name      string    `json:name`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at`
}
