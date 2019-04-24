package bloom

import (
	"bytes"
	"fmt"

	"github.com/elastos/Elastos.ELA.SideChain/bloom"
	"github.com/elastos/Elastos.ELA.SideChain/filter"
	"github.com/elastos/Elastos.ELA.SideChain/types"
	"github.com/elastos/Elastos.ELA/p2p/msg"
)

type txFilter struct {
	filter *bloom.Filter
}

func (f *txFilter) Load(filter []byte) error {
	var fl msg.FilterLoad
	err := fl.Deserialize(bytes.NewReader(filter))
	if err != nil {
		return err
	}

	f.filter = bloom.LoadFilter(&fl)

	return nil
}

func (f *txFilter) Add(filter []byte) error {
	if f.filter == nil || !f.filter.IsLoaded() {
		return fmt.Errorf("filter not loaded")
	}

	f.filter.Add(filter)

	return nil
}

func (f *txFilter) Match(tx *types.Transaction) bool {
	// Match all register asset transactions.
	if tx.IsRegisterAssetTx() {
		return true
	}
	return f.filter.MatchTxAndUpdate(tx)
}

func NewTxFilter() filter.TxFilter {
	return &txFilter{}
}
