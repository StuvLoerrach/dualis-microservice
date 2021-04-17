package endpoint

import (
	"fmt"

	"dhbw-loerrach.de/dualis/microservice/internal/api/models"
	"dhbw-loerrach.de/dualis/microservice/internal/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// HandleStudents finds students by various criteria
func HandleStudents(params operations.StudentsParams) middleware.Responder {

	var (
		id           int64
		email        string
		organization int64
	)

	if params.StudentID != nil {
		id = *params.StudentID
	}

	if params.Email != nil {
		email = *params.Email
	}

	if params.Organization != nil {
		organization = *params.Organization
	}

	/*var errMsg *string
	students := []*models.Student{}*/
	studentList := models.StudentList{}

	filterableValues := map[string]interface{}{"id": id, "email": email, "organization_fk": organization}
	var values []interface{}
	var where []string
	for a, b := range filterableValues {
		if b != nil {
			fmt.Println("Not nil")
			fmt.Println(b)
			where = append(where, fmt.Sprintf("%s = ?", a))
			values = append(values, b)
		} else {
			fmt.Println("Nil")
		}
	}

	fmt.Println(where)
	fmt.Println(values...)
	/*rows, err := db.Query("SELECT * FROM dualis.student WHERE "+strings.Join(where, " AND "), values...)

	if err != nil {
		*errMsg = err.Error()
		return operations.NewStudentsInternalServerError().WithPayload(&models.SimpleError{Error: errMsg})
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&id, &email, &organization)

		if err != nil {
			*errMsg = err.Error()
			return operations.NewStudentsInternalServerError().WithPayload(&models.SimpleError{Error: errMsg})
		}

		students = append(students, &models.Student{ID: &id, Email: &email, Organization: &organization})

	}

	studentList = models.StudentList{Students: students}*/

	return operations.NewStudentsOK().WithPayload(&studentList)
}
