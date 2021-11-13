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

package validations_test

import (
	"strings"
	"testing"

	"github.com/BBVA/apicheck/tools/jwt-checker/jwtvalidator/validations"
)

var (
	badSecret                   = "bXlTZWNyZXRQYXNzd29yZA"
	secret                      = "bXlTZWNyZXRQYXNzd29yZG15U2VjcmV0UGFzc3dvcmQK"
	noClaimsSignedToken         = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJncm91cHMiOiJncnB1cDEsZ3JvdXAyIiwiZW1haWwiOiJ1bmVtYWlsQGNvbXBhbnkuY29tIn0.UEyAkQeusnw18O0tNnT3UL1VweE1uITHPUIrL3MYFyc")
	unsignedToken               = []byte("eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJpc3MiOiJiYnZhLWlhbSIsInN1YiI6InN1YmplY3QtaWQifQ.UEyAkQeusnw18O0tNnT3UL1VweE1uITHPUIrL3MYFyc")
	unsignedBadTypeToken        = []byte("eyJhbGciOiJub25lIiwidHlwIjoiSldTIiwiY3R5IjoiSldTIn0.eyJpc3MiOiJiYnZhLWlhbSIsInN1YiI6InN1YmplY3QtaWQifQ.UEyAkQeusnw18O0tNnT3UL1VweE1uITHPUIrL3MYFyc")
	unsignedBadContenttypeToken = []byte("eyJhbGciOiJub25lIiwidHlwIjoiSldUIiwiY3R5IjoiSldTIn0.eyJpc3MiOiJiYnZhLWlhbSIsInN1YiI6InN1YmplY3QtaWQifQ.UEyAkQeusnw18O0tNnT3UL1VweE1uITHPUIrL3MYFyc")
	issuerClaimsToken           = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJiYnZhLWlhbSIsImdyb3VwcyI6ImdycHVwMSxncm91cDIiLCJlbWFpbCI6InVuZW1haWxAY29tcGFueS5jb20ifQ.01YlrnAuZPL4Ok02B2viGl2EAm2A4p5mJrV7_PZzcAU")
	issuerSubjectClaimsToken    = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJiYnZhLWlhbSIsInN1YiI6InN1YmplY3QtaWQiLCJncm91cHMiOiJncnB1cDEsZ3JvdXAyIiwiZW1haWwiOiJ1bmVtYWlsQGNvbXBhbnkuY29tIn0.Omfue9tqAjTOKRR0dxgA0wXQr6ilE0bhkO__8zw1IVc")
	fullStdClaimsToken          = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsib25lIiwib3RoZXIiXSwiaXNzIjoiYmJ2YS1pYW0iLCJzdWIiOiJzdWJqZWN0LWlkIiwiZ3JvdXBzIjoiZ3JwdXAxLGdyb3VwMiIsImVtYWlsIjoidW5lbWFpbEBjb21wYW55LmNvbSJ9.bJTts4nlMgELdw2WSfzc8CN-hwFcnxs-1YBtxKB1Yn0")
)

func TestKOWhenInvalidTokenGiven(t *testing.T) {
	o := validations.Options{}

	if errs := validations.ValidateToken([]byte("FOO.BAR.TOKEN"), o); len(errs) == 0 || errs[0].Error() != "jwt: token format is not valid" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "jwt: token format is not valid", errs[0].Error())
	}
}

func TestKOWhenBadTokenTypeGiven(t *testing.T) {
	o := validations.Options{}

	if errs := validations.ValidateToken(unsignedBadTypeToken, o); len(errs) == 0 || errs[0].Error() != "JWT type (JWS) not supported" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "JWT type (JWS) not supported", errs[0].Error())
	}
}

func TestKOWhenBadTokenContenttypeGiven(t *testing.T) {
	o := validations.Options{}

	if errs := validations.ValidateToken(unsignedBadContenttypeToken, o); len(errs) == 0 || errs[0].Error() != "JWT content type (JWS) not supported" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "JWT content type (JWS) not supported", errs[0].Error())
	}
}

func TestKOWhenUnsignedTokenGiven(t *testing.T) {
	o := validations.Options{Secret: badSecret}

	if errs := validations.ValidateToken(unsignedToken, o); len(errs) == 0 || errs[0].Error() != "JWT is unsigned" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "JWT is unsigned", errs[0].Error())
	}
}

func TestKOWhenNotmatchingAlgorithmTokenGiven(t *testing.T) {
	o := validations.Options{Secret: badSecret, AllowedSignAlgs: []string{"RS256"}}

	if errs := validations.ValidateToken(noClaimsSignedToken, o); len(errs) == 0 || errs[0].Error() != "JWT signing algorithm (HS256) is not allowed" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "JWT signing algorithm (HS256) is not allowed", errs[0].Error())
	}
}

func TestKOWhenValidTokenAndEmptySecretGiven(t *testing.T) {
	o := validations.Options{Secret: "", AllowedSignAlgs: []string{"HS256"}}

	if errs := validations.ValidateToken(noClaimsSignedToken, o); len(errs) == 0 || errs[0].Error() != "No secret provided" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "No secret provided", errs[0].Error())
	}
}

func TestKOWhenValidTokenAndIncorrectEncodedSecretGiven(t *testing.T) {
	o := validations.Options{Secret: "FOO_SECRET?FOO_SECRET?FOO_SECRET"}

	if errs := validations.ValidateToken(noClaimsSignedToken, o); len(errs) == 0 || !strings.HasPrefix(errs[0].Error(), "Error base64decoding") {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "Error base64decoding...", errs[0].Error())
	}
}

func TestKOWhenValidTokenAndWrongSecretGiven(t *testing.T) {
	o := validations.Options{Secret: badSecret}

	if errs := validations.ValidateToken(noClaimsSignedToken, o); len(errs) == 0 || errs[0].Error() != "jwt: signature is not valid" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "jwt: signature is not valid", errs[0].Error())
	}
}

func TestKOWhenNoIssuerClaim(t *testing.T) {
	o := validations.Options{Secret: secret}

	if errs := validations.ValidateToken(noClaimsSignedToken, o); len(errs) == 0 || errs[0].Error() != "No issuer claim found" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "No issuer claim found", errs[0].Error())
	}
}

func TestKOWhenNotmachingIssuerClaim(t *testing.T) {
	o := validations.Options{Secret: secret, Issuer: "oher-issuer"}

	if errs := validations.ValidateToken(issuerClaimsToken, o); len(errs) == 0 || errs[0].Error() != "Issuer claim doesn't match. Expected: oher-issuer, got: bbva-iam" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "Issuer claim doesn't match. Expected: oher-issuer, got: bbva-iam", errs[0].Error())
	}
}

func TestKOWhenNoSubjectClaim(t *testing.T) {
	o := validations.Options{Secret: secret}

	if errs := validations.ValidateToken(issuerClaimsToken, o); len(errs) == 0 || errs[0].Error() != "No subject claim found" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "No subject claim found", errs[0].Error())
	}
}

func TestKOWhenNotmachingSubjectClaim(t *testing.T) {
	o := validations.Options{Secret: secret, Subject: "other-subject"}

	if errs := validations.ValidateToken(issuerSubjectClaimsToken, o); len(errs) == 0 || errs[0].Error() != "Subject claim doesn't match. Expected: other-subject, got: subject-id" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "Subject claim doesn't match. Expected: other-subject, got: subject-id", errs[0].Error())
	}
}

func TestKOWhenNoAudienceClaim(t *testing.T) {
	o := validations.Options{Secret: secret}

	if errs := validations.ValidateToken(issuerSubjectClaimsToken, o); len(errs) == 0 || errs[0].Error() != "No audience claim found" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "No audience claim found", errs[0].Error())
	}
}

func TestKOWhenNotmachingAudienceClaim(t *testing.T) {
	o := validations.Options{Secret: secret, PermittedFor: "another"}

	if errs := validations.ValidateToken(fullStdClaimsToken, o); len(errs) == 0 || errs[0].Error() != "Audience claim doesn't match. Expected: another, got: one,other" {
		t.Errorf("ValidateToken didn't return the expected error. Expected: %q; got %q", "Audience claim doesn't match. Expected: another, got: one,other", errs[0].Error())
	}
}

func TestOKWhenCompleteTokenGiven(t *testing.T) {
	o := validations.Options{Secret: secret, Issuer: "bbva-iam", Subject: "subject-id", PermittedFor: "one"}

	if errs := validations.ValidateToken(fullStdClaimsToken, o); len(errs) > 0 {
		t.Errorf("ValidateToken returns unexpected error: %v", errs)
	}
}
