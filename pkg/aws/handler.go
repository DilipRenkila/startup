package aws

import (
	"encoding/json"
	"fmt"

	. "github.com/mlabouardy/komiser/services/aws"
	. "github.com/mlabouardy/komiser/services/cache"
)

type AWSHandler struct {
	cache    Cache
	multiple bool
	aws      AWS
}

func NewAWSHandler(cache Cache, multiple bool) *AWSHandler {
	awsHandler := AWSHandler{
		cache:    cache,
		multiple: multiple,
		aws:      AWS{},
	}
	return &awsHandler
}

func respondWithError( code int, msg string) {
	respondWithJSON( code, map[string]string{"error": msg})
}

func respondWithJSON( code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(code)
	fmt.Println(response)
}
