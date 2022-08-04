package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"os"
	"time"
)

func main() {
	k := krogo.NewCMD()

	k.GET("", func(c *krogo.Context) (interface{}, error) {
		var out interface{}
		for {
			_, err := c.Subscribe(&out)
			if err != nil {
				return nil, err
			}

			jsonData, err := json.Marshal(out)
			if err != nil {
				c.Logger.Error("Cannot marshal IMF record")
				return nil, err
			}

			buf := bytes.NewBuffer(jsonData)
			filename := fmt.Sprintf("output/IMF_%v.json", time.Now().Unix())

			err = os.WriteFile(filename, buf.Bytes(), 0677)
			if err != nil {
				c.Logger.Error("Error writing file")
				return nil, err
			}
		}

		return nil, nil
	})

	k.Start()
}
