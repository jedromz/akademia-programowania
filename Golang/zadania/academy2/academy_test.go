package academy_test

import (
	"fmt"
	academy "github.com/grupawp/akademia-programowania/Golang/zadania/academy2"
	"github.com/grupawp/akademia-programowania/Golang/zadania/academy2/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGradeStudent(t *testing.T) {
	invalidGrades := []int{-1, 0, 6}
	passingGrades := []int{2, 3, 4, 5}

	for _, grade := range invalidGrades {
		t.Run(fmt.Sprintf("should return error when student grade is %d", grade), func(t *testing.T) {

			mockRepo := mocks.NewRepository(t)
			mockStudent := mocks.NewStudent(t)
			mockStudent.On("FinalGrade").Return(grade)
			mockRepo.On("Get", "Alice").Return(mockStudent, nil)

			err := academy.GradeStudent(mockRepo, "Alice")

			assert.Equal(t, academy.ErrInvalidGrade, err)
		})
	}

	t.Run("should return nil when student not found", func(t *testing.T) {

		mockRepo := mocks.NewRepository(t)
		mockRepo.On("Get", "Alice").Return(nil, academy.ErrStudentNotFound)

		err := academy.GradeStudent(mockRepo, "Alice")

		assert.NoError(t, err)
	})

	t.Run("should not promote student when grade is 1", func(t *testing.T) {

		mockRepo := mocks.NewRepository(t)
		mockStudent := mocks.NewStudent(t)

		mockStudent.On("FinalGrade").Return(1)
		mockStudent.On("Name").Return("Alice")
		mockStudent.On("Year").Return(uint8(1))
		mockRepo.On("Get", "Alice").Return(mockStudent, nil)
		mockRepo.On("Save", "Alice", uint8(1)).Return(nil)

		err := academy.GradeStudent(mockRepo, mockStudent.Name())

		assert.NoError(t, err)
		mockRepo.AssertCalled(t, "Save", mockStudent.Name(), mockStudent.Year())
	})

	t.Run("should promote student when grade is more than 1", func(t *testing.T) {

		mockRepo := mocks.NewRepository(t)
		mockStudent := mocks.NewStudent(t)

		mockStudent.On("FinalGrade").Return(2)
		mockStudent.On("Name").Return("Alice")
		mockStudent.On("Year").Return(uint8(1))
		mockRepo.On("Get", "Alice").Return(mockStudent, nil)
		mockRepo.On("Save", "Alice", uint8(2)).Return(nil)

		err := academy.GradeStudent(mockRepo, mockStudent.Name())

		assert.NoError(t, err)
		mockRepo.AssertCalled(t, "Save", mockStudent.Name(), mockStudent.Year()+1)
	})

	for _, grade := range passingGrades {
		t.Run(fmt.Sprintf("should graduate student when grade is %d", grade), func(t *testing.T) {

			mockRepo := mocks.NewRepository(t)
			mockStudent := mocks.NewStudent(t)

			mockStudent.On("FinalGrade").Return(grade)
			mockStudent.On("Name").Return("Alice")
			mockStudent.On("Year").Return(uint8(3))
			mockRepo.On("Get", "Alice").Return(mockStudent, nil)
			mockRepo.On("Graduate", "Alice").Return(nil)

			err := academy.GradeStudent(mockRepo, mockStudent.Name())

			assert.NoError(t, err)
			mockRepo.AssertCalled(t, "Graduate", mockStudent.Name())
		})
	}
}
func TestGradeYear(t *testing.T) {
	t.Run("should return error when repository returns error", func(t *testing.T) {

		mockRepo := mocks.NewRepository(t)
		mockRepo.On("List", uint8(1)).Return(nil, fmt.Errorf("failed to list students"))

		err := academy.GradeYear(mockRepo, uint8(1))

		assert.Error(t, err)
	})

	t.Run("should return nil when repository returns no students", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)
		mockRepo.On("List", uint8(1)).Return(nil, nil)

		err := academy.GradeYear(mockRepo, uint8(1))

		assert.NoError(t, err)
	})
	t.Run("should return error when GradeYear returns error", func(t *testing.T) {
		mockRepo := mocks.NewRepository(t)

		mockRepo.On("List", uint8(1)).Return([]string{"Alice"}, nil)
		mockRepo.On("Get", "Alice").Return(nil, errors.New("failed to get student"))

		err := academy.GradeYear(mockRepo, uint8(1))

		assert.Error(t, err)
	})

}
