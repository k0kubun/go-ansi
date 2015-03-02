# go-ansi

Windows-portable ANSI escape sequence utility for Go language

## What's this?

This library converts ANSI escape sequences to Windows API calls on Windows environment.

![](http://i.gyazo.com/12ecc4e1b4387f5c56d3e6ae319ab6c4.png)
![](http://i.gyazo.com/c41072712ee05e28565ca92b416675e2.png)

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

## Notes

This is just a cursor and display supported version of [mattn/go-colorable](https://github.com/mattn/go-colorable).
I used almost the same implementation as it for coloring. Many thanks for [@mattn](https://github.com/mattn).

## License

MIT License
