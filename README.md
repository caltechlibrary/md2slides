
# mkslides

Convert a Markdown file into a sequence of HTML5 slides.

+ Use Markdown to write your presentation in one file
+ Separate slides by "--" and a new line (e.g. \n versus \r\n)
+ Apply the simple default template or use your own
+ Control Layout and display with HTML5 and CSS

Here's an example of a three slide presentation

```
    Welcome to [mkslides](../)
    by R. S. Doiel, <rsdoiel@caltech.edu>

    --

    # mkslides

    mkslides can generate multiple HTML5 pages from
    one markdown file.  It splits the markdown file
    on each "--" 

    --

    Thank you

    Hope you enjoy [mkslides](https://github.com/caltechlbrary/mkslides)

```

If you save this [file](example/presentation.md) as presentation.md it would
generate the following webpages

+ [00-presentation01.html](example/00-presentation.html)
+ [01-presentation02.html](example/01-presentation.html)
+ [02-presentation03.html](example/02-presentation.html)


## Releases and cross compilation

The script [mk-release.sh](./mk-release.sh) cross compiles *mkslides* for Windows, Max OS X, Linux (amd64) and Raspberry Pi (Raspbian/ARM6 and ARM7).
It places all the resulting executable programs in the *dist* folders. See [INSTALL.md](install.html) for instructions.

## windows issues

*mkslides* has very limited Windows testing done.  *mkslides.exe* 
presumes the Unix style new line only and not the old DOS/Windows CR/LF type endings.


