# Contributing To `gh-dash`

Thank you for investing your time in contributing to our project!

In this guide you will get an overview of the contribution workflow from opening an issue, creating a PR, reviewing, and merging the PR.

## The critical rules

- The most important rule: you must understand your code. If you can't explain what your changes do and how they interact with the greater system, do not contribute to this project.
- When you submit a PR, be willing to address comments and maintain the code. Do not submit drive-by changes without the willingness to iterate.
- AI-assisted contributions are welcome and receive the same quality review as any other contribution.

## AI Usage

Please read the [AI contribution policy](AI_POLICY.md). Disclosure is optional,
but understanding, validation, and ownership are mandatory.

## Quick Guide

### I Have an Idea for a Feature

First search the fork's issues to see whether the feature has already been
requested. Otherwise, [open a feature request](https://github.com/clintoncodewell/gh-dash/issues/new?template=feature_request.md).

### I've Implemented a Feature

- If there is an issue for the feature, open a focused pull request and link it.
- For non-trivial work without an issue, open one before implementation.
- Small, obvious fixes may go straight to a pull request, but include the context
  needed to understand and verify them.

### I Have a Question Which Is Neither a Bug Report nor a Feature Request

Open a question in this fork's issue tracker. For upstream usage support, see the
[upstream documentation](https://gh-dash.dev) and community links.

## Working on the Code

### Installing Required Tooling

Our project uses [Devbox](https://github.com/jetpack-io/devbox) to manage its development environment.

Using Devbox will get your dev environment up and running easily and make sure we're all using the same tools with the same versions.

- Clone this repo

```sh
git clone git@github.com:clintoncodewell/gh-dash.git && cd gh-dash
```

- Install `devbox`

```sh
curl -fsSL https://get.jetpack.io/devbox | bash
```

- Start the `devbox` shell and run the setup (will take a while on first time)

```sh
devbox shell
```

_This will create a shell where all required tools are installed._

- _(Optional)_ Set up `direnv` so `devbox shell` runs automatically
  - [direnv](https://www.jetify.com/devbox/docs/ide_configuration/direnv/) is a tool that allows setting unique environment variables per directory in your filesystem.
    - Install `direnv` with: `brew install direnv`
    - Add the following line at the end of the `~/.bashrc` file: `eval "$(direnv hook bash)"`
      - See [direnv's installation instructions](https://direnv.net/docs/hook.html) for other shells.
    - Enable `direnv` by running `direnv allow`
- _(Optional)_ Install the VSCode Extension
  - Follow [this guide](https://www.jetify.com/devbox/docs/ide_configuration/vscode/) to set up VSCode to automatically run `devbox shell`.

#### Troubleshooting

- delete the `.devbox` directory at the project's root

### Navigating the Codebase

To navigate our codebase with confidence, familiarize yourself with:

- [Bubbletea](https://github.com/charmbracelet/bubbletea) - the TUI framework we're using
- [The Elm architecture](https://guide.elm-lang.org/architecture/)
- [charmbracelet/glow](https://github.com/charmbracelet/glow) - for parsing and presenting Markdown

#### Code Structure

- `ui/` - this is the code that's responsible for rendering the different parts of the TUI
- `data/` - the code that fetches data from GitHub's GraphQL API
- `config/` - code to parse the user's `config.yml` file
- `utils/` - various utilities

### Debugging

- Write to the log by using Charm's `log` package
- Tail the log by running `task logs`
- Run `dash` in debug mode with `task debug` in another terminal window / pane

```go
import "charm.land/log/v2"

// more code...

log.Debug("some message", "someVariable", someVariable)
```

### Linting

CI runs this check, but for PRs from forks it may require maintainer approval before it starts. Running `task lint` locally saves a round-trip.

```sh
task lint
```

To auto-fix formatting issues (line length, imports, etc.):

```sh
task lint:fix
```

## Before you open a pull request

- Link the issue with `Closes #<issue>` or `Related: #<issue>`.
- Lead with the problem and concrete user impact.
- Describe the chosen solution briefly; let the diff carry implementation detail.
- Include relevant validation evidence and disclose any verification gap.
- Run `go test ./...`, `go vet ./...`, and `go build ./...`.
- Keep the pull request focused and review your own diff before submission.
- For UI changes, include a screenshot or short recording when practical.

### Running the Docs Locally

- Run the docs site by running `task docs`

* Go to `localhost:4321` to view them
