package usecases

import "github.com/ervitis/sink-float/master/repository"

type masterSink struct {
	repository.Memcache
}

type MasterSink interface {
}

func NewMasterSinkUseCase(repo repository.Memcache) MasterSink {
	return &masterSink{repo}
}
