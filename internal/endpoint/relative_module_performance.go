package endpoint

import (
	"fmt"
	"strconv"

	"dhbw-loerrach.de/dualis/microservice/internal"
	"dhbw-loerrach.de/dualis/microservice/internal/api/models"
	"dhbw-loerrach.de/dualis/microservice/internal/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func HandleStudentRelativeModulePerformance(params operations.StudentRelativeModulePerformanceParams, principal interface{}) middleware.Responder {

	var (
		studentIdDB      int64
		moduleGrade      string
		moduleGradeFloat float32
		semesterFK       int64
		moduleFK         int64
		currentGrade     float32
	)

	relativePerformance := models.RelativePerformance{}
	betterGrades := []float32{}
	equalGrades := []float32{}
	worseGrades := []float32{}

	rows, err := db.Query("SELECT student.id, REPLACE(grade, ',', '.') as grade, semester_fk, module_fk FROM dualis.enrollment INNER JOIN dualis.student ON dualis.enrollment.student_fk = dualis.student.id WHERE dualis.enrollment.id = ?", *params.EnrollmentID)

	if err != nil {
		errMsg := err.Error()
		return operations.NewStudentRelativeModulePerformanceInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
	}

	defer rows.Close()

	entered := false

	for rows.Next() {
		entered = true
		err = rows.Scan(&studentIdDB, &moduleGrade, &semesterFK, &moduleFK)

		if err != nil {
			errMsg := err.Error()
			return operations.NewStudentRelativeModulePerformanceInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
		}

		value, err := strconv.ParseFloat(moduleGrade, 32)
		if err != nil {
			return operations.NewStudentRelativeModulePerformanceNoContent()
		}
		moduleGradeFloat = float32(value)
	}

	if !entered {
		return operations.NewStudentRelativeModulePerformanceNoContent()
	}

	claims := principal.(*internal.DualisClaims)

	if studentIdDB != claims.StudentID {
		unauthorized := "The studentId which belongs to this enrollment is not yours!"
		return operations.NewStudentRelativeModulePerformanceInternalServerError().WithPayload(&models.SimpleError{Error: &unauthorized})
	}

	//rows2, err := db.Query("SELECT CAST(REPLACE(grade, ',', '.') AS double) AS grades FROM enrollment WHERE semester_fk = ? and module_fk = ?", semesterFK, moduleFK)
	rows2, err := db.Query("SELECT CAST(REPLACE(grade, ',', '.') AS double) AS grades FROM enrollment INNER JOIN student ON enrollment.student_fk = student.id INNER JOIN course ON student.course_fk = course.id INNER JOIN organization ON course.organization_fk = organization.id WHERE semester_fk = ? AND module_fk = ? AND organization.id = ?", semesterFK, moduleFK, claims.Organization)

	if err != nil {
		errMsg := err.Error()
		return operations.NewStudentRelativeModulePerformanceInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
	}

	defer rows2.Close()

	for rows2.Next() {
		err = rows2.Scan(&currentGrade)

		if err != nil {
			errMsg := err.Error()
			return operations.NewStudentRelativeModulePerformanceInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
		}

		currentGradeVal := currentGrade

		switch {
		case currentGradeVal < moduleGradeFloat:
			betterGrades = append(betterGrades, currentGradeVal)
		case currentGradeVal == moduleGradeFloat:
			equalGrades = append(equalGrades, currentGradeVal)
		case currentGradeVal > moduleGradeFloat:
			worseGrades = append(worseGrades, currentGradeVal)
		}

	}

	gradeCount := len(betterGrades) + len(equalGrades) + len(worseGrades)

	better := fmt.Sprintf("%v%%", (float32(len(betterGrades))/float32(gradeCount))*100)
	equal := fmt.Sprintf("%v%%", (float32(len(equalGrades))/float32(gradeCount))*100)
	worse := fmt.Sprintf("%v%%", (float32(len(worseGrades))/float32(gradeCount))*100)

	relativePerformance = models.RelativePerformance{Better: &better, Equal: &equal, Worse: &worse}

	return operations.NewStudentRelativeModulePerformanceOK().WithPayload(&relativePerformance)

}
