package database

import (
	"time"

	"github.com/hackedu/backend/model"
)

const clubGetByIDStmt = `SELECT id, created, updated, school_id, name FROM
clubs WHERE id=$1`

const clubGetAllStmt = `SELECT id, created, updated, school_id, name FROM
clubs ORDER BY id`

const clubCreateStmt = `INSERT INTO clubs (created, updated, school_id, name)
VALUES ($1, $2, $3, $4) RETURNING id`

// GetClub gets a club from the database with the provided ID
func GetClub(id int64) (*model.Club, error) {
	c := model.Club{}
	row := db.QueryRow(clubGetByIDStmt, id)
	if err := row.Scan(&c.ID, &c.Created, &c.Updated, &c.SchoolID,
		&c.Name); err != nil {
		return nil, err
	}
	return &c, nil
}

// GetClubs gets all of the clubs from the database ordered by id.
func GetClubs() ([]*model.Club, error) {
	clubs := []*model.Club{}
	rows, err := db.Query(clubGetAllStmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := model.Club{}
		if err := rows.Scan(&c.ID, &c.Created, &c.Updated, &c.SchoolID,
			&c.Name); err != nil {
			return nil, err
		}

		clubs = append(clubs, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clubs, nil
}

// SaveClub saves the provided club to the database. If the club is a new club,
// then the club.Created field is set to the current time. The club.Updated
// field is set to the current time regardless.
func SaveClub(c *model.Club) error {
	if c.ID == 0 {
		c.Created = time.Now()
	}
	c.Updated = time.Now()

	row := db.QueryRow(clubCreateStmt, c.Created, c.Updated, c.SchoolID, c.Name)
	if err := row.Scan(&c.ID); err != nil {
		return err
	}

	return nil
}
