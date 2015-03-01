# go-ansi

Windows-portable ANSI escape sequence utility for Go language

## What's this?

This library converts ANSI escape sequences to Windows API calls on Windows environment.

```go
gore> fmt.Println("\x1b[31mRED\x1b[0m")
\x1b[31mRED\x1b[0m
gore> ansi.Println("\x1b[31mRED\x1b[0m")
RED
```

### Coloring

Many coloring libraries for Go just use ANSI escape sequences, which don't work on Windows.
If you use go-ansi, you can use these libraries' nice APIs for Windows too.

- [fatih/color](https://github.com/fatih/color)
- [mitchellh/colorstring](https://github.com/mitchellh/colorstring)

```go
import (
  "github.com/fatih/color"
  "github.com/k0kubun/go-ansi"
  "github.com/mitchellh/colorstring"
)

func main() {
  color.Output = ansi.NewAnsiStdout()
  c := color.New(color.FgCyan, color.Bold)
  c.Println("fatih/color")

  colorstring.Fprintln(ansi.NewAnsiStdout(), "[red]mitchellh/colorstring")
}
```

### Cursor

You can control cursor in your terminal.

### Display

You can easily control your terminal display. You can easily provide unix-like
shell functionarities for display, such as C-k or C-l.

## License

MIT License
