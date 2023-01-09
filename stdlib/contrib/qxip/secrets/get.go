package secrets

import (
	"context"
	"os"

	"github.com/influxdata/flux/interpreter"
	"github.com/influxdata/flux/runtime"
	"github.com/influxdata/flux/values"
)

const GetKind = "get"

func init() {
	runtime.RegisterPackageValue("contrib/qxip/secrets", GetKind, GetFunc)
}

// GetFunc is a function that calls Get.
var GetFunc = makeGetFunc()

func makeGetFunc() values.Function {
	sig := runtime.MustLookupBuiltinType("contrib/qxip/secrets", "get")
	return values.NewFunction("get", sig, Get, false)
}

// Get retrieves the key variable for a given ENV.
func Get(ctx context.Context, args values.Object) (values.Value, error) {
	fargs := interpreter.NewArguments(args)
	key, err := fargs.GetRequiredString("key")
	if err != nil {
		return nil, err
	}

	value := os.Getenv(key)
	if len(value) < 1 {
		return nil, err
	}
	return values.NewString(value), nil
}
