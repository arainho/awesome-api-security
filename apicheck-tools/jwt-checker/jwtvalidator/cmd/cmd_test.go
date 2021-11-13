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

package cmd_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/BBVA/apicheck/tools/jwt-checker/jwtvalidator/cmd"
	"github.com/BBVA/apicheck/tools/jwt-checker/jwtvalidator/validations"
)

func TestReturnOKWhenNoOptionsOrArgumentsGiven(t *testing.T) {
	o := cmd.ProgramOptions{}

	if e := cmd.Run(o); 2 != e.Code || "" == e.Error() {
		t.Errorf("Run raised an unexpected error. Expected: (2, %q); got: (%d, %q)", "", e.Code, e.Error())
	}
}

func TestReturnOKWhenHelpRequested(t *testing.T) {
	o := cmd.ProgramOptions{HelpRequested: true, FlagsProcessed: 1}

	if e := cmd.Run(o); 2 != e.Code || "" == e.Error() {
		t.Errorf("Run raised an unexpected error. Expected: (2, %q); got: (%d, %q)", "", e.Code, e.Error())
	}
}

func TestReturnOKWhenVersionRequested(t *testing.T) {
	expectMsg := "Version 1.0.0\n"
	o := cmd.ProgramOptions{VersionRequested: true, FlagsProcessed: 1}

	if e := cmd.Run(o); 0 != e.Code || expectMsg != e.Error() {
		t.Errorf("Run raised an unexpected error. Expected: (0, %q); got: (%d, %q)", expectMsg, e.Code, e.Error())
	}
}

func TestKOWhenNotAllowUnsignedAndNoSecretGiven(t *testing.T) {
	expectMsg := "Errors: \nSecret for validating signature must be provided\n"
	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false}, FlagsProcessed: 1}

	if e := cmd.Run(o); 1 != e.Code || expectMsg != e.Error() {
		t.Errorf("Run didn't raise the expected error. Expected: (1, %q); got: (%d, %q)", expectMsg, e.Code, e.Error())
	}
}

func TestKOWhenNotAllowUnsignedSecretAndInvalidNotBeforeGiven(t *testing.T) {
	expectMsg := "Errors: \nInvalid date format for issuedBefore. Should be YYYYMMDDThh:mm:ss\n"
	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false, Secret: "FOO_secret", IssuedBefore: "20200301T12:30"}, FlagsProcessed: 2}

	if e := cmd.Run(o); 1 != e.Code || expectMsg != e.Error() {
		t.Errorf("Run didn't raise the expected error. Expected: (1, %q); got: (%d, %q)", expectMsg, e.Code, e.Error())
	}
}

func TestKOWhenNotAllowUnsignedSecretAndInvalidExpiresAtGiven(t *testing.T) {
	expectMsg := "Errors: \nInvalid date format for expiresAt. Should be YYYYMMDDThh:mm:ss\n"
	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false, Secret: "FOO_secret", ExpiresAt: "20200301T12:30"}, FlagsProcessed: 2}

	if e := cmd.Run(o); 1 != e.Code || expectMsg != e.Error() {
		t.Errorf("Run didn't raise the expected error. Expected: (1, %q); got: (%d, %q)", expectMsg, e.Code, e.Error())
	}
}

func TestKOWhenNotAllowUnsignedAndSecretGivenButTokenNotGiven(t *testing.T) {
	expectMsg := "Errors: \nOne and only one token must be provided\n"
	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false, Secret: "FOO_secret"}, FlagsProcessed: 2}

	if e := cmd.Run(o); 1 != e.Code || expectMsg != e.Error() {
		t.Errorf("Run didn't raise the expected error. Expected: (1, %q); got: (%d, %q)", expectMsg, e.Code, e.Error())
	}
}

func TestKOWhenNotAllowUnsignedAndSecretGivenButMoreThanATokenGiven(t *testing.T) {
	expectMsg := "Errors: \nOne and only one token must be provided\n"
	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false, Secret: "FOO_secret"}, FlagsProcessed: 2, Arguments: []string{"FOO_BAR_TOKEN", "FOO_BAR_TOKEN"}}

	if e := cmd.Run(o); 1 != e.Code || expectMsg != e.Error() {
		t.Errorf("Run didn't raise the expected error. Expected: (1, %q); got: (%d, %q)", expectMsg, e.Code, e.Error())
	}
}

func TestKOWhenNotAllowUnsignedAndSecretGivenButInvalidTokenGiven(t *testing.T) {
	expectMsg := "Errors: \njwt: token format is not valid\n"
	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false, Secret: "FOO_secret"}, FlagsProcessed: 2, Arguments: []string{"FOO_BAR_TOKEN"}}

	if e := cmd.Run(o); 1 != e.Code || expectMsg != e.Error() {
		t.Errorf("Run didn't raise the expected error. Expected: (1, %q); got: (%d, %q)", expectMsg, e.Code, e.Error())
	}
}

func TestKOWhenInputExpectedAndNotGiven(t *testing.T) {
	expectMsg := "Errors: \nError reading standard Input"
	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false, Secret: "FOO_secret"}, FlagsProcessed: 2, Arguments: []string{"-"}, StdIn: strings.NewReader("")}

	if e := cmd.Run(o); 1 != e.Code || !strings.HasPrefix(e.Error(), expectMsg) {
		t.Errorf("Run didn't raise the expected error. Expected: (1, %q); got: (%d, %q)", expectMsg+"...", e.Code, e.Error())
	}
}

func TestKOWhenInvalidJSONInputGiven(t *testing.T) {
	expectMsg := "Errors: \nJSON Error: "
	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false, Secret: "FOO_secret"}, FlagsProcessed: 2, Arguments: []string{"-"}, StdIn: strings.NewReader(`{"":"}`)}

	if e := cmd.Run(o); 1 != e.Code || !strings.HasPrefix(e.Error(), expectMsg) {
		t.Errorf("Run didn't raise the expected error. Expected: (1, %q); got: (%d, %q)", expectMsg+"...", e.Code, e.Error())
	}
}

func TestKOWhenNoAuthorizationHeaderExists(t *testing.T) {
	expectMsg := "Errors: \nNo Authorization header was found"
	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false, Secret: "FOO_secret"}, FlagsProcessed: 2, Arguments: []string{"-"}, StdIn: strings.NewReader(`{"request":{"headers":{"Accept":"*/*"}}}`)}

	if e := cmd.Run(o); 1 != e.Code || !strings.HasPrefix(e.Error(), expectMsg) {
		t.Errorf("Run didn't raise the expected error. Expected: (1, %q); got: (%d, %q)", expectMsg+" ...", e.Code, e.Error())
	}
}

func TestKOWhenAuthorizationHeaderDontHaveBearer(t *testing.T) {
	expectMsg := "Errors: \nNo Authorization bearer was found\n"
	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false, Secret: "FOO_secret"}, FlagsProcessed: 2, Arguments: []string{"-"}, StdIn: strings.NewReader(`{"request":{"headers":{"Authorization":"foo-auth"}}}`)}

	if e := cmd.Run(o); 1 != e.Code || expectMsg != e.Error() {
		t.Errorf("Run didn't raise the expected error. Expected: (1, %q); got: (%d, %q)", expectMsg, e.Code, e.Error())
	}
}

func TestKOWhenInvalidBearerTokenGiven(t *testing.T) {
	expectMsg := "Errors: \njwt: token format is not valid\n"
	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false, Secret: "FOO_secret"}, FlagsProcessed: 2, Arguments: []string{"-"}, StdIn: strings.NewReader(`{"request":{"headers":{"Authorization":"Bearer FOO_BAR_TOKEN"}}}`)}

	if e := cmd.Run(o); 1 != e.Code || expectMsg != e.Error() {
		t.Errorf("Run didn't raise the expected error. Expected: (1, %q); got: (%d, %q)", expectMsg, e.Code, e.Error())
	}
}

func TestOKWhenValidBearerTokenGiven(t *testing.T) {
	tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsib25lIiwib3RoZXIiXSwiaXNzIjoiYmJ2YS1pYW0iLCJzdWIiOiJzdWJqZWN0LWlkIiwiZ3JvdXBzIjoiZ3JwdXAxLGdyb3VwMiIsImVtYWlsIjoidW5lbWFpbEBjb21wYW55LmNvbSJ9.bJTts4nlMgELdw2WSfzc8CN-hwFcnxs-1YBtxKB1Yn0"
	s := "bXlTZWNyZXRQYXNzd29yZG15U2VjcmV0UGFzc3dvcmQK"
	jsonDoc := fmt.Sprintf(`{"request":{"headers":{"Authorization":"Bearer %s"}}}`, tk)
	expectMsg := jsonDoc + "\n"

	o := cmd.ProgramOptions{Options: validations.Options{AllowUnsignedTokens: false, Secret: s}, FlagsProcessed: 2, Arguments: []string{"-"}, StdIn: strings.NewReader(jsonDoc)}

	if e := cmd.Run(o); 0 != e.Code || expectMsg != e.Error() {
		t.Errorf("Run didn't raise the expected error. Expected: (0, %q); got: (%d, %q)", expectMsg, e.Code, e.Error())
	}
}
