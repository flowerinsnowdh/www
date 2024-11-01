package service

func (s *Service) AddToBlacklist(address string) error {
	return s.DAO().InsertBlacklistAddress(address)
}

func (s *Service) IsBlacklistAddress(address string) (bool, error) {
	return s.DAO().IsBlacklistAddress(address)
}
