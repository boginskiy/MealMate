package db

import (
	"errors"
	"mealmate/internal/warnings"
)

// Errors:
// Переменная для создания будущих error
var someErr = errors.New("")

// Warnings:
// Не найден объект для обновления
var notFoundFoodWarn = warnings.New("not found food for update, check fields 'name'")

// Добавляемый объект не уникален, нужно изменить 'name'
var notUnicFoodWarn = warnings.New("added food is not unic, check fields 'name'")
