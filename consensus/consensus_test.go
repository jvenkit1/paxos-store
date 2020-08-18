package consensus

import(
	"github.com/jvenkit1/paxos-go"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestConsensusWithSingleData(t *testing.T){
	network, acceptorList := NewPaxosEnvironment()

	inputString := "Hello World"

	// Create Proposer
	proposer := paxos.NewProposer(100, inputString, network.GetNodeNetwork(100), 1, 2, 3)
	go proposer.Run()

	for index :=range acceptorList {
		go acceptorList[index].Accept()
	}

	// Create learner.
	learner := paxos.NewLearner(200, network.GetNodeNetwork(200), 1, 2, 3)
	learnedValue := learner.Learn()

	logrus.Infof("Learner picked up value %s", learnedValue)

	if learnedValue != inputString {
		t.Errorf("Learner learned wrong proposal")
	}
}