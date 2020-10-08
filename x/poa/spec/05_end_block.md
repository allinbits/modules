# End-Block

At each ABCI `EndBlock` call, the operations to update validator sets are executed.

## Validator set changes

### Adding validators

The POA validator set is updated as a part of the `EndBlock` process, any updated validators are also returned back to Tendermint for inclusion in the Tendermint validator set. The operations are as follows:

- The previous validator set is then compared with the new validator set and:
    1. Accepted Validators are added to the set
    1. Jailed Validators are removed from the set
    1. All validators are given a `VotingPower` of 10

Any validators leaving or entering the validator set are to send a message which is passed back to Tendermint.

### Removing validators

If a validator has been acting maliciously it will be removed from the validator set.

