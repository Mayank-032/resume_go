package usecase

import (
	"context"
	"fmt"
	"log"

	// "fmt"
	"resume-go/domain"
	"resume-go/domain/entity"

	"github.com/jung-kurt/gofpdf"
)

func (ruc ResumeUCase) processRawDataToPDF(ctx context.Context, resumeData entity.Resume) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// set page margin
	leftMargin, topMargin, rightMargin, _ := pdf.GetMargins()
	pdf.SetMargins(leftMargin, topMargin, rightMargin)

	// add heading
	pdf.CellFormat(0, 10, fmt.Sprintf("%v", resumeData.Name), "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 10, fmt.Sprintf("%v", resumeData.Title), "", 1, "C", false, 0, "")

	// add summary
	pdf.CellFormat(0, 10, "Summary", "", 1, "L", false, 0, "")
	pdf.MultiCell(0, 10, fmt.Sprintf("%v", resumeData.Summary), "", "L", false)

	// add education details
	pdf.CellFormat(0, 10, "Education", "", 1, "L", false, 0, "")
	addEducationalDetails(pdf, resumeData.EducationDetails)

	// add work experience
	pdf.CellFormat(0, 10, "Work Experience", "", 1, "L", false, 0, "")
	addWorkExperience(pdf, resumeData.WorkExperienceDetails)

	// add skills
	pdf.CellFormat(0, 10, "Skills", "", 1, "L", false, 0, "")
	addSkills(pdf, resumeData.Skills)

	// add Social Media Links
	pdf.MultiCell(0, 10, fmt.Sprintf("LinkedIn: %v", resumeData.LinkedInURL), "", "L", false)
	pdf.MultiCell(0, 10, fmt.Sprintf("GitHub: %v", resumeData.GitHubUrl), "", "L", false)

	// save pdf to google cloud bucket
	publicUrl, err := ruc.savePdfToCloud(ctx, pdf, resumeData.Name)
	if err != nil {
		log.Printf("Error %v\n, unable to save pdf to cloud\n\n", err)
	}
	return publicUrl, nil
}

func addSkills(pdf *gofpdf.Fpdf, skills entity.UserSkills) {
	pdf.MultiCell(0, 10, fmt.Sprintf("Programming languages: %v\nDatabase Knowledge: %v\nSoftware Tools: %v\nOther Skills: %v", skills.ProgrammingLanguages, skills.Database, skills.SoftwareTools, skills.Other), "", "L", false)
}

func addWorkExperience(pdf *gofpdf.Fpdf, workExperienceDetails []entity.UserWorkExperience) {
	for _, data := range workExperienceDetails {
		pdf.CellFormat(60, 10, fmt.Sprintf("%v", data.Designation), "B", 0, "L", false, 0, "")
		pdf.CellFormat(0, 10, fmt.Sprintf("%v", data.Organization), "", 1, "L", false, 0, "")
		pdf.CellFormat(60, 10, fmt.Sprintf("%v - %v", data.JobStartDate, data.JobEndingDate), "", 0, "L", false, 0, "")
		pdf.CellFormat(0, 10, fmt.Sprintf("%v", data.Responsibilities), "", 1, "L", false, 0, "")
	}
}

func addEducationalDetails(pdf *gofpdf.Fpdf, educationDetails []entity.UserEducationField) {
	for _, data := range educationDetails {
		pdf.CellFormat(60, 10, fmt.Sprintf("%v", data.Degree), "B", 0, "L", false, 0, "")
		pdf.CellFormat(0, 10, fmt.Sprintf("%v", data.Major), "", 1, "L", false, 0, "")
		pdf.CellFormat(60, 10, fmt.Sprintf("%v, %v", data.College, data.City), "", 0, "L", false, 0, "")
		pdf.CellFormat(0, 10, fmt.Sprintf("%v - %v", data.EducationStartDate, data.EducationEndingDate), "", 1, "L", false, 0, "")
	}
}

func (ruc ResumeUCase) savePdfToCloud(ctx context.Context, pdf *gofpdf.Fpdf, pdfName string) (string, error) {
	client := ruc.Client
	objName := fmt.Sprintf("%v.pdf", pdfName)
	bucketName := "resume_store2002"

	bucket := client.Bucket(bucketName)
	wc := bucket.Object(objName).NewWriter(ctx)
	if err := pdf.Output(wc); err != nil {
		log.Printf("Error %v", err)
		return domain.EmptyString, err
	}
	if err := wc.Close(); err != nil {
		log.Printf("Error %v", err)
		return domain.EmptyString, err
	}
	return fmt.Sprintf("https://storage.googleapis.com/%v/%v", bucketName, objName), nil
}
