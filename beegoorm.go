package dotsql

import (
	"github.com/astaxie/beego/orm"
)

type Dot struct {
	Ormer  orm.Ormer
	DbType string
}

func MakeDotContext(o orm.Ormer, dbType string) Dot {
	return Dot{Ormer: o, DbType: dbType}
}

func (t Dot) Raw(id string, args ...interface{}) orm.RawSeter {
	sql := SqlByDbType(id, t.DbType)
	return t.Ormer.Raw(sql, args...)
}
