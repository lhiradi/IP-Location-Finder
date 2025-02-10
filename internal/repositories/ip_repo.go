package repositories

import (
	"ip-api/internal/models"

	"gorm.io/gorm"
)

type IPRepo struct {
	db *gorm.DB
}

func NewIPRepo(db *gorm.DB) IPRepo {
	return IPRepo{db: db}

}

// SaveIP saves the given IP data to the database
func (r *IPRepo) SaveIP(ip models.IP) error {
	result := r.db.Create(&ip)
	return result.Error
}

// GetIP retrieves the IP data from the database based on the provided IP address.
func (r *IPRepo) GetIP(ip string) (retrievedIP *models.IP, err error) {
	var dbResponse models.IP
	result := r.db.Where("ip = ?", ip).First(&dbResponse)
	return &dbResponse, result.Error
}
