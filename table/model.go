package table

import (
	"errors"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID         string
	content    any
	Created_At time.Time
	Updated_At time.Time
	Deleted_At time.Time
}

func Init(content any) (*Model, error) {
	model := Model{
		ID: uuid.New().String(),
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}
	err := model.SetContent(content)
	if err != nil {
		return nil, err
	} else {
		return &model, nil
	}
}

func (m *Model) GetContent() any {
	return m.content
}

func (m *Model) SetContent(content any) error {
	t := reflect.TypeOf(content)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return errors.New("content isn't valid struct")
	}
	m.content = content
	return nil
}