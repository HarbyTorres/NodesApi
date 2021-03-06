package models

import (
	"time"
)

type CreateDrawflow struct {
	Name  string      `json:"name"`
	Body  DrawflowMap `json:"body"`
	DType []string    `json:"dgraph.type,omitempty"`
}

type DrawflowMap struct {
	Drawflow struct {
		Home struct {
			Data map[string]interface{} `json:"data"`
		} `json:"Home"`
	} `json:"drawflow"`
}

type DgraphMapping struct {
	Data struct {
		QueryDrawflow []struct {
			Body      DrawflowMap `json:"body"`
			ID        string      `json:"id"`
			Name      string      `json:"name"`
			CreatedAt time.Time   `json:"createdAt,omitempty"`
		} `json:"queryDrawflow"`
	} `json:"data"`
	Extensions struct {
		TouchedUids int `json:"touched_uids"`
		Tracing     struct {
			Version   int       `json:"version"`
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
			Duration  int       `json:"duration"`
			Execution struct {
				Resolvers []struct {
					Path        []string `json:"path"`
					ParentType  string   `json:"parentType"`
					FieldName   string   `json:"fieldName"`
					ReturnType  string   `json:"returnType"`
					StartOffset int      `json:"startOffset"`
					Duration    int      `json:"duration"`
					Dgraph      []struct {
						Label       string `json:"label"`
						StartOffset int    `json:"startOffset"`
						Duration    int    `json:"duration"`
					} `json:"dgraph"`
				} `json:"resolvers"`
			} `json:"execution"`
		} `json:"tracing"`
	} `json:"extensions"`
}
