// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package primitives

import "github.com/ethereum/go-ethereum/common"

const (
	RootLength = 32
)

type (
	// ExecutionAddress represents an address on the execution layer
	// which is derived via secp256k1 w/recovery bit.
	//
	// Related: https://eips.ethereum.org/EIPS/eip-55
	ExecutionAddress = common.Address

	// ExecutionHash represents a hash on the execution layer which is
	// currently a Keccak256 hash.
	ExecutionHash = common.Hash
)

//nolint:lll
type (
	// Slot as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	Slot = SSZUInt64

	// Epoch as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	Epoch = SSZUInt64

	// CommitteeIndex as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	CommitteeIndex = SSZUInt64

	// ValidatorIndex as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	ValidatorIndex = SSZUInt64

	// Root as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	Root Bytes32

	// Hash32 as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	Hash32 Bytes32

	// DomainType as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	DomainType Bytes4

	// Version as per the Ethereum 2.0 specification.
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	Version Bytes4

	// ForkDigest as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	ForkDigest Bytes4

	// Domain as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	Domain Bytes32

	// ValidatorIndex as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	BLSPubkey Bytes48

	// BLSSignature as per the Ethereum 2.0 Specification:
	// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
	BLSSignature Bytes96
)