package controllers

import(
	"net/http"
	"encoding/json"
)

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

//GetRegistroHandler é responsavel por obter os dados da pagina html e realizar a chamada da função getRegistro na blockchain
func (app *Application) GetRegistroHandler(w http.ResponseWriter, r *http.Request){
	//dados de resposta da transacao
	data := &struct{
		Success		bool
		Response	bool
		Alert		string
		//CAMPOS RELACIONADOS AO REGISTRO
		Nome           string                
		RG             string                
		CPF            string                
		DataNascimento string                
		Sexo           string                
		NomeMae        string                
		Naturalidade   string                
		Rua            string                
		Numero         int                   
		Complemento    string                
		Bairro         string                
		Municipio      string                
		Estado         string                
		CEP            string      
		Interessados   []listaDeInteressados          
	}{
		Success:	false,
		Response:	false,
		Alert:		"",
		Nome:		"",
		RG:			"",
		CPF:		"",
		DataNascimento:	"",
		Sexo:			"",
		NomeMae:		"",
		Naturalidade:	"",
		Rua:			"",
		Numero:			0,
		Complemento:	"",
		Bairro:			"",
		Municipio:		"",
		Estado:			"",
		CEP:			"",
		Interessados:	nil,
	}

	//array para armazenar entradas para a chamada da função
	var dados []string

	//verifica se dados foram enviados
	if r.FormValue("submitted") == "true"{
		
		//obtem dados dos respectivos campos do form e armazena no array "dados"
		cpf_paciente := r.FormValue("cpf_paciente")
		dados = append(dados, cpf_paciente)
		cpf_interessado := r.FormValue("cpf_interessado")
		dados = append(dados, cpf_interessado)

		//declara variavel que irá receber registro serializado no processo de unmarshal
		registro := &paciente{}

		if len(cpf_paciente) > 0 && len(cpf_interessado) > 0{	//caso seja informado o dado
			//realiza chamada da função
			resposta, err := app.Fabric.GetRegistro(dados)
			
			//verificação da função
			if err != nil{
				if resposta == nil{	//resposta caso paciente não for encontrado
					data.Response = true
					data.Success = false
					data.Alert = "Erro ao obter registro!\nDados do paciente/interessado incompatíveis"
				}else{	//resposta caso não for possivel invocar o chaincode
					http.Error(w, "Não foi possivel invocar a função GetRegistro na blockchain", 500)
				}
			}else{	//resposta operação seja bem sucedida
				//realiza unmarshal dos dados recebidos
				if err = json.Unmarshal(resposta, &registro); err != nil{
					data.Response = true
					data.Success = false
					data.Alert = "Erro ao realizar unmarshal dos dados obtidos"
				}

				//formata resposta
				data.Response = true
				data.Success = true
				data.Alert = "Transação concluída!"
				
				//formata dados a serem exibidos
				data.Nome = registro.Nome
				data.RG = registro.RG
				data.CPF = registro.CPF
				data.DataNascimento = registro.DataNascimento
				data.Sexo = registro.Sexo
				data.NomeMae = registro.NomeMae
				data.Naturalidade = registro.Naturalidade
				data.Rua = registro.Rua
				data.Numero = registro.Numero
				data.Complemento = registro.Complemento
				data.Bairro = registro.Bairro
				data.Municipio = registro.Municipio
				data.Estado = registro.Estado
				data.CEP = registro.CEP
				data.Interessados = registro.Interessados
			}
		}else{	//caso nenhum dado seja informado
			if len(cpf_paciente) == 0{
				data.Response = true
				data.Alert = "CPF do paciente não foi informado!"
			}else if len(cpf_interessado) == 0{
				data.Response = true
				data.Alert = "CPF do interessado não foi informado!"
			}
			
		}
	}

	//renderiza o template html
	renderTemplate(w, r, "getRegistro.html", data)
}