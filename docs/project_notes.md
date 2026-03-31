# Project Notes

## How to release

```bash
git tag v0.0.X
git push origin v0.0.X
```

The release workflow builds Mac Apple Silicon and Mac Intel binaries and attaches them to a GitHub release automatically.

## Project structure

```
suppress_success/
├── src/                   # Go source code (module root)
├── build.sh               # builds to build/suppress_success
├── test.sh                # runs go test
├── docs/
│   ├── guide.md           # user guide
│   └── project_notes.md   # this file
└── .github/workflows/
    ├── build.yml           # CI: test + build on push/PR
    └── release.yml         # release: build binaries on v* tag
```



## Key decisions

- Binary name: `suppress_success` (config dir: `~/.config/suppress_success/`)
- Go source lives in `src/` (not repo root)
- macOS only (Apple Silicon + Intel)
- Distribution via GitHub releases (binary download), no package managers
- Single dependency: `github.com/pelletier/go-toml/v2`
