package usecases_test

import (
	"errors"
	"testing"

	"github.com/SallyKinoshita/compass/internal/application/usecases"
	"github.com/SallyKinoshita/compass/internal/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStudentRepository struct {
	mock.Mock
}

func (m *MockStudentRepository) FindByFacilitatorID(facilitatorID, page, limit int, sort, order string) ([]models.Student, int, error) {
	args := m.Called(facilitatorID, page, limit, sort, order)
	return args.Get(0).([]models.Student), args.Int(1), args.Error(2)
}

func TestGetStudentsByFacilitator_Success(t *testing.T) {
	t.Parallel()

	mockRepo := new(MockStudentRepository)
	usecase := usecases.NewStudentUsecase(mockRepo)

	facilitatorID := 1
	page := 1
	limit := 10
	sort := "id"
	order := "asc"

	expectedStudents := []models.Student{
		{ID: 1, Name: "John Doe", LoginID: "john_doe", Classroom: models.Classroom{ID: 101, Name: "Classroom 1"}},
		{ID: 2, Name: "Jane Smith", LoginID: "jane_smith", Classroom: models.Classroom{ID: 102, Name: "Classroom 2"}},
	}
	expectedTotalCount := 2

	mockRepo.On("FindByFacilitatorID", facilitatorID, page, limit, sort, order).Return(expectedStudents, expectedTotalCount, nil)

	students, totalCount, err := usecase.GetStudentsByFacilitator(facilitatorID, page, limit, sort, order)

	assert.NoError(t, err)
	assert.Equal(t, expectedStudents, students)
	assert.Equal(t, expectedTotalCount, totalCount)
	mockRepo.AssertExpectations(t)
}

func TestGetStudentsByFacilitator_Error(t *testing.T) {
	t.Parallel()

	mockRepo := new(MockStudentRepository)
	usecase := usecases.NewStudentUsecase(mockRepo)

	facilitatorID := 1
	page := 1
	limit := 10
	sort := "id"
	order := "asc"

	mockRepo.On("FindByFacilitatorID", facilitatorID, page, limit, sort, order).Return([]models.Student{}, 0, errors.New("database error"))

	students, totalCount, err := usecase.GetStudentsByFacilitator(facilitatorID, page, limit, sort, order)

	// テスト結果の検証
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	assert.Empty(t, students) // 空のスライスが返ることを確認
	assert.Equal(t, 0, totalCount)
	mockRepo.AssertExpectations(t)
}
