package db

type warning string

const (
	notFoundFoodWarn warning = "Not found food for update, check fields 'name'"
	unicFoodWarn     warning = "Added food is not unic, check fields 'name'"
)
