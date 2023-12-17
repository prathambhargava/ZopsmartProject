package models

type Book struct {
  ID        int64  `json:"id,omitempty"`
  Name      string `json:"name" validate:"required"`
  Author    string `json:"author" validate:"required"`
  StudentID int64  `json:"student_id" validate:"required"`
  Status    string `json:"status" validate:"required"`
}
