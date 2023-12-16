package db

import (
	"encoding/binary"
	"fmt"
	"github.com/bubblenet/bubble/core/rawdb"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/datavalidator/types"
	"github.com/bubblenet/bubble/ethdb"
	"github.com/bubblenet/bubble/rlp"
)

const (
	ScanLogKey   = "dv_scanlog"
	DetailPrefix = "dv_detail"
	SignPrefix   = "dv_sign"
	SignIndex    = "dv_sign_hash"
	UnSignPrefix = "dv_unsign"
	UnSignIndex  = "dv_unsign_hash"
)

func Uint64toBytes(num uint64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], num)
	return buf[:]
}

func makeDetailKey(id common.Hash) string {
	return fmt.Sprintf("%s_%s", DetailPrefix, id.Hex())
}
func makeDetailKeyPrefix() string {
	return fmt.Sprintf("%s_", DetailPrefix)
}

func makeUnSignChainIdNonceIndexKey(chainId, nonce uint64) string {
	return fmt.Sprintf("%s_%x_%x", UnSignPrefix, Uint64toBytes(chainId), Uint64toBytes(nonce))
}
func makeUnSignChainIdNonceIndexKeyPrefix(chainId uint64) string {
	return fmt.Sprintf("%s_%x_", UnSignPrefix, Uint64toBytes(chainId))
}

func makeUnSignHashSequenceIndexKey(txHash common.Hash, sequence uint64) string {
	return fmt.Sprintf("%s_%s_%x", UnSignIndex, txHash.Hex(), Uint64toBytes(sequence))
}
func makeUnSignHashSequenceIndexKeyPrefix(txHash common.Hash) string {
	return fmt.Sprintf("%s_%s_", UnSignIndex, txHash.Hex())
}
func makeUnSignHashRangePrefix() string {
	return fmt.Sprintf("%s_", UnSignIndex)
}
func makeSignChainIdNonceIndexKey(chainId, nonce uint64) string {
	return fmt.Sprintf("%s_%x_%x", SignPrefix, Uint64toBytes(chainId), Uint64toBytes(nonce))
}
func makeSignChainIdNonceIndexKeyPrefix(chainId uint64) string {
	return fmt.Sprintf("%s_%x_", SignPrefix, Uint64toBytes(chainId))
}
func makeSignHashSequenceIndexKey(txHash common.Hash, sequence uint64) string {
	return fmt.Sprintf("%s_%s_%x", SignIndex, txHash.Hex(), Uint64toBytes(sequence))
}
func makeSignHashSequenceIndexKeyPrefix(txHash common.Hash) string {
	return fmt.Sprintf("%s_%s_", SignIndex, txHash.Hex())
}

type DB struct {
	db ethdb.Database
}

func NewLevelDbDataValidatorDB(path string) (*DB, error) {
	db, err := rawdb.NewLevelDBDatabase(path, 0, 0, "")
	if err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}
func NewDataValidatorDB(db ethdb.Database) *DB {
	return &DB{
		db: db,
	}
}

func (d *DB) StoreScanLog(block uint64) error {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], block)
	return d.db.Put([]byte(ScanLogKey), buf[:])
}

func (d *DB) GetScanLog() (uint64, error) {
	buf, err := d.db.Get([]byte(ScanLogKey))
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(buf), nil
}

//func (d *DB) GetStoreQuorumLog(id common.Hash) (*types.QuorumLog, error) {
//	d.logCache.Get()
//	buf, err := d.db.Get([]byte(makeDetailKey(id)))
//	if err != nil {
//		return nil, err
//	}
//	var log types.MessagePublishedDetail
//	err = rlp.DecodeBytes(buf, &log)
//	if err != nil {
//		return nil, err
//	}
//	return &types.QuorumLog{
//		MessagePublished: log.Log,
//		Signatures:       log.Signatures,
//	}, err
//}

func (d *DB) SetQuorumLog(logs []*types.MessagePublishedDetail) error {
	batch := d.db.NewBatch()
	for _, log := range logs {
		buf, err := log.Bytes()
		if err != nil {
			return err
		}
		//存储原始数据
		detailKey := []byte(makeDetailKey(log.Log.Hash()))
		batch.Put(detailKey, buf)
		d.deleteUnSignIndex(batch, log)
		//构建索引
		batch.Put([]byte(makeSignChainIdNonceIndexKey(log.Log.ChainId, log.Log.Nonce)), detailKey)
		batch.Put([]byte(makeSignHashSequenceIndexKey(log.TxHash, log.Log.Sequence)), detailKey)
	}
	return batch.Write()
}
func (d *DB) GetMessagePublished(id common.Hash) (*types.MessagePublishedDetail, error) {
	buf, _ := d.db.Get([]byte(makeDetailKey(id)))

	if buf == nil {
		return nil, nil
	}
	var log *types.MessagePublishedDetail
	err := rlp.DecodeBytes(buf, &log)
	if err != nil {
		return nil, err
	}
	return log, nil
}
func (d *DB) GetQuorumChainIdNonce(chainId, startNonce uint64) (*types.MessagePublishedDetail, error) {
	key := makeSignChainIdNonceIndexKey(chainId, startNonce)
	detailkey, _ := d.db.Get([]byte(key))
	if detailkey == nil {
		return nil, nil
	}
	value, _ := d.db.Get(detailkey)
	if value == nil {
		return nil, nil
	}
	var log *types.MessagePublishedDetail
	err := rlp.DecodeBytes(value, &log)
	if err != nil {
		return nil, err
	}
	return log, nil
}
func (d *DB) GetUnSignChainIdNonce(chainId, startNonce uint64) (*types.MessagePublishedDetail, error) {
	key := makeUnSignChainIdNonceIndexKey(chainId, startNonce)
	detailkey, _ := d.db.Get([]byte(key))
	if detailkey == nil {
		return nil, nil
	}
	value, _ := d.db.Get(detailkey)
	if value == nil {
		return nil, nil
	}
	var log *types.MessagePublishedDetail
	err := rlp.DecodeBytes(value, &log)
	if err != nil {
		return nil, err
	}
	return log, nil
}
func (d *DB) GetQuorumLogRangeNonce(chainId, startNonce uint64, limit uint64) ([]*types.MessagePublishedDetail, error) {
	it := d.db.NewIterator([]byte(makeSignChainIdNonceIndexKeyPrefix(chainId)), []byte(fmt.Sprintf("%x", Uint64toBytes(startNonce))))
	var details []*types.MessagePublishedDetail
	for it.Next() {
		if uint64(len(details)) == limit {
			break
		}
		buf, _ := d.db.Get(it.Value())
		if buf == nil {
			return nil, nil
		}
		var log *types.MessagePublishedDetail
		err := rlp.DecodeBytes(buf, &log)
		if err != nil {
			return nil, err
		}
		if log.Log.Nonce >= startNonce {
			details = append(details, log)
		}
	}
	return details, nil
}

func (d *DB) GetQuorumLogByTxHash(hash common.Hash) ([]*types.MessagePublishedDetail, error) {
	it := d.db.NewIterator([]byte(makeSignHashSequenceIndexKeyPrefix(hash)), nil)
	var details []*types.MessagePublishedDetail
	for it.Next() {
		buf, _ := d.db.Get(it.Value())
		if buf == nil {
			return nil, nil
		}
		var log *types.MessagePublishedDetail
		err := rlp.DecodeBytes(buf, &log)
		if err != nil {
			return nil, err
		}
		details = append(details, log)
	}
	return details, nil
}

func (d *DB) GetUnSignLogRangeNonce(chainId, startNonce uint64, limit uint64) ([]*types.MessagePublishedDetail, error) {
	it := d.db.NewIterator([]byte(makeUnSignChainIdNonceIndexKeyPrefix(chainId)), []byte(fmt.Sprintf("%x", Uint64toBytes(startNonce))))
	var details []*types.MessagePublishedDetail
	for it.Next() {
		if uint64(len(details)) == limit {
			break
		}
		buf, _ := d.db.Get(it.Value())
		if buf == nil {
			return nil, nil
		}
		var log *types.MessagePublishedDetail
		err := rlp.DecodeBytes(buf, &log)
		if err != nil {
			return nil, err
		}
		if log.Log.Nonce >= startNonce {
			details = append(details, log)
		}
	}
	return details, nil
}

func (d *DB) GetUnSignLogByTxHash(hash common.Hash) ([]*types.MessagePublishedDetail, error) {
	it := d.db.NewIterator([]byte(makeUnSignHashSequenceIndexKeyPrefix(hash)), nil)
	var details []*types.MessagePublishedDetail
	for it.Next() {
		buf, _ := d.db.Get(it.Value())
		if buf == nil {
			return nil, nil
		}
		var log *types.MessagePublishedDetail
		err := rlp.DecodeBytes(buf, &log)
		if err != nil {
			return nil, err
		}
		details = append(details, log)
	}
	return details, nil
}
func (d *DB) GetAllUnSignMessagePublished() ([]*types.MessagePublishedDetail, error) {
	it := d.db.NewIterator([]byte(makeUnSignHashRangePrefix()), nil)
	var details []*types.MessagePublishedDetail
	for it.Next() {
		var log *types.MessagePublishedDetail
		err := rlp.DecodeBytes(it.Value(), &log)
		if err != nil {
			return nil, err
		}
		details = append(details, log)
	}
	return details, nil
}
func (d *DB) GetAllMessagePublished() ([]*types.MessagePublishedDetail, error) {
	it := d.db.NewIterator([]byte(makeDetailKeyPrefix()), nil)
	var details []*types.MessagePublishedDetail
	for it.Next() {
		var log *types.MessagePublishedDetail
		err := rlp.DecodeBytes(it.Value(), &log)
		if err != nil {
			return nil, err
		}
		details = append(details, log)
	}
	return details, nil
}

func (d *DB) UpdateMessagePublished(log *types.MessagePublishedDetail) error {
	batch := d.db.NewBatch()
	buf, err := log.Bytes()
	if err != nil {
		return err
	}
	//存储原始数据
	detailKey := []byte(makeDetailKey(log.Log.Hash()))
	batch.Put(detailKey, buf)
	return batch.Write()
}

func (d *DB) SetMessagePublished(logs []*types.MessagePublishedDetail) error {
	batch := d.db.NewBatch()
	for _, log := range logs {
		buf, err := log.Bytes()
		if err != nil {
			return err
		}
		//存储原始数据
		detailKey := []byte(makeDetailKey(log.Log.Hash()))
		batch.Put(detailKey, buf)
		//构建索引
		batch.Put([]byte(makeUnSignChainIdNonceIndexKey(log.Log.ChainId, log.Log.Nonce)), detailKey)
		batch.Put([]byte(makeUnSignHashSequenceIndexKey(log.TxHash, log.Log.Sequence)), detailKey)
	}
	return batch.Write()
}

func (d *DB) deleteUnSignIndex(batch ethdb.Batch, log *types.MessagePublishedDetail) {
	batch.Delete([]byte(makeUnSignChainIdNonceIndexKey(log.Log.ChainId, log.Log.Nonce)))
	batch.Delete([]byte(makeUnSignHashSequenceIndexKey(log.TxHash, log.Log.Sequence)))
}
