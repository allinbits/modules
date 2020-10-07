# Messages

In this section we describe the processing of the poa messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./01_state.md) section.

## **MsgCreateValidator**

---

A validator is created using the `MsgCreateValidatorPOA` message

```go
type MsgCreateValidatorPOA struct {
    Name                   string
    Address                sdk.ValAddress
    PubKey                 crypto.PubKey
    Description            Description
}
```

This message is expected to fail if:

- Another validator with this operator address is already registered
- Another validator with this `PubKey` is already registered
- The description fields are too large

This message creates and stores the `Validator` object at appropriate indexes. The validator then sends the MsgVoteValidator message for the validator.


## **MsgVoteValidator**

---

A `Vote` can be cast as a validator using the `MsgVoteValidator`.

```go
type MsgVoteValidator struct {
    Name | Cabidate        string | sdk.ValAddress
    Voter                  sdk.ValAddress
    inFavor                bool
}
```

This message updates the `Vote` object or creates one if it does not exist. 
