package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//QueryRevogacao realiza uma transação de consulta para invocar a função correspondente no chaincode
func (setup *FabricSetup) QueryRevogacao(dados []string) (string, error){

	//prepara os argumentos
	var args []string
	args = append(args, "queryRevogacao")
	for i:=0; i<len(dados);i++{
		args = append(args, dados[i])
	}

	// adiciona dados que serão visiveis na proposta, como uma descrição da transação
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in queryConcessao invoke")

	//cria a proposta de transação e envia
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}