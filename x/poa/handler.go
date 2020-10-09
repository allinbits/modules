package poa

import (
	"fmt"

	"github.com/allinbits/modules/x/poa/keeper"
	"github.com/allinbits/modules/x/poa/msg"
	"github.com/allinbits/modules/x/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgCreateValidatorPOA:
			return handleMsgCreateValidatorPOA(ctx, msg, k)
		case types.MsgVoteValidator:
			return handleMsgVoteValidator(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreateValidatorPOA(ctx sdk.Context, msg msg.MsgCreateValidatorPOA, k keeper.Keeper) (*sdk.Result, error) {
	if _, found := k.GetValidator(ctx, msg.Name); found {
		return nil, sdkerrors.Wrap(types.ErrBadValidatorAddr, fmt.Sprintf("unrecognized %s validator already exists: %T", types.ModuleName, msg))
	}

	validator := types.NewValidator(
		msg.Name,
		msg.Address,
		msg.PubKey,
		msg.Description,
	)

	k.SetValidator(ctx, msg.Name, validator)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			stakingtypes.EventTypeCreateValidator,
			sdk.NewAttribute(stakingtypes.AttributeKeyValidator, msg.Address.String()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgVoteValidator(ctx sdk.Context, msg msg.MsgVoteValidator, k keeper.Keeper) (*sdk.Result, error) {
	_, found := k.GetValidator(ctx, msg.Name)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNoValidatorFound, fmt.Sprintf("unrecognized %s validator does not exist: %T", types.ModuleName, msg))
	}

	val, found := k.GetValidatorByAddress(ctx, msg.Voter)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrEmptyValidatorAddr, fmt.Sprintf("unrecognized %s voting validator does not exist: %T", types.ModuleName, msg))
	}

	// if the validator hasn't been accepted they cannot vote
	if !val.Accepted {
		return nil, sdkerrors.Wrap(types.ErrNoAcceptedValidatorFound, fmt.Sprintf("error %s validator is not accepted by consensus: %T", types.ModuleName, msg))
	}

	vote := types.NewVote(
		msg.Voter,
		msg.Name,
		msg.InFavor,
	)

	k.SetVote(ctx, vote)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeVote,
			sdk.NewAttribute(stakingtypes.AttributeKeyValidator, msg.Voter.String()),
			sdk.NewAttribute(types.AttributeKeyCandidate, msg.Name),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
