package model

type AbsPO interface {
	TableName() string
}

type PostgresEntity struct {
	ID         uint64 `gorm:"primarykey;AUTO_INCREMENT;comment:ID"`
	CreateTime int64  `gorm:"autoCreateTime:milli;comment:创建时间"`
	UpdateTime int64  `gorm:"autoUpdateTime:milli;comment:修改时间"`
	Deleted    uint8  `gorm:"type:int2;default:0;comment:删除标志位"`
}
