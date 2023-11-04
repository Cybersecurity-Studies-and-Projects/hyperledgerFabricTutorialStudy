package blockchain

import(
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//InitPaciente realiza uma transação para invocar a função correspondente no chaincode
func (setup *FabricSetup) InitPaciente(dados []string) (string, error){

	//Prepara os argumentos 
	var args []string
	args = append(args, "initPaciente")
	for i := 0; i < len(dados); i++{
		args = append(args, dados[i])
	}

	// adiciona dados que serão visiveis na proposta, como uma descrição
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in initPaciente invoke")


	// Cria a proposta de transação e envia
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5]), []byte(args[6]), []byte(args[7]), []byte(args[8]), []byte(args[9]), []byte(args[10]), []byte(args[11]), []byte(args[12]), []byte(args[13]), []byte(args[14]), []byte(args[15])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}

	return string(response.TransactionID), nil
}

