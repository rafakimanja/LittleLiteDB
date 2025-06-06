package orm

import (
	"fmt"
	"littlelight/controller"
	"littlelight/services"
	"littlelight/types"
)

type ORM[T any] struct{
	db string
}

var (
	control controller.DBController
)

func New[T any](db string) *ORM[T]{
	control = *controller.New()
	return &ORM[T]{db: db}
}

func (orm *ORM[T]) Select(limit int, offset int) ([]types.ResultModel[T], error){
	control.ConnectDB(orm.db)
	var rmodels []types.ResultModel[T]

	models, err := control.Select(limit, offset, false)
	if err != nil {
		return nil, err
	}

	for _, item := range(models){
		result, err := services.ToResultModel[T](&item)
		if err != nil {
			return nil, err
		}
		rmodels = append(rmodels, *result)
	}
	return rmodels, nil
}

func (orm *ORM[T]) SelectByID(id string) (*types.ResultModel[T], error){
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

func (orm *ORM[T]) Insert(data any){
	control.ConnectDB(orm.db)

	err := control.Insert(data)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("data insert successful")
	}
}

func (orm *ORM[T]) Update(id string, data any){
	control.ConnectDB(orm.db)

	err := control.Update(id, data)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("update data successful")
	}
}

func (orm *ORM[T]) Delete(id string, delete bool) {
	control.ConnectDB(orm.db)

	err := control.Delete(id, delete)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("delete data successful")
	}
}