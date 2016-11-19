package inject

import (
	"github.com/facebookgo/inject"
)

//Injector tool
func Injector(params ...interface{}) (err error) {
	var g inject.Graph
	for _, v := range params {
		if err = g.Provide(&inject.Object{Value: v}); err != nil {
			return
		}
	}
	err = g.Populate()
	return
}
