package jwt

import (
	"e.coding.net/double-j/ego/colago/common/model"
	"encoding/json"
	"github.com/cristalhq/jwt"
)

func JwtVerify(bearerToken string, salt string) (*model.TokenData, error) {
	token, errParse := jwt.Parse([]byte(bearerToken))
	if errParse != nil {
		return nil, errParse
	}

	key := []byte(salt)
	signer, _ := jwt.NewHS256(key)
	errVerify := signer.Verify(token.Payload(), token.Signature())
	if errVerify != nil {
		return nil, errParse
	}

	tokenData := new(model.TokenData)
	err := json.Unmarshal(token.RawClaims(), tokenData)
	if err != nil {
		return nil, err
	}

	return tokenData, nil
}

func JwtBuild(tokenData model.TokenData, salt string) (string, error) {
	key := []byte(salt)
	signer, _ := jwt.NewHS256(key)
	builder := jwt.NewTokenBuilder(signer)

	token, errBuild := builder.Build(tokenData)
	if errBuild != nil {
		return "", errBuild
	}
	return string(token.Raw()), nil
}
