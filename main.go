package main

import (
	"github.com/eibrunorodrigues/ms-sender-go/ftp"
)

func main() {
	ftpClient := ftp.FtpAcess{}
	ftpClient.Connect()
	err := ftpClient.Server.ChangeDir(ftpClient.Config.Path)
	if err != nil {
		panic(err)
	}
}
