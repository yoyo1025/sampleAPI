package validator

import "sampleAPI/model"

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

// taskValidator構造体を作成するためのコンストラクタ
func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

