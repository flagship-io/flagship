<p align="center">

<img  src="https://mk0abtastybwtpirqi5t.kinstacdn.com/wp-content/uploads/picture-solutions-persona-product-flagship.jpg"  width="211"  height="182"  alt="flagship"  />

</p>

<h3 align="center">Bring your features to life</h3>

A Tool to manage your Flagship resources built in [Go](https://go.dev/) using the library [Cobra](https://cobra.dev/).

[Website](https://flagship.io) | [Documentation](https://docs.developers.flagship.io/docs/flagship-command-line-interface) | [Installation Guide](https://docs.developers.flagship.io/docs/flagship-command-line-interface#download-and-install-the-flagship-cli) | [Twitter](https://twitter.com/feature_flags)

[![Apache2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/flagship-io/flagship)](https://goreportcard.com/report/github.com/flagship-io/flagship)
[![Go Reference](https://pkg.go.dev/badge/github.com/flagship-io/flagship.svg)](https://pkg.go.dev/github.com/flagship-io/flagship)
[![codecov](https://codecov.io/gh/flagship-io/flagship/branch/main/graph/badge.svg?token=nr9K5jvTlR)](https://codecov.io/gh/flagship-io/flagship)

## Overview

The Flagship CLI is a set of tools to create and manage your Flagship resources such as project, campaigns, users etc... You can use these tools to perform many common flagship platform task from the command line or through scripts and other automation.
For example, you can use the Flagship CLI to manage the following;
* Project and campaigns and other sub resources
* Users and environments
* Panic mode

## The Flagship CLI cheat sheet

For an introduction to the Flagship CLI, a list of commonly used commands, and a look at how these commands are structured, see the [Flagship cheat sheet](https://docs.developers.flagship.io/docs/cli-reference#commands).

## Download and install Flagship CLI

The Flagship CLI can be installed and deployed in your infrastructure either by downloading and running the binary, or pulling and running the docker image in your orchestration system.

### Using a binary

* Linux/Darwin

#### With wget
```bash
wget -qO- https://raw.githubusercontent.com/flagship-io/flagship/main/install.sh | bash
```

#### With curl
```bash
curl -sL https://raw.githubusercontent.com/flagship-io/flagship/main/install.sh | bash
```

#### With Homebrew
```bash
brew tap flagship-io/flagship
brew install flagship
```

* Other Operating Systems

Please download the binary from the [release page](https://github.com/flagship-io/flagship/releases)

### Using a Docker image
You can pull the latest docker image from docker hub: docker pull [flagshipio/cli](https://hub.docker.com/repository/docker/flagshipio/cli)

### Using Golang from source
You can pull the project from github and build it using golang latest stable version (+1.18): 

    git clone git@github.com:flagship-io/flagship.git
    cd flagship
    go build -o flagship

## Contributors

- Chadi Laoulaou [@Chadiii](https://github.com/chadiii)
- Guillaume Jacquart [@GuillaumeJacquart](https://github.com/guillaumejacquart)

## Contributing

Please read our [contributing guide](./CONTRIBUTING.md).

## Licence

[Apache License.](https://github.com/flagship-io/flagship/blob/main/LICENSE)

## About Flagship
​
<img src="https://www.flagship.io/wp-content/uploads/Flagship-horizontal-black-wake-AB.png" alt="drawing" width="150"/>
​
[Flagship by AB Tasty](https://www.flagship.io/) is a feature flagging platform for modern engineering and product teams. It eliminates the risks of future releases by separating code deployments from these releases :bulb: With Flagship, you have full control over the release process. You can:
​
- Switch features on or off through remote config.
- Automatically roll-out your features gradually to monitor performance and gather feedback from your most relevant users.
- Roll back any feature should any issues arise while testing in production.
- Segment users by granting access to a feature based on certain user attributes.
- Carry out A/B tests by easily assigning feature variations to groups of users.
​
<img src="https://www.flagship.io/wp-content/uploads/demo-setup.png" alt="drawing" width="600"/>
​
Flagship also allows you to choose whatever implementation method works for you from our many available SDKs or directly through a REST API. Additionally, our architecture is based on multi-cloud providers that offer high performance and highly-scalable managed services.
​
**To learn more:**
​
- [Solution overview](https://www.flagship.io/#showvideo) - A 5mn video demo :movie_camera:
- [Documentation](https://docs.developers.flagship.io/) - Our dev portal with guides, how tos, API and SDK references
- [Sign up for a free trial](https://www.flagship.io/sign-up/) - Create your free account
- [Guide to feature flagging](https://www.flagship.io/feature-flags/) - Everything you need to know about feature flag related use cases
- [Blog](https://www.flagship.io/blog/) - Additional resources about release management
