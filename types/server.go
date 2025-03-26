package types

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

type Server struct {
	Server     string
	ApiKey     string
	SkipVerify bool
}
type Project struct {
	Name        string        `json:"name"`
	Version     string        `json:"version"`
	Parent      *Parent       `json:"parent"`
	Classifier  string        `json:"classifier"`
	AccessTeams []*Team       `json:"accessTeams,omitempty"`
	Tags        []interface{} `json:"tags,omitempty"`
	Active      bool          `json:"active"`
	IsLatest    bool          `json:"isLatest"`
}

type Team struct {
	Name    string        `json:"name"`
	Uuid    string        `json:"uuid"`
	ApiKeys []interface{} `json:"apiKeys,omitempty"`
}
type Parent struct {
	Uuid string `json:"uuid"`
}

type UploadBomArg struct {
	Method         string `json:"method"`
	ProjectName    string `json:"name"`
	ProjectVersion string `json:"version"`
	ProjectId      string `json:"uuid"`
	Async          bool   `json:"async"`
	File           string `json:"bom"`
	Content        string `json:"content"`
	AutoCreate     bool   `json:"autoCreate"`
}

func (u *UploadBomArg) SetContent() error {
	content := []byte(u.Content)
	if strings.TrimSpace(u.File) != "" {
		var err error
		content, err = os.ReadFile(u.File)
		if err != nil {
			return err
		}
	}
	if strings.ToUpper(u.Method) == "PUT" {
		u.Content = base64.StdEncoding.EncodeToString(content)
	} else {
		u.Content = string(content)
	}
	if strings.TrimSpace(u.Content) == "" {
		return fmt.Errorf("content is empty")
	}
	return nil
}
