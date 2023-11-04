package controllers

import(
	"fmt"
	"github.com/pfc/healthapp/blockchain"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Application struct{
	Fabric *blockchain.FabricSetup
}

//renderTemplate é responsavel por renderizar o template associado pelo parametro "templateName"
func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}){
	
	//cria o caminho para o template de layout
	lp := filepath.Join("web", "templates", "layout.html")
	//cria o caminho para o template
	tp := filepath.Join("web", "templates", templateName)

	//retorna 404 se o template não existir
	info, err := os.Stat(tp)
	if err != nil{
		if os.IsNotExist(err){
			http.NotFound(w, r)
			return
		}
	}

	//retorna 404 se a requisição for para um diretório
	if info.IsDir(){
		http.NotFound(w, r)
		return
	}

	//analisa o caminho dos templates e retorna um tipo para os templates associados
	resultTemplate, err := template.ParseFiles(tp, lp)
	if err != nil{
		//informa o determinado erro
		fmt.Println(err.Error())

		//retorna um erro do servidor
		http.Error(w, http.StatusText(500), 500)
		return
	}

	//executa o template
	if err := resultTemplate.ExecuteTemplate(w, "layout", data); err != nil{
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}