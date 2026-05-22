# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project follows Semantic Versioning.

---

## [0.1.1] - 2026-05-21

[Release v0.1.1](https://github.com/abhi44mbd/nukeport/releases/tag/v0.1.1)

### Changed
- Renamed project from `killport` to `nukeport`
- Updated CLI command name to `nukeport`
- Updated Go module path
- Updated Homebrew formula and tap integration

### Fixed
- Fixed incorrect Homebrew release URL (404 issue)
- Fixed SHA256 mismatch in formula
- Fixed binary naming consistency across platforms

---

## [0.1.0] - 2026-05-21

[Release v0.1.0](https://github.com/abhi44mbd/nukeport/releases/tag/v0.1.0)

### Added
- Initial CLI implementation using Go + Cobra
- Find process by port
- Safe process termination with confirmation
- Force kill option (`--force`)
- Dry-run mode (`--dry-run`)
- Basic macOS and Linux support