package repository

import (
	"resume-go/domain/interfaces/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResumeRepository struct {
	DB *mongo.Collection
}

func NewResumeRepository(db *mongo.Collection) repository.IResumeRepo {
	return ResumeRepository {
		DB: db,
	}
}

func (rp ResumeRepository) Upsert() {}

func (rp ResumeRepository) Remove() {}

func (rp ResumeRepository) Fetch() {}