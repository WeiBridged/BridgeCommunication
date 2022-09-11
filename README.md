# bridgeCommunication

Backend scripts to help contracts read and write to each other on different blockchains.

## General Overview

1. Owner adds 1000 wei [ETH] (or another token) to another side of a bridge

2. User requests to bridge 1000 wei by paying 1003 wei (3 wei covers the 0.3% fee)

3. After deposit is made, user is added to the user bridge queue

4. Owner detects user in queue

5. Locally store user in queue with Golang variable  :warning: Assume bridge will not go down with the stored address removed from queue. :warning:

6. Dequeue to remove user from the queue.

7. Unlock the added wei [ETH] (or another token) to the user locally stored, then remove user locally.

:warning:

We remove the user locally between chains before the contract sends funds follow a "change state =>" then transfer pattern to prevent reentrancy attacks.

Ideally we remove the user from the queue like the mock Chainlink Keepers for simplicity.

If we had CCIP, we would be able to easily call between different contracts on different blockchains to have the same contract format.

For extra unlock security, the contract has

        msg.sender != Owner

which would revert the unlock function call .

:warning:

## GoerliBridgeToOptimism

        gvm use go1.1sss7
        cd GoerliBridgeToOptimism

### Owner Deposit Optimism

        go run ownerDeposit1000wei.go

### User Lock Goerli

:red_circle: REWRITE CONTRACT SO THAT IT:

-ONLY ACCEPTS NEW LOCK IF BRIDGE BALANCE OTHER SIDE CAN HANDLE NEW QUEUE SIZE BY 1 BASED ON CURRENT BALANCE BY 1000 WEI INCREMENTS

:red_circle:

        go run userLock1000wei.go

### Owner Unlock Optimism

:red_circle: REWRITE CONTRACT SO THAT IT:

-ONLY SENDS 1000 WEI PER QUEUE TO AVOID ISSUE WITH BALANCE AND QUEUE BEING OUT OF SYNC

:red_circle:

        go run ownerUnlock1000weiBlockListner.go

### Owner Withdraw

        go run ownerWithdraw1000wei.go
