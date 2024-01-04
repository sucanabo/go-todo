package storage

import (
	"context"
	"todo/common"
	"todo/modules/item/model"
)

func (s *sqlStore) GetItemList(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]model.ToDoItem, error) {
	var result []model.ToDoItem

	ss := model.ItemStatus(model.ItemStatusDeleted)
	db := s.db.Where("status <> ?", ss.String())

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}
	if err := db.Table(model.ToDoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
