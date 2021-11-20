package admin

import (
	"minpro_arya/bussiness/admin"

	"gorm.io/gorm"
)

type Roles struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Admins struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
	Role     Roles  `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:RESTRICT;"`
}

func (rec *Admins) toDomain() admin.Domain {
	return admin.Domain{
		Id:        rec.ID,
		Username:  rec.Username,
		Password:  rec.Password,
		RoleID:    rec.RoleID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain admin.Domain) *Admins {
	return &Admins{
		Model: gorm.Model{
			ID:        domain.Id,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		Username: domain.Username,
		Password: domain.Password,
		RoleID:   domain.RoleID,
	}
}
