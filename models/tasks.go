package models

import (
	"teste_tecnico/db"
)

type Task struct {
	Id        int
	Titulo    string
	Descricao string
	Status    string
	CridoEm   string
}

// traz todas as tasks do banco
func BuscarTodasTasks() []Task {
	// conecta com o banco de dados
	db := db.ConectaComBancoDeDados()
	// traz todas as tasks do banco de dados atraves da query sql ordenando pelo status
	selectDeTodasTasks, err := db.Query("select * from tasks order by status asc")
	if err != nil {
		panic(err)
	}

	t := Task{}
	tasks := []Task{}
	// passa por todas as linhas retornadas pela query
	for selectDeTodasTasks.Next() {
		var id int
		var titulo, descricao, status, criadoEm string
		err = selectDeTodasTasks.Scan(&id, &titulo, &descricao, &status, &criadoEm)
		if err != nil {
			panic(err)
		}
		t.Id = id
		t.Titulo = titulo
		t.Descricao = descricao
		t.Status = status
		t.CridoEm = criadoEm

		tasks = append(tasks, t)
	}
	// fecha conexao com o banco
	defer db.Close()
	return tasks
}

// busca uma task pelo seu titulo
func BuscarTaskPorTitulo(titulo string) []Task {
	db := db.ConectaComBancoDeDados()
	taskDoBanco, err := db.Query("select * from tasks where titulo=$1", titulo)
	if err != nil {
		panic(err.Error())
	}

	taskParaRetornar := Task{}
	tasks := []Task{}

	for taskDoBanco.Next() {
		var id int
		var titulo, descricao, status, criadoEm string
		err = taskDoBanco.Scan(&id, &titulo, &descricao, &status, &criadoEm)
		if err != nil {
			panic(err.Error())
		}
		taskParaRetornar.Id = id
		taskParaRetornar.Titulo = titulo
		taskParaRetornar.Descricao = descricao
		taskParaRetornar.Status = status
		taskParaRetornar.CridoEm = criadoEm

		tasks = append(tasks, taskParaRetornar)

	}
	defer db.Close()
	return tasks
}

// cria uma nova task no banco
func CriarNovaTask(titulo, descricao, status string) {
	db := db.ConectaComBancoDeDados()
	// insere os dados no banco de dados atraves da query sql
	insereDadosNoBanco, err := db.Prepare("insert into tasks(titulo, descricao, status) values($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(titulo, descricao, status)
	defer db.Close()
}

// deleta uma task do banco pelo id conseguido no controller pela url
func DeletarTask(id string) {
	db := db.ConectaComBancoDeDados()
	deletarATask, err := db.Prepare("delete from tasks where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarATask.Exec(id)
	defer db.Close()
}

// traz uma task para edicao
func EditarTask(id string) Task {
	db := db.ConectaComBancoDeDados()
	// da select em uma task pelo id
	taskDoBanco, err := db.Query("select * from tasks where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	taskParaAtualizar := Task{}

	for taskDoBanco.Next() {
		var id int
		var titulo, descricao, status, criadoEm string
		err = taskDoBanco.Scan(&id, &titulo, &descricao, &status, &criadoEm)
		if err != nil {
			panic(err.Error())
		}
		taskParaAtualizar.Id = id
		taskParaAtualizar.Titulo = titulo
		taskParaAtualizar.Descricao = descricao
		taskParaAtualizar.Status = status
		taskParaAtualizar.CridoEm = criadoEm
	}
	defer db.Close()
	return taskParaAtualizar
}

// atualiza uma task no banco de dados
func AtualizarTask(id int, titulo, descricao, status string) {
	db := db.ConectaComBancoDeDados()
	// query sql para atualizar a task no banco
	atualizaTask, err := db.Prepare("update tasks set titulo=$1, descricao=$2, status=$3 where id=$4")
	if err != nil {
		panic(err.Error())
	}
	atualizaTask.Exec(titulo, descricao, status, id)
	defer db.Close()
}
