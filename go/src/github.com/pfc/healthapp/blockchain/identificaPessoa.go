package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//IdentificaPessoa realiza uma transação de consulta para invocar a função correspondente no chaincode
func (setup *FabricSetup) IdentificaPessoa(cpf string) (string, error){

	//prepara os argumentos
	var args []string
	args = append(args, "identificaPessoa")
	args = append(args, cpf)

	//cria a proposta de transação e envia
	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}