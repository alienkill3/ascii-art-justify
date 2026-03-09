package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"ascii-art/src"
)

func TestMainProcess(t *testing.T) {

	tests := []struct {
		name     string
		args     []string
		wantExit int
	}{
		{
			name:     "empty arguments",
			args:     []string{},
			wantExit: 0,
		},
		{
			name:     "too many arguments",
			args:     []string{"sub string", "input string", "something"},
			wantExit: 1,
		},
		{
			name:     "substr without color flag",
			args:     []string{"l", "hello\nworld"},
			wantExit: 1,
		},
		{
			name:     "not ascii",
			args:     []string{"ψάρι"},
			wantExit: 1,
		},
		{
			name:     "not ascii",
			args:     []string{"--color=red", "ψάρι"},
			wantExit: 1,
		},
		{
			name:     "RGB color format with substring and font",
			args:     []string{"--color", "rgb(255, 0, 0)", "e", "hello", "shadow"},
			wantExit: 0,
		},
		{
			name:     "wrong flag with substring and font",
			args:     []string{"---color", "rgb(255, 0, 0)", "e", "hello", "shadow"},
			wantExit: 1,
		},
		{
			name:     "wrong RGB color format with substring and font",
			args:     []string{"--color", "rgb(256, 0, 0)", "e", "hello", "shadow"},
			wantExit: 1,
		},
		{
			name:     "wrong RGB color format with substring and font",
			args:     []string{"--color", "rgb(256, 0, 0", "e", "hello", "shadow"},
			wantExit: 1,
		},
		{
			name:     "wrong RGB color format with substring and font",
			args:     []string{"--color", "rgb(56, 10, 10, 30)", "e", "hello", "shadow"},
			wantExit: 1,
		},
		{
			name:     "#hex color format with substring and font",
			args:     []string{"--color", "#00ff00", "e", "hello", "shadow"},
			wantExit: 0,
		},
		{
			name:     "wrong length #hex color format with substring and font",
			args:     []string{"--color", "#00ff00f", "e", "hello", "shadow"},
			wantExit: 1,
		},
		{
			name:     "wrong #hex color format with substring and font",
			args:     []string{"--color", "#00gg00", "e", "hello", "shadow"},
			wantExit: 1,
		},
		{
			name:     "named color format with substring and font",
			args:     []string{"--color", "red", "e", "hello", "shadow"},
			wantExit: 0,
		},
		{
			name:     "wrong named color format with substring and font",
			args:     []string{"--color", "rad", "e", "hello", "shadow"},
			wantExit: 1,
		},
		{
			name:     "named color format with substring and wrong font",
			args:     []string{"--color", "red", "s", "sa", "shadw"},
			wantExit: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This trick tells the test to run the 'main' function
			// within this specific test binary.
			cmd := exec.Command(os.Args[0], "-test.run=TestActualMain")

			// For tests with flags, use environment variable
			if len(tt.args) > 0 && strings.HasPrefix(tt.args[0], "--") {
				// Join args with a separator for environment variable
				cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
				cmd.Env = append(cmd.Env, "ASCII_ART_TEST_ARGS="+strings.Join(tt.args, "\x1f"))
			} else {
				// For regular positional args, pass as command args
				cmd.Args = append(cmd.Args, tt.args...)
				cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
			}

			out, err := cmd.CombinedOutput()
			if tt.wantExit == 0 && err != nil {
				t.Fatalf("expected success, got error: %v, output: %s", err, out)
			}

			if tt.wantExit != 0 {
				if err == nil {
					t.Fatalf("expected exit code %d, but process succeeded", tt.wantExit)
				}
				if e, ok := err.(*exec.ExitError); ok {
					if e.ExitCode() != tt.wantExit {
						t.Errorf("got exit code %d, want %d. Output: %s", e.ExitCode(), tt.wantExit, out)
					}
				}
			}
		})
	}
}

func runMainAndGetStdout(t *testing.T, args ...string) []byte {
	t.Helper()

	cmd := exec.Command(os.Args[0], "-test.run=TestActualMain")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")

	out, err := cmd.Output()
	if err != nil {
		// If it exited non-zero, include stderr in message via CombinedOutput
		combined, _ := cmd.CombinedOutput()
		t.Fatalf("process failed: %v\noutput:\n%s", err, combined)
	}
	return out
}

func runMainAndGetStdoutWithFlags(t *testing.T, args ...string) []byte {
	t.Helper()

	cmd := exec.Command(os.Args[0], "-test.run=TestActualMain")
	cmd.Env = append(
		os.Environ(),
		"GO_WANT_HELPER_PROCESS=1",
		"ASCII_ART_TEST_ARGS="+strings.Join(args, "\x1f"),
	)

	out, err := cmd.Output()
	if err != nil {
		combined, _ := cmd.CombinedOutput()
		t.Fatalf("process failed: %v\noutput:\n%s", err, combined)
	}
	return out
}

func TestOutput_NewlineOnly(t *testing.T) {
	golden, err := os.ReadFile("testdata/newline.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "\\n")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_HelloDoubleNewline(t *testing.T) {
	golden, err := os.ReadFile("testdata/hello_double_newline.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "Hello\\n\\nThere")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_EmptyInput(t *testing.T) {
	golden, err := os.ReadFile("testdata/empty.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%q\n--- want (golden) ---\n%q", got, golden)
	}
}

func TestOutput_HelloTrailingNewline(t *testing.T) {
	golden, err := os.ReadFile("testdata/hello_trailing_newline.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "Hello\\n")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_HelloThere(t *testing.T) {
	golden, err := os.ReadFile("testdata/hello_there.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "Hello There")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_HelloShadow(t *testing.T) {
	golden, err := os.ReadFile("testdata/hello_shadow.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "hello", "shadow")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_HelloThinkertoy(t *testing.T) {
	golden, err := os.ReadFile("testdata/hello_thinkertoy.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "hello", "thinkertoy")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_BracketsInternalNewline(t *testing.T) {
	golden, err := os.ReadFile("testdata/brackets_internal_newline.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "{Hello\\nThere}")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_NumbersTrailingNewline(t *testing.T) {
	golden, err := os.ReadFile("testdata/numbers_trailing_newline.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "1Hello 2There\\n")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_RGBColorSubstrHelloShadow(t *testing.T) {
	golden, err := os.ReadFile("testdata/rgb_color_substr_e_hello_shadow.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdoutWithFlags(t, "--color=rgb'(255, 0, 0)'", "e", "hello", "shadow")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_ColorFullHelloRed(t *testing.T) {
	golden, err := os.ReadFile("testdata/color_full_hello_red.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdoutWithFlags(t, "--color=red", "hello", "standard")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_ColorSubstrNotFound(t *testing.T) {
	golden, err := os.ReadFile("testdata/color_substr_not_found.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdoutWithFlags(t, "--color=red", "z", "hello", "standard")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_WriteToFile(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "banner.txt")
	expected := runMainAndGetStdout(t, "hello", "standard")

	cmd := exec.Command(os.Args[0], "-test.run=TestActualMain")
	cmd.Env = append(
		os.Environ(),
		"GO_WANT_HELPER_PROCESS=1",
		"ASCII_ART_TEST_ARGS=--output="+filePath+"\x1fhello\x1fstandard",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("expected success, got error: %v, output: %s", err, out)
	}

	gotFile, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}
	if !bytes.Equal(gotFile, expected) {
		t.Fatalf("output file mismatch\n--- got ---\n%s\n--- want ---\n%s", gotFile, expected)
	}
}

func TestOutput_InvalidOutputFormat_ShowsUsage(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestActualMain")
	cmd.Env = append(
		os.Environ(),
		"GO_WANT_HELPER_PROCESS=1",
		"ASCII_ART_TEST_ARGS=--invalid\x1fhello\x1fstandard\x1fextra",
	)

	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("expected non-zero exit, got success. output: %s", out)
	}
	if !strings.Contains(string(out), src.OutputUsage) {
		t.Fatalf("expected usage message, got: %s", out)
	}
}

func TestOutput_WriteToDirectory_ShowsWriteError(t *testing.T) {
	dirPath := t.TempDir()
	cmd := exec.Command(os.Args[0], "-test.run=TestActualMain")
	cmd.Env = append(
		os.Environ(),
		"GO_WANT_HELPER_PROCESS=1",
		"ASCII_ART_TEST_ARGS=--output="+dirPath+"\x1fhello\x1fstandard",
	)

	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("expected non-zero exit, got success. output: %s", out)
	}
	if !strings.Contains(string(out), "Error writing output file") {
		t.Fatalf("expected write error message, got: %s", out)
	}
}

// Helper process that runs the real main() for integration tests.
func TestActualMain(t *testing.T) {
	// Run only when invoked as helper process
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	// Helper allows passing app args that begin with '-' via env.
	if rawArgs := os.Getenv("ASCII_ART_TEST_ARGS"); rawArgs != "" {
		args := strings.Split(rawArgs, "\x1f")
		os.Args = append([]string{os.Args[0]}, args...)
	} else {
		// For positional args, they are after -test.run=TestActualMain
		for i, arg := range os.Args {
			if arg == "-test.run=TestActualMain" {
				os.Args = append([]string{os.Args[0]}, os.Args[i+1:]...)
				break
			}
		}
	}

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Execute CLI logic
	main()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	out, _ := io.ReadAll(r)
	r.Close()

	// Write to stdout for the parent to capture
	os.Stdout.Write(out)

	// Exit immediately to avoid printing "PASS" to stdout
	os.Exit(0)
}
