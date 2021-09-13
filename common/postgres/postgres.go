package postgres

import (
	"github.com/zhanjunjie2019/colago/common/conf"
	"github.com/zhanjunjie2019/colago/common/dbcli"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

func init() {
	ioc.AppendInjection(func() dbcli.Cli {
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
		postgres := new(dbcli.DbCli)
		postgres.SetWriter(writer)
		postgres.SetReader(readers)
		return postgres
	})
}

func build(dsn string) *gorm.DB {
	db, _ := gorm.Open(
		postgres.Open(dsn),
		new(gorm.Config),
	)
	return db
}
