package env

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

func AutoLoad() {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		return
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// ParseConfigFromEnv retrieves configurations from environmental variables and
// parse them into the given struct.
func ParseConfigFromEnv(config interface{}) error {
	configVal := reflect.ValueOf(config)
	if configVal.Kind() != reflect.Ptr {
		return errors.New("config must be a pointer")
	}

	if configVal.IsNil() {
		return errors.New("config can't be nil")
	}

	elem := configVal.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("config must be a struct")
	}

	numFields := elem.NumField()
	configType := elem.Type()

	for idx := 0; idx < numFields; idx++ {
		field := configType.Field(idx)
		envName, ok := field.Tag.Lookup("env")
		if !ok {
			continue
		}
		defaultVal := field.Tag.Get("default")
		envVal := getVar(envName, defaultVal)
		err := setFieldValue(field, elem.Field(idx), envVal)
		if err != nil {
			return err
		}
	}
	return nil
}

func setFieldValue(field reflect.StructField, fieldValue reflect.Value, newValue string) error {
	kind := field.Type.Kind()
	switch kind {
	case reflect.String:
		fieldValue.SetString(newValue)
		return nil
	case reflect.Int, reflect.Int64:
		num, err := parseInt(newValue, field.Type)
		if err != nil {
			return err
		}
		fieldValue.SetInt(num)
		return nil
	case reflect.Bool:
		boolean, err := strconv.ParseBool(newValue)
		if err != nil {
			return err
		}
		fieldValue.SetBool(boolean)
		return nil
	default:
		return fmt.Errorf("unexpected field type: %s", kind)
	}
}

func parseInt(newValue string, typeOfValue reflect.Type) (int64, error) {
	pkg, kind := typeOfValue.PkgPath(), typeOfValue.Name()
	switch {
	case kind == "int":
		num, err := strconv.Atoi(newValue)
		return int64(num), err
	case kind == "Duration":
		if pkg != "time" {
			return 0, errors.New("unknown package or kind")
		}
		duration, err := parseDuration(newValue)
		return int64(duration), err
	default:
		return 0, errors.New("unknown package or kind")
	}
}

func getVar(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
