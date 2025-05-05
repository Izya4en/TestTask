package handler

import (
	"TestTask/migration"
	"TestTask/model"
	"encoding/json"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Пример обогащения данных: запрашиваем возраст, пол и национальность (пока заглушка)
	// Ты можешь здесь добавить логику для интеграции с внешними API

	// Сохраняем в базу данных
	if err := migration.DB.Create(&person).Error; err != nil {
		http.Error(w, "Failed to create person", http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

// Функция для получения всех людей
func GetPeople(w http.ResponseWriter, r *http.Request) {
	var people []model.Person
	if err := migration.DB.Find(&people).Error; err != nil {
		http.Error(w, "Failed to retrieve people", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(people)
}

// Функция для получения человека по ID
func GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person model.Person
	if err := migration.DB.First(&person, id).Error; err != nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(person)
}

// Функция для обновления данных человека
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person model.Person
	if err := migration.DB.First(&person, id).Error; err != nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := migration.DB.Save(&person).Error; err != nil {
		http.Error(w, "Failed to update person", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(person)
}

// Функция для удаления человека
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person model.Person
	if err := migration.DB.First(&person, id).Error; err != nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	if err := migration.DB.Delete(&person).Error; err != nil {
		http.Error(w, "Failed to delete person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
