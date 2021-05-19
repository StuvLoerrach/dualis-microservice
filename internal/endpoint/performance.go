package endpoint

import (
	"fmt"
	"strings"

	"dhbw-loerrach.de/dualis/microservice/internal"
	"dhbw-loerrach.de/dualis/microservice/internal/api/models"
	"dhbw-loerrach.de/dualis/microservice/internal/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func HandleStudentPerformance(params operations.StudentPerformanceParams, principal interface{}) middleware.Responder {

	var (
		isWintersemester bool
		year             string
	)

	lectureResults := []*models.LectureResult{}
	moduleResults := []*models.ModuleResult{}
	enrollments := []*models.Enrollment{}
	performancesList := models.PerformancesList{}

	claims := principal.(*internal.DualisClaims)
	filterableValues := map[string]interface{}{"student_fk": claims.StudentID}

	if params.IsWintersemester != nil {
		filterableValues["is_wintersemester"] = *params.IsWintersemester
	}

	if params.Year != nil {
		filterableValues["year"] = *params.Year
	}

	var values []interface{}
	var where []string
	for a, b := range filterableValues {
		where = append(where, fmt.Sprintf("%s = ?", a))
		var pVal *interface{}
		pVal = &b
		values = append(values, *pVal)
	}

	rows, err := db.Query("SELECT enrollment.id, grade, status, is_wintersemester, year, no, name, credits FROM enrollment INNER JOIN semester ON enrollment.semester_fk = semester.id INNER JOIN module ON enrollment.module_fk = module.id WHERE "+strings.Join(where, " AND "), values...)

	if err != nil {
		errMsg := err.Error()
		return operations.NewStudentsInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
	}

	defer rows.Close()

	for rows.Next() {

		var enrollment models.Enrollment
		var moduleResult models.ModuleResult

		err = rows.Scan(&enrollment.ID, &enrollment.Grade, &enrollment.Status, &isWintersemester, &year, &moduleResult.Number, &moduleResult.Name, &moduleResult.Credits)

		if err != nil {
			errMsg := err.Error()
			return operations.NewStudentPerformanceInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
		}

		rows2, err := db.Query("SELECT exam_type, lecture_result.grade, name, no, presence, weighting FROM lecture_result INNER JOIN enrollment ON lecture_result.enrollment_fk = enrollment.id INNER JOIN lecture ON lecture_result.lecture_fk = lecture.id WHERE enrollment.id = ?", *enrollment.ID)

		if err != nil {
			errMsg := err.Error()
			return operations.NewStudentsInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
		}

		defer rows2.Close()

		for rows2.Next() {

			var lectureResult models.LectureResult

			err = rows2.Scan(&lectureResult.ExamType, &lectureResult.Grade, &lectureResult.Name, &lectureResult.Number, &lectureResult.Presence, &lectureResult.Weighting)

			if err != nil {
				errMsg := err.Error()
				return operations.NewStudentPerformanceInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
			}

			lectureResults = append(lectureResults, &lectureResult)
		}

		moduleResult.LectureResults = lectureResults

		moduleResults = append(moduleResults, &moduleResult)
		lectureResults = nil

		isWintersemesterVal := isWintersemester
		yearVal := year

		var semester string

		if isWintersemesterVal {
			semester = "WiSe " + yearVal
		} else {
			semester = "SoSe " + yearVal
		}

		enrollment.Semester = &semester
		enrollment.ModuleResult = moduleResults

		enrollments = append(enrollments, &enrollment)
		moduleResults = nil
	}

	performancesList = models.PerformancesList{Enrollments: enrollments}

	if len(enrollments) == 0 {
		return operations.NewStudentPerformanceNoContent()
	}

	return operations.NewStudentPerformanceOK().WithPayload(&performancesList)

}
