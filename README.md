# cidetails

[![GoDoc](https://godoc.org/github.com/adamdecaf/cidetails?status.svg)](https://godoc.org/github.com/adamdecaf/cidetails)
[![AppVeyor](https://ci.appveyor.com/api/projects/status/github/adamdecaf/cidetails?branch=master&svg=true)](https://ci.appveyor.com/project/adamdecaf/cidetails)
[![CircleCI](https://circleci.com/gh/adamdecaf/cidetails.svg?style=svg)](https://circleci.com/gh/adamdecaf/cidetails)
[![Travis CI](https://travis-ci.com/adamdecaf/cidetails.svg?branch=master)](https://travis-ci.com/adamdecaf/cidetails)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/adamdecaf/cidetails/master/LICENSE)

Get information about the current Continuous Integration (CI) environment. Inspired by [watson/ci-info](https://github.com/watson/ci-info).

## Installation

```
$ go get github.com/adamdecaf/cidetails@master
```

## Usage

```go
package main

import (
	"testing"
	"github.com/adamdecaf/cidetails"
)

func TestFoo(t *testing.T) {
	if cidetails.In("TravisCI") {
		// provider specific code
	} else {
		t.Logf("skipping CI setup in %s", cidetails.Name())
	}
}
```

## Supported CI tools

Officially supported CI servers:

- [AppVeyor](http://www.appveyor.com)
- [AWS CodeBuild](https://aws.amazon.com/codebuild/)
- [Azure Pipelines](https://azure.microsoft.com/en-us/services/devops/pipelines/)
- [Bamboo](https://www.atlassian.com/software/bamboo) by Atlassian
- [Bitbucket Pipelines](https://bitbucket.org/product/features/pipelines)
- [Bitrise](https://www.bitrise.io/)
- [Buddy](https://buddy.works/)
- [Buildkite](https://buildkite.com)
- [CircleCI](http://circleci.com)
- [Cirrus CI](https://cirrus-ci.org)
- [Codeship](https://codeship.com)
- [Drone](https://drone.io)
- [dsari](https://github.com/rfinnie/dsari)
- [GitLab CI](https://about.gitlab.com/gitlab-ci/)
- [Github Actions](https://github.com/features/actions)
- [GoCD](https://www.go.cd/)
- [Hudson](http://hudson-ci.org)
- [Jenkins CI](https://jenkins-ci.org)
- [Magnum CI](https://magnum-ci.com)
- [Netlify CI](https://www.netlify.com/)
- [Nevercode](http://nevercode.io/)
- [Sail CI](https://sail.ci/)
- [Semaphore](https://semaphoreci.com)
- [Shippable](https://www.shippable.com/)
- [Solano CI](https://www.solanolabs.com/)
- [Strider CD](https://strider-cd.github.io/)
- [TaskCluster](http://docs.taskcluster.net)
- [TeamCity](https://www.jetbrains.com/teamcity/) by JetBrains
- [Travis CI](http://travis-ci.org)
- [Woodpecker CI](https://woodpecker-ci.org)

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
