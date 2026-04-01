package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue) // используем метод Quote чтобы записать строку в двойных кавычках

	return []byte(quotedJSONValue), nil
}
