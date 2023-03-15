package academy

import (
	"math"
)

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	size := len(grades)
	if size == 0 {
		return 0
	}
	sum := 0
	for _, v := range grades {
		sum += v
	}
	avg := float64(sum) / float64(len(grades))
	return int(math.Round(avg))
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from  0 to 1,
// with 2 digits of precision.
func AttendancePercentage(attendance []bool) float64 {
	presentCount := 0
	for _, present := range attendance {
		if present {
			presentCount++
		}
	}
	return float64(presentCount) / float64(len(attendance))
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.
//
//	If the student's attendance is below 80%, the final grade is
//
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	att := AttendancePercentage(s.Attendance)
	avgGrade := AverageGrade(s.Grades)
	//Project or Semester failed
	if s.Project == 1 || avgGrade == 1 {
		return 1
	}
	finalGrade := int(math.Round(float64(avgGrade+s.Project) / 2.0))
	switch {
	case att < 0.6:
		finalGrade = 1
	case att < 0.8:
		finalGrade--
	}
	return finalGrade
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	grades := make(map[string]uint8, len(students))
	for _, s := range students {
		grades[s.Name] = uint8(FinalGrade(s))
	}
	return grades
}
