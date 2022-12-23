package gov

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/rlp"
)

func GetGovernParamValueWithDataBase(module, name string, blockNumber uint64, blockHash common.Hash, db snapshotdb.DB) (string, error) {
	paramValue, err := findGovernParamValueWithDataBase(module, name, blockHash, db)
	if err != nil {
		log.Error("get govern parameter value failed", "module", module, "name", name, "blockNumber", blockNumber, "blockHash", blockHash, "err", err)
		return "", err
	}
	if paramValue == nil {
		log.Error("govern parameter value is nil", "module", module, "name", name, "blockNumber", blockNumber, "blockHash", blockHash, "err", err)
		return "", UnsupportedGovernParam
	} else {
		if blockNumber >= paramValue.ActiveBlock {
			return paramValue.Value, nil
		} else {
			return paramValue.StaleValue, nil
		}
	}
}

func findGovernParamValueWithDataBase(module, name string, blockHash common.Hash, db snapshotdb.DB) (*ParamValue, error) {
	value, err := db.Get(blockHash, KeyParamValue(module, name))
	if snapshotdb.NonDbNotFoundErr(err) {
		return nil, err
	}

	if len(value) > 0 {
		var paramValue ParamValue
		if err := rlp.DecodeBytes(value, &paramValue); err != nil {
			return nil, err
		} else {
			return &paramValue, nil
		}
	}
	return nil, nil
}
