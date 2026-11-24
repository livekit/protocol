# Releasing protocol

This repository releases with [verwalter](https://github.com/livekit/verwalter).
The commands that are the same in every repository are in the
[verwalter cheat sheet](https://github.com/livekit/verwalter/blob/main/docs/cheat-sheet.md).
This file holds what is specific to this repository.

## Packages

The two packages always share a version.

| Package                       | Tag                           | Goes to                         |
| ----------------------------- | ----------------------------- | ------------------------------- |
| `github.com/livekit/protocol` | `v<version>`                  | Go proxy. No GitHub release.    |
| `@livekit/protocol`           | `@livekit/protocol@<version>` | npm, with a GitHub release      |

The version of the Go module is in `go.mod` and the root `package.json`.

## Workflows

| File                   | Starts when               | Does                                                                         |
| ---------------------- | ------------------------- | ---------------------------------------------------------------------------- |
| `changeset-check.yaml` | A PR is opened or updated | Fails a PR that changes a package without a changeset.                       |
| `release-pr.yaml`      | A push to `main`          | Keeps the release PR `release/main` current.                                 |
| `release.yaml`         | The release PR merges     | Builds the JS package, tags both packages, publishes to npm and the Go proxy. |
| `pre-release.yaml`     | Run by hand               | Opens an rc PR against the branch it runs on.                                |

Do not rename `release.yaml`: npm trusted publishing names the file.
`generate.yaml` does not run on `release/*` branches: verwalter
regenerates them on every push to `main`.

## Release

1. Merge the release PR.
2. Watch the run of `release.yaml`: `gh run list --workflow release.yaml --limit 1`.

A release is complete when:

- both tags are on the remote
- the tag `@livekit/protocol@<version>` has a public GitHub release
- npm and the Go proxy have the version
- the `verify` job is green. It checks the three points above.

## Pre-release

`gh workflow run pre-release.yaml --ref <branch>` opens an rc PR against
`<branch>`. The merge publishes the rc to npm with the dist-tag `next`.

## A release failed

Find the cause first. The paths by cause are in the
[recovery section](https://github.com/livekit/verwalter/blob/main/docs/release-workflows.md#recovery).

| Situation                                                       | Command                                                                    |
| --------------------------------------------------------------- | -------------------------------------------------------------------------- |
| A job failed for a transient reason                             | `gh run rerun <run-id> --failed`                                           |
| The workflow file needed a fix, and the fix is on `main`        | `gh workflow run release.yaml -f ref=<tag>`                                |
| The released commit is wrong, and npm does not have the version | `verwalter release --abandon <tag>`, then a changeset and the next version |

A re-run of failed jobs keeps the merge commit. Prefer it to a run by
hand. Set `ref` on a run by hand: an empty `ref` releases the head of
the branch. If no tag exists yet, set `ref` to the merge commit of the
release PR.

`publish` skips a version that a registry has. It fails when the files
of the package changed since the tag.

A version that npm or the Go proxy has cannot be replaced. For npm,
deprecate it with `npm deprecate`. For Go, add a `retract` directive to
`go.mod`. Then release the next version.
