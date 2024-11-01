package dao

import "database/sql"

func (d *DAO) InsertAccessLog(remoteAddr *sql.NullString, method string, host string, path string, referer *sql.NullString, userAgent *sql.NullString, params *sql.NullString, blocked bool) error {
	var db *sql.DB = (*sql.DB)(d)

	_, err := db.Exec(
		"INSERT INTO `access_log` (`remote_address`, `method`, `host`, `path`, `referer`, `user_agent`, `params`, `blocked`, `time`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, now())",
		remoteAddr, method, host, path, referer, userAgent, params, blocked,
	)

	return err
}
