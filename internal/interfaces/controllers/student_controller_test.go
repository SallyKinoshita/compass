package controllers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SallyKinoshita/compass/internal/domain/models"
	openapicompass "github.com/SallyKinoshita/compass/internal/gen/openapi"
	"github.com/SallyKinoshita/compass/internal/interfaces/controllers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStudentUsecase struct {
	mock.Mock
}

func (m *MockStudentUsecase) GetStudentsByFacilitator(facilitatorID, page, limit int, sort, order string) ([]models.Student, int, error) {
	args := m.Called(facilitatorID, page, limit, sort, order)
	return args.Get(0).([]models.Student), args.Int(1), args.Error(2)
}

func TestGetStudents_Success(t *testing.T) {
	t.Parallel()

	e := echo.New()

	mockStudents := []models.Student{
		{ID: 1, Name: "John Doe", LoginID: "john.doe", Classroom: models.Classroom{ID: 0, Name: ""}},
		{ID: 2, Name: "Jane Doe", LoginID: "jane.doe", Classroom: models.Classroom{ID: 0, Name: ""}},
	}
	mockTotalCount := 2

	mockUsecase := new(MockStudentUsecase)
	mockUsecase.On("GetStudentsByFacilitator", 1, 1, 10, "id", "asc").Return(mockStudents, mockTotalCount, nil)

	controller := controllers.NewStudentController(mockUsecase)

	req := httptest.NewRequest(http.MethodGet, "/students?facilitator_id=1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	params := openapicompass.GetStudentsParams{
		FacilitatorId: 1,
	}

	err := controller.GetStudents(c, params)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedBody := `{
		"students": [
			{"id": 1, "name": "John Doe", "loginId": "john.doe", "Classroom": {"id": 0, "name": ""}},
			{"id": 2, "name": "Jane Doe", "loginId": "jane.doe", "Classroom": {"id": 0, "name": ""}}
		],
		"totalCount": 2
	}`
	assert.JSONEq(t, expectedBody, rec.Body.String())
}

func TestGetStudents_InternalServerError(t *testing.T) {
	t.Parallel()

	e := echo.New()

	mockUsecase := new(MockStudentUsecase)
	mockUsecase.On("GetStudentsByFacilitator", 1, 1, 10, "id", "asc").Return([]models.Student{}, 0, errors.New("unexpected error"))

	controller := controllers.NewStudentController(mockUsecase)

	req := httptest.NewRequest(http.MethodGet, "/students?facilitator_id=1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	params := openapicompass.GetStudentsParams{
		FacilitatorId: 1,
	}

	err := controller.GetStudents(c, params)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, "", rec.Body.String()) // レスポンスボディは空
}

func TestGetStudents_InvalidQueryParam(t *testing.T) {
	t.Parallel()

	e := echo.New()

	mockUsecase := new(MockStudentUsecase)

	controller := controllers.NewStudentController(mockUsecase)

	// 不正な値のリクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/students?facilitator_id=0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := controller.GetStudents(c, openapicompass.GetStudentsParams{
		FacilitatorId: 0,
	})

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "", rec.Body.String()) // レスポンスボディは空
}
