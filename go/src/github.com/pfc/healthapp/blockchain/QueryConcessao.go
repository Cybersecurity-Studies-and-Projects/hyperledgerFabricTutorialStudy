package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//QueryConcessao realiza uma transação para invocar a função correspondente no chaincode
func (setup *FabricSetup) QueryConcessao (dados []string) (string, error){
	
	//Prepara os argumentos
	var args []string
	args = append(args, "queryConcessao")
	for i:=0; i < len(dados); i++{
		args = append(args, dados[i])
	}

	// adiciona dados que serão visiveis na proposta, como uma descrição da transação
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in queryConcessao invoke")

	// cria a transação e envia
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}

	return string(response.TransactionID), nil

}