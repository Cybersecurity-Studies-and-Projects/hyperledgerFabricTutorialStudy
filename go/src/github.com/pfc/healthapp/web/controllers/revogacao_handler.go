package controllers

import(
	"net/http"
)

//QueryRevogacaoHandler é responsavel por obter os dados da pagina html e realizar a chamada da função queryRevogacao na blockchain
func (app *Application) QueryRevogacaoHandler(w http.ResponseWriter, r *http.Request){
	//dados de resposta da transacao
	data := &struct {
		TransactionId	string
		Success			bool
		Response		bool
	}{
		TransactionId:	"",
		Success:		false,
		Response:		false,
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

		if ( len(r.FormValue("cpf_paciente")) > 0 && len(r.FormValue("cpf_interessado")) > 0){
			//realiza chamada da função
			txid, err := app.Fabric.QueryRevogacao(dados)
			
			//verificação da transação
			if err != nil{
				if txid == ""{ //caso cpf do paciente ou interessado não estejam na blockchain
					data.TransactionId = "Dados não estão presentes nos registros"
					data.Success = false
					data.Response = true
				}else{	//resposta caso não for possivel invocar o chaincode
					http.Error(w, "Não foi possivel invocar a função QueryRevogacao na blockchain", 500)
				}
			}else{	//resposta caso transação for concluida
				data.TransactionId = txid
				data.Success = true
				data.Response = true
			}
		}else{	//caso nenhum dado seja informado
			data.Response = true
			data.TransactionId = "Faltando dados a serem informados"
		}
	}

	//renderiza o template html
	renderTemplate(w, r, "queryRevogacao.html", data)
}