package APNsGoServer

import (
	"crypto/ecdsa"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"github.com/dgrijalva/jwt-go"
	"time"
	"errors"
)

var (
	privateKey *ecdsa.PrivateKey
)

type Config struct {
	PrivateKey string `toml:"key_path"`
	BundleIdentifier string `toml:"bundle_identifier"`
	Sandbox bool `toml:"sandbox"`
	TeamId string `toml:"team_id"`
	KeyId string `toml:"key_id"`
	Topic string `toml:"topic"`
}

func buildDefaultConfig() Config {
	return Config{
		PrivateKey: "",
		BundleIdentifier: "",
		Sandbox: true,
		TeamId: "",
		KeyId: "",
	}
}

func LoadConfig(filePath *string) (Config, error) {
	var config Config
	_, err := toml.DecodeFile(*filePath, &config)
	if err != nil {
		return buildDefaultConfig(), err
	}
	return config, nil
}

func readPrivateKey() (*ecdsa.PrivateKey, error) {
	file, err := ioutil.ReadFile(Conf.PrivateKey)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(file)
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	if privateKey, ok := key.(*ecdsa.PrivateKey); ok {
		return privateKey, nil
	}
	return nil, errors.New("")
}

func CreateJWT() (string, error) {
	if privateKey == nil {
		newPrivateKey, err := readPrivateKey()
		if err != nil {
			return "", err
		}
		privateKey = newPrivateKey
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		jwt.MapClaims{
			"iat": time.Now().Unix(),
			"iss": Conf.TeamId,
		})
	token.Header["alg"] = "ES256"
	token.Header["kid"] = Conf.KeyId
	s, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return s, nil
}
