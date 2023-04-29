package responses

type FetchResume struct {
	BaseResponse
	PDFFile string `json:"pdffile"`
}

func (fr FetchResume) Success(pdfFile, msg string, status bool) FetchResume {
	fr.Message = msg
	fr.Status  = status
	fr.PDFFile = pdfFile
	return fr
}

func (fr FetchResume) Failure(msg string, status bool) FetchResume {
	fr.Message = msg
	fr.Status = status
	return fr
}