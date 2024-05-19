package gopddikti

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	getStudentDetailPath = "/detail_mhs"
)

type StudentDetail struct {
	GeneralInfo   StudentGeneralData  `json:"dataumum"`
	StudyHistory  []StudentStudyData  `json:"datastudi"`
	StudyStatuses []StudentStatusData `json:"datastatuskuliah"`
}

type StudentGeneralData struct {
	Gender             string    `json:"jk"`
	Status             string    `json:"ket_keluar"`
	ProgrammeDetailID  string    `json:"link_prodi"`
	CollegeDetailID    string    `json:"link_pt"`
	StartYear          string    `json:"mulai_smt"`
	DegreeLevel        string    `json:"namajenjang"`
	Programme          string    `json:"namaprodi"`
	College            string    `json:"namapt"`
	NIM                string    `json:"nipd"`
	RegistrationStatus string    `json:"nm_jns_daftar"`
	Name               string    `json:"nm_pd"`
	InitialProgramme   string    `json:"nm_prodi_asal"`
	InitialCollege     string    `json:"nm_pt_asal"`
	CertificateNumber  string    `json:"no_seri_ijazah"`
	RegistrationNumber string    `json:"reg_pd"`
	GraduationDate     time.Time `json:"tgl_keluar"`
}

type StudentStudyData struct {
	Semester       string `json:"id_smt"`
	SubjectCode    string `json:"kode_mk"`
	SubjectGrade   string `json:"nilai_huruf"`
	SubjectName    string `json:"nm_mk"`
	SubjectCredits int    `json:"sks_mk"`
}

type StudentStatusData struct {
	Semester        string `json:"id_smt"`
	SemesterStatus  string `json:"nm_stat_mhs"`
	SemesterCredits int    `json:"sks_smt"`
}

// GetStudentDetailByNIM is used to get the student information in detail.
// It receives NIM in the type of string as parameter.
func (c *Client) GetStudentDetailByNIM(nim string) (res StudentDetail, err error) {
	raw, err := c.searchStudents(nim)
	if err != nil {
		return
	}

	if len(raw.Students) < 1 {
		err = fmt.Errorf("empty result for id: %s", nim)
		return
	}

	match := strings.Split(raw.Students[0].WebsiteLink, "/")
	if len(match) < 3 || strings.TrimSpace(match[2]) == "" {
		err = fmt.Errorf("invalid result for id: %s", nim)
		return
	}

	res, err = c.getStudentDetail(match[2])
	if err != nil {
		return
	}

	return
}

// GetStudentDetailByDetailID is used to get the student information in detail.
// It receives detailID as parameter.
func (c *Client) GetStudentDetailByDetailID(detailID string) (res StudentDetail, err error) {
	return c.getStudentDetail(detailID)
}

func (c *Client) getStudentDetail(detailID string) (res StudentDetail, err error) {
	if err = checkParamString(detailID); err != nil {
		return
	}

	req, err := c.createRequest(http.MethodGet, getStudentDetailPath+"/"+detailID)
	if err != nil {
		return
	}

	err = c.doRequest(req, &res)
	if err != nil {
		return
	}

	res.splitStudentDetailID()

	return
}

func (r *StudentDetail) splitStudentDetailID() {
	collegeDetailID := strings.Split(r.GeneralInfo.CollegeDetailID, "/")
	if len(collegeDetailID) == 3 {
		r.GeneralInfo.CollegeDetailID = collegeDetailID[2]
	}

	programmeDetailID := strings.Split(r.GeneralInfo.ProgrammeDetailID, "/")
	if len(programmeDetailID) == 3 {
		r.GeneralInfo.ProgrammeDetailID = programmeDetailID[2]
	}
}
