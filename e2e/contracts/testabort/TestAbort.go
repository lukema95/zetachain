// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testabort

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AbortContext is an auto generated low-level Go binding around an user-defined struct.
type AbortContext struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}

// TestAbortMetaData contains all meta data concerning the TestAbort contract.
var TestAbortMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"aborted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"abortedWithMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"outgoing\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"revertMessage\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"getAbortedWithMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"outgoing\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"revertMessage\",\"type\":\"bytes\"}],\"internalType\":\"structAbortContext\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAborted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"outgoing\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"revertMessage\",\"type\":\"bytes\"}],\"internalType\":\"structAbortContext\",\"name\":\"abortContext\",\"type\":\"tuple\"}],\"name\":\"onAbort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5061197b8061001f6000396000f3fe60806040526004361061004e5760003560e01c80632d4cfb7e1461005757806372748f7d1461008057806380b62b70146100c25780639e59f463146100ed578063fe4caa641461012a57610055565b3661005557005b005b34801561006357600080fd5b5061007e60048036038101906100799190610aac565b610155565b005b34801561008c57600080fd5b506100a760048036038101906100a29190610b2b565b610531565b6040516100b996959493929190610c5d565b60405180910390f35b3480156100ce57600080fd5b506100d76106aa565b6040516100e49190610ccc565b60405180910390f35b3480156100f957600080fd5b50610114600480360381019061010f9190610e1c565b6106bd565b6040516101219190610f6c565b60405180910390f35b34801561013657600080fd5b5061013f6108b8565b60405161014c9190610ccc565b60405180910390f35b6101ba818060a001906101689190610f9d565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050826101b590611207565b6108cf565b610215818060a001906101cd9190610f9d565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506109d0565b1561052e57600080826020016020810190610230919061121a565b73ffffffffffffffffffffffffffffffffffffffff1663d9eeebed6040518163ffffffff1660e01b81526004016040805180830381865afa158015610279573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061029d9190611271565b915091508260200160208101906102b4919061121a565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614610321576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103189061130e565b60405180910390fd5b8260400135811115610368576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035f906113a0565b60405180910390fd5b600081846040013561037a91906113ef565b905083602001602081019061038f919061121a565b73ffffffffffffffffffffffffffffffffffffffff1663095ea7b33386604001356040518363ffffffff1660e01b81526004016103cd929190611423565b6020604051808303816000875af11580156103ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104109190611461565b503373ffffffffffffffffffffffffffffffffffffffff16637c0dcb5f85806000019061043d9190610f9d565b60405160200161044e9291906114bb565b60405160208183030381529060405283876020016020810190610471919061121a565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160405180602001604052806000815250815260200160008152506040518563ffffffff1660e01b81526004016104f89493929190611555565b600060405180830381600087803b15801561051257600080fd5b505af1158015610526573d6000803e3d6000fd5b505050505050505b50565b6000602052806000526040600020600091509050806000018054610554906115d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610580906115d7565b80156105cd5780601f106105a2576101008083540402835291602001916105cd565b820191906000526020600020905b8154815290600101906020018083116105b057829003601f168201915b5050505050908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030160009054906101000a900460ff1690806004015490806005018054610627906115d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610653906115d7565b80156106a05780601f10610675576101008083540402835291602001916106a0565b820191906000526020600020905b81548152906001019060200180831161068357829003601f168201915b5050505050905086565b600160009054906101000a900460ff1681565b6106c5610a26565b600080836040516020016106d9919061164f565b6040516020818303038152906040528051906020012081526020019081526020016000206040518060c0016040529081600082018054610718906115d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610744906115d7565b80156107915780601f1061076657610100808354040283529160200191610791565b820191906000526020600020905b81548152906001019060200180831161077457829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900460ff161515151581526020016004820154815260200160058201805461082f906115d7565b80601f016020809104026020016040519081016040528092919081815260200182805461085b906115d7565b80156108a85780601f1061087d576101008083540402835291602001916108a8565b820191906000526020600020905b81548152906001019060200180831161088b57829003601f168201915b5050505050815250509050919050565b6000600160009054906101000a900460ff16905090565b80600080846040516020016108e4919061164f565b604051602081830303815290604052805190602001208152602001908152602001600020600082015181600001908161091d9190611812565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030160006101000a81548160ff0219169083151502179055506080820151816004015560a08201518160050190816109ae9190611812565b5090505060018060006101000a81548160ff0219169083151502179055505050565b60006040516020016109e190611930565b6040516020818303038152906040528051906020012082604051602001610a08919061164f565b60405160208183030381529060405280519060200120149050919050565b6040518060c0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160001515815260200160008152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600060c08284031215610aa357610aa2610a88565b5b81905092915050565b600060208284031215610ac257610ac1610a7e565b5b600082013567ffffffffffffffff811115610ae057610adf610a83565b5b610aec84828501610a8d565b91505092915050565b6000819050919050565b610b0881610af5565b8114610b1357600080fd5b50565b600081359050610b2581610aff565b92915050565b600060208284031215610b4157610b40610a7e565b5b6000610b4f84828501610b16565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610b92578082015181840152602081019050610b77565b60008484015250505050565b6000601f19601f8301169050919050565b6000610bba82610b58565b610bc48185610b63565b9350610bd4818560208601610b74565b610bdd81610b9e565b840191505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c1382610be8565b9050919050565b610c2381610c08565b82525050565b6000819050919050565b610c3c81610c29565b82525050565b60008115159050919050565b610c5781610c42565b82525050565b600060c0820190508181036000830152610c778189610baf565b9050610c866020830188610c1a565b610c936040830187610c33565b610ca06060830186610c4e565b610cad6080830185610c33565b81810360a0830152610cbf8184610baf565b9050979650505050505050565b6000602082019050610ce16000830184610c4e565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d2982610b9e565b810181811067ffffffffffffffff82111715610d4857610d47610cf1565b5b80604052505050565b6000610d5b610a74565b9050610d678282610d20565b919050565b600067ffffffffffffffff821115610d8757610d86610cf1565b5b610d9082610b9e565b9050602081019050919050565b82818337600083830152505050565b6000610dbf610dba84610d6c565b610d51565b905082815260208101848484011115610ddb57610dda610cec565b5b610de6848285610d9d565b509392505050565b600082601f830112610e0357610e02610ce7565b5b8135610e13848260208601610dac565b91505092915050565b600060208284031215610e3257610e31610a7e565b5b600082013567ffffffffffffffff811115610e5057610e4f610a83565b5b610e5c84828501610dee565b91505092915050565b600082825260208201905092915050565b6000610e8182610b58565b610e8b8185610e65565b9350610e9b818560208601610b74565b610ea481610b9e565b840191505092915050565b610eb881610c08565b82525050565b610ec781610c29565b82525050565b610ed681610c42565b82525050565b600060c0830160008301518482036000860152610ef98282610e76565b9150506020830151610f0e6020860182610eaf565b506040830151610f216040860182610ebe565b506060830151610f346060860182610ecd565b506080830151610f476080860182610ebe565b5060a083015184820360a0860152610f5f8282610e76565b9150508091505092915050565b60006020820190508181036000830152610f868184610edc565b905092915050565b600080fd5b600080fd5b600080fd5b60008083356001602003843603038112610fba57610fb9610f8e565b5b80840192508235915067ffffffffffffffff821115610fdc57610fdb610f93565b5b602083019250600182023603831315610ff857610ff7610f98565b5b509250929050565b600080fd5b600080fd5b600067ffffffffffffffff82111561102557611024610cf1565b5b61102e82610b9e565b9050602081019050919050565b600061104e6110498461100a565b610d51565b90508281526020810184848401111561106a57611069610cec565b5b611075848285610d9d565b509392505050565b600082601f83011261109257611091610ce7565b5b81356110a284826020860161103b565b91505092915050565b6110b481610c08565b81146110bf57600080fd5b50565b6000813590506110d1816110ab565b92915050565b6110e081610c29565b81146110eb57600080fd5b50565b6000813590506110fd816110d7565b92915050565b61110c81610c42565b811461111757600080fd5b50565b60008135905061112981611103565b92915050565b600060c0828403121561114557611144611000565b5b61114f60c0610d51565b9050600082013567ffffffffffffffff81111561116f5761116e611005565b5b61117b8482850161107d565b600083015250602061118f848285016110c2565b60208301525060406111a3848285016110ee565b60408301525060606111b78482850161111a565b60608301525060806111cb848285016110ee565b60808301525060a082013567ffffffffffffffff8111156111ef576111ee611005565b5b6111fb8482850161107d565b60a08301525092915050565b6000611213368361112f565b9050919050565b6000602082840312156112305761122f610a7e565b5b600061123e848285016110c2565b91505092915050565b600081519050611256816110ab565b92915050565b60008151905061126b816110d7565b92915050565b6000806040838503121561128857611287610a7e565b5b600061129685828601611247565b92505060206112a78582860161125c565b9150509250929050565b600082825260208201905092915050565b7f7a72633230206973206e6f742067617320746f6b656e00000000000000000000600082015250565b60006112f86016836112b1565b9150611303826112c2565b602082019050919050565b60006020820190508181036000830152611327816112eb565b9050919050565b7f66656520616d6f756e7420697320686967686572207468616e2074686520616d60008201527f6f756e7400000000000000000000000000000000000000000000000000000000602082015250565b600061138a6024836112b1565b91506113958261132e565b604082019050919050565b600060208201905081810360008301526113b98161137d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006113fa82610c29565b915061140583610c29565b925082820390508181111561141d5761141c6113c0565b5b92915050565b60006040820190506114386000830185610c1a565b6114456020830184610c33565b9392505050565b60008151905061145b81611103565b92915050565b60006020828403121561147757611476610a7e565b5b60006114858482850161144c565b91505092915050565b600061149a8385610b63565b93506114a7838584610d9d565b6114b083610b9e565b840190509392505050565b600060208201905081810360008301526114d681848661148e565b90509392505050565b600060a0830160008301516114f76000860182610eaf565b50602083015161150a6020860182610ecd565b50604083015161151d6040860182610eaf565b50606083015184820360608601526115358282610e76565b915050608083015161154a6080860182610ebe565b508091505092915050565b6000608082019050818103600083015261156f8187610baf565b905061157e6020830186610c33565b61158b6040830185610c1a565b818103606083015261159d81846114df565b905095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806115ef57607f821691505b602082108103611602576116016115a8565b5b50919050565b600081519050919050565b600081905092915050565b600061162982611608565b6116338185611613565b9350611643818560208601610b74565b80840191505092915050565b600061165b828461161e565b915081905092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026116c87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261168b565b6116d2868361168b565b95508019841693508086168417925050509392505050565b6000819050919050565b600061170f61170a61170584610c29565b6116ea565b610c29565b9050919050565b6000819050919050565b611729836116f4565b61173d61173582611716565b848454611698565b825550505050565b600090565b611752611745565b61175d818484611720565b505050565b5b818110156117815761177660008261174a565b600181019050611763565b5050565b601f8211156117c65761179781611666565b6117a08461167b565b810160208510156117af578190505b6117c36117bb8561167b565b830182611762565b50505b505050565b600082821c905092915050565b60006117e9600019846008026117cb565b1980831691505092915050565b600061180283836117d8565b9150826002028217905092915050565b61181b82610b58565b67ffffffffffffffff81111561183457611833610cf1565b5b61183e82546115d7565b611849828285611785565b600060209050601f83116001811461187c576000841561186a578287015190505b61187485826117f6565b8655506118dc565b601f19841661188a86611666565b60005b828110156118b25784890151825560018201915060208501945060208101905061188d565b868310156118cf57848901516118cb601f8916826117d8565b8355505b6001600288020188555050505b505050505050565b7f7769746864726177000000000000000000000000000000000000000000000000600082015250565b600061191a600883611613565b9150611925826118e4565b600882019050919050565b600061193b8261190d565b915081905091905056fea2646970667358221220d20180e86da23fed34c64f60115551feacc54ec56a79385b23598a38b2a6c26364736f6c634300081a0033",
}

// TestAbortABI is the input ABI used to generate the binding from.
// Deprecated: Use TestAbortMetaData.ABI instead.
var TestAbortABI = TestAbortMetaData.ABI

// TestAbortBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestAbortMetaData.Bin instead.
var TestAbortBin = TestAbortMetaData.Bin

// DeployTestAbort deploys a new Ethereum contract, binding an instance of TestAbort to it.
func DeployTestAbort(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestAbort, error) {
	parsed, err := TestAbortMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestAbortBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestAbort{TestAbortCaller: TestAbortCaller{contract: contract}, TestAbortTransactor: TestAbortTransactor{contract: contract}, TestAbortFilterer: TestAbortFilterer{contract: contract}}, nil
}

// TestAbort is an auto generated Go binding around an Ethereum contract.
type TestAbort struct {
	TestAbortCaller     // Read-only binding to the contract
	TestAbortTransactor // Write-only binding to the contract
	TestAbortFilterer   // Log filterer for contract events
}

// TestAbortCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestAbortCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAbortTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestAbortTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAbortFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestAbortFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAbortSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestAbortSession struct {
	Contract     *TestAbort        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestAbortCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestAbortCallerSession struct {
	Contract *TestAbortCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TestAbortTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestAbortTransactorSession struct {
	Contract     *TestAbortTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestAbortRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestAbortRaw struct {
	Contract *TestAbort // Generic contract binding to access the raw methods on
}

// TestAbortCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestAbortCallerRaw struct {
	Contract *TestAbortCaller // Generic read-only contract binding to access the raw methods on
}

// TestAbortTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestAbortTransactorRaw struct {
	Contract *TestAbortTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestAbort creates a new instance of TestAbort, bound to a specific deployed contract.
func NewTestAbort(address common.Address, backend bind.ContractBackend) (*TestAbort, error) {
	contract, err := bindTestAbort(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestAbort{TestAbortCaller: TestAbortCaller{contract: contract}, TestAbortTransactor: TestAbortTransactor{contract: contract}, TestAbortFilterer: TestAbortFilterer{contract: contract}}, nil
}

// NewTestAbortCaller creates a new read-only instance of TestAbort, bound to a specific deployed contract.
func NewTestAbortCaller(address common.Address, caller bind.ContractCaller) (*TestAbortCaller, error) {
	contract, err := bindTestAbort(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestAbortCaller{contract: contract}, nil
}

// NewTestAbortTransactor creates a new write-only instance of TestAbort, bound to a specific deployed contract.
func NewTestAbortTransactor(address common.Address, transactor bind.ContractTransactor) (*TestAbortTransactor, error) {
	contract, err := bindTestAbort(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestAbortTransactor{contract: contract}, nil
}

// NewTestAbortFilterer creates a new log filterer instance of TestAbort, bound to a specific deployed contract.
func NewTestAbortFilterer(address common.Address, filterer bind.ContractFilterer) (*TestAbortFilterer, error) {
	contract, err := bindTestAbort(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestAbortFilterer{contract: contract}, nil
}

// bindTestAbort binds a generic wrapper to an already deployed contract.
func bindTestAbort(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestAbortMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAbort *TestAbortRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAbort.Contract.TestAbortCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAbort *TestAbortRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAbort.Contract.TestAbortTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAbort *TestAbortRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAbort.Contract.TestAbortTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAbort *TestAbortCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAbort.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAbort *TestAbortTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAbort.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAbort *TestAbortTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAbort.Contract.contract.Transact(opts, method, params...)
}

// Aborted is a free data retrieval call binding the contract method 0x80b62b70.
//
// Solidity: function aborted() view returns(bool)
func (_TestAbort *TestAbortCaller) Aborted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TestAbort.contract.Call(opts, &out, "aborted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Aborted is a free data retrieval call binding the contract method 0x80b62b70.
//
// Solidity: function aborted() view returns(bool)
func (_TestAbort *TestAbortSession) Aborted() (bool, error) {
	return _TestAbort.Contract.Aborted(&_TestAbort.CallOpts)
}

// Aborted is a free data retrieval call binding the contract method 0x80b62b70.
//
// Solidity: function aborted() view returns(bool)
func (_TestAbort *TestAbortCallerSession) Aborted() (bool, error) {
	return _TestAbort.Contract.Aborted(&_TestAbort.CallOpts)
}

// AbortedWithMessage is a free data retrieval call binding the contract method 0x72748f7d.
//
// Solidity: function abortedWithMessage(bytes32 ) view returns(bytes sender, address asset, uint256 amount, bool outgoing, uint256 chainID, bytes revertMessage)
func (_TestAbort *TestAbortCaller) AbortedWithMessage(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}, error) {
	var out []interface{}
	err := _TestAbort.contract.Call(opts, &out, "abortedWithMessage", arg0)

	outstruct := new(struct {
		Sender        []byte
		Asset         common.Address
		Amount        *big.Int
		Outgoing      bool
		ChainID       *big.Int
		RevertMessage []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Asset = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Outgoing = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.ChainID = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RevertMessage = *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return *outstruct, err

}

// AbortedWithMessage is a free data retrieval call binding the contract method 0x72748f7d.
//
// Solidity: function abortedWithMessage(bytes32 ) view returns(bytes sender, address asset, uint256 amount, bool outgoing, uint256 chainID, bytes revertMessage)
func (_TestAbort *TestAbortSession) AbortedWithMessage(arg0 [32]byte) (struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}, error) {
	return _TestAbort.Contract.AbortedWithMessage(&_TestAbort.CallOpts, arg0)
}

// AbortedWithMessage is a free data retrieval call binding the contract method 0x72748f7d.
//
// Solidity: function abortedWithMessage(bytes32 ) view returns(bytes sender, address asset, uint256 amount, bool outgoing, uint256 chainID, bytes revertMessage)
func (_TestAbort *TestAbortCallerSession) AbortedWithMessage(arg0 [32]byte) (struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}, error) {
	return _TestAbort.Contract.AbortedWithMessage(&_TestAbort.CallOpts, arg0)
}

// GetAbortedWithMessage is a free data retrieval call binding the contract method 0x9e59f463.
//
// Solidity: function getAbortedWithMessage(string message) view returns((bytes,address,uint256,bool,uint256,bytes))
func (_TestAbort *TestAbortCaller) GetAbortedWithMessage(opts *bind.CallOpts, message string) (AbortContext, error) {
	var out []interface{}
	err := _TestAbort.contract.Call(opts, &out, "getAbortedWithMessage", message)

	if err != nil {
		return *new(AbortContext), err
	}

	out0 := *abi.ConvertType(out[0], new(AbortContext)).(*AbortContext)

	return out0, err

}

// GetAbortedWithMessage is a free data retrieval call binding the contract method 0x9e59f463.
//
// Solidity: function getAbortedWithMessage(string message) view returns((bytes,address,uint256,bool,uint256,bytes))
func (_TestAbort *TestAbortSession) GetAbortedWithMessage(message string) (AbortContext, error) {
	return _TestAbort.Contract.GetAbortedWithMessage(&_TestAbort.CallOpts, message)
}

// GetAbortedWithMessage is a free data retrieval call binding the contract method 0x9e59f463.
//
// Solidity: function getAbortedWithMessage(string message) view returns((bytes,address,uint256,bool,uint256,bytes))
func (_TestAbort *TestAbortCallerSession) GetAbortedWithMessage(message string) (AbortContext, error) {
	return _TestAbort.Contract.GetAbortedWithMessage(&_TestAbort.CallOpts, message)
}

// IsAborted is a free data retrieval call binding the contract method 0xfe4caa64.
//
// Solidity: function isAborted() view returns(bool)
func (_TestAbort *TestAbortCaller) IsAborted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TestAbort.contract.Call(opts, &out, "isAborted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAborted is a free data retrieval call binding the contract method 0xfe4caa64.
//
// Solidity: function isAborted() view returns(bool)
func (_TestAbort *TestAbortSession) IsAborted() (bool, error) {
	return _TestAbort.Contract.IsAborted(&_TestAbort.CallOpts)
}

// IsAborted is a free data retrieval call binding the contract method 0xfe4caa64.
//
// Solidity: function isAborted() view returns(bool)
func (_TestAbort *TestAbortCallerSession) IsAborted() (bool, error) {
	return _TestAbort.Contract.IsAborted(&_TestAbort.CallOpts)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestAbort *TestAbortTransactor) OnAbort(opts *bind.TransactOpts, abortContext AbortContext) (*types.Transaction, error) {
	return _TestAbort.contract.Transact(opts, "onAbort", abortContext)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestAbort *TestAbortSession) OnAbort(abortContext AbortContext) (*types.Transaction, error) {
	return _TestAbort.Contract.OnAbort(&_TestAbort.TransactOpts, abortContext)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestAbort *TestAbortTransactorSession) OnAbort(abortContext AbortContext) (*types.Transaction, error) {
	return _TestAbort.Contract.OnAbort(&_TestAbort.TransactOpts, abortContext)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestAbort *TestAbortTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TestAbort.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestAbort *TestAbortSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestAbort.Contract.Fallback(&_TestAbort.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestAbort *TestAbortTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestAbort.Contract.Fallback(&_TestAbort.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestAbort *TestAbortTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAbort.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestAbort *TestAbortSession) Receive() (*types.Transaction, error) {
	return _TestAbort.Contract.Receive(&_TestAbort.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestAbort *TestAbortTransactorSession) Receive() (*types.Transaction, error) {
	return _TestAbort.Contract.Receive(&_TestAbort.TransactOpts)
}
