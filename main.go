package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
)

const Corefile = `
. {
	debug
	errors
	whoami
	log

	etcd {{.Domain}} {
		stubzones
		path /
		endpoint {{.Etcd}}
	}

	proxy . {{.Upstream}}
}
`

const (
	EnvDomain      = "DDOG_DOMAIN"
	EnvEtcd        = "DDOG_ETCD_ENDPOINT"
	EnvUpStream    = "DDOG_UP_STREAM"
	EnvConfPath    = "DDOG_CONF_PATH"
	EnvMongo       = "DDOG_MONGO_ENDPOINT"
	EnvMongoName   = "DDOG_MONGO_NAME"
	EnvMongoPasswd = "DDOG_MONGO_PASSWD"
	EnvMongoDB     = "DDOG_MONGO_DB"
	EnvRegion      = "DDOG_REGION"
	EnvNotFound            = " Env Not Found!"
	EnvDomainNotFound      = "The " + EnvDomain + EnvNotFound
	EnvEtcdNotFound        = "The " + EnvEtcd + EnvNotFound
	EnvUpStreamNotFound    = "The " + EnvUpStream + EnvNotFound
	EnvMongoNotFound       = "The " + EnvMongo + EnvNotFound
	EnvMongoNameNotFound   = "The " + EnvMongoName + EnvNotFound
	EnvMongoPasswdNotFound = "The " + EnvMongoPasswd + EnvNotFound
	EnvMongoDBNotFound     = "The " + EnvMongoDB + EnvNotFound
	EnvRegionNotFound      = "The " + EnvRegion + EnvNotFound
)

func main(){
	if err := genConfigure(); err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
}

// genConfigure 生成DNS配置文件
func genConfigure() error {
	type conf struct {
		Domain   string
		Etcd     string
		Upstream string
	}

	name := os.Getenv(EnvDomain)
	if name == "" {
		return errors.New(EnvDomainNotFound)
	}

	etcd := os.Getenv(EnvEtcd)
	if etcd == "" {
		return errors.New(EnvEtcdNotFound)
	}

	etcdList := strings.Split(etcd, ";")
	for i, e := range etcdList {
		etcdList[i] = "http://" + e
	}

	etcd = strings.Join(etcdList, " ")
	path := os.Getenv(EnvConfPath)
	if path == "" {
		path = "/"
	}

	var mf = conf{
		Domain: name,
		Etcd:   etcd,
	}

	upstream := os.Getenv(EnvUpStream)
	if upstream == "" {
		return errors.New(EnvUpStreamNotFound)
	}

	us := strings.Split(upstream, ";")
	for _, s := range us {
		if !strings.Contains(s, ":") && !strings.Contains(s, "/") {
			mf.Upstream += s + ":53 "
		} else if !strings.Contains(s, ":") && strings.Contains(s, "/") {
			mf.Upstream += s + " "
		} else {
			mf.Upstream += s + " "
		}
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	t := template.Must(template.New("makefile").Parse(Corefile))

	file, err := os.Create(path + "corefile")
	if err != nil {
		return err
	}

	err = t.Execute(file, mf)
	if err != nil {
		return err
	}

	return nil
}
