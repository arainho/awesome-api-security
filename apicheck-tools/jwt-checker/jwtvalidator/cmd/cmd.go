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

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"github.com/BBVA/apicheck/tools/jwt-checker/jwtvalidator/validations"
)

type CommandError struct {
	Code   int
	errors []string
}

func (c CommandError) Error() string {
	b := bytes.Buffer{}

	if c.Code > 0 {
		b.WriteString("Errors: \n")
	}

	for _, e := range c.errors {
		b.WriteString(e)
		b.WriteString("\n")
	}

	return b.String()
}

type ReqResp struct {
	Meta struct {
		Host       string            `json:"host"`
		Schema     string            `json:"schema"`
		ToolParams map[string]string `json:"jwtchk"`
	} `json:"_meta"`
	Request struct {
		Path    string            `json:"path"`
		Method  string            `json:"method"`
		Headers map[string]string `json:"headers"`
		Body    string            `json:"body"`
	} `json:"request"`
	Response struct {
		Status  int               `json:"status"`
		Reason  string            `json:"reason"`
		Headers map[string]string `json:"headers"`
		Body    string            `json:"body"`
	} `json:"response"`
}

type ProgramOptions struct {
	HelpRequested    bool
	VersionRequested bool
	Arguments        []string
	FlagsProcessed   int
	StdIn            io.Reader
	jsDoc            []byte
	validations.Options
}

const (
	VERSION = "1.0.0"
)

func Run(opt ProgramOptions) CommandError {
	// If help requested or no options or arguments provided show usage and exit
	if len(opt.Arguments) == 0 && opt.FlagsProcessed == 0 {
		return CommandError{2, []string{}}
	}
	if opt.HelpRequested {
		return CommandError{2, []string{}}
	}

	// Show version info
	if opt.VersionRequested {
		return CommandError{0, []string{fmt.Sprintf("Version %s", VERSION)}}
	}

	// If unsecure tokens are not allowed, a secret must be present
	if !opt.AllowUnsignedTokens && opt.Secret == "" {
		return CommandError{1, []string{fmt.Sprintf("Secret for validating signature must be provided")}}
	}

	if opt.IssuedBefore != "" {
		var err error
		if opt.TimeIssuedBefore, err = time.Parse(time.RFC3339, opt.IssuedBefore); err != nil {
			return CommandError{1, []string{fmt.Sprintf("Invalid date format for %s. Should be YYYYMMDDThh:mm:ss", "issuedBefore")}}
		}
	}

	if opt.ExpiresAt != "" {
		var err error
		if opt.TimeExpiresAt, err = time.Parse(time.RFC3339, opt.ExpiresAt); err != nil {
			return CommandError{1, []string{fmt.Sprintf("Invalid date format for %s. Should be YYYYMMDDThh:mm:ss", "expiresAt")}}
		}
	}

	switch len(opt.Arguments) {
	case 1:
		var (
			tokenArg string
			err      error
		)
		if opt.Arguments[0] == "-" {
			if tokenArg, err = readTokenFromReqResp(&opt); err != nil {
				return CommandError{1, []string{err.Error()}}
			}
		} else {
			tokenArg = opt.Arguments[0]
		}

		if errList := validations.ValidateToken(([]byte)(tokenArg), opt.Options); errList != nil {
			er := CommandError{Code: 1}
			for _, e := range errList {
				er.errors = append(er.errors, fmt.Sprint(e))
			}
			return er
		}

		// If called with a json document attach it to be printed
		if len(opt.jsDoc) > 0 {
			return CommandError{0, []string{string(opt.jsDoc)}}
		}
		return CommandError{0, []string{}}
	default: /*Only one token must be provided or - to read from standard input*/
		return CommandError{1, []string{fmt.Sprint("One and only one token must be provided")}}
	}
}

func readTokenFromReqResp(opt *ProgramOptions) (string, error) {
	rr := ReqResp{}

	if jsonStr, err := ioutil.ReadAll(opt.StdIn); err != nil || len(jsonStr) == 0 {
		return "", fmt.Errorf("Error reading standard Input (bytes readed: %d; error: %v)", len(jsonStr), err)
	} else if err := json.Unmarshal(jsonStr, &rr); err != nil {
		return "", fmt.Errorf("JSON Error: %v", err)
	} else {
		opt.jsDoc = jsonStr
	}

	if val, ok := rr.Request.Headers["Authorization"]; ok {
		if !strings.HasPrefix(val, "Bearer ") {
			return "", fmt.Errorf("No Authorization bearer was found")
		}
		// Remove the "Bearer " prefix and return the rest of the string
		return strings.TrimPrefix(val, "Bearer "), nil
	} else {
		return "", fmt.Errorf("No Authorization header was found")
	}
}
