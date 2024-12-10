// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package components

import (
	"os"

	"github.com/berachain/beacon-kit/config/spec"
	"github.com/berachain/beacon-kit/primitives/common"
)

const (
	ChainSpecTypeEnvVar  = "CHAIN_SPEC"
	DevnetChainSpecType  = "devnet"
	BetnetChainSpecType  = "betnet"
	BoonetChainSpecType  = "boonet"
	TestnetChainSpecType = "testnet"
)

// ProvideChainSpec provides the chain spec based on the environment variable.
func ProvideChainSpec() (common.ChainSpec, error) {
	// TODO: This is hood as fuck needs to be improved
	// but for now we ball to get CI unblocked.
	var (
		chainSpec common.ChainSpec
		err       error
	)
	switch os.Getenv(ChainSpecTypeEnvVar) {
	case DevnetChainSpecType:
		chainSpec, err = spec.DevnetChainSpec()
	case BetnetChainSpecType:
		chainSpec, err = spec.BetnetChainSpec()
	case BoonetChainSpecType:
		chainSpec, err = spec.BoonetChainSpec()
	case TestnetChainSpecType:
		fallthrough
	default:
		chainSpec, err = spec.TestnetChainSpec()
	}
	return chainSpec, err
}