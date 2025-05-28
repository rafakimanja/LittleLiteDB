package orm

import (
	"littlelight/controller"
	"littlelight/services"
	"littlelight/types"
)

func SelectByID[T any](id string, db string) (*types.ResultModel[T], error){
	control := controller.New()
	control.ConnectDB(db)

	model, err := control.Select(id)
	if err != nil {
		return nil, err
	}

	result, err := services.ToResultModel[T](model)
	if err != nil {
		return nil, err
	}

	return result, nil
}