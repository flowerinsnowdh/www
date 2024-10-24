package dao

import "database/sql"

func (d *DAO) InsertAccessLog(remoteAddr *sql.NullString, path string, referer *sql.NullString, userAgent *sql.NullString) error {
    var db *sql.DB = (*sql.DB)(d)

    _, err := db.Exec(
        "INSERT INTO `access_log` (`remote_address`, `path`, `referer`, `user_agent`, `time`) VALUES (?, ?, ?, ?, now()",
        remoteAddr, path, referer, userAgent,
    )

    return err
}
