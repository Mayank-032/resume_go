package usecase

import (
	"resume-go/domain/interfaces/usecase"
	"resume-go/domain/interfaces/repository"
)

type ResumeUCase struct {
	ResumeRepo repository.IResumeRepo
}

func NewResumeUseCase(resumeRepo repository.IResumeRepo) usecase.IResumeUCase {
	return ResumeUCase {
		ResumeRepo: resumeRepo,
	}
}

func (ruc ResumeUCase) UpsertPDF() {}

func (ruc ResumeUCase) RemovePDF() {}

func (ruc ResumeUCase) FetchPDF() {}