package controllers

import (
	"net/http"
)

//QueryConcessaoHandler é responsavel por obter os dados da pagina html e realizar a chamada da função queryConcessao na blockchain
func (app *Application) QueryConcessaoHandler(w http.ResponseWriter, r *http.Request){
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
		nome_interessado := r.FormValue("nome_interessado")
		dados = append(dados, nome_interessado)
		cpf_interessado := r.FormValue("cpf_interessado")
		dados = append(dados, cpf_interessado)
		dados_permitidos := r.FormValue("dados")
		dados = append(dados, dados_permitidos)

		if ( len(r.FormValue("cpf_paciente")) > 0 && len(r.FormValue("cpf_interessado")) > 0 && len(r.FormValue("nome_interessado")) > 0 && len(r.FormValue("dados")) > 0){
			//realiza chamada da função
			txid, err := app.Fabric.QueryConcessao(dados)
			
			//verificação da transação
			if err != nil{
				if txid == ""{ //caso paciente não esteja presente na blockchain
					data.TransactionId = "Paciente não encontrado"
					data.Success = false
					data.Response = true
				}else{	//resposta caso não for possivel invocar o chaincode
					http.Error(w, "Não foi possivel invocar a função QueryConcessao da blockchain", 500)
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
	renderTemplate(w, r, "queryConcessao.html", data)
}