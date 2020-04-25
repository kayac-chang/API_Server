package game

import (
	"api/model"
	"fmt"
)

// Store store game in db
func (it Repo) Store(game *model.Game) error {

	query := fmt.Sprintf(`
		INSERT INTO %s 
			(game_id, name, href, category, created_at, updated_at) 
		VALUES 
			(:game_id, :name, :href, :category, :created_at, :updated_at)
	`, table)

	_, err := it.db.NamedExec(query, game)

	return err
}

// Update update game in db
func (it Repo) Update(game *model.Game) error {

	query := fmt.Sprintf(`
		UPDATE %s 
		SET	
			game_id = :game_id, 
			name = :name, 
			href = :href, 
			category = :category, 
			created_at = :created_at, 
			updated_at = :updated_at 
		WHERE 
			game_id = :game_id
	`, table)

	_, err := it.db.NamedExec(query, game)

	return err
}
