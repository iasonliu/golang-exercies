# Notes
## Authenticate and Authorization
## HTTP Basic Authentication
- Basic authentication part of the specification of HTTP
    - send username/password with every request
    - users authorization header & keyword "basic"
        - Put `username:password` covert to Base64
            - `base64.StdEncoding.EncodeToString([]byte("user:password"))`
        - into `Authorization` HTTP Header

## Storing Passwords
- nerver store passwords, keep on-way "hash" values for it
- hashing algorithms
    - bcrypt
    - scrypt

## Bearer Tokens & Hmac
- bearer tokens
    - added to http spec with OAUTH2
    - uses `Authoriztion` header & keyword `Bearer `
- to prevent faked bearer tokens, requier using `cryptographic signing`
    - Keyed-Hash Message Authentication Code (HMAC) using to crypto signing
- https://godoc.org/crypto/hmac

## JSON Web Token (JWT)
```
{JWT standard fields}.{Your fields}.Signature
```
github.com/dgrijalva/jwt-go

```
# list versions for go mod
go list -m -versions github.com/dgrijalva/jwt-go
```
- Hashing
    - MD5
    - SHA
    - Bcrypt
    - Scrypt
- Signing
    - Symmetric Key
        - HMAC
            - same key to sign (encrypt) / verify(decrypt)
    - Asymmetric Key
        - RSA
        - ECDSA - better than RSA; faster; smaller keys
        - private key to sign (encrypt) / public key to verify (decrypt)
- JWT
    - JWT requier Singing like HMAC
    - HMAC singing requier Hashing like SHA