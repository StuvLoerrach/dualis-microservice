package endpoint

import (
	"fmt"
	"reflect"
	"strings"

	"dhbw-loerrach.de/dualis/microservice/internal/api/models"
	"dhbw-loerrach.de/dualis/microservice/internal/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// HandleStudents finds students by various criteria
func HandleStudents(params operations.StudentsParams, principal interface{}) middleware.Responder {

	var (
		id     int64
		email  string
		course int64
	)

	var errMsg *string
	students := []*models.Student{}
	studentList := models.StudentList{}

	filterableValues := map[string]interface{}{"id": params.StudentID, "email": params.Email, "course_fk": params.Course}
	var values []interface{}
	var where []string
	for a, b := range filterableValues {
		if !reflect.ValueOf(b).IsNil() {
			where = append(where, fmt.Sprintf("%s = ?", a))
			var pVal *interface{}
			pVal = &b
			values = append(values, *pVal)
		}
	}

	if len(where) == 0 {
		where = append(where, "id IS NOT NULL")
	}

	rows, err := db.Query("SELECT id, email, course_fk FROM dualis.student WHERE "+strings.Join(where, " AND "), values...)

	if err != nil {
		*errMsg = err.Error()
		return operations.NewStudentsInternalServerError().WithPayload(&models.SimpleError{Error: errMsg})
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&id, &email, &course)

		if err != nil {
			*errMsg = err.Error()
			return operations.NewStudentsInternalServerError().WithPayload(&models.SimpleError{Error: errMsg})
		}

		idVal := id
		emailVal := email
		courseVal := course

		students = append(students, &models.Student{ID: &idVal, Email: &emailVal, Course: &courseVal})

	}

	studentList = models.StudentList{Students: students}

	return operations.NewStudentsOK().WithPayload(&studentList)
}
