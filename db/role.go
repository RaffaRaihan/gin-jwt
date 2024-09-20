package db

type Roles struct{
	Id		int		`gorm:"primaryKey"`
	Role	string
}