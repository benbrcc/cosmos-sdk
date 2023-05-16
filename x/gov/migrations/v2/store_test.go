package v2_test

import (
	"bytes"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/migrations/v1"
	v2 "github.com/cosmos/cosmos-sdk/x/gov/migrations/v2"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func TestMigrateStore(t *testing.T) {
	cdc := moduletestutil.MakeTestEncodingConfig().Codec
	govKey := storetypes.NewKVStoreKey("gov")
	ctx := testutil.DefaultContext(govKey, storetypes.NewTransientStoreKey("transient_test"))
	store := ctx.KVStore(govKey)

	_, _, addr1 := testdata.KeyTestPubAddr()
	proposalID := uint64(6)
	now := time.Now()
	// Use dummy value for keys where we don't test values.
	dummyValue := []byte("foo")
	// Use real values for votes, as we're testing weighted votes.
	oldVote := v1beta1.Vote{ProposalId: 1, Voter: "foobar", Option: v1beta1.OptionNoWithVeto}
	oldVoteValue := cdc.MustMarshal(&oldVote)
	newVote := v1beta1.Vote{ProposalId: 1, Voter: "foobar", Options: v1beta1.WeightedVoteOptions{{Option: v1beta1.OptionNoWithVeto, Weight: math.LegacyNewDec(1)}}}
	newVoteValue := cdc.MustMarshal(&newVote)

	testCases := []struct {
		name                               string
		oldKey, oldValue, newKey, newValue []byte
	}{
		{
			"ProposalKey",
			v1.ProposalKey(proposalID), dummyValue,
			types.ProposalKey(proposalID), dummyValue,
		},
		{
			"ActiveProposalQueue",
			v1.ActiveProposalQueueKey(proposalID, now), dummyValue,
			types.ActiveProposalQueueKey(proposalID, now), dummyValue,
		},
		{
			"InactiveProposalQueue",
			v1.InactiveProposalQueueKey(proposalID, now), dummyValue,
			types.InactiveProposalQueueKey(proposalID, now), dummyValue,
		},
		{
			"ProposalIDKey",
			v1.ProposalIDKey, dummyValue,
			types.ProposalIDKey, dummyValue,
		},
		{
			"DepositKey",
			v1.DepositKey(proposalID, addr1), dummyValue,
			depositKey(proposalID, addr1), dummyValue,
		},
		{
			"VotesKeyPrefix",
			v1.VoteKey(proposalID, addr1), oldVoteValue,
			voteKey(proposalID, addr1), newVoteValue,
		},
	}

	// Set all the old keys to the store
	for _, tc := range testCases {
		store.Set(tc.oldKey, tc.oldValue)
	}

	// Run migratio
	storeService := runtime.NewKVStoreService(govKey)
	err := v2.MigrateStore(ctx, storeService, cdc)
	require.NoError(t, err)

	// Make sure the new keys are set and old keys are deleted.
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if !bytes.Equal(tc.oldKey, tc.newKey) {
				require.Nil(t, store.Get(tc.oldKey))
			}
			require.Equal(t, tc.newValue, store.Get(tc.newKey))
		})
	}
}

// depositKey key of a specific deposit from the store.
// NOTE(tip): legacy, eventually remove me.
func depositKey(proposalID uint64, depositorAddr sdk.AccAddress) []byte {
	return append(append(types.DepositsKeyPrefix, sdk.Uint64ToBigEndian(proposalID)...), address.MustLengthPrefix(depositorAddr.Bytes())...)
}

func voteKey(proposalID uint64, addr sdk.AccAddress) []byte {
	return append(append(types.VotesKeyPrefix, sdk.Uint64ToBigEndian(proposalID)...), address.MustLengthPrefix(addr.Bytes())...)
}
