package service

import (
	"database/sql"
	"github.com/flowerinsnowdh/www/dao"
)

func (service *Service) LogAccess(remoteAddr string, method string, host string, path string, referer string, userAgent string, params string, blocked bool) error {
	var d *dao.DAO = (*dao.DAO)(service)

	return d.InsertAccessLog(
		&sql.NullString{
			String: remoteAddr,
			Valid:  remoteAddr != "",
		},
		method,
		host,
		path,
		&sql.NullString{
			String: referer,
			Valid:  referer != "",
		},
		&sql.NullString{
			String: userAgent,
			Valid:  userAgent != "",
		},
		&sql.NullString{
			String: params,
			Valid:  params != "",
		},
		blocked,
	)
}
