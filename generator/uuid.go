package generator

import (
	"github.com/google/uuid"
	"github.com/zrcoder/ttoy/util"
)

func UUID() error {
	res := uuid.NewString()
	util.Show(res)
	return nil
}
