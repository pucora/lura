// SPDX-License-Identifier: Apache-2.0

/*
Package core contains some basic constants and variables
*/
package core

import (
	"fmt"
	"runtime"
	"strings"
)

// VeloneticsHeaderName is the name of the custom Velonetics response header.
const VeloneticsHeaderName = "X-VELONETICS"

// VeloneticsVersion is the version of the build.
var VeloneticsVersion = "undefined"

// GoVersion is the version of the go compiler used at build time
var GoVersion = strings.TrimPrefix(runtime.Version(), "go")

// GlibcVersion is the version of the glibc used by CGO at build time
var GlibcVersion = "undefined"

// VeloneticsHeaderValue is the value of the custom Velonetics response header.
var VeloneticsHeaderValue = fmt.Sprintf("Version %s", VeloneticsVersion)

// VeloneticsUserAgent is the value of the user agent header sent to the backends.
var VeloneticsUserAgent = fmt.Sprintf("Velonetics Version %s", VeloneticsVersion)
