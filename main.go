package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"os"
	"time"
)

// 使用eth_newPendingTransactionFilter获取filterId
// 该filterId存储在以太坊节点内存中(重启后无效)
// 使用eth_getFilterChanges获取pendingTransactionHash

var (
	client   *rpc.Client
	filterId string = "0x42de2bf5a88b642b51cf50d786328b21"
	url      string = "http://127.0.0.1:8545"
)

func main() {
	var err error

	client, err = rpc.Dial(url)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	if filterId == "" {
		if err := client.Call(&filterId, "eth_newPendingTransactionFilter"); err != nil {
			fmt.Errorf(err.Error())
			os.Exit(1)
		}
	}
	fmt.Printf("filterId:%s\r\n", filterId)

	for {
		var logs []string
		if err := client.Call(&logs, "eth_getFilterChanges", filterId); err != nil {
			fmt.Errorf(err.Error())
		}
		for _, v := range logs {
			fmt.Printf("transaction hash:%s\r\n", v)
		}
		time.Sleep(1 * time.Second)
	}
}
