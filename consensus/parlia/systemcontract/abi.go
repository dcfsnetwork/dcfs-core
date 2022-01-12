package systemcontract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"strings"
)

// ValidatorsInteractiveABI contains all methods to interactive with validator contracts.
const ValidatorsInteractiveABI = `
[
	{
		"inputs": [
		  {
			"internalType": "address[]",
			"name": "vals",
			"type": "address[]"
		  }
		],
		"name": "initialize",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "distributeBlockReward",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "getTopValidators",
		"outputs": [
		  {
			"internalType": "address[]",
			"name": "",
			"type": "address[]"
		  }
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
		  {
			"internalType": "address[]",
			"name": "newSet",
			"type": "address[]"
		  },
		  {
			"internalType": "uint256",
			"name": "epoch",
			"type": "uint256"
		  }
		],
		"name": "updateActiveValidatorSet",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
    "inputs": [
      {
        "internalType": "address",
        "name": "val",
        "type": "address"
      }
    ],
    "name": "getValidatorInfo",
    "outputs": [
      {
        "internalType": "address payable",
        "name": "",
        "type": "address"
      },
      {
        "internalType": "enum Validators.Status",
        "name": "",
        "type": "uint8"
      },
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      },
      {
        "internalType": "address[]",
        "name": "",
        "type": "address[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]
`

const PunishInteractiveABI = `
[
	{
		"inputs": [],
		"name": "initialize",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
		  {
			"internalType": "address",
			"name": "val",
			"type": "address"
		  }
		],
		"name": "punish",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
		  {
			"internalType": "uint256",
			"name": "epoch",
			"type": "uint256"
		  }
		],
		"name": "decreaseMissedBlocksCounter",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	  }
]
`

const ProposalInteractiveABI = `
[
	{
		"inputs": [
		  {
			"internalType": "address[]",
			"name": "vals",
			"type": "address[]"
		  }
		],
		"name": "initialize",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
`

const SysGovInteractiveABI = `
[
    {
		"inputs": [
			{
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			}
		],
		"name": "finishProposalById",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "index",
				"type": "uint32"
			}
		],
		"name": "getPassedProposalByIndex",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
        		"internalType": "uint256",
        		"name": "action",
        		"type": "uint256"
        	},
			{
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			},
			{
				"internalType": "bytes",
				"name": "data",
				"type": "bytes"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "getPassedProposalCount",
		"outputs": [
			{
				"internalType": "uint32",
				"name": "",
				"type": "uint32"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_admin",
				"type": "address"
			}
		],
		"name": "initialize",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

const AddrListInteractiveABI = `
[
    {
        "inputs": [],
        "name": "devVerifyEnabled",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getBlacksFrom",
        "outputs": [
            {
                "internalType": "address[]",
                "name": "",
                "type": "address[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getBlacksTo",
        "outputs": [
            {
                "internalType": "address[]",
                "name": "",
                "type": "address[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_admin",
                "type": "address"
            }
        ],
        "name": "initialize",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "addr",
                "type": "address"
            }
        ],
        "name": "isDeveloper",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    }
]`

const ValidatorsV1InteractiveABI = `[
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "name": "activeValidators",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "distributeBlockReward",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getTopValidators",
        "outputs": [
            {
                "internalType": "address[]",
                "name": "",
                "type": "address[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address[]",
                "name": "_candidates",
                "type": "address[]"
            },
            {
                "internalType": "address[]",
                "name": "_manager",
                "type": "address[]"
            },
            {
                "internalType": "address",
                "name": "_admin",
                "type": "address"
            }
        ],
        "name": "initialize",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address[]",
                "name": "newSet",
                "type": "address[]"
            },
            {
                "internalType": "uint256",
                "name": "epoch",
                "type": "uint256"
            }
        ],
        "name": "updateActiveValidatorSet",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    }
]`

const PunishV1InteractiveABI = `[
   {
     "inputs": [],
     "name": "initialize",
     "outputs": [],
     "stateMutability": "nonpayable",
     "type": "function"
   }
]`

const RewardSwitchABI = `[
    {
      "inputs": [],
      "name": "admin",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "blockNumber",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "called",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_admin",
          "type": "address"
        }
      ],
      "name": "initialize",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "initialized",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "start",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "startReward",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }
  ]`
const RewardHomeABI = `[
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "activeCount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "stakingAmount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "gasDestory",
          "type": "uint256"
        }
      ],
      "name": "AddMargin",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "manager",
          "type": "address"
        }
      ],
      "name": "ChangeManager",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "poolAddress",
          "type": "address"
        }
      ],
      "name": "CreateVotePool",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "Punish",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "SetAcceptPower",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "SetOriginPower",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "stakingCost",
          "type": "uint256"
        },
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "gasCost",
          "type": "uint256"
        }
      ],
      "name": "UpdateCost",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "WithdrawMargin",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "WithdrawRewrad",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "DayBlockCount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "MINSUPERNODE",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "StakingLockPeriod",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "acceptDebt",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "acceptPerShare",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "acceptPower",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "activeCount",
          "type": "uint256"
        }
      ],
      "name": "addMargin",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "admin",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_manager",
          "type": "address"
        }
      ],
      "name": "changeManager",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "curVotePool",
      "outputs": [
        {
          "internalType": "contract INodeVote",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "distributeReward",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "gasDestory",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "gasPerPower",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "user",
          "type": "address"
        }
      ],
      "name": "getPendingReward",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getTopNode",
      "outputs": [
        {
          "internalType": "address[]",
          "name": "",
          "type": "address[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_admin",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "_stakingPerPower",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "_gasPerPower",
          "type": "uint256"
        }
      ],
      "name": "initialize",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "initialized",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "lastPayPerShare",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "manager",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "normalPerShare",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "obtainReward",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "originPower",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "remainPower",
          "type": "uint256"
        }
      ],
      "name": "punishAmount",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "setAcceptPower",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "setOriginPower",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "stakingPerPower",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "stakingPower",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "startSuperVote",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "superPerShare",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "superVotePerShare",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "totalAcceptPower",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "totalDestory",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "totalNodeCount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "totalNormalPower",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "totalPunishCount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "stake",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "gas",
          "type": "uint256"
        }
      ],
      "name": "updatePowerCost",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "userAcceptState",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "userCanReedmeBalance",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "userDebt",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "userLockAmount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "userLockEndDate",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "votePoolCount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "votePools",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "withdrawMargin",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`

// DevMappingPosition is the position of the state variable `devs`.
// Since the state variables are as follow:
//    bool public initialized;
//    bool public enabled;
//    address public admin;
//    address public pendingAdmin;
//    mapping(address => bool) private devs;
//
// according to [Layout of State Variables in Storage](https://docs.soliditylang.org/en/v0.8.4/internals/layout_in_storage.html),
// and after optimizer enabled, the `initialized`, `enabled` and `admin` will be packed, and stores at slot 0,
// `pendingAdmin` stores at slot 1, so the position for `devs` is 2.
const DevMappingPosition = 2

var (
	ValidatorsContractName   = "validators"
	PunishContractName       = "punish"
	ProposalContractName     = "proposal"
	SysGovContractName       = "governance"
	AddressListContractName  = "address_list"
	ValidatorsV1ContractName = "validators_v1"
	PunishV1ContractName     = "punish_v1"
	RewardSwitchContractName = "rewardSwitch"
	RewardHomeContractName   = "rewardHome"
	ValidatorsContractAddr   = common.HexToAddress("0x000000000000000000000000000000000000c000")
	PunishContractAddr       = common.HexToAddress("0x000000000000000000000000000000000000c001")
	ProposalAddr             = common.HexToAddress("0x000000000000000000000000000000000000c002")
	SysGovContractAddr       = common.HexToAddress("0x000000000000000000000000000000000000c003")
	AddressListContractAddr  = common.HexToAddress("0x000000000000000000000000000000000000c004")
	ValidatorsV1ContractAddr = common.HexToAddress("0x000000000000000000000000000000000000c005")
	PunishV1ContractAddr     = common.HexToAddress("0x000000000000000000000000000000000000c006")
	RewardSwitchContractAddr = common.HexToAddress("0x000000000000000000000000000000000000c007")
	RewardHomeContractAddr   = common.HexToAddress("0x000000000000000000000000000000000000c008")

	// SysGovToAddr is the To address for the system governance transaction, NOT contract address
	SysGovToAddr = common.HexToAddress("0x000000000000000000000000000000000000cccc")

	abiMap map[string]abi.ABI
)

func init() {
	abiMap = make(map[string]abi.ABI, 0)
	tmpABI, _ := abi.JSON(strings.NewReader(ValidatorsInteractiveABI))
	abiMap[ValidatorsContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(PunishInteractiveABI))
	abiMap[PunishContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(ProposalInteractiveABI))
	abiMap[ProposalContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(SysGovInteractiveABI))
	abiMap[SysGovContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(AddrListInteractiveABI))
	abiMap[AddressListContractName] = tmpABI

	tmpABI, _ = abi.JSON(strings.NewReader(ValidatorsV1InteractiveABI))
	abiMap[ValidatorsV1ContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(PunishV1InteractiveABI))
	abiMap[PunishV1ContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(RewardSwitchABI))
	abiMap[RewardSwitchContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(RewardHomeABI))
	abiMap[RewardHomeContractName] = tmpABI

}

func GetInteractiveABI() map[string]abi.ABI {
	return abiMap
}

func GetValidatorAddr(blockNum *big.Int, config *params.ChainConfig) *common.Address {
	if config.IsRedCoast(blockNum) {
		return &ValidatorsV1ContractAddr
	}
	return &ValidatorsContractAddr
}

func GetPunishAddr(blockNum *big.Int, config *params.ChainConfig) *common.Address {
	if config.IsRedCoast(blockNum) {
		return &PunishV1ContractAddr
	}
	return &PunishContractAddr
}
