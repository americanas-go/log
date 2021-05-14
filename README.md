
log
=======

A simple, fast and consistent way for instantianting and using your favorite logging library in Golang. By a few changes in your config you can change the version or switch to a different library in seconds.

Installation
------------

	go get -u github.com/americanas-go/log


Supported libs
--------
* [Logrus](https://github.com/sirupsen/logrus) - Is a structured logger for Go (golang), completely API compatible with the standard library logger.
* [Zap](https://github.com/uber-go/zap) - Blazing fast, structured, leveled logging in Go.
* [Zerolog](https://github.com/rs/zerolog) - The zerolog package provides a fast and simple logger dedicated to JSON output.

Examples
--------
### Simple Logging Example

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	// start logrus
	loggerLogrus := logrus.NewLogger()
	loggerLogrus.Infof("Hi, I'm Logrus")
}
// Output: INFO[2021/05/14 16:22:17.634] Hi, I'm Logrus
```

Contributing
--------
Every help is always welcome. Fell free do throw us a pull request, we'll do our best to check it out as soon as possible. But before that, let us establish some guidelines:

1. This is an open source project so please do not add any proprietary code or infringe any copyright of any sort.
2. Avoid unnecessary dependencies or messing up go.mod file.
3. Be aware of golang coding style. Use a lint to help you out.
4.  Add tests to cover your contribution.
5. Use meaningful [messages](https://medium.com/@menuka/writing-meaningful-git-commit-messages-a62756b65c81) to your commits.
6. Use [pull requests](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests).
7. At last, but also important, be kind and polite with the community.

Any submitted issue which disrespect one or more guidelines above, will be discarded and closed.


<hr>

Released under the [MIT License](LICENSE.txt).