package usecase

import (
	"context"
	"errors"
	"resume-go/domain"
	"resume-go/domain/entity"
	"resume-go/domain/interfaces/repository"
	"resume-go/domain/interfaces/usecase"

	"cloud.google.com/go/storage"
)

type ResumeUCase struct {
	Client     *storage.Client
	ResumeRepo repository.IResumeRepo
}

func NewResumeUseCase(resumeRepo repository.IResumeRepo, client *storage.Client) usecase.IResumeUCase {
	return ResumeUCase{
		Client: client,
		ResumeRepo: resumeRepo,
	}
}

func (ruc ResumeUCase) UpsertPDF(ctx context.Context, resumeData entity.Resume) (string, error) {
	pdfFile, err := ruc.processRawDataToPDF(ctx, resumeData)
	if err != nil {
		return domain.EmptyString, errors.New("unable to process data to pdf")
	}
	userId, err := ruc.ResumeRepo.Upsert(ctx, pdfFile)
	if err != nil {
		return domain.EmptyString, errors.New("unable to upsert data")
	}
	return userId, nil
}

func (ruc ResumeUCase) RemovePDF() {}

func (ruc ResumeUCase) FetchPDF() {}
