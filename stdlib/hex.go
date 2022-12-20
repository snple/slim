package stdlib

import (
	"encoding/hex"

	"github.com/snple/slim"
)

var hexModule = map[string]slim.Object{
	"encode": &slim.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &slim.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
