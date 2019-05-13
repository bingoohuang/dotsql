package dotsql_test

import (
	"dotsql"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	// 参数1   driverName
	// 参数2   数据库类型
	// 这个用来设置 driverName 对应的数据库类型
	// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", ":memory:")

	orm.Debug = true
}

// refer : https://beego.me/docs/mvc/model/rawsql.md
func TestBeegoOrmRaw(t *testing.T) {
	o1 := orm.NewOrm()

	dot := dotsql.MakeDotContext(o1, "sqlite")
	result, err := dot.Raw("/demo/createTeams").Exec()
	a := assert.New(t)
	a.Nil(err)
	a.NotNil(result)

	var p orm.RawPreparer
	p, err = dot.Raw("/demo/addTeam").Prepare()
	a.Nil(err)

	res, err := p.Exec("bingoo")
	a.Nil(err)
	affected, _ := res.RowsAffected()
	a.Equal(int64(1), affected)

	res, err = p.Exec("huang")
	a.Nil(err)
	affected, _ = res.RowsAffected()
	a.Equal(int64(1), affected)

	p.Close()

	type Team struct {
		Id   int
		Name string
	}

	var team Team
	err = dot.Raw("/demo/findTeam", 1).QueryRow(&team)
	a.Nil(err)
	a.Equal(Team{1, "bingoo"}, team)

	var teams []Team
	num, err := dot.Raw("/demo/selectTeams").QueryRows(&teams)
	a.Nil(err)
	a.Equal(int64(2), num)
	a.Equal([]Team{{1, "bingoo"}, {2, "huang"}}, teams)
}
