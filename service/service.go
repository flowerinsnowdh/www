package service

import "github.com/flowerinsnowdh/www/dao"

type Service dao.DAO

func (s *Service) DAO() *dao.DAO {
	return (*dao.DAO)(s)
}
