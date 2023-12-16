package mock

import (
	"github.com/bubblenet/bubble/datavalidator/db"
	"github.com/bubblenet/bubble/datavalidator/types"
)

type DataCheck struct {
	*db.DB
	types.FilterMessage
}
