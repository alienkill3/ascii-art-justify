## Example Usage

### Basic String
```bash
go run . "Hello"
```
Renders "Hello" in the standard ASCII font.

### With Shadow Font
```bash
go run . "Hello" shadow
```
Renders "Hello" using the shadow font.

### Color - Entire Output
```bash
go run . --color=red "Hello"
```
Renders "Hello" in red color.

### Color - HEX Format
```bash
go run . --color=#FF5733 "Hello"
```
Renders "Hello" in the HEX color `#FF5733`.

### Color - RGB Format
```bash
go run . --color='rgb(255, 87, 51)' "Hello"
```
Renders "Hello" using a custom RGB color.

### Color - Substring Match
```bash
go run . --color=blue --substr=ll "Hello"
go run . --color=blue ll "hello"
```
Renders "Hello" with only the "ll" substring colored blue.

### Alignment
```bash
go run . --align=right "Hello"
```
Renders "Hello" aligned to the right.

### Alignment - Justify
```bash
go run . --align=justify "Hello There"
```
Renders the text with justified spacing between words.

### Combined: Color + Alignment + Font
```bash
go run . --color=green --align=center "Hello" thinkertoy
```
Renders "Hello" in green, centered, using thinkertoy font.

### With Newlines
```bash
go run . "Hello\nWorld"
```
Renders "Hello" and "World" on separate blocks with proper vertical spacing.

### With Multiple Newlines
```bash
go run . "Hello\n\nWorld"
```
Renders an empty line block between "Hello" and "World".

### Numbers and Special Characters
```bash
go run . "-123"
```
Renders "-123" as ASCII art (negative numbers are treated as input text, not flags).

### File Output
```bash
go run . --output=result.txt "Hello"
cat result.txt
```
Writes the rendered ASCII art to `result.txt` instead of displaying it.

### Complex Example
```bash
go run . --align=center --color='rgb(255,100,50)' Pro "Project" shadow
```
Renders "Project" with only the substring "Pro" colored in a custom orange RGB color, centered, using shadow font

## Error Examples

### Substring Without Color
```bash
go run . -s=ll "Hello"
```
Returns an error because substring coloring requires `--color`.

### Invalid Color Name
```bash
go run . --color=rad "Hello"
```
Returns an error for unsupported color name.

### Invalid Font
```bash
go run . "Hello" shadw
```
Returns an error and shows valid fonts: `standard`, `shadow`, `thinkertoy`.

## Examples from Audit Cases
*   **Numbers and Newlines:** `go run . "1Hello 2There\n"` correctly renders numbers, spaces, and the trailing newline.
*   **Internal Newlines:** `go run . "Hello\n\nThere"` ensures the correct vertical space is generated between the ASCII words.
*   **Empty Input:** `go run . ""` results in no output or a single empty line.
*   **Negative Numbers:** `go run . "-123"` renders "-123" as ASCII art without errors.
