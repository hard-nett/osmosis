package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgInboundTransfer{}

func NewMsgInboundTransfer(
	sender string,
	destAddr string,
	assetID AssetID,
	amount math.Int,
) *MsgInboundTransfer {
	return &MsgInboundTransfer{
		Sender:   sender,
		DestAddr: destAddr,
		AssetId:  assetID,
		Amount:   amount,
	}
}

func (m MsgInboundTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(m.DestAddr)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid destination address (%s)", err)
	}

	err = m.AssetId.Validate()
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidAssetID, err.Error())
	}

	// check if amount > 0
	if !m.Amount.IsPositive() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "Amount should be positive: %s", m.Amount.String())
	}

	return nil
}

func (m MsgInboundTransfer) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgOutboundTransfer{}

func NewMsgOutboundTransfer(
	sender string,
	destAddr string,
	assetID AssetID,
	amount math.Int,
) *MsgOutboundTransfer {
	return &MsgOutboundTransfer{
		Sender:   sender,
		DestAddr: destAddr,
		AssetId:  assetID,
		Amount:   amount,
	}
}

func (m MsgOutboundTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(m.DestAddr)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid destination address (%s)", err)
	}

	err = m.AssetId.Validate()
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidAssetID, err.Error())
	}

	// check if amount > 0
	if !m.Amount.IsPositive() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "Amount should be positive: %s", m.Amount.String())
	}

	return nil
}

func (m MsgOutboundTransfer) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgUpdateParams{}

func NewMsgUpdateParams(
	sender string,
	newParams Params,
) *MsgUpdateParams {
	return &MsgUpdateParams{
		Sender:    sender,
		NewParams: newParams,
	}
}

func (m MsgUpdateParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	err = m.NewParams.Validate()
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidParams, err.Error())
	}
	return nil
}

func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgChangeAssetStatus{}

func NewMsgChangeAssetStatus(
	sender string,
	assetID AssetID,
	newStatus AssetStatus,
) *MsgChangeAssetStatus {
	return &MsgChangeAssetStatus{
		Sender:    sender,
		AssetId:   assetID,
		NewStatus: newStatus,
	}
}

func (m MsgChangeAssetStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	err = m.AssetId.Validate()
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidAssetID, err.Error())
	}

	err = m.NewStatus.Validate()
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidAssetStatus, err.Error())
	}
	return nil
}

func (m MsgChangeAssetStatus) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}
