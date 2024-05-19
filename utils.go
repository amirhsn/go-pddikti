package gopddikti

import (
	"errors"
	"regexp"
	"strings"
	"sync"
)

var (
	reStudent, reLecturer, reCollege, reProgramme *regexp.Regexp
)

func initRegex() {
	reStudent = regexp.MustCompile(`([\w\s]+)\((\d*)\s*\),\s*PT\s*:\s*(.*?),\s*Prodi\s*:\s*(.*)`)
	reLecturer = regexp.MustCompile(`(\w+(?:\s+\w+)*),\s*NIDN\s*:\s*(\d+),?\s*PT\s*:\s*(.*?),\s*Prodi\s*:\s*(.*)`)
	reCollege = regexp.MustCompile(`Nama PT:\s*(.*?),?\s*NPSN:\s*(\d+)?\s*,?\s*Singkatan:\s*(.*?),?\s*Alamat PT:\s*(.*)`)
	reProgramme = regexp.MustCompile(`Nama Prodi:\s*(.*?),?\s*Jenjang Prodi:\s*(.*?),?\s*Nama Lembaga:\s*(.*)`)
}

func appendError(stack *[]error, err error, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	*stack = append(*stack, err)
}

func checkParamString(param string) error {
	if strings.TrimSpace(param) == "" {
		return errors.New("param is empty")
	}
	return nil
}

func parseStudentResponse(param []SearchResponse) []Student {
	students := make([]Student, len(param))
	for i, v := range param {
		infoMatches := reStudent.FindStringSubmatch(v.Text)
		linkMatches := strings.Split(v.WebsiteLink, "/")
		if len(infoMatches) < 5 || len(linkMatches) < 3 {
			continue
		}
		students[i] = Student{
			Name:      infoMatches[1],
			ID:        infoMatches[2],
			College:   infoMatches[3],
			Programme: infoMatches[4],
			DetailID:  linkMatches[2],
		}
	}
	return students
}

func parseLecturerResponse(param []SearchResponse) []Lecturer {
	lecturers := make([]Lecturer, len(param))
	for i, v := range param {
		infoMatches := reLecturer.FindStringSubmatch(v.Text)
		linkMatches := strings.Split(v.WebsiteLink, "/")
		if len(infoMatches) < 5 || len(linkMatches) < 3 {
			continue
		}
		lecturers[i] = Lecturer{
			Name:      infoMatches[1],
			ID:        infoMatches[2],
			College:   infoMatches[3],
			Programme: infoMatches[4],
			DetailID:  linkMatches[2],
		}
	}
	return lecturers
}

func parseCollegeResponse(param []SearchResponse) []College {
	colleges := make([]College, len(param))
	for i, v := range param {
		infoMatches := reCollege.FindStringSubmatch(v.Text)
		linkMatches := strings.Split(v.WebsiteLink, "/")
		if len(infoMatches) < 5 || len(linkMatches) < 3 {
			continue
		}
		colleges[i] = College{
			Name:     infoMatches[1],
			ID:       infoMatches[2],
			Nickname: infoMatches[3],
			Address:  infoMatches[4],
			DetailID: linkMatches[2],
		}
	}
	return colleges
}

func parseProgrammeResponse(param []SearchResponse) []Programme {
	programmes := make([]Programme, len(param))
	for i, v := range param {
		infoMatches := reProgramme.FindStringSubmatch(v.Text)
		linkMatches := strings.Split(v.WebsiteLink, "/")
		if len(infoMatches) < 4 || len(linkMatches) < 3 {
			continue
		}
		programmes[i] = Programme{
			Name:     infoMatches[1],
			Level:    infoMatches[2],
			College:  infoMatches[3],
			DetailID: linkMatches[2],
		}
	}
	return programmes
}
