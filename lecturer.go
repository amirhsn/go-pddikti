package gopddikti

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	getLecturerDetailPath = "/detail_dosen"
)

type LecturerDetail struct {
	GeneralInfo      LecturerGeneralData     `json:"dataumum"`
	TeachingHistory  []LecturerTeachingData  `json:"datamengajar"`
	EducationHistory []LecturerEducationData `json:"datapendidikan"`
}

type LecturerGeneralData struct {
	Functional         string `json:"fungsional"`
	SDMID              string `json:"id_sdm"`
	TeachingStatus     string `json:"ikatankerja"`
	Gender             string `json:"jk"`
	ProgrammeDetailID  string `json:"linkprodi"`
	CollegeDetailID    string `json:"linkpt"`
	Programme          string `json:"namaprodi"`
	College            string `json:"namapt"`
	Name               string `json:"nm_sdm"`
	LastEducationLevel string `json:"pend_tinggi"`
	Status             string `json:"status_keaktifan"`
	BirthCity          string `json:"tmpt_lahir"`
}

type LecturerTeachingData struct {
	Semester        string `json:"id_smt"`
	SubjectCode     string `json:"kode_mk"`
	CollegeDetailID string `json:"linkpt"`
	College         string `json:"namapt"`
	SubjectName     string `json:"nm_mk"`
}

type LecturerEducationData struct {
	DegreeLevel    string `json:"namajenjang"`
	College        string `json:"nm_sp_formal"`
	Degree         string `json:"singkat_gelar"`
	GraduationYear string `json:"thn_lulus"`
}

// GetLecturerDetailByNIDN is used to get the lecturer information in detail.
// It receives NIDN in string type as parameter.
func (c *Client) GetLecturerDetailByNIDN(nidn string) (res LecturerDetail, err error) {
	raw, err := c.searchAll(nidn)
	if err != nil {
		return
	}

	if len(raw.Lecturers) < 1 {
		err = fmt.Errorf("empty result for id: %s", nidn)
		return
	}

	match := strings.Split(raw.Lecturers[0].WebsiteLink, "/")
	if len(match) < 3 || strings.TrimSpace(match[2]) == "" {
		err = fmt.Errorf("invalid result for id: %s", nidn)
		return
	}

	res, err = c.getLecturerDetail(match[2])
	if err != nil {
		return
	}

	return
}

// GetLecturerDetailByDetailID is used to get the lecturer information in detail.
// It receives detailID as parameter.
func (c *Client) GetLecturerDetailByDetailID(detailID string) (res LecturerDetail, err error) {
	return c.getLecturerDetail(detailID)
}

func (c *Client) getLecturerDetail(detailID string) (res LecturerDetail, err error) {
	if err = checkParamString(detailID); err != nil {
		return
	}

	req, err := c.createRequest(http.MethodGet, getLecturerDetailPath+"/"+detailID)
	if err != nil {
		return
	}

	err = c.doRequest(req, &res)
	if err != nil {
		return
	}

	res.splitLecturerDetailID()

	return
}

func (r *LecturerDetail) splitLecturerDetailID() {
	collegeDetailID := strings.Split(r.GeneralInfo.CollegeDetailID, "/")
	if len(collegeDetailID) == 3 {
		r.GeneralInfo.CollegeDetailID = collegeDetailID[2]
	}

	programmeDetailID := strings.Split(r.GeneralInfo.ProgrammeDetailID, "/")
	if len(programmeDetailID) == 3 {
		r.GeneralInfo.ProgrammeDetailID = programmeDetailID[2]
	}

	for _, v := range r.TeachingHistory {
		collegeDetailID := strings.Split(v.CollegeDetailID, "/")
		if len(collegeDetailID) == 3 {
			v.CollegeDetailID = collegeDetailID[2]
		}
	}

}
