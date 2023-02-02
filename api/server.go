package api

import (
	"context"
	"net/http"
)

type Config struct {
	Host string
	Port int
	//Listenlimit
	//KeepAlive
	//ReadTimeout
	//WriteTimeout
	//ShutdownTimeout
	//AuthEnabled
	//ServerUUID
	//Username
	//Password
	//RBACEnabled
	//RBACPolicyFile
	//Path
	//Directory
	//TLSEnabled
	//TLSPort
	//CertPath
}

func NewServer(ctx context.Context, config *Config, handler http.Handler) (*http.Server, func(), error) {

}
