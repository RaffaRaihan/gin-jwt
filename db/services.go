package db

type Services struct{
	Id				int		`gorm:"primaryKey"`
	Technician_id	int
	Client_id		int
	Ac_id			int
	Date			string
	Status			string
}