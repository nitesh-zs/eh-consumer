package main

import (
	"github.com/krogertechnology/data-loader/pkg/loader"
	"github.com/krogertechnology/krogo/pkg/krogo"
)

func main() {
	k := krogo.NewCMD()

	k.GET("load", load)

	k.Start()
}

func load(c *krogo.Context) (i interface{}, err error) {
	configPath := c.PathParam("config")

	// create new loader
	l, err := loader.New(c, configPath)
	if err != nil {
		return "Failed", err
	}

	// start loader
	if err := l.Start(c); err != nil {
		return "Failed", err
	}

	c.Logger.Info("Data Loader Completed Successfully")

	return "Done!", nil
}
