package state

import (
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"
)

var DefaultEthereumOracleEndpoint string = "http://127.0.0.1:1234"

type ethereumOracle struct {
	endPoint string
}

func newDefaultEthereumOracle() ethereumOracle {
	return ethereumOracle{
		endPoint: DefaultEthereumOracleEndpoint,
	}
}

func (oracle ethereumOracle) getLatestBlockNumber() (uint64, error) {
	client, err := rpc.Dial(oracle.endPoint)
	if err != nil {
		return 0, err
	}

	var result string
	err = client.Call(&result, "eth_blockNumber")
	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseUint(strings.TrimPrefix(result, "0x"), 16, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}
