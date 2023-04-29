package routes

import (
	"errors"
	"resume-go/domain/dto/requests"
	"resume-go/domain/entity"
)

func Validate(req requests.UpsertResume) error {
	if len(req.Name) == 0 || len(req.Email) == 0 {
		return errors.New("invalid_details")
	}
	return nil
}

func ToResumeDto(req requests.UpsertResume) entity.Resume {
	return req.Resume
}