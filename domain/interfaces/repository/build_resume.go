package repository

import "context"

type IResumeRepo interface {
	Upsert(ctx context.Context, pdfFile string) (string, error)
	Remove()
	Fetch()
}