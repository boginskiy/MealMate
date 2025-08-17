package service

import (
	"errors"
	"mealmate/internal/warnings"
)

// Errors:
// Некорректное тело запроса
var badBodyErr = errors.New("the body for the update is bad.")

// Отсутствует поле ID для обновления
var notIdFieldErr = errors.New("missing ID field")

// Не верный формат данных у поля ID
var notValidIdErr = errors.New("invalid ID format")

// Warnings:
var Warn = warnings.New("")

// Info:
var notIdFromPathInfo = "[INFO]: the 'id' is not in url path"
