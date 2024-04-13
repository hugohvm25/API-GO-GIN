/*
Inclusão do import de validação dos dados para impedir de ser criado um cadastro em branco
*/

package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"` // não pode ser 0
	CPF  string `json:"cpf" validate:"len=9"`    // deve ter um tamanho especifico
	RG   string `json:"rg" validate:"len=11"`    // deve ter um tamanho especifico
}

func ValidaDadosAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
