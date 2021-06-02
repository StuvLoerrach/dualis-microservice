package endpoint

import (
	"fmt"
	"strconv"

	"dhbw-loerrach.de/dualis/microservice/internal"
	"dhbw-loerrach.de/dualis/microservice/internal/api/models"
	"dhbw-loerrach.de/dualis/microservice/internal/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func HandleStudentModuleStatistics(params operations.StudentModuleStatisticsParams, principal interface{}) middleware.Responder {

	var (
		studentIdDB      int64
		moduleGrade      string
		moduleGradeFloat float32
		semesterFK       int64
		moduleFK         int64
		currentGrade     float32
		amountFailures   int64
	)

	relativePerformance := models.ModuleStatistics{}
	betterGrades := []float32{}
	equalGrades := []float32{}
	worseGrades := []float32{}

	rows, err := db.Query("SELECT student.id, REPLACE(grade, ',', '.') as grade, semester_fk, module_fk FROM dualis.enrollment INNER JOIN dualis.student ON dualis.enrollment.student_fk = dualis.student.id WHERE dualis.enrollment.id = ?", params.EnrollmentID)

	if err != nil {
		errMsg := err.Error()
		return operations.NewStudentModuleStatisticsInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
	}

	defer rows.Close()

	entered := false

	for rows.Next() {
		entered = true
		err = rows.Scan(&studentIdDB, &moduleGrade, &semesterFK, &moduleFK)

		if err != nil {
			errMsg := err.Error()
			return operations.NewStudentModuleStatisticsInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
		}

		value, err := strconv.ParseFloat(moduleGrade, 32)
		if err != nil {
			return operations.NewStudentModuleStatisticsNoContent()
		}
		moduleGradeFloat = float32(value)
	}

	if !entered {
		return operations.NewStudentModuleStatisticsNoContent()
	}

	claims := principal.(*internal.DualisClaims)

	if studentIdDB != claims.StudentID {
		unauthorized := "The studentId which belongs to this enrollment is not yours!"
		return operations.NewStudentModuleStatisticsInternalServerError().WithPayload(&models.SimpleError{Error: &unauthorized})
	}

	//rows2, err := db.Query("SELECT CAST(REPLACE(grade, ',', '.') AS double) AS grades FROM enrollment WHERE semester_fk = ? and module_fk = ?", semesterFK, moduleFK)
	rows2, err := db.Query("SELECT CAST(REPLACE(grade, ',', '.') AS double) AS grades FROM enrollment INNER JOIN student ON enrollment.student_fk = student.id INNER JOIN course ON student.course_fk = course.id INNER JOIN organization ON course.organization_fk = organization.id WHERE semester_fk = ? AND module_fk = ? AND organization.id = ?", semesterFK, moduleFK, claims.Organization)

	if err != nil {
		errMsg := err.Error()
		return operations.NewStudentModuleStatisticsInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
	}

	defer rows2.Close()

	for rows2.Next() {
		err = rows2.Scan(&currentGrade)

		if err != nil {
			errMsg := err.Error()
			return operations.NewStudentModuleStatisticsInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
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

		if currentGradeVal > 4.0 {
			amountFailures++
		}

	}

	gradeCount := len(betterGrades) + len(equalGrades) + len(worseGrades)

	better := "0"
	equal := "0"
	worse := "0"
	failureRate := fmt.Sprintf("%.2f", (float32(amountFailures)/float32(gradeCount))*100)

	if gradeCount > 1 {
		better = fmt.Sprintf("%.2f", (float32(len(betterGrades))/float32(gradeCount-1))*100)
		equal = fmt.Sprintf("%.2f", (float32(len(equalGrades)-1)/float32(gradeCount-1))*100)
		worse = fmt.Sprintf("%.2f", (float32(len(worseGrades))/float32(gradeCount-1))*100)
	}

	relativePerformance = models.ModuleStatistics{Better: &better, Equal: &equal, Worse: &worse, FailureRate: &failureRate}

	return operations.NewStudentModuleStatisticsOK().WithPayload(&relativePerformance)

}
