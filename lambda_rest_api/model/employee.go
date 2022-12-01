package model

const (
	STATUS_PRE_ADMITIDO = "PRE_ADMITIDO"
	STATUS_ATIVO        = "ATIVO"
	STATUS_EM_FERIAS    = "EM_FERIAS"
	STATUS_AFASTADO     = "AFASTADO"
	STATUS_BLOQUEADO    = "BLOQUEADO"
	STATUS_RESCISAO     = "RESCISAO"
	STATUS_INATIVO      = "INATIVO"
)

type EmployeeModel struct {
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
