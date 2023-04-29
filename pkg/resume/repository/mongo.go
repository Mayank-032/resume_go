package repository

import (
	"context"
	"errors"
	"resume-go/domain"
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

func (rp ResumeRepository) Upsert(ctx context.Context, pdfFile string) (string, error) {
	result, err := rp.DB.InsertOne(ctx, pdfFile)
	if err != nil {
		return domain.EmptyString, errors.New("unable to make connection to db")
	}
	
	last_inserted_id := result.InsertedID
	if err != nil {
		return domain.EmptyString, errors.New("unable to fetch last inserted row id")
	}
	return last_inserted_id.(string), nil
}

func (rp ResumeRepository) Remove() {}

func (rp ResumeRepository) Fetch() {}