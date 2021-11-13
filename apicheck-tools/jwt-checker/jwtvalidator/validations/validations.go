/*
 * Copyright 2020 Banco Bilbao Vizcaya Argentaria, S.A.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package validations

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	jwt2 "github.com/cristalhq/jwt/v2"
)

type ListFlag []string

func (l *ListFlag) String() string { return strings.Join(*l, ", ") }

func (l *ListFlag) Set(value string) error {
	*l = append(*l, value)
	return nil
}

type Options struct {
	AllowUnsignedTokens bool
	AllowedSignAlgs     ListFlag
	Secret              string
	Issuer              string
	Subject             string
	PermittedFor        string
	IssuedBefore        string
	TimeIssuedBefore    time.Time
	ExpiresAt           string
	TimeExpiresAt       time.Time
}

var base64Decode = base64.RawURLEncoding.Decode

func ValidateToken(rawToken []byte, opt Options) []error {
	valError := []error{}

	token, errParse := jwt2.Parse(rawToken)
	if errParse != nil {
		return []error{errParse}
	}

	// fmt.Printf("Algorithm %v\n", token.Header().Algorithm)
	// fmt.Printf("Type      %v\n", token.Header().Type)
	// fmt.Printf("Claims    %v\n", string(token.RawClaims()))
	// fmt.Printf("Payload   %v\n", string(token.Payload()))
	// fmt.Printf("Token     %v\n", string(token.Raw()))

	// Validate header
	if err := validateJWTHeader(token.Header(), opt); len(err) > 0 {
		valError = append(valError, err...)
	}

	// Validate signature
	if v, err := getSignValidator(token.Header().Algorithm, opt.Secret); err != nil {
		return append(valError, err)
	} else {
		token, errSign := jwt2.ParseAndVerify(rawToken, v)
		if errSign != nil {
			valError = append(valError, errSign)
		} else {
			// Validate claims
			claims := &jwt2.StandardClaims{}
			_ = json.Unmarshal(token.RawClaims(), claims)
			if err := validateJWTStdClaims(*claims, opt); err != nil {
				valError = append(valError, err...)
			}
		}
	}

	if len(valError) > 0 {
		return valError
	} else {
		return nil
	}
	// Output:
	// Algorithm HS256
	// Type      JWT
	// Claims    {"aud":"admin","jti":"random-unique-string"}
	// Payload   eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhZG1pbiIsImp0aSI6InJhbmRvbS11bmlxdWUtc3RyaW5nIn0
	// Token     eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhZG1pbiIsImp0aSI6InJhbmRvbS11bmlxdWUtc3RyaW5nIn0.dv9-XpY9P8ypm1uWQwB6eKvq3jeyodLA7brhjsf4JVs
}

func getSignValidator(alg jwt2.Algorithm, secret string) (jwt2.Verifier, error) {

	if secret == "" {
		return nil, fmt.Errorf("No secret provided")
	}
	buf := make([]byte, len(secret))
	if n, err := base64Decode(buf, []byte(secret)); err != nil {
		return nil, fmt.Errorf("Error base64decoding (%s)", err.Error())
	} else {
		buf = buf[:n-1]
	}

	switch alg {
	//	case jwt2.EdDSA:
	//		// Create public key from byte array
	//		var pubKey ed25519.PrivateKey
	//		if , err := jwt2.NewVerifierEdDSA(pubKey); err != nil {
	//			return v, nil
	//		} else {
	//			return nil, err
	//		}
	//	case jwt2.ES256, jwt2.ES384, jwt2.ES512:
	//		// Create public key from byte array
	//		var pubKey *ecdsa.PublicKey
	//    if s, err := jwt2.NewVerifierES(alg, ); err != nil {
	//			return s, nil
	//		} else {
	//			return nil, err
	//		}
	case jwt2.HS256, jwt2.HS384, jwt2.HS512:
		if v, err := jwt2.NewVerifierHS(alg, buf); err == nil {
			return v, nil
		} else {
			return nil, err
		}
	//	case jwt2.PS256, jwt2.PS384, jwt2.PS512:
	//		// Create public key from byte array
	//		var pubKey *rsa.PublicKey
	//    if v, err := jwt2.NewVerifierPS(alg, pubKey); err != nil {
	//			return v, nil
	//		} else {
	//			return nil, err
	//		}
	//	case jwt2.RS256, jwt2.RS384, jwt2.RS512:
	//		// Create public key from byte array
	//		var pubKey *rsa.PublicKey
	//		if v, err := jwt2.NewVerifierRS(alg, pubKey); err != nil {
	//			return s, nil
	//		} else {
	//			return nil, err
	//		}
	default:
		return nil, fmt.Errorf("Unsupported Algorithm: %s", alg)
	}
}

func validateJWTHeader(header jwt2.Header, opt Options) []error {
	errors := []error{}

	if header.Type != "JWT" {
		errors = append(errors, fmt.Errorf("JWT type (%s) not supported", header.Type))
	}
	if header.ContentType != "" && header.ContentType != "JWT" {
		errors = append(errors, fmt.Errorf("JWT content type (%s) not supported", header.ContentType))
	}
	if !opt.AllowUnsignedTokens && header.Algorithm == "none" {
		errors = append(errors, fmt.Errorf("JWT is unsigned"))
	}
	if len(opt.AllowedSignAlgs) > 0 {
		found := false
		for _, alg := range opt.AllowedSignAlgs {
			if header.Algorithm == jwt2.Algorithm(alg) {
				found = true
				break
			}
		}
		if !found {
			errors = append(errors, fmt.Errorf("JWT signing algorithm (%s) is not allowed", header.Algorithm))
		}
	}

	return errors
}

func validateJWTStdClaims(claims jwt2.StandardClaims, opt Options) []error {
	errors := make([]error, 0)

	if claims.Issuer == "" {
		errors = append(errors, fmt.Errorf("No issuer claim found"))
	} else if opt.Issuer != "" && !claims.IsIssuer(opt.Issuer) {
		errors = append(errors, fmt.Errorf("Issuer claim doesn't match. Expected: %s, got: %s", opt.Issuer, claims.Issuer))
	}
	if claims.Subject == "" {
		errors = append(errors, fmt.Errorf("No subject claim found"))
	} else if opt.Subject != "" && !claims.IsSubject(opt.Subject) {
		errors = append(errors, fmt.Errorf("Subject claim doesn't match. Expected: %s, got: %s", opt.Subject, claims.Subject))
	}
	if len(claims.Audience) == 0 {
		errors = append(errors, fmt.Errorf("No audience claim found"))
	} else if opt.PermittedFor != "" && !claims.IsForAudience(opt.PermittedFor) {
		errors = append(errors, fmt.Errorf("Audience claim doesn't match. Expected: %s, got: %s", opt.PermittedFor, strings.Join(claims.Audience, ",")))
	}
	if opt.IssuedBefore != "" && !claims.IsValidNotBefore(opt.TimeIssuedBefore) {
		errors = append(errors, fmt.Errorf("NotBefore claim doesn't match. Expected: %s, got: %v", opt.IssuedBefore, claims.NotBefore))
	}
	if opt.ExpiresAt != "" && !claims.IsValidExpiresAt(opt.TimeExpiresAt) {
		errors = append(errors, fmt.Errorf("ExpiresAt claim doesn't match. Expected: %s, got: %v", opt.ExpiresAt, claims.ExpiresAt))
	}

	//  jwt2.AudienceChecker([]string{"admin"}),
	//  jwt2.IDChecker("random-unique-string"),

	return errors
}

// mySecretPassword = bXlTZWNyZXRQYXNzd29yZAo=
// mySecretPasswordmySecretPassword = bXlTZWNyZXRQYXNzd29yZG15U2VjcmV0UGFzc3dvcmQK
/*******************************************************************************
/* header.typ === "JWT"
/* header.alg !=="none" except Options.allowUnsignedTokens
/* header.alg in Options.allowedSignAlgs
token is signed
/* payload.issuer !=== ""
/* payload.issuer ==== Options.issuer
/* payload.subject !=== ""
/* payload.subject ==== Options.subject
/* payload.audience !=== ""
/* payload.audience ==== Options.permittedFor
¿¿payload.IsValidNow === ok??
payload.NotBeforeChecker(Options.issuedBefore) === ok
payload.ExpirationTimeChecker(Options.expiresAt) === ok
*******************************************************************************/
