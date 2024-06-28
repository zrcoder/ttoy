package generator

import (
	"github.com/google/uuid"
	"github.com/zrcoder/ttoy/util"
)

func UUID(input []byte) error {
	res := uuid.NewString()
	return util.Show([]byte(res))
}
