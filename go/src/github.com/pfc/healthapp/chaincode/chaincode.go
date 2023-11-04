/*
	Implementação do chaincode da rede de apoio, temporariamente
	denominada HealthNet.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	_"strings"
	"time"
	"encoding/base64"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Chaincode struct {
}

/*
	A estrutura paciente determina as informações armazenadas
	do paciente, nesta versão segue o padrão de prontuários
	utilizados em clínicas e hospitais, com informações básicas,
	tais como:
	- Nome completo
	- RG
	- CPF
	- Data de nascimento
	- Sexo
	- Nome da mãe
	- Naturalidade(Cidade e estado de nascimento)
	- Endereço Completo: Nome da via; número; complemento;
			     bairro/distrito; município; estado;
			     CEP;
*/
/*type endereco struct{
	Rua	       string   `json:"nome"`
	Numero	       int      `json:"numero"`
	Complemento    string   `json:"complemento"`
	Bairro         string   `json:"bairro"`
	Municipio      string   `json:"municipio"`
	Estado         string   `json:"estado"`
	CEP    	       string   `json:"cep"`
}
*/

type listaDeInteressados struct {
	ObjectType      string `json:"docType"`
	Nome            string `json:"nome"`
	CPF             string `json:"cpf"`
	DadosPermitidos string `json:"dados"`
}

type paciente struct {
	ObjectType     string                `json:"docType"`
	Nome           string                `json:"nome"`
	RG             string                `json:"rg"`
	CPF            string                `json:"cpf"`
	DataNascimento string                `json:"data_nasc"`
	Sexo           string                `json:"sexo"`
	NomeMae        string                `json:"nome_mae"`
	Naturalidade   string                `json:"naturalidade"`
	Rua            string                `json:"rua"`
	Numero         int                   `json:"numero"`
	Complemento    string                `json:"complemento"`
	Bairro         string                `json:"bairro"`
	Municipio      string                `json:"municipio"`
	Estado         string                `json:"estado"`
	CEP            string                `json:"cep"`
	Interessados   []listaDeInteressados `json:"interessados"`
	Dados          []string              `json:"dadosInseridos"`
}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Erro ao ininiciar Chaincode: %s\n", err)
	}
}

/*
*	Init inicializa o chaincode
 */
func (t *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

/*
*	A função Invoke é utilizada pelos peers(pares)
*	para invocar os métodos relacionados ao chaincode,
*	dessa forma, serve como porta de entrada para a
*	execução do contrato inteligente
 */
func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcao, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoke esta rodando " + funcao)

	switch funcao {
	case "initPaciente":
		return t.initPaciente(stub)
	case "identificaPessoa":
		return t.identificaPessoa(stub, args)
	case "UpdatePerson":
		return t.UpdatePerson(stub, args)
	case "getHistoryForPerson":
		return t.getHistoryForPerson(stub, args)
	case "queryConcessao":
		return t.queryConcessao(stub, args)
	case "queryRevogacao":
		return t.queryRevogacao(stub, args)
	case "queryInsercao":
		return t.queryInsercao(stub, args)
	case "getRegistro":
		return t.getRegistro(stub, args)
	default:
		fmt.Println("Invoke nao encontrou a funcao: " + funcao)
		return shim.Error("Invocacao de funcao desconhecida")
	}
}

func (t *Chaincode) getRegistro(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	//verifica quantidade de argumentos
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	//armazena argumentos nas respectivas variáveis 
	cpfPaciente := args[0]
	cpfInteressado := args[1]

	//obtem registro na blockchain
	pacienteAsBytes, err := stub.GetState(cpfPaciente)
	if err != nil {
		return shim.Error("Failed to get person:" + err.Error())
	}else if pacienteAsBytes == nil {
		return shim.Error("person does not exist")
	}

	pacienteConcedente := &paciente{}
	err = json.Unmarshal(pacienteAsBytes, pacienteConcedente) //realiza unmarshal do registro do paciente
	if err != nil {
		return shim.Error(err.Error())
	}

	//verifica se cpf do interessado corresponde a lista de interessados informada
	index := 0
	for i, elem := range pacienteConcedente.Interessados {
		if elem.CPF == cpfInteressado {
			index = i
			break
		}
	}
	if pacienteConcedente.Interessados[index].CPF != cpfInteressado {
		return shim.Error("Interessado " + cpfInteressado + " does not exist")
	}

	return shim.Success(pacienteAsBytes)
}

func (t *Chaincode) identificaPessoa(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	cpfPessoa := args[0]
	fmt.Println("- start identificacao " + args[0])
	pessoaAsBytes, err := stub.GetState(cpfPessoa)
	if err != nil {
		return shim.Error("Failed to get person:" + err.Error())
	}
	if pessoaAsBytes == nil {
		return shim.Error("person does not exist")
	}
	result, _ := json.Marshal(cpfPessoa + " SUCCESS!!!")
	return shim.Success(result)

}
func (t *Chaincode) queryConcessao(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	cpfPaciente := args[0]
	fmt.Println("- start concederpermissao ", args[2], args[1])
	pacienteAsBytes, err := stub.GetState(cpfPaciente)
	if err != nil {
		return shim.Error("Failed to get paciente:" + err.Error())
	} else if pacienteAsBytes == nil {
		return shim.Error("paciente does not exist")
	}
	pacienteConcedente := &paciente{}
	err = json.Unmarshal(pacienteAsBytes, pacienteConcedente) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	for _, elem := range pacienteConcedente.Interessados {
		if elem.CPF == args[2] {
			return shim.Success([]byte("paciente já está na lista de interessados"))
		}
	}

	var ListanovoInteressado listaDeInteressados
	ListanovoInteressado.ObjectType = "interessado"
	ListanovoInteressado.Nome = args[1]
	ListanovoInteressado.CPF = args[2]
	ListanovoInteressado.DadosPermitidos = args[3]

	pacienteConcedente.Interessados = append(pacienteConcedente.Interessados, ListanovoInteressado)
	pacienteJSONasBytes, _ := json.Marshal(pacienteConcedente)
	//pacienteJSONasBytes = []byte(strings.Replace(string(pacienteJSONasBytes), "\"", `\"`, -1))

	err = stub.DelState(pacienteConcedente.CPF) //reescreve o paciente
	if err != nil {
		return shim.Error(err.Error())
	}

	//NOVO TRECHO ADICIONADO
	err = stub.PutState(pacienteConcedente.CPF, pacienteJSONasBytes) //reescreve o paciente
	if err != nil {
		return shim.Error(err.Error())
	}

	encoded := base64.StdEncoding.EncodeToString(pacienteJSONasBytes)
	return shim.Success([]byte(encoded))
}

func (t *Chaincode) UpdatePerson(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 received ")
	}
	pacienteAsBytes, erro :=  base64.StdEncoding.DecodeString(args[0])
	if erro != nil {
		return shim.Error(erro.Error())
	}
	//return shim.Success(pacienteAsBytes)
	pacienteConcedente := paciente{}
	err := json.Unmarshal(pacienteAsBytes, &pacienteConcedente) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	pacienteJSONasBytes, _ := json.Marshal(pacienteConcedente)

	err = stub.PutState(pacienteConcedente.CPF, pacienteAsBytes) //reescreve o paciente
	if err != nil {
		return shim.Error(err.Error())
	}
	encoded := base64.StdEncoding.EncodeToString(pacienteJSONasBytes)
	return shim.Success([]byte(encoded))
}

func (t *Chaincode) queryRevogacao(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	cpfPaciente := args[0]
	pacienteAsBytes, err := stub.GetState(cpfPaciente)
	if err != nil {
		return shim.Error("Failed to get paciente:" + err.Error())
	} else if pacienteAsBytes == nil {
		return shim.Error("paciente does not exist")
	}
	pacienteConcedente := &paciente{}
	err = json.Unmarshal(pacienteAsBytes, pacienteConcedente) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	index := 0
	for i, elem := range pacienteConcedente.Interessados {
		if elem.CPF == args[1] {
			index = i
			break
		}
	}
	if pacienteConcedente.Interessados[index].CPF != args[1] {
		return shim.Error("Interessado " + args[1] + " does not exist")
	}
	pacienteConcedente.Interessados = append(pacienteConcedente.Interessados[:index], pacienteConcedente.Interessados[index+1:]...)
	pacienteJSONasBytes, _ := json.Marshal(pacienteConcedente)
	//pacienteJSONasBytes = []byte(strings.Replace(string(pacienteJSONasBytes), "\"", `\"`, -1))

	err = stub.DelState(cpfPaciente) //reescreve o paciente
	if err != nil {
		return shim.Error(err.Error())
	}

	//NOVO TRECHO ADICIONADO
	err = stub.PutState(cpfPaciente, pacienteJSONasBytes) //reescreve o paciente
	if err != nil {
		return shim.Error(err.Error())
	}
	
	encoded := base64.StdEncoding.EncodeToString(pacienteJSONasBytes)
	return shim.Success([]byte(encoded))
}

func (t *Chaincode) queryInsercao(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	cpfPaciente := args[0]
	pacienteAsBytes, err := stub.GetState(cpfPaciente)
	if err != nil {
		return shim.Error("Failed to get paciente:" + err.Error())
	} else if pacienteAsBytes == nil {
		return shim.Error("paciente does not exist")
	}
	pacienteConcedente := &paciente{}
	err = json.Unmarshal(pacienteAsBytes, pacienteConcedente) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	if pacienteConcedente.Dados == nil {
		var Data []string
		pacienteConcedente.Dados = Data
	}
	pacienteConcedente.Dados = append(pacienteConcedente.Dados, args[1])

	pacienteJSONasBytes, _ := json.Marshal(pacienteConcedente)
	//pacienteJSONasBytes = []byte(strings.Replace(string(pacienteJSONasBytes), "\"", `\"`, -1))
	err = stub.DelState(pacienteConcedente.CPF) //reescreve o paciente
	if err != nil {
		return shim.Error(err.Error())
	}
	encoded := base64.StdEncoding.EncodeToString(pacienteJSONasBytes)
	return shim.Success([]byte(encoded))
}

/*
	A função initPaciente cria um novo paciente e armazena seus dados no banco de dados,
*/

func (t *Chaincode) initPaciente(stub shim.ChaincodeStubInterface) pb.Response {

	_, args := stub.GetFunctionAndParameters()
	var err error
	var entradaPaciente paciente

	// checagem de entrada
	fmt.Println("- iniciar initPaciente")
	str := " "
	for i := 0; i < len(args); i++ {
		str = str + " " + args[i]
	}

	if len(args) != 15 {
		return shim.Error("Numero incorreto de argumentos " + str)
	}
	Atuacao := args[0]
	entradaPaciente.Nome = args[1]
	entradaPaciente.RG = args[2]
	entradaPaciente.CPF = args[3]
	entradaPaciente.DataNascimento = args[4]
	entradaPaciente.Sexo = args[5]
	entradaPaciente.NomeMae = args[6]
	entradaPaciente.Naturalidade = args[7]
	entradaPaciente.Rua = args[8]
	entradaPaciente.Numero, err = strconv.Atoi(args[9])
	entradaPaciente.Complemento = args[10]
	entradaPaciente.Bairro = args[11]
	entradaPaciente.Municipio = args[12]
	entradaPaciente.Estado = args[13]
	entradaPaciente.CEP = args[14]
	var inT []listaDeInteressados
	entradaPaciente.Interessados = inT
	entradaPaciente.Interessados = append(entradaPaciente.Interessados, listaDeInteressados{ObjectType: "interessado", Nome: entradaPaciente.Nome, CPF: entradaPaciente.CPF, DadosPermitidos: "11111111111111"})
	var Str []string
	entradaPaciente.Dados = Str
	if len(entradaPaciente.Nome) == 0 {
		return shim.Error("o campo nome deve ser uma cadeia nao vazia")
	}
	/* *********************************************** */
	/* Cria um objeto Paciente, faz marshal para JSON e armazena no banco de dados */
	paciente := &paciente{
		ObjectType:     Atuacao,
		Nome:           entradaPaciente.Nome,
		RG:             entradaPaciente.RG,
		CPF:            entradaPaciente.CPF,
		DataNascimento: entradaPaciente.DataNascimento,
		Sexo:           entradaPaciente.Sexo,
		NomeMae:        entradaPaciente.NomeMae,
		Naturalidade:   entradaPaciente.Naturalidade,
		Rua:            entradaPaciente.Rua,
		Numero:         entradaPaciente.Numero,
		Complemento:    entradaPaciente.Complemento,
		Bairro:         entradaPaciente.Bairro,
		Municipio:      entradaPaciente.Municipio,
		Estado:         entradaPaciente.Estado,
		CEP:            entradaPaciente.CEP,
		Interessados:   entradaPaciente.Interessados,
		Dados:          entradaPaciente.Dados,
	}
	pacienteJSONBytes, err := json.Marshal(paciente)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(entradaPaciente.CPF, pacienteJSONBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- fim de inicializacao de paciente")

	_, err = json.Marshal("Init data works!")
	if err == nil {
		return shim.Success(nil)
	}
	encoded := base64.StdEncoding.EncodeToString(pacienteJSONBytes)
	return shim.Success([]byte(encoded))
}

/*
	Operações de leitura dados do paciente
*/
//-----------------------------------------------------------------

func (t *Chaincode) getHistoryForPerson(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	cpfPessoa := args[0]

	fmt.Printf("- start getHistoryForPerson: %s\n", cpfPessoa)
	//first let's get the person current state value so we can see if the request are valid
	personAsBytes, erro := stub.GetState(cpfPessoa)
	if erro != nil {
		return shim.Error(erro.Error())
	} else if personAsBytes == nil {
		return shim.Error("person does not exist")
	}
	person := paciente{}
	erro = json.Unmarshal(personAsBytes, &person)
	if erro != nil {
		return shim.Error(erro.Error())
	}

	personInteressados := person.Interessados
	found := false
	for _, elem := range personInteressados {
		if elem.CPF == args[1] {
			found = true
		}
	}
	if found == false {
		return shim.Error(args[1] + " nao tem permissao de acesso")
	}

	resultsIterator, err := stub.GetHistoryForKey(cpfPessoa)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForMarble returning:\n%s\n", buffer.String())
	return shim.Success([]byte(buffer.Bytes()))
}
