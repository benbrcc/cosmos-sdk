package types

import (
	"encoding/json"
	fmt "fmt"
	strings "strings"

	"github.com/cosmos/gogoproto/proto"
	protov2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"cosmossdk.io/x/tx/signing"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

type (
	// Msg is an empty tag interface for transaction messages.
	Msg interface{}

	// LegacyMsg defines the interface a transaction message needed to fulfill up through
	// v0.47.
	LegacyMsg interface {
		proto.Message

		// ValidateBasic does a simple validation check that
		// doesn't require access to any other information.
		ValidateBasic() error

		// GetSigners returns the addrs of signers that must sign.
		// CONTRACT: All signatures must be present to be valid.
		// CONTRACT: Returns addrs in some deterministic order.
		GetSigners() []AccAddress
	}

	// Fee defines an interface for an application application-defined concrete
	// transaction type to be able to set and return the transaction fee.
	Fee interface {
		GetGas() uint64
		GetAmount() Coins
	}

	// Signature defines an interface for an application application-defined
	// concrete transaction type to be able to set and return transaction signatures.
	Signature interface {
		GetPubKey() cryptotypes.PubKey
		GetSignature() []byte
	}

	// Tx defines the interface a transaction must fulfill.
	Tx interface {
		// GetMsgs gets the all the transaction's messages.
		GetMsgs() []Msg

		// ValidateBasic does a simple and lightweight validation check that doesn't
		// require access to any other information.
		ValidateBasic(getSignersCtx *signing.GetSignersContext) error
	}

	// FeeTx defines the interface to be implemented by Tx to use the FeeDecorators
	FeeTx interface {
		Tx
		GetGas() uint64
		GetFee() Coins
		FeePayer() AccAddress
		FeeGranter() AccAddress
	}

	// TxWithMemo must have GetMemo() method to use ValidateMemoDecorator
	TxWithMemo interface {
		Tx
		GetMemo() string
	}

	// TxWithTimeoutHeight extends the Tx interface by allowing a transaction to
	// set a height timeout.
	TxWithTimeoutHeight interface {
		Tx

		GetTimeoutHeight() uint64
	}
)

// TxDecoder unmarshals transaction bytes
type TxDecoder func(txBytes []byte) (Tx, error)

// TxEncoder marshals transaction to bytes
type TxEncoder func(tx Tx) ([]byte, error)

// MsgTypeURL returns the TypeURL of a `sdk.Msg`.
func MsgTypeURL(msg Msg) string {
	if msg, ok := msg.(protov2.Message); ok {
		return "/" + string(msg.ProtoReflect().Descriptor().FullName())
	} else if msg, ok := msg.(proto.Message); ok {
		return "/" + proto.MessageName(msg)
	} else {
		panic(fmt.Errorf("%T is not a proto message", msg))
	}
}

// GetMsgFromTypeURL returns a `sdk.Msg` message type from a type URL
func GetMsgFromTypeURL(cdc codec.Codec, input string) (Msg, error) {
	var msg Msg
	bz, err := json.Marshal(struct {
		Type string `json:"@type"`
	}{
		Type: input,
	})
	if err != nil {
		return nil, err
	}

	if err := cdc.UnmarshalInterfaceJSON(bz, &msg); err != nil {
		return nil, fmt.Errorf("failed to determine sdk.Msg for %s URL : %w", input, err)
	}

	return msg, nil
}

// GetModuleNameFromTypeURL assumes that module name is the second element of the msg type URL
// e.g. "cosmos.bank.v1beta1.MsgSend" => "bank"
// It returns an empty string if the input is not a valid type URL
func GetModuleNameFromTypeURL(input string) string {
	moduleName := strings.Split(input, ".")
	if len(moduleName) > 1 {
		return moduleName[1]
	}

	return ""
}

func ValidateBasic(msg Msg) error {
	if validateBasic, ok := msg.(interface{ ValidateBasic() error }); ok {
		err := validateBasic.ValidateBasic()
		if err != nil {
			return err
		}
	}

	return nil
}

func GetSigners(msg Msg, getSignersCtx *signing.GetSignersContext) ([]string, error) {
	if legacyMsg, ok := msg.(LegacyMsg); ok {
		var signers []string
		for _, addr := range legacyMsg.GetSigners() {
			signer := addr.String()
			signers = append(signers, signer)
		}
		return signers, nil
	} else if msgv2, ok := msg.(protov2.Message); ok {
		return getSignersCtx.GetSigners(msgv2)
	} else if msgv1, ok := msg.(proto.Message); ok {
		bz, err := proto.Marshal(msgv1)
		if err != nil {
			return nil, err
		}

		return getSignersCtx.GetSignersForAny(&anypb.Any{
			TypeUrl: MsgTypeURL(msgv1),
			Value:   bz,
		})
	} else {
		return nil, fmt.Errorf("%T is not a proto message", msg)
	}
}
