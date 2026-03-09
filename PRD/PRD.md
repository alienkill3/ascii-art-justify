### 1. Problem Statement
The goal of this project is to create a program that receives a string as an argument and outputs a graphic representation of that string using ASCII characters.

### 2. User / Use Case
The primary use case is for a user to generate large, stylized ASCII text in a terminal environment. The user interacts with the program via a command-line interface, providing strings that may include letters, numbers, and special characters to be converted into banner-style art.

### 3. CLI Contract
The program is executed through the terminal using the Go command:
```
go run . [FLAGS] [SUBSTRING] [INPUT STRING] [BANNER FONT]
```
#### Flags:
*   `--color=<value>`: Apply color to output (named color, hex: RRGGBB, or RGB: rgb(R,G,B))
*   `--align=<value>`: Set text alignment (default: `left`, options: `left`, `right`, `center`, `justify`)
*   `-s <value>`: Apply color only to matching substring (requires `--color`)
*   `--output=<file>`: Write output to file instead of stdout

#### Positional Arguments:
*   **1 argument:** Input string (uses default font: `standard`)
*   **2 arguments:** Input string + font name (`standard`, `shadow`, `thinkertoy`)
*   **3 arguments (with `--color`):** Substring + input string + font name

The banner font defaults to `standard` if omitted.

### 4. Functional Requirements (Rules)
*   **Input Handling:** The program must handle strings containing numbers, uppercase and lowercase letters, spaces, special characters, and the newline character (`\n`). It also supports arguments starting with `-` (e.g., `-123`, `-randomstring`) as valid input text.
*   **Template Usage:** The program must use specific banner files provided as graphical templates: `shadow`, `standard`, and `thinkertoy`.
*   **Banner Format:**
    *   Each character in the banner files has a fixed height of **8 lines**.
    *   Characters in the banner files are separated by a newline character.
    *   The banner files are formatted correctly and do not require modification.
*   **Output logic:** The program must output the string graphically, respecting the 8-line height for each character and correctly interpreting `\n` in the input to move to a new block of ASCII art.
*   **Color Support:** The program must support coloring the entire output or a substring match using:
    *   Named colors (e.g., `--color=red`)
    *   Hex format (e.g., `--color=FF5733`)
    *   RGB format (e.g., `--color=rgb(255,87,51)`)
*   **Text Alignment:** The program must support left, right, center, and justify alignment via the `--align` flag.
*   **File Output:** The program must support writing the rendered ASCII art to a file using the `--output` flag.

### 5. Non-Goals (Out of Scope)
*   The use of any external packages; only **standard Go packages** are allowed.

### 6. Acceptance Criteria
#### 6.1 Audit Cases
*   Words and Numbers with trailing newline: Providing "1Hello 2There\n" must correctly render the numbers and the space between the words.
*   Brackets and internal newline: Providing "{Hello\nThere}" must correctly render the curly braces and the internal space.
*   Mixed Case and Spaces: Verify that strings like "HeLlO" and "Hello There" are rendered correctly, maintaining the distinct graphical differences between uppercase and lowercase letters and respecting space width

#### 6.2 Additional Golden Tests
*   **Empty String:** `go run . ""` should result in no output or a single empty line.
*   **Newline Only:** `go run . "\n"` should be handled correctly.
*   **Multiple Newlines:** Handling inputs like "Hello\n\nThere" to ensure the correct amount of vertical space is generated between the ASCII words.
*   **Banner Loading**
*   **Character Mapping**
*   **Full Input Validation**

### 7. Implementation Approach (High Level)
##### Decision Summary
The program will load the selected banner file or the default one and map the input string's characters to their corresponding 8-line blocks in the template. It will then print these blocks line-by-line (first line of all characters, then second line, etc.) to form the words.

##### Sketch
1.  **Parse Argument:** Receive the input string from the CLI.
2.  **Load Banner:** Read the chosen banner file (e.g., `standard`) into memory.
3.  **Mapping:** Create a data structure to map ASCII characters to their 8-line graphical representation found in the banner file.
4.  **Render:** Iterate through the input string, handle `\n` if present, and print the 8 lines for each character block.

### 8. Milestones
*   **Milestone 1:** Successfully parse and load a banner file into a Go data structure.
*   **Milestone 2:** Implement the logic to print a single character in ASCII art (8 lines high).
*   **Milestone 3:** Implement logic to print a full string without newlines.
*   **Milestone 4:** Finalize handling of `\n` and multiple spaces to match the required output.