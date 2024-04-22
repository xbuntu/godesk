package model

const TableNameConfig = "config"

type Config struct {
	ID    int    `gorm:"column:id;primaryKey" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	Field string `gorm:"column:field" json:"field"`
	Value string `gorm:"column:value" json:"value"`
}

func (*Config) TableName() string {
	return TableNameConfig
}
