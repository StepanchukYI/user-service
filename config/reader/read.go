package reader

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	viperDefaultDelimiter = "."
	defaultTagName        = "default"
	squashTagValue        = ",squash"
	mapStructureTagName   = "mapstructure"
)

func Read(config interface{}, opts ...viper.DecoderConfigOption) error {
	viperLogger := viper.New()
	viperLogger.SetEnvKeyReplacer(strings.NewReplacer(viperDefaultDelimiter, "_")) // replace default viper delimiter for env vars
	viperLogger.AutomaticEnv()
	viperLogger.SetTypeByDefaultValue(true)
	err := setDefaults("", viperLogger, reflect.StructField{}, reflect.ValueOf(config).Elem())
	if err != nil {
		return errors.WithMessage(err, "failed to apply defaults")
	}
	err = viperLogger.Unmarshal(config, opts...)
	if err != nil {
		return errors.WithMessage(err, "failed to parse configuration")
	}

	return nil
}

// setDefaults sets default values for struct fields based using value from default tag
func setDefaults(parentName string, vip *viper.Viper, structField reflect.StructField, val reflect.Value) error {
	if val.Kind() == reflect.Struct {
		value, ok := structField.Tag.Lookup(mapStructureTagName)
		if ok && value != squashTagValue {
			if parentName != "" {
				parentName += viperDefaultDelimiter
			}
			parentName += strings.ToUpper(value)
		}
		for i := 0; i < val.NumField(); i++ {
			if err := setDefaults(parentName, vip, val.Type().Field(i), val.Field(i)); err != nil {
				return err
			}
		}

		return nil
	}
	value, _ := structField.Tag.Lookup(defaultTagName)
	fieldName, ok := structField.Tag.Lookup(mapStructureTagName)
	if ok && fieldName != squashTagValue {
		if parentName != "" {
			fieldName = parentName + viperDefaultDelimiter + strings.ToUpper(fieldName)
		}
		vip.SetDefault(strings.ToUpper(fieldName), value)
	}

	return nil
}
