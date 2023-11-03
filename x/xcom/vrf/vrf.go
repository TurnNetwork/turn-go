package vrf

import (
	"encoding/hex"
	"errors"
	"github.com/bubblenet/bubble/common/math"
	"github.com/bubblenet/bubble/common/sort"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/x/xcom"
	"math/big"
	"math/rand"
	"reflect"
	"strconv"
)

// VRFItem is the element of the VRFQueue
type VRFItem struct {
	V interface{}
	X int64
	W *big.Int
}

// VRFQueue Wrap any slice to support VRF
type VRFQueue []*VRFItem

func (vq VRFQueue) Len() int {
	return len(vq)
}

func (vq VRFQueue) Less(i, j int) bool {
	return vq[i].X > vq[j].X // It's actually bigger
}

func (vq VRFQueue) Swap(i, j int) {
	vq[i], vq[j] = vq[j], vq[i]
}

// VRFQueueWrapper Wrap any slice to be VRFQueue
func VRFQueueWrapper(slice interface{}, wrapper func(interface{}) *VRFItem) (VRFQueue, error) {
	// convert slice to an interface queue, to Supports running wrapper
	s := reflect.ValueOf(slice)
	//fmt.Println(kind)
	queue := make([]interface{}, 0)
	if s.Kind() == reflect.Slice {
		for i := 0; i < s.Len(); i++ {
			queue = append(queue, s.Index(i).Interface())
		}
	} else {
		return nil, errors.New("the first parameter must be slice")
	}
	// wrap interface queue to an vrfQueue
	vrfQueue := make(VRFQueue, 0)
	for _, item := range queue {
		if vrfItem := wrapper(item); vrfItem == nil {
			return nil, errors.New("failed to convert the slice element to VRFItem")
		} else {
			vrfQueue = append(vrfQueue, vrfItem)
		}
	}

	return vrfQueue, nil
}

// VRF randomly pick number of elements from vrfQueue, it achieves randomness through the nonces
func VRF(vrfQueue VRFQueue, number uint, curNonce []byte, preNonces [][]byte) (VRFQueue, error) {
	// check params
	if len(curNonce) == 0 || len(preNonces) == 0 || len(vrfQueue) != len(preNonces) {
		log.Error("Failed to VRF", "vrfQueue Size", len(vrfQueue), "curNonceSize", len(curNonce), "preNoncesSize", len(preNonces))
		return nil, errors.New("vrf param is invalid")
	}

	totalWeights := new(big.Int)
	totalSqrtWeights := new(big.Int)
	for _, vrfer := range vrfQueue {
		totalWeights.Add(totalWeights, vrfer.W)
		totalSqrtWeights.Add(totalSqrtWeights, new(big.Int).Sqrt(vrfer.W))
	}

	var maxValue float64 = (1 << 256) - 1
	totalWeightsFloat, err := strconv.ParseFloat(totalWeights.Text(10), 64)
	if nil != err {
		return nil, err
	}
	totalSqrtWeightsFloat, err := strconv.ParseFloat(totalSqrtWeights.Text(10), 64)
	if nil != err {
		return nil, err
	}

	p := xcom.CalcP(totalSqrtWeightsFloat)
	shuffleSeed := new(big.Int).SetBytes(preNonces[0]).Int64()
	log.Debug("Call VRF parameter", "queueSize", len(vrfQueue), "p", p, "totalWeights", totalWeightsFloat, "totalSqrtWeightsFloat",
		totalSqrtWeightsFloat, "number", number, "shuffleSeed", shuffleSeed)

	rd := rand.New(rand.NewSource(shuffleSeed))
	rd.Shuffle(len(vrfQueue), func(i, j int) {
		vrfQueue[i], vrfQueue[j] = vrfQueue[j], vrfQueue[i]
	})

	for i, vrfer := range vrfQueue {
		resultStr := new(big.Int).Xor(new(big.Int).SetBytes(curNonce), new(big.Int).SetBytes(preNonces[i])).Text(10)
		xorValue, err := strconv.ParseFloat(resultStr, 64)
		if nil != err {
			return nil, err
		}

		xorP := xorValue / maxValue
		bd := math.NewBinomialDistribution(vrfer.W.Int64(), p)
		if x, err := bd.InverseCumulativeProbability(xorP); err != nil {
			return nil, err
		} else {
			vrfer.X = x
		}

		log.Debug("Call VRF finished", "index", i, "node", vrfer.V, "curNonce", hex.EncodeToString(curNonce), "preNonce",
			hex.EncodeToString(preNonces[i]), "xorValue", xorValue, "xorP", xorP, "weight", vrfer.W, "X", vrfer.X)
	}

	sort.Sort(vrfQueue)

	return vrfQueue[:number], nil
}
