package main

import (
	"github.com/krogertechnology/data-loader/pkg/loader"
	"github.com/krogertechnology/krogo/pkg/krogo"
)

func main() {
	k := krogo.NewCMD()

	k.GET("", IMFHandler)

	k.Start()
}

func IMFHandler(c *krogo.Context) (i interface{}, err error) {
	// create new loader
	l, err := loader.New(c, "config.json")
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
