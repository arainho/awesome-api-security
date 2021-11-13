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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BBVA/apicheck/tools/jwt-checker/jwtvalidator/cmd"
)

func main() {
	programOptions := ParseOptions()

	if e := cmd.Run(programOptions); e.Code == 2 { // Print help message
		flag.Usage()
		os.Exit(e.Code)
	} else if e.Code == 1 { // Print error messages
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(e.Code)
	} else {
		fmt.Fprintln(os.Stdin, e.Error())
		os.Exit(e.Code)
	}
}

func ParseOptions() cmd.ProgramOptions {
	opt := cmd.ProgramOptions{}
	flag.BoolVar(&opt.HelpRequested, "h", false, "show help message and exit")
	flag.BoolVar(&opt.VersionRequested, "V", false, "show version info and exit")

	flag.BoolVar(&opt.AllowUnsignedTokens, "unsig", false, "don't raise an error if token is unsigned (algorithm \"none\")")
	flag.Var(&opt.AllowedSignAlgs, "allowAlg", "raise an error if token has different signing algorithm. Provide several options if you want to allow more than one algorithm")
	flag.StringVar(&opt.Secret, "secret", "", "use this secret to validate the sign. Base64 encoded (mandatory when validating signed tokens)")
	flag.StringVar(&opt.Issuer, "issuer", "", "raise an error if token has different issuer")
	flag.StringVar(&opt.Subject, "subject", "", "raise an error if token has different subject")
	flag.StringVar(&opt.PermittedFor, "audience", "", "raise an error if token has different audience")
	flag.StringVar(&opt.IssuedBefore, "notBefore", "", "raise an error if the not before date is priot to this. Format: YYYY-MM-DDThh:mm:ss")
	flag.StringVar(&opt.ExpiresAt, "expiresAt", "", "raise an error if the expiration date is after this. Format: YYYY-MM-DDThh:mm:ss")

	flag.Parse()

	opt.Arguments = flag.Args()
	opt.FlagsProcessed = flag.NFlag()
	opt.StdIn = os.Stdin

	return opt
}
