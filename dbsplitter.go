package dbsplitter

import (
	"github.com/revel/revel"
	"log"
	"os"
	"strings"
)

var (
	conf *revel.MergedConfig
)

func loadConfig(filename string) {
	separator := "/"
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {
		separator = "/"
	}

	// add .conf extension
	if !strings.HasSuffix(filename, ".conf") {
		filename += ".conf"
	}
	c, err := revel.LoadConfig(revel.RunMode + separator + filename)
	if err != nil {
		// fallback default setting
		c, err = revel.LoadConfig(filename)
	}
	if err != nil {
		log.Fatalln("error load config: "+filename, err)
	}
	c.SetSection("database")
	conf = c
}

// spec情報を返却する
func getConnectionStrings(name string) map[string]string {
	params := getConnectionString(name + ".master")
	return params
}

// spec情報を返却する
func getConnectionString(name string) map[string]string {
	host := getDatabaseParamString(name, "host", "")
	port := getDatabaseParamString(name, "port", "3306")
	user := getDatabaseParamString(name, "user", "")
	pass := getDatabaseParamString(name, "password", "")
	dbname := getDatabaseParamString(name, "name", "")
	protocol := getDatabaseParamString(name, "protocol", "tcp")
	dbargs := getParamString("db.args", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return map[string]string{
		"user":     user,
		"pass":     pass,
		"protocol": protocol,
		"host":     host,
		"port":     port,
		"dbname":   dbname,
		"dbargs":   dbargs,
	}
}

func getDatabaseParamString(name string, param string, defaultValue string) string {
	value := getParamString("db."+name+"."+param, "")
	if value != "" {
		return value
	}
	value = getParamString("db.common."+param, defaultValue)

	if value != "" {
		return value
	}
	return ""
}

func getParamString(param string, defaultValue string) string {
	if conf == nil {
		loadConfig("database")
	}
	p, found := conf.String(param)
	if !found {
		return defaultValue
	}
	return p
}
