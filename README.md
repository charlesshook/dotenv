# dotenv

Go (Golang) does not autoload environment variables like other languages so I created a package to do just that!

## Install
```
go get -u github.com/charlesshook/dotenv
```

## Usage

Loading .env 
```
package main

import (
	"github.com/charlesshook/dotenv"
)

func main() {
	dotenv.Load()
}
```

Loading files other than .env
```
package main

import (
	"github.com/charlesshook/dotenv"
)

func main() {
	dotenv.Load("dev.env", "vars.txt")
}
```

A environment variable can look like the following:
```
PORT_TEST=8080
# This is a comment
DB_USER_TEST=root
DB_PASS_TEST="mysecretpassword" # This is a inline comment
APP_NAME_TEST='My App'
DEBUG_TEST=true
```

Single line comments and inline comments are both supported.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Some things that would be nice to have:
- [ ] Support for multiline values
- [ ] Tests to ensure parsing is correct
- [ ] Ensure that there are no edge cases that are not handled

