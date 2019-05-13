package dotsql

import (
	"github.com/bingoohuang/statical/fs"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

func LoadSqls(fs *fs.StaticalFS) map[string]string {
	sqlmap := make(map[string]string)

	for filename, file := range fs.Files {
		ext := filepath.Ext(filename)
		ext = strings.ToLower(ext)
		if ext != ".sql" {
			continue
		}

		sqlFileConetent := string(file.Data)
		sqls, err := LoadFromString(sqlFileConetent)
		if err != nil {
			logrus.Warnf("fail to load sqls from file %s, error %v", filename, err)
			continue
		}

		var pkg = filename[0 : len(filename)-4]

		for k, v := range sqls.QueryMap() {
			sqlKey := pkg + "/" + k
			if _, ok := sqlmap[sqlKey]; ok {
				logrus.Warnf("sqlid %s duplicated found in file %s, it will be overwritten", sqlKey, filename)
			}
			logrus.Debugf("load sqlid %s from file %s", sqlKey, filename)
			sqlmap[sqlKey] = v
		}
	}

	return sqlmap
}

var gSqlmap map[string]string

func RegisterSqls(fs *fs.StaticalFS) {
	gSqlmap = LoadSqls(fs)
}

func Sql(key string) string {
	if raw, ok := gSqlmap[key]; ok {
		return strings.TrimRight(raw, ";")
	}

	panic("unable to find sql by name " + key)
}

func SqlByDbType(key, dbType string) string {
	if raw, ok := gSqlmap[key+"-"+dbType]; ok {
		return strings.TrimRight(raw, ";")
	}

	return Sql(key)
}
