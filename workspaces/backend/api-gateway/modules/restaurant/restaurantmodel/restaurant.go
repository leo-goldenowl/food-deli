package restaurantmodel

import (
	"errors"
	"strings"

	"api-gateway/common"

	"github.com/google/uuid"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;type:varchar(255);not null"`
	UserId          uuid.UUID          `json:"-" gorm:"column:owner_id;type:uuid;not null;"`
	Address         string             `json:"address" gorm:"column:address;type:varchar(255);not null"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false"`
	LikedCount      int                `json:"likedCount" gorm:"-"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name    *string        `json:"name" gorm:"column:name;"`
	Address *string        `json:"address" gorm:"column:address;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	OwnerId         uuid.UUID      `json:"-" gorm:"column:owner_id;"`
	Address         string         `json:"address" gorm:"column:address;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can not be blank")
	}

	return nil
}
