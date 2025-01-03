package usecases

import (
	"github.com/SallyKinoshita/compass/internal/domain/models"
	"github.com/SallyKinoshita/compass/internal/domain/repositories"
)

type StudentUsecase interface {
	GetStudentsByFacilitator(facilitatorID, page, limit int, sort, order string) ([]models.Student, int, error)
}

type StudentUsecaseImpl struct {
	Repo repositories.StudentRepository
}

func NewStudentUsecase(repo repositories.StudentRepository) StudentUsecase {
	return &StudentUsecaseImpl{Repo: repo}
}

func (u *StudentUsecaseImpl) GetStudentsByFacilitator(facilitatorID, page, limit int, sort, order string) ([]models.Student, int, error) {
	return u.Repo.FindByFacilitatorID(facilitatorID, page, limit, sort, order)
}
