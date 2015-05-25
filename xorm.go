package dbsplitter

import (
	"fmt"

	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

type DbSplitterXorm struct{}

func (d *DbSplitterXorm) GetEngines(dbs []string) (map[string]*xorm.Engine, error) {
	engines := map[string]*xorm.Engine{}
	for _, db := range dbs {
		params := getConnectionStrings(db)
		engines[db] = d.getEngine(params)
	}
	return engines, nil
}

func (d *DbSplitterXorm) getEngine(params map[string]string) *xorm.Engine {
	dsn := d.parseParameter(params)
	db, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic("cannot get engines. params=" + dsn)
	}
	db.SetMapper(core.NewCacheMapper(new(core.GonicMapper)))
	return db
}

func (d *DbSplitterXorm) parseParameter(params map[string]string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", params["user"], params["pass"], params["host"], params["port"], params["dbname"])
}
