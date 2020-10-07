!! This module is still in beta and not fit for production use !!

### How to find more information on the module

---

Please view to Proof of Authority module spec for more details on the commands defined in the `How to` sections

> **[Link to Spec](./spec/README.md)**

<br />

### How to install module

---

#### How to add the module to your application

1. See the diff on [this PR](https://github.com/PaddyMc/authority-chain/pull/1) and emulate the changes in your application

1. (OPTIONAL) Set up local network to test changes, can be copy/paste'd from [this PR](https://github.com/PaddyMc/authority-chain/pull/2)

<br />

### How to use the module

---

#### How to add a validator to the validator set

1. Create a CreateValidatorPOA transaction and submit it

```sh
appcli tx poa create-validator val1 `hex-encoded-public-key-of-the-validator` --trust-node --from validator --chain-id cash
```

<br />

2. Verfiy that the CreateValidatorPOA tranaction was correctly processed

```sh
appcli query poa validator-poa val1 --trust-node --chain-id cash
```

<br />

#### How to vote on a validator to allow the validator to be added to the validator set

1. Vote for a validator to join the validator set

```sh
appcli tx poa vote-validator val1 --trust-node --from validator --chain-id cash
```

<br />

2. Query for the vote to see if the transaction was successful

```sh
appcli query poa vote-poa val1 `validator address` --trust-node --chain-id cash
```

<br />

#### How to query all votes and validators

1. Query all votes

```sh
appcli query poa votes
```

<br />

2. Query all validators

```sh
appcli query poa validators
```
