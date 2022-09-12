//Step 1: Create abi file by running: solc --abi Store.sol > Store_sol_Store.abi
//Step 2: Create bin file by running: solc --bin Store.sol > Store_sol_Store.bin
//Step 3: Remove comments above the abi and bin files.
//Step 4: Generate Go contract interaction file by running:  abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=main --out=Store.go
//Step 5: Run: getSetEvent.go Store.go
package main

import (
    "fmt"
    "log"
    "os"
    "context"
    "crypto/ecdsa"
    "math/big"

    weth "testProject/contracts/WETH"
    goerliBridge "testProject/contracts/GoerliBridge"
    // optimismBridge "testProject/contracts/OptimismBridge"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/crypto"
)

func main() {

  // Use this endpoint when you are running your own node on a specific chain (no events)
  // client, chainID := clientSetup("http://localhost/8545")

  // Use this endpoint when you are running your own node on a specific chain (events allowed)
  // client, chainID := clientSetup("ws://localhost/8546")

  client, chainID := clientSetup(os.Getenv("goerliWebSocketSecureEventsInfuraAPIKey"))
  fmt.Println("chainID: ", chainID)

  contractAddress := common.HexToAddress("0xe33EE68Fc5477Ea95F4897b67d3E763b7F74FC52")
  contract := connectContractAddress(client,contractAddress)

  auth, fromAddress := connectWallet(os.Getenv("devTestnetPrivateKeyTwo"),client,chainID)

  Owner := getOwner(contract)
  fmt.Println("Owner:", Owner)

  clientCrossChain, chainIDCrossChain := clientSetup(os.Getenv("mumbaiQuicknodeWSS"))
  fmt.Println("chainIDCrossChain: ", chainIDCrossChain)

  contractAddressWETH := common.HexToAddress("0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa")
  contractWETH := connectContractAddressWETH(clientCrossChain,contractAddressWETH)

  contractAddressCrossChain := common.HexToAddress("0xb7307DDD7C370A309DB38243258318CbB5E1860C")
  // contractCrossChain := connectContractAddressCrossChain(clientCrossChain,contractAddressCrossChain)


  balanceBridgeWETH := getBalanceOfBridgeWETH(contractWETH,contractAddressCrossChain)
  fmt.Println("Bridge WETH Balance:", balanceBridgeWETH)

  First := getFirst(contract)
  fmt.Println("First:", First)

  Last := getLast(contract)
  fmt.Println("Last:", Last)

  bufferLast := big.NewInt(0).Add(Last, big.NewInt(1))
  newUserInQueue := big.NewInt(0).Add(bufferLast, big.NewInt(1))
  getNewQueueSize := big.NewInt(0).Sub(newUserInQueue, First)
  contractBalanceForNewQueue := big.NewInt(0).Mul(getNewQueueSize, big.NewInt(1000))

  // ((Last+1)-First))*1000 >= contractbalance
  fmt.Println("CONTRACT MUST HAVE THIS BALANCE FOR NEW QUEUE:", contractBalanceForNewQueue)

  // BigInt0 := big.NewInt(0)
  if  balanceBridgeWETH.Cmp(contractBalanceForNewQueue) == -1 {
    log.Fatal("BRIDGE DOES NOT HAVE ANY FUNDS LEFT FOR NEXT USER IN QUEUE!!")
  }


  // CHECK THE OTHER CHAIN CONTRACT TO SEE IF THE OTHER SIDE HAS ENOUGH TOKENS TO BRIDGE!!!!!!!

  LockTokensForOptimismTx(client,auth,fromAddress,contract);

}

func clientSetup(wssConnectionURL string) (client *ethclient.Client, chainID *big.Int) {

  client, err := ethclient.Dial(wssConnectionURL)
  if err != nil {
      log.Fatal(err)
  }

  chainID, err = client.NetworkID(context.Background())
  if err != nil {
     log.Fatal(err)
  }
  return
}

func connectContractAddress(client *ethclient.Client, contractAddress common.Address) (contract *goerliBridge.GoerliBridge) {

  contract, err := goerliBridge.NewGoerliBridge(contractAddress, client)
  if err != nil {
      log.Fatal(err)
  }
  return
}

func connectWallet(privateKeyString string, client *ethclient.Client, chainID *big.Int) (auth *bind.TransactOpts, fromAddress common.Address) {

   privateKey, err := crypto.HexToECDSA(privateKeyString)
   if err != nil {
       log.Fatal(err)
   }

   publicKey := privateKey.Public()
   publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
   if !ok {
       log.Fatal("error casting public key to ECDSA")
   }

   fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)

   auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
   if err != nil {
       log.Fatal(err)
   }

   return

}

func getOwner(contract *goerliBridge.GoerliBridge) (storedData common.Address) {

  storedData, err := contract.Owner(&bind.CallOpts{})
  if err != nil {
        log.Fatal(err)
  }
  return

}

func LockTokensForOptimismTx(client *ethclient.Client, auth *bind.TransactOpts, fromAddress common.Address, contract *goerliBridge.GoerliBridge) {

  gasPrice, err := client.SuggestGasPrice(context.Background())
  if err != nil {
      log.Fatal(err)
  }

  nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
  if err != nil {
      log.Fatal(err)
  }

  auth.Nonce = big.NewInt(int64(nonce))
  auth.GasLimit = uint64(300000) // in units
  auth.GasPrice = gasPrice
  auth.Value = big.NewInt(1003)     // in wei

  tx, err := contract.LockTokensForOptimism(auth)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println("Tx hash:", tx.Hash().Hex()) // tx sent

  return
}


// func connectContractAddressCrossChain(client *ethclient.Client, contractAddress common.Address) (contract *goerliBridge.GoerliBridge) {
//
//   contract, err := goerliBridge.NewGoerliBridge(contractAddress, client)
//   if err != nil {
//       log.Fatal(err)
//   }
//   return
// }

func getFirst(contract *goerliBridge.GoerliBridge) (storedData *big.Int) {

  storedData, err := contract.First(&bind.CallOpts{})
  if err != nil {
        log.Fatal(err)
  }
  return

}

func getLast(contract *goerliBridge.GoerliBridge) (storedData *big.Int) {

  storedData, err := contract.Last(&bind.CallOpts{})
  if err != nil {
        log.Fatal(err)
  }
  return

}

func getBalanceOfBridgeWETH(contract *weth.Weth, bridgeAddress common.Address) (storedData *big.Int) {

  storedData, err := contract.BalanceOf(&bind.CallOpts{},bridgeAddress)
  if err != nil {
        log.Fatal(err)
  }
  return

}

func connectContractAddressWETH(client *ethclient.Client, contractAddress common.Address) (contract *weth.Weth) {

  contract, err := weth.NewWeth(contractAddress, client)
  if err != nil {
      log.Fatal(err)
  }
  return
}
