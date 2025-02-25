package aws

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func (handler *AWSHandler) EC2InstancesHandler() {
	profile := "default"
	cfg, err := external.LoadDefaultAWSConfig()

	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError( http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	key := fmt.Sprintf("aws.%s.ec2.instances", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON( 200, response)
	} else {
		response, err := handler.aws.DescribeInstances(cfg)
		if err != nil {
			respondWithError( http.StatusInternalServerError, "ec2:DescribeInstances is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON( 200, response)
		}
	}
}

