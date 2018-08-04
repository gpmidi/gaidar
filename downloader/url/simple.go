package url

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	url2 "net/url"
	"net/http"
	"io/ioutil"
	"crypto/sha512"
	"fmt"
)

type URLImageSource struct {
	url       string
	urlParsed *url2.URL
	sha512    string
	filename  string
	data      *[]byte
	kvs       *map[string]string
	log       *logrus.Logger
	cfg       *viper.Viper
}

func (uis URLImageSource) GetName() (string) {
	return "URLImageSource"
}

func (uis URLImageSource) GetKVs() (*map[string]string) {
	return uis.kvs
}

func (uis URLImageSource) GetFileName() (string) {
	return uis.filename
}

func (uis URLImageSource) GetSHA512() (string) {
	return uis.sha512
}

func (uis URLImageSource) GetData() (*[]byte) {
	return uis.data
}

func (uis URLImageSource) GetURL() (*url2.URL) {
	return uis.urlParsed
}

func (uis URLImageSource) Setup(log *logrus.Logger, cfg *viper.Viper) (error) {
	uis.log = log
	uis.cfg = cfg
	uis.kvs = &map[string]string{}
	return nil
}

func (uis URLImageSource) SetURL(url string) {
	uis.url = url
}

func (uis URLImageSource) download() error {
	res, err := http.Get(uis.url)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	uis.data = &data
	return nil
}

func (uis URLImageSource) hash() error {
	h := sha512.New()
	_, err := h.Write(*uis.data)
	if err != nil {
		return err
	}
	uis.sha512 = fmt.Sprintf("%x", h.Sum(nil))
	return nil
}

func (uis URLImageSource) Run() error {
	ret, err := url2.Parse(uis.url)
	if err != nil {
		return err
	}
	uis.urlParsed = ret

	err = uis.download()
	if err != nil {
		return err
	}

	err = uis.hash()
	if err != nil {
		return err
	}

	return nil
}

func NewURLImageSource(url string, log *logrus.Logger, cfg *viper.Viper) (*URLImageSource, error) {
	ret := URLImageSource{}
	err := ret.Setup(log, cfg)
	if err != nil {
		return nil, err
	}
	ret.SetURL(url)
	return &ret, nil
}
