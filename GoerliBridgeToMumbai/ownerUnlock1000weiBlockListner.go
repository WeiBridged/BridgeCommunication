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
    "time"

    weth "testProject/contracts/WETH"
    goerliBridge "testProject/contracts/GoerliBridge"
    mumbaiBridge "testProject/contracts/MumbaiBridge"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/core/types"
)

func main() {

  client, chainID := clientSetup(os.Getenv("mumbaiQuicknodeWSS"))
  fmt.Println("chainID: ", chainID)

  contractAddressWETH := common.HexToAddress("0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa")
  contractWETH := connectContractAddressWETH(client,contractAddressWETH)

  contractAddress := common.HexToAddress("0xb7307DDD7C370A309DB38243258318CbB5E1860C")
  contract := connectContractAddress(client,contractAddress)

  auth, fromAddress := connectWallet(os.Getenv("devTestnetPrivateKey"),client,chainID)

  Owner := getOwner(contract)
  fmt.Println("storedData:", Owner)

  // clientCrossChain, chainIDCrossChain := clientSetup(os.Getenv("goerliWebSocketSecureEventsInfuraAPIKey"))
  // fmt.Println("chainIDCrossChain: ", chainIDCrossChain)

  // Use this endpoint when you are running your own node on a specific chain (events allowed)
  clientCrossChain, chainIDCrossChain := clientSetup("ws://localhost:8546")
  fmt.Println("chainIDCrossChain: ", chainIDCrossChain)

  authCrossChain, fromAddress := connectWalletCrossChain(os.Getenv("devTestnetPrivateKey"),clientCrossChain,chainIDCrossChain)

  contractAddressCrossChain := common.HexToAddress("0xe33EE68Fc5477Ea95F4897b67d3E763b7F74FC52")
  contractCrossChain := connectContractAddressCrossChain(clientCrossChain,contractAddressCrossChain)

  //Listen for new blocks.on Goerli to see if queue changed
  headers := make(chan *types.Header)
  sub, err := clientCrossChain.SubscribeNewHead(context.Background(), headers)
  if err != nil {
     log.Fatal(err)
  }

  for {
     select {
     case err := <-sub.Err():
         log.Fatal(err)
     case header := <-headers:
         fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

         block, err := clientCrossChain.BlockByHash(context.Background(), header.Hash())
         if err != nil {
             log.Fatal(err)
         }

         fmt.Println("Goerli Block Number:", block.Number().Uint64())   // 3477413


         First := getFirst(contractCrossChain)
         fmt.Println("First:", First)

         Last := getLast(contractCrossChain)
         fmt.Println("Last:", Last)

         if Last.Cmp(First) == -1 {
           // log.Fatal("QUEUE IS EMPTY!")
           fmt.Println("QUEUE IS EMPTY!", Last)
           continue
         }

         UserInQueue := getUserInQueue(Last,contractCrossChain)
         fmt.Println("UserInQueue:", UserInQueue)

         balanceBridgeWETH := getBalanceOfBridgeWETH(contractWETH,contractAddress)
         fmt.Println("Bridge WETH Balance:", balanceBridgeWETH)

         DequeueTx(clientCrossChain,authCrossChain,fromAddress,contractCrossChain);

         time.Sleep(30 * time.Second)

         OwnerUnlockOptimismETHTx(UserInQueue,client,auth,fromAddress,contract);

         time.Sleep(30 * time.Second)

         // fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
         // fmt.Println(block.Time().Uint64())     // 1529525947
         // fmt.Println(block.Nonce())             // 130524141876765836
         // fmt.Println(len(block.Transactions())) // 7
     }
  }



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

func connectContractAddress(client *ethclient.Client, contractAddress common.Address) (contract *mumbaiBridge.MumbaiBridge ) {

  contract, err := mumbaiBridge.NewMumbaiBridge(contractAddress, client)
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

func getOwner(contract *mumbaiBridge.MumbaiBridge) (storedData common.Address) {

  storedData, err := contract.Owner(&bind.CallOpts{})
  if err != nil {
        log.Fatal(err)
  }
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

func getUserInQueue(Last *big.Int , contract *goerliBridge.GoerliBridge) (storedData common.Address) {

  storedData, err := contract.Queue(&bind.CallOpts{}, Last)
  if err != nil {
        log.Fatal(err)
  }
  return

}

func DequeueTx(client *ethclient.Client, auth *bind.TransactOpts, fromAddress common.Address, contract *goerliBridge.GoerliBridge) {

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

  tx, err := contract.Dequeue(auth)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println("Tx hash:", tx.Hash().Hex()) // tx sent
  return
}


func OwnerUnlockOptimismETHTx(queueAddress common.Address, client *ethclient.Client, auth *bind.TransactOpts, fromAddress common.Address, contract *mumbaiBridge.MumbaiBridge) {

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

  tx, err := contract.OwnerUnlockOptimismETH(auth,queueAddress)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println("Tx hash:", tx.Hash().Hex()) // tx sent
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
