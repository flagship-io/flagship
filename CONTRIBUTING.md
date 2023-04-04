# Contributing
We encourage any form of contribution, whether that be issues, comments, or pull requests.

## Requirements
- Golang `1.18+`

### Custom Regex

In order to contribute your custom regex make sure to:

- Fork the project
- Read the [documentation](https://docs.developers.flagship.io/docs/codebase-analyzer) on the codebase analyzer integration.
- Create a predefined regex using the template `template-regexes.json` following the format origin_platform-regexes.json in case of a new feature flag provider, if you want to improve or fix an existing one just commit your changes.
- Adapt the analyze command with your changes (if necessary).
- Create a Pull Request.

If you have any questions while implementing a fix or feature, feel free to create an issue and ask us. We're happy to help!

