package repository

type IResumeRepo interface {
	Upsert()
	Remove()
	Fetch()
}