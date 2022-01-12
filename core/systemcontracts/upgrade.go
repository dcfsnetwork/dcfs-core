package systemcontracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
)

type UpgradeConfig struct {
	BeforeUpgrade upgradeHook
	AfterUpgrade  upgradeHook
	ContractAddr  common.Address
	CommitUrl     string
	Code          string
}

type Upgrade struct {
	UpgradeName string
	Configs     []*UpgradeConfig
}

type upgradeHook func(blockNumber *big.Int, contractAddr common.Address, statedb *state.StateDB) error

const (
	mainNet    = "Mainnet"
	chapelNet  = "Chapel"
	rialtoNet  = "Rialto"
	defaultNet = "Default"
)

var (
	GenesisHash common.Hash
	//upgrade config
	ramanujanUpgrade = make(map[string]*Upgrade)

	nielsUpgrade = make(map[string]*Upgrade)

	mirrorUpgrade = make(map[string]*Upgrade)
)
