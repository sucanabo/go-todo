package biz

import (
	"context"
	"todo/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.ToDoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error
}

type updateItemBiz struct {
	store UpdateItemStorage
}

func NewUpdateItemBiz(store UpdateItemStorage) *updateItemBiz {

	return &updateItemBiz{store}
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrItemDeleted
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}

	return nil
}
