# ASCII-Art Generator

## Description
This project is a command-line tool written in **Go** that converts a given string into a large, stylized **ASCII graphical representation**. It maps input characters to 8-line high blocks from specific banner templates to create banner-style art in the terminal. The tool also supports **color output**, **alignment**, and **file output** through flags.

## Requirements
*   Golang installed (recommended version `1.25.6`).
*   A terminal that supports ANSI color output.

## Usage
The program is executed through the terminal with one of the following forms:

```bash
go run . [FLAGS] 'INPUT_STRING' [BANNER_FONT]
go run . [FLAGS] [SUBSTRING] 'INPUT_STRING' [BANNER_FONT]
```

### Keep In Mind
*   Put your input string inside quotes, especially when it contains spaces or special characters.
*   For complex inputs that may look like flags, use the standard UNIX separator `--`.
*   `BANNER_FONT` is optional and defaults to `standard`.

### Flags
*   `--color=<value>`: Apply color to the output.
    * Named colors: `red`, `green`, `blue`, `yellow`, `cyan`, `purple`, `orange`, `pink`
    * Hex format: `#RRGGBB` (example: `--color=#FF5733`)
    * RGB format: `rgb(R, G, B)` (example: `--color='rgb(255, 87, 51)'`)
*   `--align=<value>`: Set text alignment.
    * Options: `left` (default), `right`, `center`, `justify`
*   `-s <substring>`: Color only matching substring (requires `--color`).
*   `--output=<file>`: Write output to a file instead of stdout.

### Positional Arguments
*   **`INPUT_STRING`**: Text to render as ASCII art.
    * Supports letters, numbers, spaces, special characters, and newline sequences (`\n`).
*   **`BANNER_FONT`**: Optional font.
    * Options: `standard`, `shadow`, `thinkertoy`
*   **`SUBSTRING`**: Optional positional substring when color is enabled.
    * With color enabled, the parser accepts `[SUBSTRING] [INPUT_STRING] [BANNER_FONT]`.
    * If substring is omitted, the whole input is colorized.

## Features
*   **Standard Library Only**: Built entirely using standard Go packages.
*   **Template Support**: Includes three styles: `standard`, `shadow`, and `thinkertoy`.
*   **Character Mapping**: Maps each character to a fixed **8-line height** graphical representation.
*   **Color Support**: Supports named colors, HEX, and RGB.
*   **Selective Coloring**: Supports substring-based coloring with `-s` or positional substring.
*   **Text Alignment**: Supports `left`, `right`, `center`, and `justify`.
*   **File Output**: Can write rendered output directly to a file via `--output`.

## Quick Examples
```bash
go run . "Hello There"
go run . "Hello" shadow
go run . --color=red "Hello"
go run . --color=blue ll "Hello" shadow
go run . --align=right "Hello"
go run . --output=result.txt "Hello"
go run . "Hello\nWorld"
```

## Implementation Approach
*   Pipeline approach. For more details refer to [Architecture](/docs/architecture.md) and [Diagram](/docs/diagram.md).

## More Examples
*   See [Examples](/docs/examples.md).

