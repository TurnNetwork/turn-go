package rpc

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/datavalidator/db"
	"github.com/bubblenet/bubble/datavalidator/mock"
	"github.com/bubblenet/bubble/datavalidator/p2p"
	"github.com/bubblenet/bubble/datavalidator/types"
	"github.com/bubblenet/bubble/rpc"
)

type DataValidatorRpc struct {
	db      *db.DB
	network *p2p.Network
}

func NewDataValidatorRpc(db *db.DB, network *p2p.Network) *DataValidatorRpc {
	return &DataValidatorRpc{
		db:      db,
		network: network,
	}
}

func (d *DataValidatorRpc) Api() rpc.API {
	return rpc.API{
		Namespace: "datavalidator",
		Version:   "1.0.0",
		Service: &DataValidatorApi{
			db:      d.db,
			network: d.network,
		},
		Public: true,
	}
}

type DataValidatorApi struct {
	db      *db.DB
	network *p2p.Network
}

type DataValidatorStatus struct {
	Block uint64
	Peers []*types.PeerInfo
}

func (d *DataValidatorApi) RangeNonce(chainId, startNonce uint64, limit uint64) []*types.QuorumLog {
	details, err := d.db.GetQuorumLogRangeNonce(chainId, startNonce, limit)
	if err != nil {
		return nil
	}
	var res []*types.QuorumLog
	for _, d := range details {
		res = append(res, &types.QuorumLog{
			Log:        d.Log,
			Signatures: d.Signatures,
		})
	}
	return res
}

func (d *DataValidatorApi) LogByTransaction(hash common.Hash) []*types.QuorumLog {
	details, err := d.db.GetQuorumLogByTxHash(hash)
	if err != nil {
		return nil
	}
	var res []*types.QuorumLog
	for _, d := range details {
		res = append(res, &types.QuorumLog{
			Log:        d.Log,
			Signatures: d.Signatures,
		})
	}
	return res
}

func (d *DataValidatorApi) Status() *DataValidatorStatus {
	scanBlock, err := d.db.GetScanLog()
	if err != nil {
		return nil
	}
	return &DataValidatorStatus{
		Block: scanBlock,
		Peers: d.network.PeersInfo(),
	}
}

func (d *DataValidatorApi) TestAddLog(chainId, limit uint64) bool {
	mock.SystemBlockFilter.AddMessagePublished(chainId, int(limit))
	return true
}
