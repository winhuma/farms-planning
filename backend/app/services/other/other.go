package service

import "farms-planning/app/repo"

type other struct {
	r repo.Repo
}

func NewServices(r repo.Repo) Services {
	return &other{
		r: r,
	}
}

// ################### OTHER ###################
func (s *other) ProvinceGetAll() (result interface{}, err error) {
	return s.r.ProvinceGetAll()
}
