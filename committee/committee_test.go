// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package committee

import (
	"math/big"
	"testing"
	"time"

	"github.com/iotexproject/iotex-election/types"

	"github.com/stretchr/testify/require"
)

func TestVoteFilter(t *testing.T) {
	require := require.New(t)
	c := &committee{voteThreshold: big.NewInt(10)}
	vote1, err := types.NewVote(
		time.Now(),
		time.Hour,
		big.NewInt(3),
		big.NewInt(3),
		[]byte{},
		[]byte{},
		true,
	)
	require.NoError(err)
	require.True(c.voteFilter(vote1))

	vote2, err := types.NewVote(
		time.Now(),
		time.Hour,
		big.NewInt(30),
		big.NewInt(3),
		[]byte{},
		[]byte{},
		true,
	)
	require.NoError(err)
	require.False(c.voteFilter(vote2))
}
func TestCandidateFilter(t *testing.T) {
	require := require.New(t)
	cfg := getCfg()
	commp, err := NewCommittee(nil, cfg)
	require.NoError(err)
	// candidate1 selfStaking and score both smaller than committee's threshold
	candidate1 := types.NewCandidate(
		[]byte("candidate1"),
		[]byte("candidate1addr"),
		[]byte("operatorPubKey1"),
		[]byte("rewardPubKey1"),
		1,
	)
	candidate1.SetScore(big.NewInt(9))
	candidate1.SetSelfStakingTokens(big.NewInt(9))
	require.True(commp.(*committee).candidateFilter(candidate1))
	// candidate2 selfStaking is below committee's threshold,score is bigger than committee's threshold
	candidate2 := types.NewCandidate(
		[]byte("candidate2"),
		[]byte("candidate2addr"),
		[]byte("operatorPubKey2"),
		[]byte("rewardPubKey2"),
		1,
	)
	candidate2.SetScore(big.NewInt(11))
	candidate2.SetSelfStakingTokens(big.NewInt(9))
	require.True(commp.(*committee).candidateFilter(candidate2))
	// candidate3 selfStaking is bigger than committee's threshold,score is smaller than committee's threshold
	candidate3 := types.NewCandidate(
		[]byte("candidate3"),
		[]byte("candidate3addr"),
		[]byte("operatorPubKey3"),
		[]byte("rewardPubKey3"),
		1,
	)
	candidate3.SetScore(big.NewInt(9))
	candidate3.SetSelfStakingTokens(big.NewInt(11))
	require.True(commp.(*committee).candidateFilter(candidate3))
	// candidate3 selfStaking and score both bigger than committee's threshold
	candidate4 := types.NewCandidate(
		[]byte("candidate4"),
		[]byte("candidate4addr"),
		[]byte("operatorPubKey4"),
		[]byte("rewardPubKey4"),
		1,
	)
	candidate4.SetScore(big.NewInt(11))
	candidate4.SetSelfStakingTokens(big.NewInt(11))
	require.False(commp.(*committee).candidateFilter(candidate4))
}
func getCfg() (cfg Config) {
	cfg.NumOfRetries = 8
	cfg.BeaconChainAPIs = []string{"https://kovan.infura.io"}
	cfg.BeaconChainHeightInterval = 100
	cfg.BeaconChainStartHeight = 7368630
	cfg.RegisterContractAddress = "0xb4ca6cf2fe760517a3f92120acbe577311252663"
	cfg.StakingContractAddress = "0xdedf0c1610d8a75ca896d8c93a0dc39abf7daff4"
	cfg.PaginationSize = 100
	cfg.VoteThreshold = "10"
	cfg.ScoreThreshold = "10"
	cfg.SelfStakingThreshold = "10"
	cfg.CacheSize = 100
	return
}