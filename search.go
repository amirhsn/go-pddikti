package gopddikti

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

const (
	searchAllPath        = "/hit"
	searchAllPathStudent = "/hit_mhs"
)

type SearchAllResponse struct {
	Students   []Student
	Lecturers  []Lecturer
	Colleges   []College
	Programmes []Programme
}

type Student struct {
	Name      string
	ID        string
	College   string
	Programme string
	DetailID  string
}

type Lecturer struct {
	Name      string
	ID        string
	College   string
	Programme string
	DetailID  string
}

type College struct {
	Name     string
	Nickname string
	ID       string
	Address  string
	DetailID string
}

type Programme struct {
	Name     string
	Level    string
	College  string
	DetailID string
}

type RawStudentsResponse struct {
	Students []SearchResponse `json:"mahasiswa"`
}

type RawAllResponse struct {
	Lecturers  []SearchResponse `json:"dosen"`
	Colleges   []SearchResponse `json:"pt"`
	Programmes []SearchResponse `json:"prodi"`
}

type SearchResponse struct {
	Text        string `json:"text"`
	WebsiteLink string `json:"website-link"`
}

// Search is used to search entities list based on given param.
// It will returns all four entities in the type of array.
func (c *Client) Search(param string) (res SearchAllResponse, err error) {
	if err = checkParamString(param); err != nil {
		return
	}

	var (
		errStack []error
		mu       sync.Mutex
		wg       sync.WaitGroup

		resRawStudent RawStudentsResponse
		resRawAll     RawAllResponse
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		reqStudent, err := c.createRequest(http.MethodGet, searchAllPathStudent+"/"+url.QueryEscape(param))
		if err != nil {
			appendError(&errStack, err, &mu)
		} else {
			err = c.doRequest(reqStudent, &resRawStudent)
			if err != nil {
				appendError(&errStack, err, &mu)
			}
		}
	}()

	go func() {
		defer wg.Done()
		resRawAll, err = c.searchAll(param)
		if err != nil {
			appendError(&errStack, err, &mu)
		}
	}()

	wg.Wait()

	if len(errStack) > 0 {
		combinedErr := errStack[0]
		for _, v := range errStack[1:] {
			combinedErr = fmt.Errorf("%v; %w", combinedErr, v)
		}
		err = combinedErr
		return
	}

	wg.Add(4)

	go func() {
		defer wg.Done()
		res.Students = parseStudentResponse(resRawStudent.Students)
	}()
	go func() {
		defer wg.Done()
		res.Lecturers = parseLecturerResponse(resRawAll.Lecturers)
	}()
	go func() {
		defer wg.Done()
		res.Colleges = parseCollegeResponse(resRawAll.Colleges)
	}()
	go func() {
		defer wg.Done()
		res.Programmes = parseProgrammeResponse(resRawAll.Programmes)
	}()

	wg.Wait()
	return
}

// SearchStudents is used to search students based on given param.
// It will returns all the list of students.
func (c *Client) SearchStudents(param string) (res []Student, err error) {
	raw, err := c.searchStudents(param)

	res = parseStudentResponse(raw.Students)
	return
}

// SearchLecturers is used to search lecturers based on given param.
// It will returns all the list of lecturers.
func (c *Client) SearchLecturers(param string) (res []Lecturer, err error) {
	resp, err := c.searchAll(param)

	res = parseLecturerResponse(resp.Lecturers)
	return
}

// SearchColleges is used to search colleges based on given param.
// It will returns all the list of colleges.
func (c *Client) SearchColleges(param string) (res []College, err error) {
	resp, err := c.searchAll(param)

	res = parseCollegeResponse(resp.Colleges)
	return
}

// SearchProgrammes is used to search programmes based on given param.
// It will returns all the list of programmes.
func (c *Client) SearchProgrammes(param string) (res []Programme, err error) {
	resp, err := c.searchAll(param)

	res = parseProgrammeResponse(resp.Programmes)
	return
}

func (c *Client) searchStudents(param string) (res RawStudentsResponse, err error) {
	if err = checkParamString(param); err != nil {
		return
	}

	req, err := c.createRequest(http.MethodGet, searchAllPathStudent+"/"+url.QueryEscape(param))
	if err != nil {
		return
	}

	err = c.doRequest(req, &res)
	if err != nil {
		return
	}

	return
}

func (c *Client) searchAll(param string) (res RawAllResponse, err error) {
	if err = checkParamString(param); err != nil {
		return
	}

	req, err := c.createRequest(http.MethodGet, searchAllPath+"/"+url.QueryEscape(param))
	if err != nil {
		return
	}

	err = c.doRequest(req, &res)
	if err != nil {
		return
	}

	return
}
