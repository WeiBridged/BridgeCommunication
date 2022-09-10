//Step 1: Create abi file by running: solc --abi Store.sol > Store_sol_Store.abi
//Step 2: Create bin file by running: solc --bin Store.sol > Store_sol_Store.bin
//Step 3: Remove comments above the abi and bin files.
//Step 4: Generate Go contract interaction file by running:  abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=main --out=Store.go
//Step 5: Run: StoreTest.go Store.go
package main

import (
    "fmt"
    "log"
    "os"
    "context"
    "crypto/ecdsa"
    "math/big"

    // "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    // "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
)

func main() {
    //Get smart contract starting state.
    client, err := ethclient.Dial(os.Getenv("goerliWebSocketSecureEventsInfuraAPIKey"))
    // client, err := ethclient.Dial("http://localhost/8545")

    if err != nil {
        log.Fatal(err)
    }

    chainID, err := client.NetworkID(context.Background())
     if err != nil {
         log.Fatal(err)
     }
     fmt.Println("Chain id: ", chainID)

     contractAddress := common.HexToAddress("0xd00FcF4B79D6911F54989280b132aAd21b0d2438")
     contract, err := NewMain(contractAddress, client)
     if err != nil {
         log.Fatal(err)
     }
     fmt.Println("Contract is loaded at address", contractAddress)

     storedData, err := contract.Owner(&bind.CallOpts{})
       if err != nil {
           log.Fatal(err)
     }
     fmt.Println("storedData:", storedData)

     //Set new value for smart contract uint storage slot.
     privateKey, err := crypto.HexToECDSA(os.Getenv("devTestnetPrivateKey"))
      if err != nil {
          log.Fatal(err)
      }

      publicKey := privateKey.Public()
      publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
      if !ok {
          log.Fatal("error casting public key to ECDSA")
      }

      fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
      nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
      if err != nil {
          log.Fatal(err)
      }

      gasPrice, err := client.SuggestGasPrice(context.Background())
      if err != nil {
          log.Fatal(err)
      }

      auth := bind.NewKeyedTransactor(privateKey)
      auth.Nonce = big.NewInt(int64(nonce))
      auth.GasLimit = uint64(300000) // in units
      auth.GasPrice = gasPrice
      auth.Value = big.NewInt(1000)     // in wei

      tx, err := contract.OwnerAddBridgeLiqudity(auth)
      if err != nil {
          log.Fatal(err)
      }
      fmt.Println("Tx hash:", tx.Hash().Hex()) // tx sent


}
