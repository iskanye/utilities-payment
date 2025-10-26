package models

import "time"

type Bill struct {
	ID      int64
	Address string
	Amount  int
	DueDate time.Time
}
