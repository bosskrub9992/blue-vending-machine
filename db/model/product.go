package model

type Product struct {
	ProductID int64 `gorm:"primaryKey;autoIncrement"`
	Name      string
	Stock     int64
	Price     int64
}

func (p *Product) TableName() string {
	return "product"
}

func init() {
	Migrations = append(Migrations, &Product{})
}
