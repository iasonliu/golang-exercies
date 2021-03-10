package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

type key struct {
	key     []byte
	created time.Time
}

var currentKid = ""
var keys = map[string]key{}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expires")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("Invalid session ID")
	}
	return nil
}

func main() {
	pass := "12345956"
	hashedPass, err := hashPassword(pass)
	if err != nil {
		log.Fatalln(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("you not login")
	}

	log.Println("you login")
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while hash from passwrod: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashPass []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), hashPass)
	if err != nil {
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, keys[currentKid].key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error in signMessage while hashing massage: %w", err)
	}
	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg []byte, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, err
	}
	same := hmac.Equal(newSig, sig)
	return same, nil
}

func createToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(keys[currentKid].key)
	if err != nil {
		return "", fmt.Errorf("Error in createToken when signing token : %w", err)
	}
	return signedToken, nil
}

func generateNewKey() error {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil {
		return fmt.Errorf("Error in generateNewKey while %w", err)
	}
	uid := uuid.Must(uuid.NewRandom())
	keys[uid.String()] = key{
		key:     newKey,
		created: time.Now(),
	}
	currentKid = uid.String()
	return nil
}

func parseToken(signedToken string) (*UserClaims, error) {
	claims := &UserClaims{}
	t, err := jwt.ParseWithClaims(signedToken, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm")
		}
		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("Invalid key ID")
		}
		k, ok := keys[kid]
		if !ok {
			return nil, fmt.Errorf("Inalid key ID")
		}
		return k, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error in parseToken While parsing token: %w", err)
	}
	if !t.Valid {
		return nil, fmt.Errorf("Error in parseTOken, token is not valid")
	}
	return t.Claims.(*UserClaims), nil
}
