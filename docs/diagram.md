```
┌──────────────────────────────────┐
│  Command-Line Input              │
│  (ParseCLIArgs)                  │
│  ├─ --color flag (optional)      │
│  ├─ --align flag (optional)      │
│  ├─ -s/--substr flag (optional)  │
│  ├─ --output flag (optional)     │
│  └─ Positional Args:             │
│     ├─ Input String (REQUIRED)   │
│     ├─ Font (standard/shadow/    │
│     │     thinkertoy)            │
│     └─ Substring (if --color)    │
└──────────────┬───────────────────┘
               │
               ▼
┌──────────────────────────────────────┐
│  Color Parsing Stage (if --color)    │
│  (ParseNamedColor/ParseHexColor/     │
│   ParseRGBColor)                     │
│  → RGB integer value (or nil)        │
└──────────────┬───────────────────────┘
               │
               ▼
┌──────────────────────────────────────┐
│  Font Processing Stage               │
│  BuildFontMap(font file)             │
│  → map[rune][]string                 │
│  (fully buffered in memory)          │
└──────────────┬───────────────────────┘
               │
               ▼
┌──────────────────────────────────────┐
│  Banner Rendering Stage              │
│  RenderInputString(input,            │
│                   fontMap,           │
│                   options,           │
│                   terminal_width)    │
│  → rendered ASCII art (uncolored)    │
└──────────────┬───────────────────────┘
               │
               ▼
┌──────────────────────────────────────┐
│  Argument Separation                 │
│  separateArgsAndFlags()              │
│  → Splits flags from positional args │
│  (handles args like "-123")          │
└──────────────┬───────────────────────┘
│  (Colorize with Substring Match) │
│  ├─ Apply color to full output   │
│  └─ Or only to substring matches │
└────────────┬─────────────────────┘
             │
             ▼
┌──────────────────────────────────┐
│  Output Stage                    │
│  ├─ stdout (default)             │
│  └─ File (-o/--output option)    │
└──────────────────────────────────┘
```
