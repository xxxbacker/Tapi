package internal

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Получить всех пользователей из базы данных
	users, err := getAllUsers()
	if err != nil {
		http.Error(w, "Ошибка получения пользователей", http.StatusInternalServerError)
		return
	}

	// Отправить ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	// Получить данные пользователя из тела запроса
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	// Создать пользователя в базе данных
	err = createUser(user)
	if err != nil {
		http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError)
		return
	}

	// Отправить ответ клиенту
	w.WriteHeader(http.StatusCreated)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Получить ID пользователя из параметра маршрута
	vars := mux.Vars(r)
	id := vars["id"]

	// Получить пользователя из базы данных
	user, err := getUser(id)
	if err != nil {
		http.Error(w, "Ошибка получения пользователя", http.StatusInternalServerError)
		return
	}

	// Отправить ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Получить ID пользователя из параметра маршрута
	vars := mux.Vars(r)
	id := vars["id"]

	// Получить данные пользователя из тела запроса
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	// Обновить пользователя в базе данных
	err = updateUser(id, user)
	if err != nil {
		http.Error(w, "Ошибка обновления пользователя", http.StatusInternalServerError)
		return
	}

	// Отправить ответ клиенту
	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Получить ID пользователя из параметра маршрута
	vars := mux.Vars(r)
	id := vars["id"]

	// Удалить пользователя из базы данных
	err := deleteUser(id)
	if err != nil {
		http.Error(w, "Ошибка удаления пользователя", http.StatusInternalServerError)
		return
	}

	// Отправить ответ клиенту
	w.WriteHeader(http.StatusOK)
}
