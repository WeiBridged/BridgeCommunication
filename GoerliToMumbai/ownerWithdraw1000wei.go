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
    mumbaiBridge "testProject/contracts/MumbaiBridge"

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

  client, chainID := clientSetup(os.Getenv("mumbaiQuicknodeWSS"))
  fmt.Println("chainID: ", chainID)

  contractAddressWETH := common.HexToAddress("0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa")
  contractWETH := connectContractAddressWETH(client,contractAddressWETH)

  contractAddress := common.HexToAddress("0xb7307DDD7C370A309DB38243258318CbB5E1860C")
  contract := connectContractAddress(client,contractAddress)

  auth, fromAddress := connectWallet(os.Getenv("devTestnetPrivateKey"),client,chainID)

  Owner := getOwner(contract)
  fmt.Println("storedData:", Owner)

  clientCrossChain, chainIDCrossChain := clientSetup(os.Getenv("goerliWebSocketSecureEventsInfuraAPIKey"))
  fmt.Println("chainIDCrossChain: ", chainIDCrossChain)

  contractAddressCrossChain := common.HexToAddress("0xe33EE68Fc5477Ea95F4897b67d3E763b7F74FC52")
  contractCrossChain := connectContractAddressCrossChain(clientCrossChain,contractAddressCrossChain)

  First := getFirst(contractCrossChain)
  fmt.Println("First:", First)

  Last := getLast(contractCrossChain)
  fmt.Println("Last:", Last)

  // if Last.Cmp(First) > -1 {
  //   log.Fatal("QUEUE IS NOT EMPTY!")
  // }


  // User for MSG.VALUE
  // ContractBridgeTokens, err := client.BalanceAt(context.Background(), contractAddress, nil)
  // if err != nil {
  //   log.Fatal(err)
  // }
  //
  // fmt.Println("ContractBridgeTokens", ContractBridgeTokens) // 25893180161173005034

  //USED FOR ERC-20 BALANCES.
  balanceBridgeWETH := getBalanceOfBridgeWETH(contractWETH,contractAddress)
  fmt.Println("Bridge WETH Balance:", balanceBridgeWETH)

  BigInt0 := big.NewInt(0)
  if  balanceBridgeWETH.Cmp(BigInt0) == 0 {
    log.Fatal("BRIDGE DOES NOT HAVE ANY FUNDS LEFT!!")
  }

  // CHECK THE OTHER CHAIN CONTRACT TO SEE IF QUEUE IS EMPTY!!!

  OwnerRemoveBridgeLiqudityTx(client,auth,fromAddress,contract);

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

func connectContractAddress(client *ethclient.Client, contractAddress common.Address) (contract *mumbaiBridge.MumbaiBridge) {

  contract, err := mumbaiBridge.NewMumbaiBridge(contractAddress, client)
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

func getOwner(contract *mumbaiBridge.MumbaiBridge) (storedData common.Address) {

  storedData, err := contract.Owner(&bind.CallOpts{})
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

func OwnerRemoveBridgeLiqudityTx(client *ethclient.Client, auth *bind.TransactOpts, fromAddress common.Address, contract *mumbaiBridge.MumbaiBridge) {

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

  tx, err := contract.OwnerRemoveBridgeLiqudity(auth)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println("Tx hash:", tx.Hash().Hex()) // tx sent
  return
}

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

func connectContractAddressCrossChain(client *ethclient.Client, contractAddress common.Address) (contract *goerliBridge.GoerliBridge) {

  contract, err := goerliBridge.NewGoerliBridge(contractAddress, client)
  if err != nil {
      log.Fatal(err)
  }
  return
}

func connectWalletCrossChain(privateKeyString string, client *ethclient.Client, chainID *big.Int) (auth *bind.TransactOpts, fromAddress common.Address) {

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
