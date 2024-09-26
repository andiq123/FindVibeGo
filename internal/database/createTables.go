package database

func (s *service) CreateTables() error {
	err := createUsersTable()
	if err != nil {
		return err
	}

	err = createSongsTable()
	if err != nil {
		return err
	}
	return nil
}

func createUsersTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL
			);`
	_, err := dbInstance.db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func createSongsTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS songs (
			id UUID PRIMARY KEY,
			artist VARCHAR(255) NOT NULL,
			title VARCHAR(255) NOT NULL,
			image VARCHAR(255),
			link VARCHAR(255),
			song_order INT,
			userId UUID,
			FOREIGN KEY (userId) REFERENCES users(id)
		);`

	_, err := dbInstance.db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}
