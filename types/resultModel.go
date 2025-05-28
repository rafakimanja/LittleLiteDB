package types

import "time"

type ResultModel[T any] struct {
	ID         string     `json:"id"`
	Content    T          `json:"content"`
	Created_At time.Time  `json:"created_at"`
	Updated_At time.Time  `json:"updated_at"`
	Deleted_At *time.Time `json:"deleted_at"`
}