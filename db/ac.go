package db

type Ac struct{
	Id		int			`gorm:"primaryKey"`
	Nama_Ac	string
	Brand	string
	Pk		float64
	Price	int
}