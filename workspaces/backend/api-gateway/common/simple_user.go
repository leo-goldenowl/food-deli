package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	LastName  string `json:"lastName" gorm:"column:last_name;"`
	FirstName string `json:"firstName" gorm:"column:first_name;"`
	Role      string `json:"role " gorm:"column:role;"`
	Avatar    *Image  `json:"avatar,omitempty" gorm:"column:avatar;"`
}

func (SimpleUser) TableName() string {
	return "users"
}
