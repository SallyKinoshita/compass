package persistence

import (
	"database/sql"
	"fmt"

	"github.com/SallyKinoshita/compass/internal/domain/models"
	"github.com/SallyKinoshita/compass/internal/domain/repositories"
)

type StudentRepo struct {
	DB *sql.DB
}

func NewStudentRepo(db *sql.DB) repositories.StudentRepository {
	return &StudentRepo{DB: db}
}

func (r *StudentRepo) FindByFacilitatorID(facilitatorID int, page, limit int, sort, order string) ([]models.Student, int, error) {
	offset := (page - 1) * limit
	query := `
		SELECT s.id, s.name, s.login_id, c.id, c.name
		FROM students s
		JOIN classrooms c ON s.classroom_id = c.id
		WHERE c.facilitator_id = ?
		ORDER BY s.` + sort + ` ` + order + `
		LIMIT ? OFFSET ?`
	rows, err := r.DB.Query(query, facilitatorID, limit, offset)
	fmt.Println(err)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.LoginID, &student.Classroom.ID, &student.Classroom.Name); err != nil {
			return nil, 0, err
		}
		students = append(students, student)
	}

	var totalCount int
	err = r.DB.QueryRow(`SELECT COUNT(*) FROM students s JOIN classrooms c ON s.classroom_id = c.id WHERE c.facilitator_id = ?`, facilitatorID).Scan(&totalCount)
	return students, totalCount, err
}
