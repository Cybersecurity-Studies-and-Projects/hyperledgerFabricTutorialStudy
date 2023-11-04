package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//GetRecord realiza uma transação de consulta para invocar a função correspondente no chaincode
func (setup *FabricSetup) GetRegistro(dados []string) ([]byte, error){

	//prepara os argumentos
	var args []string
	args = append(args, "getRegistro")
	for i:=0; i < len(dados); i++{
		args = append(args, dados[i])
	}

	//cria a proposta de transação e envia
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return nil, fmt.Errorf("failed to query: %v", err)
	}

	return response.Payload, nil
}