package models

type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LoginID   string `json:"loginId"`
	Classroom Classroom
}
