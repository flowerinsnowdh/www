package service

import (
    "database/sql"
    "github.com/flowerinsnowdh/www/dao"
)

func (service *Service) LogAccess(remoteAddr string, host string, path string, referer string, userAgent string) error {
    var d *dao.DAO = (*dao.DAO)(service)

    return d.InsertAccessLog(
        &sql.NullString{
            String: remoteAddr,
            Valid:  remoteAddr != "",
        },
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
    )
}
