package services

import (
	"ip-api/internal/models"
	"ip-api/internal/repositories"
	"log"
)

type IPServ struct {
	ipRepo repositories.IPRepo
}

func NewIPServ(ipRepo repositories.IPRepo) IPServ {
	return IPServ{ipRepo: ipRepo}
}

func (s *IPServ) SaveIPService(ip models.IP) error {
	err := s.ipRepo.SaveIP(ip)
	if err != nil {
		log.Println("Couldn't save IP! something went wrong.")
		return err
	}
	return nil
}

func (s *IPServ) GetIPService(ip string) (retrivedIP *models.IP, err error) {
	rIP, err := s.ipRepo.GetIP(ip)
	if err != nil {
		log.Println("Couldn't find " + ip + " in database!")
		return nil, err
	}
	return rIP, nil
}
