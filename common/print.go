package common

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func Println(cmd *cobra.Command, a ...interface{}) error {
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

	_, err := fmt.Fprintln(cmd.OutOrStdout(), objects...)
	return err
}

func Printf(cmd *cobra.Command, format string, a ...interface{}) error {
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

	_, err := fmt.Fprintf(cmd.OutOrStdout(), format, objects...)
	return err
}
