package endpoint

import (
	"dhbw-loerrach.de/dualis/microservice/internal"
	"dhbw-loerrach.de/dualis/microservice/internal/api/models"
	"dhbw-loerrach.de/dualis/microservice/internal/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func HandleStudentRelativeModulePerformance(params operations.StudentRelativeModulePerformanceParams, principal interface{}) middleware.Responder {

	var (
		studentIdDB int64
		moduleGrade string
	)

	var errMsg *string
	relativePerformance := models.RelativePerformance{}

	rows, err := db.Query("SELECT student.id, grade FROM dualis.enrollment INNER JOIN dualis.student ON dualis.enrollment.student_fk = dualis.student.id WHERE dualis.enrollment.id = ?", *params.EnrollmentID)

	if err != nil {
		*errMsg = err.Error()
		return operations.NewStudentRelativeModulePerformanceInternalServerError().WithPayload(&models.SimpleError{Error: errMsg})
	}

	defer rows.Close()

	entered := false

	for rows.Next() {
		entered = true
		err = rows.Scan(&studentIdDB, &moduleGrade)

		if err != nil {
			*errMsg = err.Error()
			return operations.NewStudentRelativeModulePerformanceInternalServerError().WithPayload(&models.SimpleError{Error: errMsg})
		}
	}

	if !entered {
		return operations.NewStudentRelativeModulePerformanceNoContent()
	}

	claims := principal.(*internal.DualisClaims)

	if studentIdDB != claims.StudentID {
		unauthorized := "The studentId which belongs to this enrollment is not yours!"
		return operations.NewStudentRelativeModulePerformanceInternalServerError().WithPayload(&models.SimpleError{Error: &unauthorized})
	}

	better := "20%"
	equal := "10%"
	worse := "70%"

	relativePerformance = models.RelativePerformance{Better: &better, Equal: &equal, Worse: &worse}

	return operations.NewStudentRelativeModulePerformanceOK().WithPayload(&relativePerformance)

}
