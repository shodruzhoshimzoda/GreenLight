package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// readIDParams - данная вспомогательная функция используется для извлечения ID из тело URL и возвршает его
func (app *application) readIDParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

type envelope map[string]any

// writeJson - вспомогательная функция которая помогает для переобразование объекта в JSON
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, header http.Header) error {

	js, err := json.MarshalIndent(data, " ", "\t") // Для более удобочитаемости используется MarshalIndent()
	if err != nil {
		return err
	}

	for k, v := range header {
		w.Header()[k] = v
	}

	js = append(js, '\n') // Добавления в конце спика байтов символ перехода на новую строку
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// readJSON - это функция прочитает тело запроса декодирует значения в структур в случае ошибки возвращает ошибку
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	// Декодируем тело запроса в целевую структуру dst
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		// Объявляем переменные для конкретных типов ошибок JSON
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalFieldError *json.InvalidUnmarshalError

		switch {
		// Ошибка синтаксиса: например, пропущена запятая или кавычка
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly formed JSON (at character %d)", syntaxError.Offset)

		// Ошибка неожиданного конца файла: JSON прерван на полпути
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly formed JSON")

		// Ошибка несоответствия типов: например, передана строка вместо ожидаемого числа
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		// Ошибка пустого тела: если клиент вообще ничего не прислал в запросе
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		// Ошибка разработчика: передача неверного типа в dst (не указатель)
		// Вызываем panic, так как это критическая ошибка в логике самого сервера
		case errors.As(err, &invalidUnmarshalFieldError):
			panic(err)
		}

		// Возвращаем все остальные типы ошибок без изменений
		return err
	}

	return nil
}
