package address

type Address struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement;column:id_address"`
	Line1      string `json:"line1" gorm:"type:varchar(150);not null"`
	Line2      string `json:"line2" gorm:"type:varchar(150)"`
	City       string `json:"city" gorm:"type:varchar(100)"`
	State      string `json:"state" gorm:"type:varchar(100)"`
	Country    string `json:"country" gorm:"type:varchar(50)"`
	PostalCode string `json:"postal_code" gorm:"type:varchar(20)"`
}

func (Address) TableName() string { return "addresses" }
