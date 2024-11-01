package dao

import "database/sql"

func (d *DAO) InsertBlacklistAddress(address string) error {
	var (
		tx  *sql.Tx
		err error
	)

	if tx, err = d.DB().Begin(); err != nil {
		return err
	}
	if _, err = tx.Exec("REPLACE INTO `blacklist_address` (`address`, `time`) VALUES(?, NOW())", address); err != nil {
		_ = tx.Rollback()
		return err
	} else {
		return tx.Commit()
	}
}

func (d *DAO) IsBlacklistAddress(address string) (bool, error) {
	if rows, err := d.DB().Query("SELECT COUNT(*) > 0 FROM `blacklist_address` WHERE `address`= ?", address); err != nil {
		return false, err
	} else {
		defer func(rows *sql.Rows) {
			_ = rows.Close()
		}(rows)

		var exists bool
		_ = rows.Next()
		if err := rows.Scan(&exists); err != nil {
			return false, err
		}
		return exists, nil
	}
}
