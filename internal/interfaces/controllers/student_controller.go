package controllers

import (
	"log"
	"net/http"

	"github.com/SallyKinoshita/compass/internal/application/usecases"
	openapicompass "github.com/SallyKinoshita/compass/internal/gen/openapi"
	"github.com/labstack/echo/v4"
)

type StudentController struct {
	Usecase usecases.StudentUsecase
}

func NewStudentController(usecase usecases.StudentUsecase) *StudentController {
	return &StudentController{Usecase: usecase}
}

func (c *StudentController) GetStudents(ec echo.Context, params openapicompass.GetStudentsParams) error {
	// 必須パラメータチェック
	if params.FacilitatorId <= 0 {
		log.Printf("Invalid facilitator_id: %d", params.FacilitatorId)
		return ec.NoContent(http.StatusBadRequest) // ボディなしで400を返す
	}

	page := 1
	if params.Page != nil {
		page = *params.Page
	}

	limit := 10
	if params.Limit != nil {
		limit = *params.Limit
	}

	sort := "id"
	if params.Sort != nil {
		sort = string(*params.Sort)
	}

	order := "asc"
	if params.Order != nil {
		order = string(*params.Order)
	}

	students, totalCount, err := c.Usecase.GetStudentsByFacilitator(params.FacilitatorId, page, limit, sort, order)
	if err != nil {
		return ec.NoContent(http.StatusInternalServerError) // ボディなしで500を返す
	}

	response := struct {
		Students   interface{} `json:"students"`
		TotalCount int         `json:"totalCount"`
	}{
		Students:   students,
		TotalCount: totalCount,
	}

	return ec.JSON(http.StatusOK, response)
}
