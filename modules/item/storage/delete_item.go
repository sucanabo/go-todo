package storage

import (
	"context"
	"todo/modules/item/model"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	itStatus := model.ItemStatus(model.ItemStatusDeleted)

	if err := s.db.Table(model.
		ToDoItem{}.TableName()).
		Where(cond).Updates(map[string]interface{}{
		"status": itStatus.String(),
	}).Error; err != nil {
		return err
	}
	return nil
}
