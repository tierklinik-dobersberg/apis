package cli

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
)

type OutputFunc func(res any)

func (root *Root) defaultPrintFunc(res any) {
	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")

	if err := enc.Encode(res); err != nil {
		logrus.Fatal(err)
	}
	var blob []byte

	if root.OutputYAML {
		var err error
		blob, err = yaml.JSONToYAML(buf.Bytes())
		if err != nil {
			logrus.Fatal(err)
		}
	} else {
		blob = buf.Bytes()
	}

	if _, err := io.Copy(os.Stdout, bytes.NewReader(blob)); err != nil {
		logrus.Fatal(err)
	}
}
