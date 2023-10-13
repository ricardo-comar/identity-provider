package entity

type EmployeeRecordEntity struct {

	// Identificação do Funcionário - ID
	ID string `dynamodbav:"id"`

	TTL string `dynamodbav:"ttl"`

	// Dados cadastrais
	CadastroFuncionario EmployeeData `dynamodbav:"cadastro_funcionario"`

	DataHoraCriacao string `dynamodbav:"data_hora_criacao"`

	DataHoraAlteracao string `dynamodbav:"data_hora_alteracao"`
}

type EmployeeData struct {
	ID             string `json:"id,omitempty"`
	NumFunc        string `json:"codigo_cadastro,omitempty"`
	Nome           string `json:"nome_funcionario,omitempty"`
	Sobrenome      string `json:"sobrenome_funcionario,omitempty"`
	Email          string `json:"email_funcionario,omitempty"`
	Sexo           string `json:"sexo_funcionario,omitempty"`
	Departamento   string `json:"departamento,omitempty"`
	Cargo          string `json:"cargo,omitempty"`
	DataNascimento string `json:"data_nascimento,omitempty"`
	DataAdmissao   string `json:"data_admissao,omitempty"`
	Status         string `json:"status,omitempty"`
}
