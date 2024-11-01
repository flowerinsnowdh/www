package dao

import "database/sql"

type DAO sql.DB

func (d *DAO) DB() *sql.DB {
	return (*sql.DB)(d)
}
