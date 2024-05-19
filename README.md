[![Library Status](https://img.shields.io/badge/status-unofficial-yellow.svg)]()
[![Go Reference](https://pkg.go.dev/badge/github.com/amirhsn/go-pddikti.svg)](https://pkg.go.dev/github.com/amirhsn/go-pddikti)
[![Go Report Card](https://goreportcard.com/badge/github.com/amirhsn/go-pddikti)](https://goreportcard.com/report/github.com/amirhsn/go-pddikti)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

# go-pddikti

Unofficial Go Client SDK for retrieving higher education data from PDDikti (Pangkalan Data Pendidikan Tinggi). Currently, the data is public, but the official API and its documentation are not yet available.

## Installing

```sh
$ go get -u github.com/amirhsn/go-pddikti
```

## Usage

```go
// Init client
client, err := pddikti.InitClient(&pddikti.ClientConfig{
	ContextTimeout: 5000,
	UseCorsPolicy:  false,
})
if err != nil {
	return err
}

searchResult, err := client.Search("181344007")
if err != nil {
	return err
}

if len(searchResult.Students) == 0 {
	return errors.New("empty data")
}

for _, student := searchResult.Students {
	fmt.Printf("Name: %s", student.Name)
	fmt.Printf("College: %s", student.College)
	fmt.Printf("Study Programme: %s", student.Programme)
	fmt.Printf("Student ID: %s", student.ID)
	fmt.Printf("Detail ID: %s \n", student.DetailID)
}
```

## Features
- Retrieve higher education data in real-time, as displayed on the [PDDikti website](https://pddikti.kemdikbud.go.id/).
- Configurable context timeout and the use of Origin and Referer headers for CORS.
- Supported entities:
  - Students
  - Lecturers
  - Universities / Colleges
  - Study Programs

### Client
The client must be initialized first using the `InitClient()` method. This method receives one parameter, which is a pointer to the configuration for the client. If the parameter is set to `nil`, the default configuration will be used. The default configuration consist of:
- No CORS policy
- The context timeout follows the default settings of the Go HTTP client

### Method List
#### Available Methods

| Method                         | Available            | Description                                                  |
| ------------------------------ | -------------------- | ------------------------------------------------------------ |
| `Search()`                       | ✅                   | Search all entities                                          |
| `SearchStudents()`               | ✅                   | Search students by given name or ID (NIM)                    |
| `SearchLecturers()`              | ✅                   | Search lecturers by given name or ID (NIDN)                  |
| `SearchColleges()`               | ✅                   | Search colleges by given name                                |
| `SearchProgrammes()`             | ✅                   | Search programmes by given name                              |
| `GetStudentDetailByNIM()`        | ✅                   | Get student information in more detail by NIM                |
| `GetStudentDetailByDetailID()`   | ✅                   | Get student information in more detail by Detail ID          |
| `GetLecturerDetailByNIDN()`      | ✅                   | Get lecturer information in more detail by NIDN              |
| `GetLecturerDetailByDetailID()`  | ✅                   | Get lecturer information in more detail by Detail ID         |
| `GetCollegeGeneralInfoByDetailID()` | ✅                | Get college or university general information by Detail ID   |
| `GetProgrammeDetailByDetailID()` | ✅                   | Get programme information in more detail by Detail ID        |

### Feature Not Yet Available
- Currently, information retrieved from universities using the detail ID only returns general information about the institution. Features such as the list of teaching lecturers, lecturer-to-student ratios, and other specific details are not yet handled. These features will be the focus of future development.

## Current Limitations
- PDDikti often returns a 403 error code due to an expired certificate, which can persist for more than a day. Currently, if an error occurs, it is directly returned from the PDDikti API. A circuit breaker and caching needs to be implemented in the future.
- The return data from PDDikti needs to be parsed first using regex. The regex initialization occurs in the first step with `InitClient()`, but the parsing process using regex still takes some time.

## Contributing

Open for contributions. Raise your PR regarding an issue or by initiative.

## Authors

- **Amir Husein** - _Initial work_ - [amirhsn](https://github.com/amirhsn)

## License
[MIT](https://github.com/amirhsn/go-pddikti/blob/main/LICENSE) © Amir Husein