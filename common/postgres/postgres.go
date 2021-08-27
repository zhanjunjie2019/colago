package postgres

import (
	"e.coding.net/double-j/ego/colago/common/conf"
	"e.coding.net/double-j/ego/colago/common/dbcli"
	"e.coding.net/double-j/ego/colago/common/ioc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

func init() {
	_ = ioc.InjectSimpleBean(new(Postgres))
}

type Postgres struct {
	*dbcli.DbCli
}

func (p *Postgres) New() ioc.AbsBean {
	writer := build(conf.ConfigMap("writer.dsn"))
	readerDsns := conf.ConfigMap("readers.dsn")
	readers := make([]*gorm.DB, 0)
	if readerDsns == "" {
		readers = append(readers, writer)
	} else {
		for _, v := range strings.Split(readerDsns, ",") {
			readers = append(readers, build(v))
		}
	}
	p.DbCli = &dbcli.DbCli{
		Writer:  writer,
		Readers: readers,
	}
	return p
}

func build(dsn string) *gorm.DB {
	db, _ := gorm.Open(
		postgres.Open(dsn),
		new(gorm.Config),
	)
	return db
}
