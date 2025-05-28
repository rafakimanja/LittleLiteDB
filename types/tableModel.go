package types

import (
	"errors"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID         string     `json:"id"`
	Content    any        `json:"content"`
	Created_At time.Time  `json:"created_at"`
	Updated_At time.Time  `json:"updated_at"`
	Deleted_At *time.Time `json:"deleted_at"`
}

func Init(content any) (*Model, error) {
	model := Model{
		ID:         uuid.New().String(),
		Created_At: time.Now(),
		Updated_At: time.Now(),
		Deleted_At: nil,
	}
	err := model.SetContent(content)
	if err != nil {
		return nil, err
	} else {
		return &model, nil
	}
}

func (m *Model) GetContent() any {
	return m.Content
}

func (m *Model) SetContent(content any) error {
	t := reflect.TypeOf(content)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return errors.New("content isn't valid struct")
	}
	m.Content = content
	return nil
}