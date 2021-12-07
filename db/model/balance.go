package model

type Balance struct {
	BalanceId int64 `gorm:"primaryKey;autoIncrement"`
	Coin1     int
	Coin5     int
	Coin10    int
}

func (b *Balance) TableName() string {
	return "balance"
}

func init() {
	Migrations = append(Migrations, &Balance{})
}
