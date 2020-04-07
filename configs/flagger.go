package configs

import (
	"github.com/spf13/viper"
	"reflect"
)

type FlaggerConfiguration struct {
	FlaggerHost string `json:"PENNANT_FLAGGER_HOST"`
	FlaggerPort string `json:"PENNANT_FLAGGER_PORT"`
}

func GetFlaggerConfig(vipe viper.Viper) FlaggerConfiguration {
	var newFlaggerConfiguration FlaggerConfiguration
	t := reflect.TypeOf(newFlaggerConfiguration)

	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get("json")

		if tag == "" { continue }
		v := reflect.ValueOf(&newFlaggerConfiguration).Elem().FieldByName(field.Name)
		if v.IsValid() {
			tagValue := vipe.GetString(tag)
			v.Set(reflect.ValueOf(tagValue))
		}
	}

	return newFlaggerConfiguration
}
