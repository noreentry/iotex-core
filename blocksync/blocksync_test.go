// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package blocksync

import (
	"net"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/iotexproject/iotex-core/actpool"
	bc "github.com/iotexproject/iotex-core/blockchain"
	"github.com/iotexproject/iotex-core/common"
	"github.com/iotexproject/iotex-core/config"
	"github.com/iotexproject/iotex-core/network"
	pb "github.com/iotexproject/iotex-core/proto"
	"github.com/iotexproject/iotex-core/state"
	"github.com/iotexproject/iotex-core/test/mock/mock_blockchain"
	"github.com/iotexproject/iotex-core/test/mock/mock_blocksync"
	"github.com/iotexproject/iotex-core/test/mock/mock_delegate"
)

func TestSyncTaskInterval(t *testing.T) {
	assert := assert.New(t)

	interval := time.Duration(0)

	cfgLightWeight := &config.Config{
		NodeType: config.LightweightType,
	}
	lightWeight := SyncTaskInterval(cfgLightWeight)
	assert.Equal(interval, lightWeight)

	cfgDelegate := &config.Config{
		NodeType: config.DelegateType,
		BlockSync: config.BlockSync{
			Interval: interval,
		},
	}
	delegate := SyncTaskInterval(cfgDelegate)
	assert.Equal(interval, delegate)

	cfgFullNode := &config.Config{
		NodeType: config.FullNodeType,
		BlockSync: config.BlockSync{
			Interval: interval,
		},
	}
	interval <<= 2
	fullNode := SyncTaskInterval(cfgFullNode)
	assert.Equal(interval, fullNode)
}

func generateP2P() *network.Overlay {
	c := &config.Network{
		Addr: "127.0.0.1:10001",
		MsgLogsCleaningInterval: 2 * time.Second,
		MsgLogRetention:         10 * time.Second,
		HealthCheckInterval:     time.Second,
		SilentInterval:          5 * time.Second,
		PeerMaintainerInterval:  time.Second,
		NumPeersLowerBound:      5,
		NumPeersUpperBound:      5,
		AllowMultiConnsPerIP:    true,
		RateLimitEnabled:        false,
		PingInterval:            time.Second,
		BootstrapNodes:          []string{"127.0.0.1:10001", "127.0.0.1:10002"},
		MaxMsgSize:              1024 * 1024 * 10,
		PeerDiscovery:           true,
	}
	return network.NewOverlay(c)
}

func TestNewBlockSyncer(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mPool := mock_delegate.NewMockPool(ctrl)
	mPool.EXPECT().AllDelegates().Times(3).Return([]net.Addr{common.NewNode("", "123")}, nil)
	mPool.EXPECT().AnotherDelegate(gomock.Any()).Times(2).Return(common.NewNode("", "123"))

	p2p := generateP2P()

	// Lightweight
	cfgLightWeight := &config.Config{
		NodeType: config.LightweightType,
	}

	bsLightWeight, err := NewBlockSyncer(cfgLightWeight, nil, nil, p2p, mPool)
	assert.NotNil(err)
	assert.Nil(bsLightWeight)

	// Delegate
	cfgDelegate := &config.Config{
		NodeType: config.DelegateType,
	}

	bsDelegate, err := NewBlockSyncer(cfgDelegate, nil, nil, p2p, mPool)
	assert.Nil(err)
	assert.Equal("123", bsDelegate.(*blockSyncer).fnd)

	// FullNode
	cfgFullNode := &config.Config{
		NodeType: config.FullNodeType,
	}

	bsFullNode, err := NewBlockSyncer(cfgFullNode, nil, nil, p2p, mPool)
	assert.Nil(err)
	assert.Equal("123", bsFullNode.(*blockSyncer).fnd)
}

func TestBlockSyncer_P2P(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mPool := mock_delegate.NewMockPool(ctrl)
	mPool.EXPECT().AllDelegates().Times(1).Return([]net.Addr{common.NewNode("", "123")}, nil)
	mPool.EXPECT().AnotherDelegate(gomock.Any()).Times(1).Return(common.NewNode("", "123"))

	p2p := generateP2P()

	cfgFullNode := &config.Config{
		NodeType: config.FullNodeType,
	}
	bs, err := NewBlockSyncer(cfgFullNode, nil, nil, p2p, mPool)
	assert.Nil(err)
	assert.Equal(p2p, bs.P2P())
}

func TestBlockSyncer_Start(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mBs := mock_blocksync.NewMockBlockSync(ctrl)
	mBs.EXPECT().Start().Times(1)
	assert.Nil(mBs.Start())
}

func TestBlockSyncer_Stop(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mBs := mock_blocksync.NewMockBlockSync(ctrl)
	mBs.EXPECT().Stop().Times(1)
	assert.Nil(mBs.Stop())
}

func TestBlockSyncer_ProcessSyncRequest(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mPool := mock_delegate.NewMockPool(ctrl)
	mPool.EXPECT().AllDelegates().Times(1).Return([]net.Addr{common.NewNode("", "123")}, nil)
	mPool.EXPECT().AnotherDelegate(gomock.Any()).Times(1).Return(common.NewNode("", "123"))

	mBc := mock_blockchain.NewMockBlockchain(ctrl)
	mBc.EXPECT().GetBlockByHeight(gomock.Any()).AnyTimes()
	p2p := generateP2P()

	cfgFullNode := &config.Config{
		NodeType: config.FullNodeType,
	}

	bs, err := NewBlockSyncer(cfgFullNode, mBc, nil, p2p, mPool)
	assert.Nil(err)

	pbBs := &pb.BlockSync{
		Start: 1,
		End:   1,
	}

	bs.(*blockSyncer).ackSyncReq = false
	assert.Nil(bs.ProcessSyncRequest("", pbBs))
}

func TestBlockSyncer_ProcessBlock_TipHeightError(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mPool := mock_delegate.NewMockPool(ctrl)
	mPool.EXPECT().AllDelegates().Times(1).Return([]net.Addr{common.NewNode("", "123")}, nil)
	mPool.EXPECT().AnotherDelegate(gomock.Any()).Times(1).Return(common.NewNode("", "123"))

	mBc := mock_blockchain.NewMockBlockchain(ctrl)
	// TipHeight return ERROR
	mBc.EXPECT().TipHeight().Times(1).Return(uint64(0), errors.New("Error"))
	p2p := generateP2P()

	cfgFullNode := &config.Config{
		NodeType: config.FullNodeType,
	}

	bs, err := NewBlockSyncer(cfgFullNode, mBc, nil, p2p, mPool)
	assert.Nil(err)
	blk := bc.NewBlock(uint32(123), uint64(4), common.Hash32B{}, nil, nil)
	bs.(*blockSyncer).ackBlockCommit = false
	assert.Nil(bs.ProcessBlock(blk))

	bs.(*blockSyncer).ackBlockCommit = true
	assert.Error(bs.ProcessBlock(blk))
}

func TestBlockSyncer_ProcessBlock_TipHeight(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mPool := mock_delegate.NewMockPool(ctrl)
	mPool.EXPECT().AllDelegates().Times(1).Return([]net.Addr{common.NewNode("", "123")}, nil)
	mPool.EXPECT().AnotherDelegate(gomock.Any()).Times(1).Return(common.NewNode("", "123"))

	mBc := mock_blockchain.NewMockBlockchain(ctrl)
	mBc.EXPECT().TipHeight().AnyTimes().Return(uint64(5), nil)
	mBc.EXPECT().CommitBlock(gomock.Any()).AnyTimes()

	sf, err := state.NewFactory(nil, state.InMemTrieOption())
	assert.NotNil(sf)
	apConfig := config.ActPool{8192, 256}
	ap, err := actpool.NewActPool(sf, apConfig)

	p2p := generateP2P()

	cfgFullNode := &config.Config{
		NodeType: config.FullNodeType,
	}

	bs, err := NewBlockSyncer(cfgFullNode, mBc, ap, p2p, mPool)
	assert.Nil(err)
	blk := bc.NewBlock(uint32(123), uint64(4), common.Hash32B{}, nil, nil)

	bs.(*blockSyncer).ackBlockCommit = true
	// less than tip height
	assert.Error(bs.ProcessBlock(blk))

	// special case
	bs.(*blockSyncer).state = Idle
	blkHeightSpecial := bc.NewBlock(uint32(123), uint64(6), common.Hash32B{}, nil, nil)
	assert.Nil(bs.ProcessBlock(blkHeightSpecial))

	// < block height
	blkHeightLess := bc.NewBlock(uint32(123), uint64(4), common.Hash32B{}, nil, nil)
	assert.Error(bs.ProcessBlock(blkHeightLess))

	// > block height
	blkHeightMore := bc.NewBlock(uint32(123), uint64(7), common.Hash32B{}, nil, nil)
	assert.Nil(bs.ProcessBlock(blkHeightMore))
}

func TestBlockSyncer_ProcessBlockSync(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mPool := mock_delegate.NewMockPool(ctrl)
	mPool.EXPECT().AllDelegates().Times(1).Return([]net.Addr{common.NewNode("", "123")}, nil)
	mPool.EXPECT().AnotherDelegate(gomock.Any()).Times(1).Return(common.NewNode("", "123"))

	mBc := mock_blockchain.NewMockBlockchain(ctrl)
	mBc.EXPECT().TipHeight().Times(1).Return(uint64(0), errors.New("Error"))
	mBc.EXPECT().TipHeight().Times(1).Return(uint64(5), nil)
	mBc.EXPECT().TipHeight().Times(1).Return(uint64(6), nil)

	sf, err := state.NewFactory(nil, state.InMemTrieOption())
	assert.NotNil(sf)
	apConfig := config.ActPool{8192, 256}
	ap, err := actpool.NewActPool(sf, apConfig)

	p2p := generateP2P()

	cfgFullNode := &config.Config{
		NodeType: config.FullNodeType,
	}

	bs, err := NewBlockSyncer(cfgFullNode, mBc, ap, p2p, mPool)
	assert.Nil(err)
	blk := bc.NewBlock(uint32(123), uint64(4), common.Hash32B{}, nil, nil)
	bs.(*blockSyncer).ackBlockSync = false
	assert.Nil(bs.ProcessBlockSync(blk))

	bs.(*blockSyncer).ackBlockSync = true
	assert.Error(bs.ProcessBlockSync(blk))
	assert.Nil(bs.ProcessBlockSync(blk))
	assert.Nil(bs.ProcessBlockSync(blk))
}
