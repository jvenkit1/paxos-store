package consensus

import (
	"github.com/jvenkit1/paxos-go"
	"github.com/sirupsen/logrus"
	"paxos-store/db"
)

type Request struct{
	Operation string // Can be Read or Write
	Key       string
	Value     string
}

func NewPaxosEnvironment() (*paxos.Environment, []paxos.Acceptor){
	/*
	 Creates a new paxos environment to which all decisions are forwarded.
	 Creating 3 proposer network
	*/
	network := paxos.NewPaxosEnvironment(100, 1, 2, 3, 200)
	var acceptorList []paxos.Acceptor
	acceptorID := 1
	for acceptorID < 4 {
		node := network.GetNodeNetwork(acceptorID)
		acceptorList = append(acceptorList, *paxos.NewAcceptor(acceptorID, node, 200))
		acceptorID+=1
	}

	return network, acceptorList
}

func (r *Request) StartConsensus() (string, error){
	/*
	 Starts a consensus cycle by proposing a Value to the list of proposers and
	*/
	network, acceptorList := NewPaxosEnvironment()

	command := r.Operation + ":" + r.Key + ":" + r.Value


	proposer := paxos.NewProposer(100, command, network.GetNodeNetwork(100), 1, 2, 3)

	logrus.Info("Starting Proposer")

	go proposer.Run()

	logrus.Info("Starting acceptor")
	for index :=range acceptorList {
		go acceptorList[index].Accept()
	}

	logrus.Info("Starting Learner")
	learner := paxos.NewLearner(200, network.GetNodeNetwork(200), 1, 2, 3)
	learnedValue := learner.Learn()

	logrus.Infof("Learned value %s", learnedValue)

	return learnedValue, nil
}

func (r *Request) Perform(){
	/*
		Perform Operation specified by object r
	*/

	if r.Operation == "Read" {
		data := db.Data{
			Key: r.Key,
		}
		data.GetData()
		r.Value = data.Value
	}else if r.Operation == "Write" {
		data := db.Data{
			Key: r.Key,
			Value: r.Value,
		}
		data.InsertData()
	}
}
