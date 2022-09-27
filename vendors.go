// Copyright 2019 Adam Shannon
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package cidetails

import (
	"os"
	"strconv"
	"strings"
)

type vendor struct {
	// case-insensitive names to describe a CI service
	// names should be from the current to oldest
	names  []string
	envVar func() bool
	pr     func() bool
}

func n(names ...string) []string {
	return names
}

func nonempty(envVar string) func() bool {
	return func() bool {
		return strings.TrimSpace(os.Getenv(envVar)) != ""
	}
}

func match(envVar, value string) func() bool {
	return func() bool {
		return strings.TrimSpace(os.Getenv(envVar)) == value
	}
}

func parseBool(envVar string) func() bool {
	return func() bool {
		v, _ := strconv.ParseBool(strings.TrimSpace(os.Getenv(envVar)))
		return v
	}
}

// func any(envVars ...string) bool { // TODO(adam): which CI provider was this for?
// 	for i := range envVars {
// 		if nonempty(envVars[i])() {
// 			return true
// 		}
// 	}
// 	return false
// }

var vendors = []vendor{
	{
		names:  n("AppVeyor"),
		envVar: nonempty("APPVEYOR"),
		pr:     nonempty("APPVEYOR_PULL_REQUEST_NUMBER"),
	},
	{
		names:  n("AWS CodeBuild"),
		envVar: nonempty("CODEBUILD_BUILD_ARN"),
		pr:     func() bool { return false },
	},
	{
		names:  n("Azure Pipelines"),
		envVar: nonempty("SYSTEM_TEAMFOUNDATIONCOLLECTIONURI"),
		pr:     nonempty("SYSTEM_PULLREQUEST_PULLREQUESTID"),
	},
	{
		names:  n("Bamboo"),
		envVar: nonempty("bamboo_planKey"),
		pr:     func() bool { return false },
	},
	{
		names:  n("Bitbucket Pipelines"),
		envVar: nonempty("BITBUCKET_COMMIT"),
		pr:     nonempty("BITBUCKET_PR_ID"),
	},
	{
		names:  n("Bitrise"),
		envVar: nonempty("BITRISE_IO"),
		pr:     nonempty("BITRISE_PULL_REQUEST"),
	},
	{
		names:  n("Buddy"),
		envVar: nonempty("BUDDY_WORKSPACE_ID"),
		pr:     nonempty("BUDDY_EXECUTION_PULL_REQUEST_ID"),
	},
	{
		names:  n("Buildkite"),
		envVar: nonempty("BUILDKITE"),
		pr:     parseBool("BUILDKITE_PULL_REQUEST"),
	},
	{
		names:  n("CircleCI"),
		envVar: nonempty("CIRCLECI"),
		pr:     nonempty("CIRCLE_PULL_REQUEST"),
	},
	{
		names:  n("Cirrus CI"),
		envVar: nonempty("CIRRUS_CI"),
		pr:     nonempty("CIRRUS_PR"),
	},
	{
		names:  n("Codeship"),
		envVar: nonempty("CI_NAME"),
		pr:     func() bool { return false },
	},
	{
		names:  n("Drone"),
		envVar: nonempty("DRONE"),
		pr:     match("DRONE_BUILD_EVENT", "pull_request"),
	},
	{
		names:  n("Woodpecker CI"),
		envVar: nonempty("CI"),
		pr:     match("CI_BUILD_EVENT", "pull_request"),
	},
	{
		names:  n("dsari"),
		envVar: nonempty("DSARI"),
		pr:     func() bool { return false },
	},
	{
		names:  n("GitLab CI"),
		envVar: nonempty("GITLAB_CI"),
		pr:     func() bool { return false },
	},
	{
		names:  n("GoCD"),
		envVar: nonempty("GO_PIPELINE_LABEL"),
		pr:     func() bool { return false },
	},
	{
		names:  n("Hudson"),
		envVar: nonempty("HUDSON_URL"),
		pr:     func() bool { return false },
	},
	{
		names:  n("Jenkins"),
		envVar: nonempty("JENKINS_URL"),
		pr:     nonempty("ghprbPullId"),
	},
	{
		names:  n("Magnum CI"),
		envVar: nonempty("MAGNUM"),
		pr:     func() bool { return false },
	},
	{
		names:  n("Netlify CI"),
		envVar: nonempty("NETLIFY_BUILD_BASE"),
		pr:     parseBool("PULL_REQUEST"),
	},
	{
		names:  n("Nevercode"),
		envVar: nonempty("NEVERCODE"),
		pr:     parseBool("NEVERCODE_PULL_REQUEST"),
	},
	{
		names:  n("Sail CI"),
		envVar: nonempty("SAILCI"),
		pr:     nonempty("SAIL_PULL_REQUEST_NUMBER"),
	},
	{
		names:  n("Semaphore"),
		envVar: nonempty("SEMAPHORE"),
		pr:     nonempty("PULL_REQUEST_NUMBER"),
	},
	{
		names:  n("Shippable"),
		envVar: nonempty("SHIPPABLE"),
		pr:     parseBool("IS_PULL_REQUEST"),
	},
	{
		names:  n("Solano CI"),
		envVar: nonempty("TDDIUM"),
		pr:     nonempty("TDDIUM_PR_ID"),
	},
	{
		names:  n("Strider CD"),
		envVar: nonempty("STRIDER"),
		pr:     func() bool { return false },
	},
	{
		names:  n("TeamCity"),
		envVar: nonempty("TEAMCITY_VERSION"),
		pr:     func() bool { return false },
	},
	{
		names:  n("Travis CI", "TravisCI"),
		envVar: nonempty("TRAVIS"),
		pr:     parseBool("TRAVIS_PULL_REQUEST"),
	},
}
