package gopddikti

import (
	"net/http"
	"strings"
)

const (
	getProgrammeDetailPath = "/detail_prodi"
)

type ProgrammeDetail struct {
	GeneralInfo    ProgrammeGeneralData         `json:"detailumum"`
	Ratio          []ProgrammeRatioData         `json:"rasio"`
	Students       []ProgrammeStudentData       `json:"datamhs"`
	Lecturers      []ProgrammeLecturerData      `json:"datadosen"`
	LecturerRatios []ProgrammeLecturerRatioData `json:"datadosenrasio"`
}

type ProgrammeGeneralData struct {
	Accreditation        string  `json:"akreditas"`
	Achievements         string  `json:"capaian"`
	Address              string  `json:"jln"`
	Competencies         string  `json:"kompetensi"`
	Programme            string  `json:"nm_lemb"`
	Description          string  `json:"deskripsi"`
	Email                string  `json:"email"`
	EstablishmentDate    string  `json:"tgl_berdiri"`
	FaxNumber            string  `json:"no_fax"`
	DetailID             string  `json:"id_sms"`
	CollegeDetailID      string  `json:"linkpt"`
	Latitude             float64 `json:"lintang"`
	Longitude            float64 `json:"bujur"`
	Mission              string  `json:"misi"`
	College              string  `json:"namapt"`
	NPSN                 string  `json:"npsn"`
	OperatingLicense     string  `json:"sk_selenggara"`
	OperatingLicenseDate string  `json:"tgl_sk_selenggara"`
	PostalCode           string  `json:"kode_pos"`
	DegreeLevel          string  `json:"namajenjang"`
	ProgrammeCode        string  `json:"kode_prodi"`
	ProgrammeStatus      string  `json:"stat_prodi"`
	TelephoneNumber      string  `json:"no_tel"`
	Vision               string  `json:"visi"`
	Website              string  `json:"website"`
}

type ProgrammeRatioData struct {
	LecturerCount int    `json:"jmldosen"`
	Semester      string `json:"smt"`
	ProgrammeCode string `json:"kode_program_studi"`
	StudentCount  int    `json:"jmlmhs"`
}

type ProgrammeStudentData struct {
	StartSemester     string `json:"mulai_smt"`
	TotalStudentCount int    `json:"jml"`
}

type ProgrammeLecturerData struct {
	DegreeLevel    string `json:"pendidikan"`
	DetailID       string `json:"id"`
	Name           string `json:"nama"`
	RegistrationID string `json:"idreg"`
	Degree         string `json:"gelar"`
}

type ProgrammeLecturerRatioData struct {
	Degree         string `json:"gelar_dosen"`
	RegistrationID string `json:"idreg"`
	DegreeLevel    string `json:"jenjang_dosen"`
	DetailID       string `json:"id"`
	Name           string `json:"nama"`
	NIDN           string `json:"nidn"`
	Programme      string `json:"prodi_homebase"`
	College        string `json:"pt"`
}

// GetProgrammeDetailByDetailID is used to get the study programme information in detail.
// It receives detailID as parameter.
func (c *Client) GetProgrammeDetailByDetailID(detailID string) (res ProgrammeDetail, err error) {
	return c.getProgrammeDetail(detailID)
}

func (c *Client) getProgrammeDetail(detailID string) (res ProgrammeDetail, err error) {
	if err = checkParamString(detailID); err != nil {
		return
	}

	req, err := c.createRequest(http.MethodGet, getProgrammeDetailPath+"/"+detailID)
	if err != nil {
		return
	}

	err = c.doRequest(req, &res)
	if err != nil {
		return
	}

	res.splitProgrammeDetailID()

	return
}

func (r *ProgrammeDetail) splitProgrammeDetailID() {
	collegeDetailID := strings.Split(r.GeneralInfo.CollegeDetailID, "/")
	if len(collegeDetailID) == 3 {
		r.GeneralInfo.CollegeDetailID = collegeDetailID[2]
	}
}
