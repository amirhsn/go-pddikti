package gopddikti

import (
	"net/http"
	"time"
)

const (
	getCollegeDetailPath = "/v2/detail_pt"
)

type CollegeGeneralInfo struct {
	NPSN                               string                `json:"npsn"`
	Status                             string                `json:"stat_sp"`
	Name                               string                `json:"nm_lemb"`
	EstablishmentDate                  string                `json:"tgl_berdiri"`
	EstablishmentCertificateNumber     string                `json:"sk_pendirian_sp"`
	EstablishmentCertificateNumberDate time.Time             `json:"tgl_sk_pendirian_sp"`
	Address                            string                `json:"jln"`
	City                               string                `json:"nama_wil"`
	PostalCode                         string                `json:"kode_pos"`
	TelephoneNumber                    string                `json:"no_tel"`
	FaxNumber                          string                `json:"no_fax"`
	Email                              string                `json:"email"`
	Website                            string                `json:"website"`
	Latitude                           float64               `json:"lintang"`
	Longitude                          float64               `json:"bujur"`
	DetailID                           string                `json:"id_sp"`
	SurfaceArea                        int                   `json:"luas_tanah"`
	TotalLaboratory                    int                   `json:"laboratorium"`
	TotalClassroom                     int                   `json:"ruang_kelas"`
	TotalLibrary                       int                   `json:"perpustakaan"`
	IsInternetProvided                 bool                  `json:"internet"`
	IsElectricityProvided              bool                  `json:"listrik"`
	RectorName                         string                `json:"nama_rektor"`
	AccreditationList                  []AccreditationDetail `json:"akreditasi_list"`
}

type AccreditationDetail struct {
	Accreditation           string    `json:"akreditasi"`
	AccreditationDate       time.Time `json:"tgl_akreditasi"`
	AccreditationExpiryDate time.Time `json:"tgl_berlaku"`
}

// GetCollegeGeneralInfoByDetailID is used to get the university or college information in detail.
// It receives detailID as parameter.
func (c *Client) GetCollegeGeneralInfoByDetailID(detailID string) (res CollegeGeneralInfo, err error) {
	return c.getCollegeGeneralInfo(detailID)
}

func (c *Client) getCollegeGeneralInfo(detailID string) (res CollegeGeneralInfo, err error) {
	if err = checkParamString(detailID); err != nil {
		return
	}

	req, err := c.createRequest(http.MethodGet, getCollegeDetailPath+"/"+detailID)
	if err != nil {
		return
	}

	err = c.doRequest(req, &res)
	if err != nil {
		return
	}

	return
}
