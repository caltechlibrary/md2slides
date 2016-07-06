
# Installation

*md2slides* is a command line program run from a shell like Bash. If you download the repository a compiled version is in the dist directory. The compiled binary matching your computer type and operating system can be copied to a bin directory in your PATH.

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows (amd64) and Rapsberry Pi (both ARM6 and ARM7)

## Mac OS X

1. Go to [github.com/caltechlibrary/md2slides](https://github.com/caltechlibrary/md2slides)
2. Click on the green "Clone or Download" button on the right side of the page
3. A panel will open, click on "Download Zip"
4. Open a finder window downloaded file and unzip it (e.g. md2slides-master.zip)
5. Look in the Unziped folder and find dist/maxosx-amd64/md2slides
6. Drag (or copy) the *md2slides* to a "bin" directory in your path
7. Open and "Terminal" and run `md2slides -h`

## Windows

1. Go to [github.com/caltechlibrary/md2slides](https://github.com/caltechlibrary/md2slides)
2. Click on the green "Clone or Download" button on the right side of the page
3. A panel will open, click on "Download Zip"
4. Open the file manager find the downloaded file and unzip it (e.g. md2slides-master.zip)
5. Look in the Unziped folder and find dist/windows/md2slides.exe
6. Drag (or copy) the *md2slides.exe* to a "bin" directory in your path
7. Open Bash and and run `md2slides -h`

## Linux

1. Go to [github.com/caltechlibrary/md2slides](https://github.com/caltechlibrary/md2slides)
2. Click on the green "Clone or Download" button on the right side of the page
3. A panel will open, click on "Download Zip"
4. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/md2slides-master.zip)
5. In the Unziped directory and find for dist/linux-amd64/md2slides
6. copy the *md2slides* to a "bin" directory (e.g. cp ~/Downloads/md2slides-master/dist/linux-amd64/md2slides ~/bin/)
7. From the shell prompt run `md2slides -h`

## Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Go to [github.com/caltechlibrary/md2slides](https://github.com/caltechlibrary/md2slides)
2. Click on the green "Clone or Download" button on the right side of the page
3. A panel will open, click on "Download Zip"
4. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/md2slides-master.zip)
5. In the Unziped directory and find for dist/raspberrypi-arm7/md2slides
6. copy the *md2slides* to a "bin" directory (e.g. cp ~/Downloads/md2slides-master/dist/raspberrypi-arm7/md2slides ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
7. From the shell prompt run `md2slides -h`

