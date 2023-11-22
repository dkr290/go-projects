package main

import (
	"azure-vmss/pkg"
	"azure-vmss/vmss"
	"flag"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var sshk pkg.SshKeys

func GenSSH() {

	if err := sshk.GetToken(); err != nil {
		log.Fatalln("Error generation the token", err)
	}

}

func main() {

	flag.String("envfile", "", "Pass env file name")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	envFl := viper.GetString("envfile")

	parameters, err := pkg.GetEnvs(envFl)

	if err != nil {
		log.Fatal(err)
	}
	GenSSH()

	if err := vmss.CreateVmss(parameters.Context, parameters, sshk); err != nil {
		log.Fatalln("Launch instance error", err)
	}

}
