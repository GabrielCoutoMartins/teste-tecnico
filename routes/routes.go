package routes

import "net/http"
import "teste_tecnico/controllers"


func CarregarRotas() {
	// Aqui voce define suas rotas e handlers
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/novaTask", controllers.NovaTask)
	http.HandleFunc("/insert", controllers.InserirTask)
	http.HandleFunc("/deletar", controllers.DeletarTask)
	http.HandleFunc("/editar", controllers.EditarTask)
	http.HandleFunc("/update", controllers.AtualizarTask)
	http.HandleFunc("/buscarPorTitulo", controllers.BuscarTaskPorTitulo)
	
}
