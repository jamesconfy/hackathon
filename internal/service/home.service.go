package service

import "project-name/internal/se"

type HomeService interface {
	CreateHome() (string, *se.ServiceError)
}

type homeSrv struct{}

func (h *homeSrv) CreateHome() (string, *se.ServiceError) {
	return "You have gotten to the home route of project-name", nil
}

func NewHomeService() HomeService {
	return &homeSrv{}
}
