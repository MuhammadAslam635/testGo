package students

import (
	"encoding/json"
	"errors"
	"example/hello/cmd/internal/types"
	"example/hello/cmd/internal/utils/response"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriterJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err != nil {
			response.WriterJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		// request validate
		// logic to store student data in database or other storage
		// ...

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriterJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))

			return
		}

		w.Write([]byte("Welcome To Student Api"))
		response.WriterJson(w, http.StatusCreated, map[string]string{"sucess": "ok"})
	}
}
