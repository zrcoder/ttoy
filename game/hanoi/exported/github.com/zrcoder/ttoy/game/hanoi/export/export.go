// export by github.com/goplus/ixgo/cmd/qexp

package export

import (
	q "github.com/zrcoder/ttoy/game/hanoi/export"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "export",
		Path: "github.com/zrcoder/ttoy/game/hanoi/export",
		Deps: map[string]string{
			"github.com/zrcoder/ttoy/game/hanoi": "hanoi",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"A": reflect.ValueOf(q.A),
			"B": reflect.ValueOf(q.B),
			"C": reflect.ValueOf(q.C),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
