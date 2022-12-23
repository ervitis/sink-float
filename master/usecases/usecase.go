package usecases

import "github.com/ervitis/sink-float/master/repository"

type sinkUseCase struct {
	repository.Memcache
}

type SinkUseCase interface {
}

func NewSinkUseCase(repo repository.Memcache) SinkUseCase {
	return &sinkUseCase{repo}
}
