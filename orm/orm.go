package orm

import (
	"littlelight/controller"
	"littlelight/services"
	"littlelight/types"
)

type ORM[T any] struct{
	db string
}

func New[T any](db string) *ORM[T]{
	return &ORM[T]{db: db}
}

func (orm *ORM[T]) SelectByID(id string) (*types.ResultModel[T], error){
	control := controller.New()
	control.ConnectDB(orm.db)

	model, err := control.SelectById(id, false)
	if err != nil {
		return nil, err
	}

	result, err := services.ToResultModel[T](model)
	if err != nil {
		return nil, err
	}

	return result, nil
}