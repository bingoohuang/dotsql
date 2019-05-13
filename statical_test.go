package dotsql_test

import (
	"dotsql"
	_ "dotsql/statical"
	"github.com/bingoohuang/statical/fs"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)

	staticalFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	dotsql.RegisterSqls(staticalFS)
}

func TestSql(t *testing.T) {
	sql := dotsql.Sql("/demo/FindUser")

	a := assert.New(t)
	a.Equal("select username, password, fullname, email, mobile, detail from typhon_admin_user where username = ?", sql)

	sql = dotsql.SqlByDbType("/demo/ClientLogSql", "xx")
	a.Equal(`insert into typhon_cur_client(app_id, conf_file, crc, ip, sync_time)
values (?, ?, ?, ?, ?)
on duplicate key update crc = ?, sync_time= ?`, sql)

	sql = dotsql.SqlByDbType("/demo/ClientLogSql", "postgres")
	a.Equal(`insert into typhon_cur_client(app_id, conf_file, crc, ip, sync_time)
values (?, ?, ?, ?, ?)
on conflict (app_id, conf_file, ip) do update set crc = ?, sync_time = ?`, sql)
}
