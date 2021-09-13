package dbcli

import (
	"e.coding.net/double-j/ego/colago/common/model"
	"gorm.io/gorm"
	"sync/atomic"
)

type Cli interface {
	AutoMigrate(pos ...model.AbsPO) error
	InsertOne(po model.AbsPO) error
	FindOne(po model.AbsPO, query []string, args []interface{}) error
	FindList(po model.AbsPO, dest interface{}, query []string, args []interface{}) error
	GetWriter() *gorm.DB
	SetWriter(*gorm.DB)
	GetReader() *gorm.DB
	SetReader([]*gorm.DB)
}

type DbCli struct {
	ops     uint32
	writer  *gorm.DB
	readers []*gorm.DB
}

func (cli *DbCli) AutoMigrate(pos ...model.AbsPO) error {
	for i := 0; i < len(pos); i++ {
		err := cli.GetWriter().Scopes(func(tx *gorm.DB) *gorm.DB {
			return tx.Table(pos[i].TableName())
		}).AutoMigrate(pos[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (cli *DbCli) InsertOne(po model.AbsPO) error {
	return cli.GetWriter().Scopes(func(tx *gorm.DB) *gorm.DB {
		return tx.Table(po.TableName())
	}).Create(po).Error
}

func (cli *DbCli) FindOne(po model.AbsPO, query []string, args []interface{}) error {
	db := cli.GetReader().Scopes(func(tx *gorm.DB) *gorm.DB {
		return tx.Table(po.TableName())
	})
	for i, q := range query {
		db = db.Where(q, args[i])
	}
	return db.First(po).Error
}

func (cli *DbCli) FindList(po model.AbsPO, dest interface{}, query []string, args []interface{}) error {
	db := cli.GetReader().Scopes(func(tx *gorm.DB) *gorm.DB {
		return tx.Table(po.TableName())
	})
	for i, q := range query {
		db = db.Where(q, args[i])
	}
	return db.Find(dest).Error
}

func (cli *DbCli) SetWriter(db *gorm.DB) {
	cli.writer = db
}

func (cli *DbCli) SetReader(dbs []*gorm.DB) {
	cli.readers = dbs
}

func (cli *DbCli) GetWriter() *gorm.DB {
	return cli.writer
}

func (cli *DbCli) GetReader() *gorm.DB {
	opsFinal := atomic.LoadUint32(&cli.ops)
	u := opsFinal % uint32(len(cli.readers))
	if opsFinal > 999 {
		atomic.StoreUint32(&cli.ops, 0)
	} else {
		atomic.AddUint32(&cli.ops, 1)
	}
	return cli.readers[u]
}
