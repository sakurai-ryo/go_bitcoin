package shared

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type Secret struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

func GetSecret() (*Secret, error) {
	secretName := "buy-bitcoin-secret"
	region := "ap-northeast-1"

	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(input)
	if result.SecretString == nil {
		return nil, errors.New("Secret is empty")
	}
	if err != nil {
		return nil, err
	}

	var secretString string
	var secret Secret
	secretString = *result.SecretString
	if err := json.Unmarshal([]byte(secretString), &secret); err != nil {
		return nil, err
	}
	return &secret, nil
}
