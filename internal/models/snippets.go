package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	result, err := m.DB.Exec(
		`INSERT INTO snippets (title, content, created, expires)
		 VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`,
		title,
		content,
		expires,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), err
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	var s Snippet

	if err := m.DB.QueryRow(
		`SELECT id, title, content, created, expires
		 FROM snippets
		 WHERE expires > UTC_TIMESTAMP() AND id = ?`,
		id,
	).Scan(
		&s.ID,
		&s.Title,
		&s.Content,
		&s.Created,
		&s.Expires,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		}

		return Snippet{}, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	rows, err := m.DB.Query(
		`SELECT id, title, content, created, expires
		 FROM snippets
		 WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []Snippet

	for rows.Next() {
		var s Snippet

		if err := rows.Scan(
			&s.ID,
			&s.Title,
			&s.Content,
			&s.Created,
			&s.Expires,
		); err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
