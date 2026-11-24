package main

import (
	envplus "github.com/mtb-pro/server-monitoring/src/pkg/go-plus/env"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.Info("Running server")

	checkEnvs()
	/*
		httpserver := httppr.GetInstance()

		if err := httpserver.Start(); err != nil {
			logrus.WithError(err).Panic("failed to start http server")
		}

		defer func() {
			if err := httpserver.Stop(); err != nil {
				logrus.WithError(err).Error("failed to stop http server")
			}
		}()

		grpcpr.Start()

		// wait for `Ctrl+c` or docker stop/restart signal
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGTERM)
		<-ch

		// Stop the application
		grpcpr.Stop()*/
}

func checkEnvs() {

	/* Full-Node stuff */

}
