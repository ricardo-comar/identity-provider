package model

import (
	"encoding/xml"
)

//EmployeeRegistries slice
type EmployeeRegistries struct {
	XMLName    xml.Name   `xml:"employees,omitempty" json:"-"`
	Registries []Registry `xml:"employee,omitempty" json:"funcionario,omitempty"`
}

//Registry data
type Registry struct {
	XMLName xml.Name `xml:"employee,omitempty" json:"-"`

	ID string `xml:"id,omitempty" json:"id,omitempty"`

	NumFunc string `xml:"ein,omitempty" json:"codigo_cadastro,omitempty"`

	Nome string `xml:"first_name,omitempty" json:"nome_funcionario,omitempty"`

	Sobrenome string `xml:"last_name,omitempty" json:"sobrenome_funcionario,omitempty"`

	Email string `xml:"email,omitempty" json:"email_funcionario,omitempty"`

	Sexo string `xml:"gender,omitempty" json:"sexo_funcionario,omitempty"`

	Departamento string `xml:"department,omitempty" json:"departamento,omitempty"`

	Cargo string `xml:"job-title,omitempty" json:"cargo,omitempty"`

	DataNascimento string `xml:"birthdate,omitempty" json:"data_nascimento,omitempty"`

	DataAdmissao string `xml:"admissiondate,omitempty" json:"data_admissao,omitempty"`

	Status string `xml:"status,omitempty" json:"status,omitempty"`
}
