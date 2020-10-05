package keeper

import (
	"github.com/allinbits/modules/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetValidator(ctx sdk.Context, key string, validator types.Validator) {
	// set validator by name
	k.Set(ctx, []byte(key), types.ValidatorsKey, validator)

	// set validiator by address
	k.Set(ctx, validator.Address, types.ValidatorsByAddressKey, validator)
}

// NOTE: could this be a state

// SetValidatorIsInSet after a validator has been accepted and not added to the set we need this function
func (k Keeper) SetValidatorIsInSet(ctx sdk.Context, key string, validator types.Validator, isInSet bool) {
	validator.InSet = isInSet
	k.SetValidator(ctx, validator.Name, validator)
}

// SetValidatorIsAccepted when the validator is accepeted into the consensus accepted is set to true
func (k Keeper) SetValidatorIsAccepted(ctx sdk.Context, key string, validator types.Validator, isAccepted bool) {
	validator.Accepted = isAccepted
	k.SetValidator(ctx, validator.Name, validator)
}

// SetValidatorIsAccepted when the validator is accepeted into the consensus accepted is set to true and is in the validator set
func (k Keeper) SetValidatorIsAcceptedAndInSet(ctx sdk.Context, key string, validator types.Validator, isAccepted bool, isInSet bool) {
	validator.Accepted = isAccepted
	validator.InSet = isInSet
	k.SetValidator(ctx, validator.Name, validator)
}

func (k Keeper) GetValidator(ctx sdk.Context, key string) (types.Validator, bool) {
	val, found := k.Get(ctx, []byte(key), types.ValidatorsKey, k.UnmarshalValidator)
	return val.(types.Validator), found
}

func (k Keeper) GetValidatorByAddress(ctx sdk.Context, valAddress sdk.ValAddress) (types.Validator, bool) {
	val, found := k.Get(ctx, valAddress, types.ValidatorsByAddressKey, k.UnmarshalValidator)
	return val.(types.Validator), found
}

func (k Keeper) UnmarshalValidator(value []byte) (interface{}, bool) {
	validator := types.Validator{}
	err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &validator)
	if err != nil {
		return types.Validator{}, false
	}
	return validator, true
}

// ValidatorSelectorFn allows validators to be selected by certain conditions
type ValidatorSelectorFn func(validator types.Validator) bool

func (k Keeper) GetAllValidatorsWithCondition(ctx sdk.Context, validatorSelector ValidatorSelectorFn) (validators []types.Validator) {
	val := k.GetAll(ctx, types.ValidatorsKey, k.UnmarshalValidator)

	// handle the case that there is only one validator in the set
	if len(val) == 1 {
		return append(validators, val[0].(types.Validator))
	}

	for _, value := range val {
		validator := value.(types.Validator)
		if validatorSelector(validator) {
			validators = append(validators, validator)
		}
	}

	return validators
}

func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator) {
	var selectAllValidators ValidatorSelectorFn = func(validator types.Validator) bool {
		return true
	}
	return k.GetAllValidatorsWithCondition(ctx, selectAllValidators)
}

func (k Keeper) GetAllAcceptedValidators(ctx sdk.Context) (validators []types.Validator) {
	var selectAcceptedValidators ValidatorSelectorFn = func(validator types.Validator) bool {
		return validator.Accepted
	}
	return k.GetAllValidatorsWithCondition(ctx, selectAcceptedValidators)
}

func (k Keeper) GetAllValidatorsAcceptedAndInSet(ctx sdk.Context) (validators []types.Validator) {
	var selectValidatorsAcceptedAndInSet ValidatorSelectorFn = func(validator types.Validator) bool {
		return validator.InSet || validator.Accepted
	}
	return k.GetAllValidatorsWithCondition(ctx, selectValidatorsAcceptedAndInSet)
}
