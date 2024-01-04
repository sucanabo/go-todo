package storage

import (
	"context"
	"todo/modules/item/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error {
	if err := s.db.Where(s.db, cond).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
