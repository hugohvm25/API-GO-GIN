/*
Inclusão do import de validação dos dados para impedir de ser criado um cadastro em branco

== sem restrição de dados ==
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`

== com limitação de não poder ser 0 no nome e tamanho especifico no RG e CPF ==
	Nome string `json:"nome" validate:"nonzero"` // não pode ser 0
	RG   string `json:"rg" validate:"len=9"`     // deve ter um tamanho especifico
	CPF  string `json:"cpf" validate:"len=11"`   // deve ter um tamanho especifico


*/

package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model

	Nome string `json:"nome" validate:"nonzero, regexp=^[a-zA-Z]*$` // não pode ser 0 e deve ter somente letras
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`       // deve ter um tamanho especifico e restrição somente para números
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$`      // deve ter um tamanho especifico e restrição somente para números

}

func ValidaDadosAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
