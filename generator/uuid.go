package generator

import (
	"github.com/google/uuid"
	"github.com/zrcoder/ttoy/result"
)

func genUuid() error {
	res := uuid.NewString()
	result.Show(res)

	return nil
}
