package models

type Bill struct {
	ID      int64
	Address string
	Amount  int
	UserID  int64
	DueDate string
}
