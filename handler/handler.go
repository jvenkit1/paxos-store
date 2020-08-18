package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"paxos-store/consensus"
	"strings"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	/*
		Handler for GET Operations involving reading data from the in-memory store.
	*/
	// Parse input json

	decoder := json.NewDecoder(r.Body)
	var input consensus.Request
	err := decoder.Decode(&input)
	if err != nil {
		// Return 400
		w.WriteHeader(400)
		w.Write([]byte(`{"message": "Incorrect request format"}`))
	}else{
		// input is now decoded
		learnedValue, err := input.StartConsensus()
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(`{"message": "Consensus not reached. Retry"}`))
		}

		splitString := strings.Split(learnedValue, ":")
		finalString := consensus.Request{
			Operation: "Read",
			Key: splitString[1],
		}
		finalString.Perform()
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf(`{"Key":"%s", "Value": "%s"}`, finalString.Key, finalString.Value)))
	}
}

func WeakConsistent(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var input consensus.Request
	err := decoder.Decode(&input)
	if err != nil {
		// Return 400
		w.WriteHeader(400)
		w.Write([]byte(`{"message": "Incorrect request format"}`))
	}else{
		// input is now decoded
		//learnedValue, err := input.StartConsensus()
		//if err != nil {
		//	w.WriteHeader(400)
		//	w.Write([]byte(`{"message": "Consensus not reached. Retry"}`))
		//}
		//
		//splitString := strings.Split(learnedValue, ":")
		//finalString := consensus.Request{
		//	Operation: "Read",
		//	Key: splitString[1],
		//}
		input.Operation = "Read"
		input.Perform()
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf(`{"Key":"%s", "Value": "%s"}`, input.Key, input.Value)))
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request){
	/*
		Handler for POST Operations involving inserting data into the in-memory store.
	*/
	decoder := json.NewDecoder(r.Body)
	var input consensus.Request
	err := decoder.Decode(&input)
	if err != nil {
		// Return 400
		w.WriteHeader(400)
		w.Write([]byte(`{"message": "Incorrect request format"}`))
	}else{
		// input is now decoded
		learnedValue, err := input.StartConsensus()
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(`{"message": "Consensus not reached. Retry"}`))
		}

		splitString := strings.Split(learnedValue, ":")
		finalString := consensus.Request{
			Operation: "Write",
			Key: splitString[1],
			Value: splitString[2],
		}
		finalString.Perform()
		w.WriteHeader(201)
		w.Write([]byte(fmt.Sprintf(`{"Key":"%s", "Value": "%s"}`, finalString.Key, finalString.Value)))
	}
}
