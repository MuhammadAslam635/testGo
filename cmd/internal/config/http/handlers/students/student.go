package students

import (
	"encoding/json"
	"errors"
	"example/hello/cmd/internal/storage"
	"example/hello/cmd/internal/types"
	"example/hello/cmd/internal/utils/response"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
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
		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		if err != nil {
			response.WriterJson(w, http.StatusInternalServerError, err)
		}

		response.WriterJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}
func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue(("id"))
		slog.Info("GetById Student", slog.String("id", id))
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriterJson(w, http.StatusBadRequest, response.GeneralError((err)))
			return
		}

		student, e := storage.GetStudentById(intId)
		if e != nil {
			response.WriterJson(w, http.StatusNotFound, response.GeneralError(e))
			return
		}
		response.WriterJson(w, http.StatusOK, student)
	}
}
