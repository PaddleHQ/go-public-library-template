# go-public-library-template

A template for building new public Go libraries.

The workflows here run on GitHub-hosted runners and need no secrets, so they work in any public repository. This
repository is public and runs the same workflows on its own pull requests, so a change that would break them in a public
library breaks here first.

## Checks

Each workflow provides one of the status checks the organisation requires on pull requests:

| Check | What it does |
|---|---|
| Ensure SemVer Label is added | Requires exactly one of `norelease`, `release:patch`, `release:minor`, `release:major` |
| Ensure code passes linting rules | Runs `golangci-lint` |
| Run code validation checks | Builds, verifies and tidies modules, checks `gofmt`, runs `go vet` and `govulncheck` |
| Run tests with coverage and race detector | Runs `go test -race`, counting coverage across all packages and leaving out generated files |
| Run tests and validate coverage | Fails if total coverage is below `requiredCoverage` in `build-config.yaml` (100 if unset) |

On merge to `main`, the release label decides the next version, and a GitHub release is created for it.

Every job starts with `step-security/harden-runner` in audit mode, which records the job's network egress and needs no secrets.

Generated files are recognised by Go's standard `// Code generated ... DO NOT EDIT.` header.

## Configuration

`build-config.yaml` sets the minimum test coverage:

```yaml
requiredCoverage: 80
```

The Go version comes from `go.mod`.
