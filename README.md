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

Every job starts with `step-security/harden-runner` in block mode, which allows network egress only to the hosts listed
under `allowed-endpoints` for that job and needs no secrets. When a job fails because a host was blocked, the job log
names it: add it to that job's list in this repository, keeping the list sorted. Libraries using this template get their
workflows synced from here, so a list edited in a library is overwritten by the next sync.

Generated files are recognised by Go's standard `// Code generated ... DO NOT EDIT.` header.

## Configuration

`build-config.yaml` sets the minimum test coverage:

```yaml
requiredCoverage: 80
```

The Go version comes from `go.mod`.
