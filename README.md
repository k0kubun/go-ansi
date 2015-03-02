# go-ansi

Windows-portable ANSI escape sequence utility for Go language

## What's this?

This library converts ANSI escape sequences to Windows API calls on Windows environment.  
You can easily use this feature by replacing `fmt` with `ansi`.

![](http://i.gyazo.com/12ecc4e1b4387f5c56d3e6ae319ab6c4.png)
![](http://i.gyazo.com/c41072712ee05e28565ca92b416675e2.png)

### Output redirection

Many coloring libraries for Go just use ANSI escape sequences, which don't work on Windows.

- [fatih/color](https://github.com/fatih/color)
- [mitchellh/colorstring](https://github.com/mitchellh/colorstring)

If you use go-ansi, you can use these libraries' nice APIs for Windows too.
Not only coloring, many ANSI escape sequences are available.

```go
color.Output = ansi.NewAnsiStdout()
color.Cyan("fatih/color")

colorstring.Fprintln(ansi.NewAnsiStdout(), "[green]mitchellh/colorstring")
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
