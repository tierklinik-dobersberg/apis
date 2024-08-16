package cli

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type OutputFunc func(res any)

func (root *Root) defaultPrintFunc(res any) {
	if fm := root.Config().Format; fm != "" {
		if strings.HasPrefix(fm, "@") {
			path := strings.TrimPrefix(fm, "@")

			content, err := os.ReadFile(path)
			if err != nil {
				logrus.Fatalf("failed to read format from file %s: %q", path, err)
			}

			fm = string(content)
		}

		tmp, err := template.New("").Parse(fm)
		if err != nil {
			logrus.Fatalf("failed to parse --format template: %s", err)
		}

		if err := tmp.Execute(os.Stdout, res); err != nil {
			logrus.Fatalf("failed to execute --format template: %s", err)
		}

		return
	}

	var buf []byte

	if msg, ok := res.(proto.Message); ok {
		var err error
		opts := protojson.MarshalOptions{
			Multiline: true,
			Indent:    "  ",
		}

		buf, err = opts.Marshal(msg)
		if err != nil {
			logrus.Fatal(err.Error())
		}
	} else {
		var b = new(bytes.Buffer)

		enc := json.NewEncoder(b)
		enc.SetIndent("", "  ")

		if err := enc.Encode(res); err != nil {
			logrus.Fatal(err)
		}

		buf = b.Bytes()
	}

	var blob []byte

	if root.Config().OutputYAML {
		var err error
		blob, err = yaml.JSONToYAML(buf)
		if err != nil {
			logrus.Fatal(err)
		}
	} else {
		blob = buf
	}

	if _, err := io.Copy(os.Stdout, bytes.NewReader(blob)); err != nil {
		logrus.Fatal(err)
	}
}
