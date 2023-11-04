package web

import(
	"fmt"
	"github.com/pfc/healthapp/web/controllers"
	"net/http"
)

//Serve é responsavel por iniciar o servidor http
func Serve(app *controllers.Application){

	//gerando handlers para cada pagina html criada
	http.HandleFunc("/initpaciente.html", app.InitHandler)
	http.HandleFunc("/identificaPessoa.html", app.IdentificaPessoaHandler)
	http.HandleFunc("/queryConcessao.html", app.QueryConcessaoHandler)
	http.HandleFunc("/queryRevogacao.html", app.QueryRevogacaoHandler)
	http.HandleFunc("/getRegistro.html", app.GetRegistroHandler)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/initpaciente.html", http.StatusTemporaryRedirect)
	})

	//executando servidor
	fmt.Println("Listening (http://localhost:3000/) ...")
	http.ListenAndServe(":3000", nil)
}