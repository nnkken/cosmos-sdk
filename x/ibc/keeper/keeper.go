package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	client "github.com/cosmos/cosmos-sdk/x/ibc/02-client"
	connection "github.com/cosmos/cosmos-sdk/x/ibc/03-connection"
	channel "github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	port "github.com/cosmos/cosmos-sdk/x/ibc/05-port"
	nft_transfer "github.com/cosmos/cosmos-sdk/x/ibc/17-nft_transfer"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/20-transfer"
)

// Keeper defines each ICS keeper for IBC
type Keeper struct {
	ClientKeeper      client.Keeper
	ConnectionKeeper  connection.Keeper
	ChannelKeeper     channel.Keeper
	PortKeeper        port.Keeper
	TransferKeeper    transfer.Keeper
	NFTTransferKeeper nft_transfer.Keeper
}

// NewKeeper creates a new ibc Keeper
func NewKeeper(
	cdc *codec.Codec, key sdk.StoreKey, codespace sdk.CodespaceType,
	bk transfer.BankKeeper, sk transfer.SupplyKeeper, nftk nft_transfer.NFTKeeper,
) Keeper {
	clientKeeper := client.NewKeeper(cdc, key, codespace)
	connectionKeeper := connection.NewKeeper(cdc, key, codespace, clientKeeper)
	portKeeper := port.NewKeeper(cdc, key, codespace)
	channelKeeper := channel.NewKeeper(cdc, key, codespace, clientKeeper, connectionKeeper, portKeeper)
	transferKeeper := transfer.NewKeeper(cdc, key, codespace, clientKeeper, connectionKeeper, channelKeeper, bk, sk)
	nftTransferKeeper := nft_transfer.NewKeeper(cdc, key, codespace, clientKeeper, connectionKeeper, channelKeeper, nftk)
	return Keeper{
		ClientKeeper:      clientKeeper,
		ConnectionKeeper:  connectionKeeper,
		ChannelKeeper:     channelKeeper,
		PortKeeper:        portKeeper,
		TransferKeeper:    transferKeeper,
		NFTTransferKeeper: nftTransferKeeper,
	}
}
