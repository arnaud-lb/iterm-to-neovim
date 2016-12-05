# iterm-to-neovim

Converts an iTerm2's itermcolors colorscheme to a list of `g:terminal_color_x` declarations.

Neovim uses the `g:terminal_color_x` variables in its embeded terminal.

## Usage

```
go install github.com/arnaud-lb/iterm-to-neovim
~/go/bin/iterm-to-neovim < some.itermcolors > /tmp/terminal_colors.vim
```
