package database

import (
	models "auth-service/models"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("mysql", "root:root@/local")
	if err != nil {
		panic(err.Error())
	}
}

func GetUser(email string) (models.User, error) {
	var user models.User
	err := DB.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func AddUser(user *models.RegisterRequest) error {
	_, err := DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	return err
}
