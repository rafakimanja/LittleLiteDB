package services

import (
	"encoding/json"
	"littlelight/types"
)

func ToResultModel[T any](model *types.Model) (*types.ResultModel[T], error) {
	bytes, err := json.Marshal(model.Content)
	if err != nil {
		return nil, err
	}

	var typed T
	err = json.Unmarshal(bytes, &typed)
	if err != nil {
		return nil, err
	}

	return &types.ResultModel[T]{
		ID:         model.ID,
		Content:    typed,
		Created_At: model.Created_At,
		Updated_At: model.Updated_At,
		Deleted_At: model.Deleted_At,
	}, nil
}

func ToModel(data any) (*types.Model, error){
	return types.Init(data)
}