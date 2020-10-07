# Begin-Block

At each ABCI `BeginBlock` call, the operations to count votes are executed.

## Vote changes

### Validators changes

- Votes are calculated live as the chain progresses
- Propose new validators to be added to the validator set
- If the votes on a new validator are greater than 2/3 of the validator set, set the `Validator.Accepted` value to true

### Check for malicious activities

- If votes for a validator don't match the `Validator.Accepted` value, jail the validator
