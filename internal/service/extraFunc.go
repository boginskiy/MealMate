package service

import (
	"fmt"
	"io"
	m "mealmate/internal/model"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ExFuncer interface {
	ReadRequestBody(*http.Request) ([]byte, error)
	NeedShow(m.Fooder, url.Values) bool
}

type ExtraFunc struct {
}

func NewExtraFunc() *ExtraFunc {
	return &ExtraFunc{}
}

func (ex *ExtraFunc) QueryParamsGoToLower(params url.Values) url.Values {
	max := len(params)
	for key, value := range params {
		params[strings.ToLower(key)] = value

		if len(params) > max {
			delete(params, key)
		}
	}
	return params
}

func (ex *ExtraFunc) NeedShow(food m.Fooder, queryParams url.Values) bool {
	// Go to lower key for good compare continue ...
	queryParams = ex.QueryParamsGoToLower(queryParams)

	// Get fields of struct of interface of Fooder
	_value := reflect.ValueOf(food)
	_type := reflect.TypeOf(food)
	numFields := _type.NumField()

	var flagCheck uint8

	// Go to every field of struct
	for i := 0; i < numFields; i++ {
		field := _type.Field(i)
		name := field.Name
		kind := field.Type.Kind()
		value := _value.FieldByName(name).Interface()

		// Check in field in params
		paramStr := queryParams.Get(strings.ToLower(name))
		flagCheck = 0

		// If the name is absent, go to continue
		if paramStr == "" {
			continue
		}

		// Filtering only easy kind of paramStr. Check only ==
		switch kind {

		case reflect.Int:
			paramInt, err := strconv.Atoi(paramStr)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%v\n", err)
			}
			// Check Int
			if value.(int) == paramInt {
				flagCheck = 1 << 0
			}

		case reflect.String:
			// Check String
			if strings.EqualFold(value.(string), paramStr) {
				flagCheck = 1 << 0
			}

		case reflect.Float64:
			paramFlt, err := strconv.ParseFloat(paramStr, 64)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%v\n", err)
			}
			// Check Float
			if value.(float64) == paramFlt {
				flagCheck = 1 << 0
			}
		}

		// Common Check
		if flagCheck == uint8(0) {
			return false
		}
	}
	return true
}

func (ex *ExtraFunc) ReadRequestBody(req *http.Request) ([]byte, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
