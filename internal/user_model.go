package internal

import (
	"database/sql"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getAllUsers() ([]User, error) {
	// Получить всех пользователей из базы данных
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Преобразовать строки в структуры User
	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func createUser(user User) error {
	// Создать пользователя в базе данных
	_, err := db.Exec("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func getUser(id string) (User, error) {
	// Получить пользователя из базы данных
	var user User
	err := db.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1", id).Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil
		} else {
			return User{}, err
		}
	}

	return user, nil
}

func updateUser(id string, user User) error {
	// Обновить пользователя в базе данных
	_, err := db.Exec("UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4", user.Name, user.Email, user.Password, id)
	if err != nil {
		return err
	}

	return nil
}

func deleteUser(id string) error {
	// Удалить пользователя из базы данных
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
