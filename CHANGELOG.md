# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/)
and this project adheres to [Semantic Versioning](https://semver.org/).

## [v0.2.0] - 2021-09-15

### Added

- feat: Add group.AddString support
- feat: Implement type and type alias support

### Changed

- function without `AddBody` will not generate empty body anymore.

## [v0.1.0] - 2021-09-08

### Added

- feat: Add interface support (#4)
- feat: Implement AppendFile
- feat: Field accept both string and Node as input
- feat: Add typed field support
- feat: Export logic statement for group
- feat: Add function call support
- feat: Allow all body accept interface instead
- feat: Add defer and function call support
- Add var decl support
- feat: Add test case for struct
- API Redesign

### Fixed

- fix: Struct should not have NamedLineComment
- fix: struct Field should not return nil
- fix: Insert a new line before closing struct
- fix: Insert new line after switch case

### Refactor

- Change return to accept input instead of function call
- refactor: Make comment as float just like go ast (#6)

## [v0.0.2] - 2021-09-03

### Added

- feat: Implement struct support
- feat: Implement named line comment support

### Fixed

- fix: Comment should be renamed to LineComment

## v0.0.1 - 2021-09-03

### Added

- Implement basic functions
- feat: Add omit wrap support

[v0.2.0]: https://github.com/Xuanwo/gg/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/Xuanwo/gg/compare/v0.0.2...v0.1.0
[v0.0.2]: https://github.com/Xuanwo/gg/compare/v0.0.1...v0.0.2
