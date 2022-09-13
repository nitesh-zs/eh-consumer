package main

import (
	"encoding/json"
	"github.com/krogertechnology/data-loader/pkg/loader"
	"github.com/krogertechnology/data-loader/pkg/models"
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

	_, _ = l.AddConfigLayerWithNeeds(write, 1)

	// start loader
	if err := l.Start(c); err != nil {
		return "Failed", err
	}

	c.Logger.Info("Data Loader Completed Successfully")

	return "Done!", nil
}

func write(c *krogo.Context, record models.Record) (rec *models.Record, err error) {
	rawData := record.Value["raw_payload"].(string)

	data := make(map[string]interface{})

	err = json.Unmarshal([]byte(rawData), &data)
	if err != nil {
		c.Logger.Error(err)
		return nil, nil
	}

	record.Value["data"] = data["data"]
	record.Value["eventHeader"] = data["eventHeader"]

	return &record, nil
}
