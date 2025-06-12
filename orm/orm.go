package orm

import (
	"log/slog"

	"github.com/rafakimanja/LittleLiteDB/controller"
	"github.com/rafakimanja/LittleLiteDB/services"
	"github.com/rafakimanja/LittleLiteDB/types"
)

type ORM[T any] struct{
	db string
}

var (
	control controller.DBController
	logger *slog.Logger
)

func New[T any](db string) *ORM[T]{
	logger = slog.Default()
	control = *controller.New()
	return &ORM[T]{db: db}
}

func (orm *ORM[T]) MigrateTable(table any){
	err := control.ConnectDB(orm.db)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = control.Migrate(table)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("migration successfully completed")
	}
}

func (orm *ORM[T]) Select(limit int, offset int, delete bool) ([]types.ResultModel[T], error){
	err := control.ConnectDB(orm.db)
	if err != nil {
		return nil, err
	}

	var rmodels []types.ResultModel[T]

	models, err := control.Select(limit, offset, delete)
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

func (orm *ORM[T]) SelectByID(id string, delete bool) (*types.ResultModel[T], error){
	err := control.ConnectDB(orm.db)
	if err != nil {
		return nil, err
	}

	model, err := control.SelectById(id, delete)
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
	err := control.ConnectDB(orm.db)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = control.Insert(data)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("data insert successful")
	}
}

func (orm *ORM[T]) Update(id string, data any){
	err := control.ConnectDB(orm.db)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = control.Update(id, data)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("update data successful")
	}
}

func (orm *ORM[T]) Delete(id string, delete bool) {
	err := control.ConnectDB(orm.db)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = control.Delete(id, delete)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("delete data successful")
	}
}