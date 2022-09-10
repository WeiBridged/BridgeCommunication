# bridgeCommunication

Backend scripts to help contracts read and write to each other on different blockchains.

## General Overview

1. Owner adds 1000 wei [ETH] (or another token) to another side of a bridge

2. User requests to bridge 1000 wei by paying 1003 wei (3 wei covers the 0.3% fee)

3. After deposit is made, user is added to the user bridge queue

4. Owner detects user in queue, then transfers the added wei [ETH] (or another token) to the user

:warning: 5. User is removed from the queue :warning:

:warning: Ideally we remove the user from the queue like the mock Chainlink Keepers example to prevent potenetial reentrancy attacks. If we had CCIP, we would be able to easily call between different contracts on different blockchains to have the same contract format. We protect against this by having the transfer unlock with "msg.sender != Owner" check. :warning:

## GoerliBridgeToOptimism

        cd GoerliBridgeToOptimism

### Owner Deposit

        go run ownerDeposit1000wei.go GoerliBridge.go

### Owner Withdraw

        go run ownerWithdraw1000wei.go GoerliBridge.go
