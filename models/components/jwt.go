// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

type Algorithm string

const (
	AlgorithmHs256 Algorithm = "HS256"
	AlgorithmHs384 Algorithm = "HS384"
	AlgorithmHs512 Algorithm = "HS512"
	AlgorithmRs256 Algorithm = "RS256"
	AlgorithmRs384 Algorithm = "RS384"
	AlgorithmRs512 Algorithm = "RS512"
	AlgorithmEs256 Algorithm = "ES256"
	AlgorithmEs384 Algorithm = "ES384"
	AlgorithmEs512 Algorithm = "ES512"
	AlgorithmPs256 Algorithm = "PS256"
	AlgorithmPs384 Algorithm = "PS384"
	AlgorithmPs512 Algorithm = "PS512"
	AlgorithmEdDsa Algorithm = "EdDSA"
)

func (e Algorithm) ToPointer() *Algorithm {
	return &e
}
func (e *Algorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HS256":
		fallthrough
	case "HS384":
		fallthrough
	case "HS512":
		fallthrough
	case "RS256":
		fallthrough
	case "RS384":
		fallthrough
	case "RS512":
		fallthrough
	case "ES256":
		fallthrough
	case "ES384":
		fallthrough
	case "ES512":
		fallthrough
	case "PS256":
		fallthrough
	case "PS384":
		fallthrough
	case "PS512":
		fallthrough
	case "EdDSA":
		*e = Algorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Algorithm: %v", v)
	}
}

type JWTConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *JWTConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type Jwt struct {
	Algorithm *Algorithm `default:"HS256" json:"algorithm"`
	// Unix epoch when the resource was created.
	CreatedAt    *int64       `json:"created_at,omitempty"`
	ID           *string      `json:"id,omitempty"`
	Key          *string      `json:"key,omitempty"`
	RsaPublicKey *string      `json:"rsa_public_key,omitempty"`
	Secret       *string      `json:"secret,omitempty"`
	Tags         []string     `json:"tags,omitempty"`
	Consumer     *JWTConsumer `json:"consumer,omitempty"`
}

func (j Jwt) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *Jwt) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Jwt) GetAlgorithm() *Algorithm {
	if o == nil {
		return nil
	}
	return o.Algorithm
}

func (o *Jwt) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Jwt) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Jwt) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *Jwt) GetRsaPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.RsaPublicKey
}

func (o *Jwt) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *Jwt) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Jwt) GetConsumer() *JWTConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}
