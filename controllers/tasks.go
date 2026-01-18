package controllers

import (
	"net/http"
	"strconv"
	"teste_tecnico/models"
	"text/template"
)

// traz todos os arquivos html da pasta templates
var temp = template.Must(template.ParseGlob("templates/*.html"))

// é a funcao que carrega a pagina inicial do sistema e traz todas as tasks
func Index(w http.ResponseWriter, r *http.Request) {
	todasAsTasks := models.BuscarTodasTasks()
	temp.ExecuteTemplate(w, "index.html", todasAsTasks)
}

// carrega a pagina para criar uma nova task
func NovaTask(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "novaTask.html", nil)
}

// insere uma nova task no banco de dados
func InserirTask(w http.ResponseWriter, r *http.Request) {
	// verifica se o metodo da requisicao é POST
	if r.Method == "POST" {
		titulo := r.FormValue("titulo")
		descricao := r.FormValue("descricao")
		status := r.FormValue("status")

		models.CriarNovaTask(titulo, descricao, status)
	}
	// redireciona para a pagina inicial
	http.Redirect(w, r, "/", 301)
}

// deleta uma task do banco de dados
func DeletarTask(w http.ResponseWriter, r *http.Request) {
	// pega o id da task que vem pela url
	idDaTask := r.URL.Query().Get("id")
	models.DeletarTask(idDaTask)
	http.Redirect(w, r, "/", 301)
}

// carrega a pagina de edicao de uma task
func EditarTask(w http.ResponseWriter, r *http.Request) {
	idDaTask := r.URL.Query().Get("id")
	task := models.EditarTask(idDaTask)
	temp.ExecuteTemplate(w, "editarTask.html", task)
}

// atualiza uma task no banco de dados
func AtualizarTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		titulo := r.FormValue("titulo")
		descricao := r.FormValue("descricao")
		status := r.FormValue("status")
		// Convertendo o id de string para int
		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		models.AtualizarTask(idConvertido, titulo, descricao, status)
	}
	http.Redirect(w, r, "/", 301)
}

// busca uma task pelo seu titulo
func BuscarTaskPorTitulo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//titulo := r.URL.Query().Get("titulo")
		titulo := r.FormValue("titulo")

		temp.ExecuteTemplate(w, "index.html", models.BuscarTaskPorTitulo(titulo))

	}
	http.Redirect(w, r, "/", 301)

}
