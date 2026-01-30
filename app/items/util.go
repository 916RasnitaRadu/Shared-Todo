package items

import "shared-todo/db"

func mapperItemDTO(item db.Item) db.ItemDTO {
	var itemDTO db.ItemDTO = db.ItemDTO{
		ID:          item.ID,
		CategoryID:  item.CategoryID.Int32,
		Name:        item.Name.String,
		Description: item.Description.String,
		Done:        item.Done.Bool,
		CreatedAt:   item.CreatedAt.Time,
	}

	return itemDTO
}
