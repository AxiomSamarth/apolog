
![Apolog Build](https://github.com/AxiomSamarth/apolog/workflows/workflow/badge.svg)

# Apolog
Apolog is a Golang package for logging. Apolog supports different levels of logging and also supports multiple logging backends like syslog and file logging. 

## Prequisites
- Go=1.13.8 darwin/amd64

## Installation and Usage
Download the repository and navigate to `apolog` folder. Run the below commands - 

```
export GOPATH=$(pwd)
cd src
go install apolog
```

Import the `apolog` package from `github.com/axiomsamarth/apolog` and employ the exported functions `Info`, `Error`, `Warn`, `Debug`, `Trace`, `Fatal` that implements the logging levels accordingly.

```
package main

import (
	apolog "github.com/axiomsamarth/apolog"
)

func main() {
	apolog.Info(`Log this at INFO level`)
}
```

## Current Status
As of now different level of logging have been implemented with backend being syslog. Following features are being worked upon and would be pushed soon:

- File based logging backend
- Context based logging with common initialization of the logging path
- Log filters like `ACCEPT`, `DENY` and `NEUTRAL`.

## Contribute
Thank you for that thought! Contribution guidelines will be soon put up. No PRs will be merged till then. However, feel free to raise the issues in the meantime.

## Contact
For anything related to the development of the apolog Golang logging package, this repository, concerns, queries and so, drop an email. My email id is - deyagondsamarth@gmail.com
