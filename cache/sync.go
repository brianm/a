package cache

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Cache struct {
	db *sql.DB
}

func Open(path string) (Cache, error) {
	c := Cache{}
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return c, err
	}
	c.db = db
	c.migrate()
	return c, err
}

func (c Cache) migrate() error {
	stmt, err := c.db.Prepare(`SELECT name 
                                   FROM sqlite_master 
                                   WHERE type='table' 
                                     AND name=?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var name sql.NullString
	r := stmt.QueryRow("tasks")
	err = r.Scan(&name)
	if err == sql.ErrNoRows {
		// no rows!
		_, err := c.db.Exec(`create table tasks (
                               id integer primary key,
                               name text
                             )`)
		if err != nil {
			return err
		}
		log.Println("created tasks table")
	}
	if err != nil {
		return err
	}

	return nil
}
