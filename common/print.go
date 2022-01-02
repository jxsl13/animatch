package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	Stdout io.Writer = os.Stdout
	Stderr io.Writer = os.Stderr
	Stdin  io.Writer = os.Stdin
)

func Println(a ...interface{}) error {
	buf := bytes.Buffer{}

	for _, v := range a {
		b, err := json.MarshalIndent(v, "", " ")
		if err != nil {
			return err
		}
		buf.Write(b)
		buf.WriteString("\n")
	}

	_, err := fmt.Fprintf(Stdout, "%s", buf.Bytes())
	return err
}

func Printf(format string, a ...interface{}) error {
	objects := make([]interface{}, 0, len(a))
	for _, v := range a {
		switch x := v.(type) {
		case string:
			objects = append(objects, x)
		case float64:
			objects = append(objects, FormatFloat64(x))
		case int:
			objects = append(objects, strconv.Itoa(x))
		default:
			b, err := json.MarshalIndent(v, "", " ")
			if err != nil {
				return err
			}
			objects = append(objects, string(b))
		}

	}

	_, err := fmt.Fprintf(Stdout, format, objects...)
	return err
}
