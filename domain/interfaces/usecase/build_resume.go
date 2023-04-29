package usecase

import (
	"context"
	"resume-go/domain/entity"
)

type IResumeUCase interface {
	UpsertPDF(ctx context.Context, resumeData entity.Resume) (string, error)
	RemovePDF()
	FetchPDF()
}