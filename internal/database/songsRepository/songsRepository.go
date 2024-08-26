package songsRepository

import (
	"FindVibeGo/cmd/models"
	"database/sql"
)

var DB *sql.DB

func getAllSongsForUserId(userId string) ([]models.Song, error) {
	query := `SELECT * FROM SONGS WHERE user_id = ?`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.Id, &song.Title, &song.Artist, &song.Image, &song.Link, &song.UserId); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}
