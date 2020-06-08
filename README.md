
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

<b> Example syslog logging </b>

```
package main

import (
	apolog "github.com/axiomsamarth/apolog"
)

func main() {
	apolog.Info(`Log this at INFO level`)
}
```

<b> Example file based logging </b>
```
package main

import (
	"fmt"

	apolog "github.com/axiomsamarth/apolog"
)

func main() {

	var logFile := "./logs/logfile.log"
	var loginContext = &apolog.Apolog{Filepath: logFile}
	loginContext.SetContext()

	var data = []byte(`Log this!`)
	apolog.Info(data)
}
```

## Current Status
As of now:

- Different level of logging have been implemented.
- File based and syslog logging backends are supported.

Following features are being worked upon and would be pushed soon:

- Context based logging with common initialization of the logging path.
- Log filters.

## Contributing guidelines
We welcome community contribution to this project. Please visit the [Contributing.md](Contributing.md) 
