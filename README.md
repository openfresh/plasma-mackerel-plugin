plasma-mackerel-plugin
===========

[![Circle CI](https://circleci.com/gh/openfresh/plasma-mackerel-plugin.svg?style=shield&circle-token=1687597decf262a817315be69433f66b240862be)](https://circleci.com/gh/openfresh/plasma-mackerel-plugin)
[![Language](https://img.shields.io/badge/language-go-brightgreen.svg?style=flat)](https://golang.org/)
[![issues](https://img.shields.io/github/issues/openfresh/plasma-mackerel-plugin.svg?style=flat)](https://github.com/openfresh/plasma-mackerel-plugin/issues?state=open)
[![License: MIT](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)


This mackerel plugin provides [Plasma](https://github.com/openfresh/plasma) connection metrics.

# Usage

## Install

```bash
$ go get github.com/openfresh/plasma-mackerel-plugin
```

## setting

You should locate binary of plasma-mackerel-plugin to any directory.

```
[plugin.metrics.plasma]
command = "/path/to/plasma-mackerel-plugin -host=127.0.0.1:9999"
```

## Graph

![Graph](metrics.png)

License
===
See [LICENSE](LICENSE).

Copyright © CyberAgent, Inc. All Rights Reserved.
