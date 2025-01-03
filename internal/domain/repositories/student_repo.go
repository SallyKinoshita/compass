package repositories

import "github.com/SallyKinoshita/compass/internal/domain/models"

type StudentRepository interface {
	FindByFacilitatorID(facilitatorID int, page, limit int, sort, order string) ([]models.Student, int, error)
}
