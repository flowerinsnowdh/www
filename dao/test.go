package dao

import "database/sql"

func (d *DAO) SQLInitTest() error {
    var db *sql.DB = (*sql.DB)(d)

    if rows, err := db.Query("SELECT 1"); err != nil {
        return err
    } else {
        _ = rows.Close()
        return nil
    }
}
