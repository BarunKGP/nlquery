package ports

import (
	"encoding/json"
	"io"
	// "github.com/BarunKGP/nlquery/internal/database"
)

type ThreadFile struct {
	Columns []string `json:"columns"`
	Types   []string `json:"types"`
}

type Thread struct {
	Query         string      `json:"query"`
	Id            int64       `json:"id"`
	UserId        int64       `json:"userId"` //! We will have to get the user's identity separately
	ThreadFileId  int64       `json:"threadFileId"`
	FileDetails   *ThreadFile `json:"file,omitempty"`
	ParsedColumns []string    `json:"parsedColumns,omitempty"`
	ParsedTypes   []string    `json:"parsedTypes,omitempty"`
}

func (t *Thread) ReadJson(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(t); err != nil {
		return err
	}
	return nil
}

// func (t *Thread) ReadDb(r database.Thread) error {
// 	t.Id = r.ID
// 	t.Query = r.Query
// 	t.UserId = r.Userid.Int64
// 	t.FileDetails = r
// }
