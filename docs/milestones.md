# ASCII-Art Milestones (Audit Checklist)

## Milestone 1: Core CLI + Validation + Pipeline
- ✅ CLI argument parsing with flags: `--color`, `--align`, `--output`, `--s/--substr`.
- ✅ Validation for argument count and ASCII-only input.
- ✅ End-to-end flow: parse input → load font → build map → render banner.
- ✅ Proper handling of positional arguments and flag separation.
Proof: `main.go`, `src/cli.go`, `src/cli_test.go`, `src/fonts.go`, `src/fontMap.go`, `src/render.go`

## Milestone 2: Font Support
- ✅ Support for `standard`, `shadow`, `thinkertoy`.
- ✅ Embedded font files with internal reader selection.
- ✅ Rune-to-glyph mapping for printable ASCII (8-line height).
Proof: `src/fonts.go`, `src/fonts/standard.txt`, `src/fonts/shadow.txt`, `src/fonts/thinkertoy.txt`, `src/fontMap.go`, `src/fontMap_test.go`

## Milestone 3: Newline Handling and Rendering Behavior
- ✅ Escaped newline handling (`\\n`) in input.
- ✅ Correct behavior for empty input.
- ✅ Correct behavior for newline-only input.
- ✅ Correct behavior for internal/double newlines.
- ✅ Correct behavior for trailing newline case.
Proof: `main.go`, `src/render.go`, `src/render_test.go`, `testdata/*.golden`

## Milestone 4: Color Support
- ✅ Text colorization with multiple formats: color names, hex (`#RRGGBB`), RGB (`rgb(R, G, B)`).
- ✅ Substring-specific coloring via `--s` or `--substr` flags.
- ✅ Proper ANSI color code generation and terminal output.
Proof: `src/colorize.go`, `src/colorize_test.go`, `src/cutSubString.go`, `src/cutSubString_test.go`

## Milestone 5: Alignment and Terminal Integration
- ✅ Text alignment support: `--align` flag (left, center, right).
- ✅ Terminal width detection via ioctl.
- ✅ Adaptive layout based on terminal dimensions.
Proof: `src/terminal.go`, `src/render.go`, `tests/align_test.go`

## Milestone 6: File Output
- ✅ Output redirection to file via `--output` flag.
- ✅ Proper file writing and error handling.
Proof: `main.go`, `src/cli.go`

## Milestone 7: Tests and Regression Safety
- ✅ Unit tests for ASCII validation.
- ✅ Unit tests for banner rendering and colorization.
- ✅ Unit tests for alignment and substring matching.
- ✅ Integration tests for main execution and exit codes.
- ✅ Golden tests for critical output cases.
Proof: `tests/isValidAscii_test.go`, `tests/align_test.go`, `src/*_test.go`, `main_test.go`, `testdata/`

## Current Status
- ✅ `go test ./...` passes.
- ✅ All core features implemented and tested.
- ✅ Production-ready with comprehensive error handling.
