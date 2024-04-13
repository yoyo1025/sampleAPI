package validator

import "sampleAPI/model"

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}