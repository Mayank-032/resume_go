package usecase

type IResumeUCase interface {
	UpsertPDF()
	RemovePDF()
	FetchPDF()
}