package common

import (
	"net/url"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ImageSource interface {
	GetName() (string)
	GetKVs() (*map[string]string)
	GetFileName() (string)
	GetSHA512() (string)
	GetData() (*[]byte)
	GetURL() (*url.URL)
	Setup(log *logrus.Logger, cfg *viper.Viper) (error)
	Run() error
}
