package systemcontract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/parlia/vmcaller"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"math"
	"math/big"
)

const (
	RewardSwitchCode = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c8063746c8ae11161005b578063746c8ae1146100c8578063be9a6555146100dc578063c4d66de8146100e6578063f851a440146100f957600080fd5b8063158ef93e1461008257806350f9b6cd146100a457806357e871e7146100b1575b600080fd5b60005461008f9060ff1681565b60405190151581526020015b60405180910390f35b60025461008f9060ff1681565b6100ba60015481565b60405190815260200161009b565b60005461008f90600160a81b900460ff1681565b6100e4610129565b005b6100e46100f43660046102bf565b610223565b6000546101119061010090046001600160a01b031681565b6040516001600160a01b03909116815260200161009b565b60005461010090046001600160a01b0316331461018d5760405162461bcd60e51b815260206004820152600a60248201527f41646d696e206f6e6c790000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b60025460ff16156101e05760405162461bcd60e51b815260206004820152600f60248201527f416c7265616479207374617274656400000000000000000000000000000000006044820152606401610184565b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff16600160a81b1790554360019081556002805460ff19169091179055565b60005460ff16156102765760405162461bcd60e51b815260206004820152601360248201527f416c726561647920696e697469616c697a6564000000000000000000000000006044820152606401610184565b6000805460ff196001600160a01b0390931661010002929092167fffffffffffffffffffffff000000000000000000000000000000000000000000909216919091176001179055565b6000602082840312156102d157600080fd5b81356001600160a01b03811681146102e857600080fd5b939250505056fea164736f6c6343000809000a"
)

var (
	rewardSwitchAdmin        = common.HexToAddress("0xffe642ecf5f51caa422e5f814483f6d989a6cc1b")
	rewardSwitchAdminTestnet = common.HexToAddress("0x83c4a11940ebcc04ac8a03cc2d290e177cccdf0f")
)

type hardForkRewardSwitch struct {
}

func (s *hardForkRewardSwitch) GetName() string {
	return RewardSwitchContractName
}
func (s *hardForkRewardSwitch) getAdminByChainId(chainId *big.Int) common.Address {
	if chainId.Cmp(params.MainnetChainConfig.ChainID) == 0 {
		return rewardSwitchAdmin
	}

	return rewardSwitchAdminTestnet
}

func (s *hardForkRewardSwitch) Update(config *params.ChainConfig, height *big.Int, state *state.StateDB) (err error) {
	contractCode := common.FromHex(RewardSwitchCode)

	//write code to sys contract
	state.SetCode(RewardSwitchContractAddr, contractCode)
	log.Debug("Write code to system contract account", "addr", RewardSwitchContractAddr.String(), "code", contractCode)

	return
}

func (s *hardForkRewardSwitch) Execute(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (err error) {
	// initialize v1 contract
	method := "initialize"
	data, err := GetInteractiveABI()[s.GetName()].Pack(method, s.getAdminByChainId(config.ChainID))
	if err != nil {
		log.Error("Can't pack data for initialize", "error", err)
		return err
	}

	msg := types.NewMessage(header.Coinbase, &RewardSwitchContractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, nil, false)
	_, err = vmcaller.ExecuteMsg(msg, state, header, chainContext, config)

	return
}
