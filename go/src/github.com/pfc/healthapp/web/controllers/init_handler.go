package controllers

import (
	"net/http"
)

//InitHandler é responsavel por obter os dados da pagina html e realizar a chamada da função initPaciente na blockchain
func (app *Application) InitHandler(w http.ResponseWriter, r *http.Request) {
	//dados de resposta da transacao
	data := &struct {
		TransactionId string
		Success       bool
		Response      bool
	}{
		TransactionId: "",
		Success:       false,
		Response:      false,
	}
	
	//array para armazenar entradas para a chamada da função
	var dados []string

	//verifica se dados foram enviados
	if r.FormValue("submitted") == "true"{
		//obtem dados dos respectivos campos do form e armazena no array "dados"
		atuacao := r.FormValue("atuacao")
		dados = append(dados, atuacao)
		nome := r.FormValue("nome")
		dados = append(dados, nome)
		rg := r.FormValue("rg")
		dados = append(dados, rg)
		cpf := r.FormValue("cpf")
		dados = append(dados, cpf)
		data_nasc := r.FormValue("data_nasc")
		dados = append(dados, data_nasc)
		sexo := r.FormValue("sexo")
		dados = append(dados, sexo)
		nome_mae := r.FormValue("mae")
		dados = append(dados, nome_mae)
		naturalidade := r.FormValue("naturalidade")
		dados = append(dados, naturalidade)
		rua := r.FormValue("rua")
		dados = append(dados, rua)
		num := r.FormValue("num")
		dados = append(dados, num)
		comp := r.FormValue("comp")
		dados = append(dados, comp)
		bairro := r.FormValue("bairro")
		dados = append(dados, bairro)
		municipio := r.FormValue("municipio")
		dados = append(dados, municipio)
		estado := r.FormValue("estado")
		dados = append(dados, estado)
		cep := r.FormValue("cep")
		dados = append(dados, cep)

		//verifica se alguma entrada não foi informada
		var input_vazio bool
		for _, input := range dados{
			if len(input) == 0{
				input_vazio = true
			}
		}

		if input_vazio == false{	//caso todos os inputs forem informados
			//realiza chamada da função
			txid, err := app.Fabric.InitPaciente(dados)
			
			//verificação da transação
			if err != nil {	
				if txid == ""{
					data.TransactionId = "Erro ao inserir paciente"
					data.Success = false
					data.Response = true
				}else{	//resposta caso não for possivel invocar o chaincode
					http.Error(w, "Não foi possivel invocar a função initPaciente na blockchain", 500)
				}
			}else{	//resposta caso transação for concluida
				data.TransactionId = txid
				data.Success = true
				data.Response = true
			}
		}else{	//caso algum dado não seja informado
			data.Response = true
			data.TransactionId = "Faltando dados a serem inseridos"
		}
	}
		
	//renderiza o template html
	renderTemplate(w, r, "initpaciente.html", data)
}