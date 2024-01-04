package biz

import (
	"context"
	"todo/common"
	"todo/modules/item/model"
)

type ListItemStorage interface {
	GetItemList(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging, moreKey ...string,
	) ([]model.ToDoItem, error)
}

type itemListBiz struct {
	store ListItemStorage
}

func NewItemListBiz(store ListItemStorage) *itemListBiz {
	return &itemListBiz{store: store}
}

func (biz *itemListBiz) GetItemListWithFilter(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.ToDoItem, error) {
	data, err := biz.store.GetItemList(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil

}
