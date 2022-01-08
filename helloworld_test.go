package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
)

// Encapsulates test state for re-use between unit tests
type HelloworldTestSuite struct {
	suite.Suite
	auth       *bind.TransactOpts
	address    common.Address
	gAlloc     core.GenesisAlloc
	sim        *backends.SimulatedBackend
	helloworld *Helloworld
}

func TestRunHelloworldSuite(t *testing.T) {
	suite.Run(t, new(HelloworldTestSuite))
}

// Configure blockchain simulator for helloworld contract
func (s *HelloworldTestSuite) SetupTest() {
	var e error

	// create a private key to populate a signer function for authorizing transactions
	key, _ := crypto.GenerateKey()

	// return *TransactOpts that points to an address from which we can make transactions
	s.auth, e = bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	s.Nil(e)

	// create genesis block with an account allocated with funds
	s.address = s.auth.From
	s.gAlloc = map[common.Address]core.GenesisAccount{
		s.address: {Balance: big.NewInt(100000000000000)},
	}

	// create a fully mocked simulated blockchain backed by in-memory database
	s.auth.GasLimit = uint64(0)
	s.auth.GasPrice = big.NewInt(875000000)
	s.sim = backends.NewSimulatedBackend(s.gAlloc, s.auth.GasLimit)

	// deploy helloworld contract
	_, _, hw, e := DeployHelloworld(s.auth, s.sim)
	s.helloworld = hw
	s.Nil(e)
	s.sim.Commit()
}

// Test helloworld contract Say function
func (s *HelloworldTestSuite) TestSay() {
	str, err := s.helloworld.Say(nil)
	s.Equal("hello etherworld", str)
	s.Nil(err)
}
