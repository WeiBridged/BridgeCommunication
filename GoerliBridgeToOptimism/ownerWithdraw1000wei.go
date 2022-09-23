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

    goerliBridge "testProject/contracts/GoerliBridge"
    optimismBridge "testProject/contracts/OptimismBridge"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/crypto"
)

func main() {

  client, chainID := clientSetup(os.Getenv("optimismAlchemyWSS"))
  fmt.Println("chainID: ", chainID)

  contractAddress := common.HexToAddress("0xf5f1e4510B7c1645491285eBb9F762E371884B45")
  contract := connectContractAddress(client,contractAddress)

  auth, fromAddress := connectWallet(os.Getenv("devTestnetPrivateKey"),client,chainID)

  Owner := getOwner(contract)
  fmt.Println("storedData:", Owner)

  // clientCrossChain, chainIDCrossChain := clientSetup(os.Getenv("goerliWebSocketSecureEventsInfuraAPIKey"))
  // fmt.Println("chainIDCrossChain: ", chainIDCrossChain)

  // Use this endpoint when you are running your own node on a specific chain (no events)
  clientCrossChain, chainIDCrossChain := clientSetup("http://localhost:8545")
  fmt.Println("chainIDCrossChain: ", chainIDCrossChain)

  contractAddressCrossChain := common.HexToAddress("0xaED1aC1429EAB4569e218b2aD1A585146fCdE061")
  contractCrossChain := connectContractAddressCrossChain(clientCrossChain,contractAddressCrossChain)

  First := getFirst(contractCrossChain)
  fmt.Println("First:", First)

  Last := getLast(contractCrossChain)
  fmt.Println("Last:", Last)

  if Last.Cmp(First) > -1 {
    log.Fatal("QUEUE IS NOT EMPTY!")
  }

  ContractBridgeTokens, err := client.BalanceAt(context.Background(), contractAddress, nil)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("ContractBridgeTokens", ContractBridgeTokens) // 25893180161173005034

  BigInt0 := big.NewInt(0)
  if  ContractBridgeTokens.Cmp(BigInt0) == 0 {
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

func connectContractAddress(client *ethclient.Client, contractAddress common.Address) (contract *optimismBridge.OptimismBridge) {

  contract, err := optimismBridge.NewOptimismBridge(contractAddress, client)
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

func getOwner(contract *optimismBridge.OptimismBridge) (storedData common.Address) {

  storedData, err := contract.Owner(&bind.CallOpts{})
  if err != nil {
        log.Fatal(err)
  }
  return

}

func OwnerRemoveBridgeLiqudityTx(client *ethclient.Client, auth *bind.TransactOpts, fromAddress common.Address, contract *optimismBridge.OptimismBridge) {

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
