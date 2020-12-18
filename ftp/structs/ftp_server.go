package ftp

import (
	"reflect"
	"strconv"

	fto "github.com/jlaffaye/ftp"

	"github.com/eibrunorodrigues/ms-sender-go/utils"
)

type FtpAcess struct {
	Config *config
	Server *fto.ServerConn
}

type config struct {
	Server string
	User   string
	Pass   string
	Port   int
	Path   string
}

func (ftp *FtpAcess) Connect() error {
	if ftp.Config == nil || ftp.Config.User == "" {
		ftp.Config = &config{
			Server: utils.GetTypedEnvVariable("FTP_SERVER", "", reflect.String).(string),
			User:   utils.GetTypedEnvVariable("FTP_USER", "", reflect.String).(string),
			Pass:   utils.GetTypedEnvVariable("FTP_PASS", "", reflect.String).(string),
			Port:   utils.GetTypedEnvVariable("FTP_PORT", "", reflect.String).(int),
			Path:   utils.GetTypedEnvVariable("FTP_BASE_PATH", "", reflect.String).(string),
		}
	}

	client, err := fto.Dial(ftp.Config.Server + ":" + strconv.Itoa(ftp.Config.Port))
	if err != nil {
		return err
	}

	if err := client.Login(ftp.Config.User, ftp.Config.Pass); err != nil {
		return err
	}

	ftp.Server = client

	return nil
}
