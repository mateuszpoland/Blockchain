package main

import (
	"os"
	"path/filepath"
)

type Tx struct {
	From  Account `json: "from"`
	To    Account `json: "to"`
	Value uint    `json: "value"`
	Data  string  `json: "data"`
}

type State struct {
	Balances  map[Account]uint
	txMempool []Tx
	dbFile    *os.File
}

func (t Tx) IsReward() bool {
	return t.Data == "reward"
}

func newStateFromDisk() (*State, error) {
	currentWorkingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	genesisFilePath := filepath.Join(currentWorkingDir, "database", "genesis.json")
	genesis, err = loadGenesis(genesisFilePath)

	if err != nil {
		return nil, err
	}

	balances := make(map[Account]uint)
	for account, balance := range genesis.Balances {
		balances[account] = balance
	}
}

// stan kont w tablicy State jest odtwarzany przez sekwencyjne dołączenie wszystkich eventów
txDbFilePath := filepath.Join(cwd, "database", "tx.db")
dbFile, err := os.OpenFile(txDbFilePath, os.O_APPEND|os.O_RDWR, 0600)
if err != nil {
	return nil, err
}

scanner = bufio.NewScanner(dbFile)
// alokujemy pusty stan - nowy obiekt za pomocą 'composite literal - czyli cos a la named constructor w php'
state = &State{balances, make([]Tx, 0), dbFile}
//lecimy po wpisach w pliku i transformujemy je na zmienne stanu:
for scanner.Scan() {
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	//deserializacja wpisu typu '{ "from": "Mati", "to": "Wróbel", "value": 20, "data": "reward"}' na obiekt typu Tx
	tx := new(Tx)
	json.Unmarshal(scanner.Bytes(), &tx)

	if err := state.apply(tx); err != nil {
		return nil, err
	}

	return state, nil
}
