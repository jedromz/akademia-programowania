package academy

import (
	"github.com/grupawp/akademia-programowania/Golang/zadania/academy2/mocks"
	"testing"
)

func TestGradeStudent(t *testing.T) {

	t.Run("Student not found", func(t *testing.T) {
		repository := mocks.NewRepository(t)
		repository.On("Get", "John").Return(nil, ErrStudentNotFound)
		err := GradeStudent(repository, "John")
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
	})

}
