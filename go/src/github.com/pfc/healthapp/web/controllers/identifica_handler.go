package controllers

import(
	"net/http"
)

//identificaPessoaHandler é responsavel por obter os dados da pagina html e realizar a chamada da função identificaPessoa na blockchain
func (app *Application) IdentificaPessoaHandler(w http.ResponseWriter, r *http.Request){
	//dados de resposta da transacao
	data := &struct{
		TransactionId	string
		Success			bool
		Response		bool
	}{
		TransactionId:	"",
		Success:		false,
		Response:		false,
	}

	//verifica se dados foram enviados
	if r.FormValue("submitted") == "true"{
		//obtem dado do respectivo campo do form
		cpf := r.FormValue("cpf_paciente")
		
		if len(cpf) > 0{	//caso seja informado um valor de cpf
			//realiza chamada da função
			txid, err := app.Fabric.IdentificaPessoa(cpf)
			
			//verifição da transação
			if err != nil{
				if txid == ""{	//resposta caso paciente não for encontrado
					data.TransactionId = "Paciente não encontrado"
					data.Success = false
					data.Response = true
				}else{	//resposta caso não for possivel invocar o chaincode
					http.Error(w, "Não foi possivel invocar a função IdentificaPessoa na blockchain", 500)
				}
			}else{	//resposta caso paciente for encontrado
				data.TransactionId = txid
				data.Success = true
				data.Response = true
			}
		}else{	//caso nenhum dado seja informado
			data.Response = true
			data.TransactionId = "CPF não informado"
		}
	}
	
	//renderiza o template html
	renderTemplate(w, r, "identificaPessoa.html", data)
}