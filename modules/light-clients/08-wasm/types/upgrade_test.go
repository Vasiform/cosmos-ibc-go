package types_test

import (
	"encoding/json"
	"errors"

	wasmvm "github.com/CosmWasm/wasmvm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"

	wasmtesting "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/testing"
	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
	solomachine "github.com/cosmos/ibc-go/v8/modules/light-clients/06-solomachine"
)

// TestVerifyUpgrade currently only tests the interface into the contract.
// Test code is used in the grandpa contract.
// New client state, consensus state, and client metadata is expected to be set in the contract on success
func (suite *TypesTestSuite) TestVerifyUpgradeGrandpa() {
	var (
		upgradedClient         exported.ClientState
		upgradedConsState      exported.ConsensusState
		proofUpgradedClient    []byte
		proofUpgradedConsState []byte
		err                    error
	)

	testCases := []struct {
		name    string
		setup   func()
		expPass bool
	}{
		// TODO: fails with check upgradedClient.GetLatestHeight().GT(lastHeight) in VerifyUpgradeAndUpdateState
		// {
		// 	"successful upgrade",
		// 	func() {},
		// 	true,
		// },
		{
			"unsuccessful upgrade: invalid new client state",
			func() {
				upgradedClient = &solomachine.ClientState{}
			},
			false,
		},
		{
			"unsuccessful upgrade: invalid new consensus state",
			func() {
				upgradedConsState = &solomachine.ConsensusState{}
			},
			false,
		},
		{
			"unsuccessful upgrade: invalid client state proof",
			func() {
				proofUpgradedClient = []byte("invalid client state proof")
			},
			false,
		},
		{
			"unsuccessful upgrade: invalid consensus state proof",
			func() {
				proofUpgradedConsState = []byte("invalid consensus state proof")
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// reset suite
			suite.SetupWasmGrandpaWithChannel()
			clientState, ok := suite.chainA.App.GetIBCKeeper().ClientKeeper.GetClientState(suite.ctx, defaultWasmClientID)
			suite.Require().True(ok)
			upgradedClient = clientState
			upgradedConsState, ok = suite.chainA.App.GetIBCKeeper().ClientKeeper.GetLatestClientConsensusState(suite.ctx, defaultWasmClientID)
			suite.Require().True(ok)
			proofUpgradedClient = []byte("upgraded client state proof")
			proofUpgradedConsState = []byte("upgraded consensus state proof")

			tc.setup()

			err = clientState.VerifyUpgradeAndUpdateState(
				suite.ctx,
				suite.chainA.Codec,
				suite.store,
				upgradedClient,
				upgradedConsState,
				proofUpgradedClient,
				proofUpgradedConsState,
			)

			if tc.expPass {
				suite.Require().NoError(err)
				clientStateBz := suite.store.Get(host.ClientStateKey())
				suite.Require().NotEmpty(clientStateBz)
				newClientState := clienttypes.MustUnmarshalClientState(suite.chainA.Codec, clientStateBz)
				// Stubbed code will increment client state
				suite.Require().Equal(clientState.GetLatestHeight().Increment(), newClientState.GetLatestHeight())
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *TypesTestSuite) TestVerifyUpgradeAndUpdateState() {
	// TODO: PR 4934 exposes a ErrMockContract that can be re-used
	contractError := errors.New("contract error")
	var (
		endpoint               *wasmtesting.WasmEndpoint
		upgradedClient         exported.ClientState
		upgradedConsState      exported.ConsensusState
		proofUpgradedClient    []byte
		proofUpgradedConsState []byte
	)

	testCases := []struct {
		name      string
		malleate  func()
		expResult func(err error)
	}{
		{
			"success: successful upgrade",
			func() {
				suite.mockVM.RegisterSudoCallback(types.VerifyUpgradeAndUpdateStateMsg{}, func(_ wasmvm.Checksum, _ wasmvmtypes.Env, sudoMsg []byte, store wasmvm.KVStore, _ wasmvm.GoAPI, _ wasmvm.Querier, _ wasmvm.GasMeter, _ uint64, _ wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
					var payload types.SudoMsg

					err := json.Unmarshal(sudoMsg, &payload)
					suite.Require().NoError(err)

					// verify payload values
					suite.Require().Equal(upgradedClient, &payload.VerifyUpgradeAndUpdateState.UpgradeClientState)
					suite.Require().Equal(upgradedConsState, &payload.VerifyUpgradeAndUpdateState.UpgradeConsensusState)
					suite.Require().Equal(proofUpgradedClient, payload.VerifyUpgradeAndUpdateState.ProofUpgradeClient)
					suite.Require().Equal(proofUpgradedConsState, payload.VerifyUpgradeAndUpdateState.ProofUpgradeConsensusState)

					// verify other Sudo fields are nil
					suite.Require().Nil(payload.CheckSubstituteAndUpdateState)
					suite.Require().Nil(payload.UpdateState)
					suite.Require().Nil(payload.UpdateStateOnMisbehaviour)
					suite.Require().Nil(payload.VerifyMembership)
					suite.Require().Nil(payload.VerifyNonMembership)

					data, err := json.Marshal(types.EmptyResult{})
					suite.Require().NoError(err)

					// set new client state and consensus state
					store.Set(host.ClientStateKey(), clienttypes.MustMarshalClientState(suite.chainA.App.AppCodec(), upgradedClient))
					store.Set(host.ConsensusStateKey(upgradedClient.GetLatestHeight()), clienttypes.MustMarshalConsensusState(suite.chainA.App.AppCodec(), upgradedConsState))

					return &wasmvmtypes.Response{Data: data}, types.DefaultGasUsed, nil
				})
			},
			func(err error) {
				suite.Require().NoError(err)

				// verify new client state and consensus state
				clientStateBz := suite.store.Get(host.ClientStateKey())
				suite.Require().NotEmpty(clientStateBz)
				suite.Require().Equal(upgradedClient, clienttypes.MustUnmarshalClientState(suite.chainA.Codec, clientStateBz))

				consStateBz := suite.store.Get(host.ConsensusStateKey(upgradedClient.GetLatestHeight()))
				suite.Require().NotEmpty(consStateBz)
				suite.Require().Equal(upgradedConsState, clienttypes.MustUnmarshalConsensusState(suite.chainA.Codec, consStateBz))
			},
		},
		{
			"failure: upgraded client state is not wasm client state",
			func() {
				// set upgraded client state to solomachine client state
				upgradedClient = &solomachine.ClientState{}
			},
			func(err error) {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, clienttypes.ErrInvalidClient)

				// verify client state is unchanged
				clientStateBz := suite.store.Get(host.ClientStateKey())
				suite.Require().Equal(clienttypes.MustUnmarshalClientState(suite.chainA.Codec, clientStateBz), endpoint.GetClientState().(*types.ClientState))
			},
		},
		{
			"failure: upgraded consensus state is not wasm client state",
			func() {
				// set upgraded consensus state to solomachine consensus state
				upgradedConsState = &solomachine.ConsensusState{}
			},
			func(err error) {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, clienttypes.ErrInvalidConsensus)

				// verify client state is unchanged
				clientStateBz := suite.store.Get(host.ClientStateKey())
				suite.Require().Equal(clienttypes.MustUnmarshalClientState(suite.chainA.Codec, clientStateBz), endpoint.GetClientState().(*types.ClientState))
			},
		},
		{
			"failure: contract returns error",
			func() {
				suite.mockVM.RegisterSudoCallback(types.VerifyUpgradeAndUpdateStateMsg{}, func(_ wasmvm.Checksum, _ wasmvmtypes.Env, _ []byte, _ wasmvm.KVStore, _ wasmvm.GoAPI, _ wasmvm.Querier, _ wasmvm.GasMeter, _ uint64, _ wasmvmtypes.UFraction) (*wasmvmtypes.Response, uint64, error) {
					return nil, 0, contractError
				})
			},
			func(err error) {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, contractError)

				// verify client state is unchanged
				clientStateBz := suite.store.Get(host.ClientStateKey())
				suite.Require().Equal(clienttypes.MustUnmarshalClientState(suite.chainA.Codec, clientStateBz), endpoint.GetClientState().(*types.ClientState))
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// reset suite
			suite.SetupWasmWithMockVM()

			endpoint = wasmtesting.NewWasmEndpoint(suite.chainA)
			err := endpoint.CreateClient()
			suite.Require().NoError(err)

			clientState := endpoint.GetClientState().(*types.ClientState)

			upgradedClient = types.NewClientState([]byte("client state"), clientState.CodeHash, clienttypes.NewHeight(0, clientState.GetLatestHeight().GetRevisionHeight()+1))
			upgradedConsState = &types.ConsensusState{[]byte("consensus state")}

			tc.malleate()

			proofUpgradedClient = []byte("proof upgraded client")
			proofUpgradedConsState = []byte("proof upgraded consensus state")

			err = clientState.VerifyUpgradeAndUpdateState(
				suite.ctx,
				suite.chainA.Codec,
				suite.store,
				upgradedClient,
				upgradedConsState,
				proofUpgradedClient,
				proofUpgradedConsState,
			)

			tc.expResult(err)
		})
	}
}
