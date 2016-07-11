
# Installation

*md2slides* is a command line program run from a shell like Bash. If you download the repository a compiled version is in the dist directory. The compiled binary matching your computer type and operating system can be copied to a bin directory in your PATH.

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows (amd64) and Rapsberry Pi (both ARM6 and ARM7)

## Mac OS X

1. Download md2slides-binary-release.zip from [github.com/caltechlibrary/md2slides/releases/lastest](https://github.com/caltechlibrary/md2slides/releases/lastest)
2. Open a new finder window, find and unzip **mdslides-binary-release.zip**
3. Look in the unziped folder and find *dist/maxosx-amd64/md2slides*
4. Drag (or copy) *md2slides* to a "bin" directory in your path
5. Open and "Terminal" and run `md2slides -h` to confirm your success

## Windows

1. Download md2slides-binary-release.zip from [github.com/caltechlibrary/md2slides/releases/lastest](https://github.com/caltechlibrary/md2slides/releases/lastest)
2. Open the file manager find and unzip **mdslides-binary-release.zip**
3. Look in the unziped folder and find *dist/windows/md2slides.exe*
4. Drag (or copy) the *md2slides.exe* to a "bin" directory in your path
5. Open Bash and and run `md2slides -h` to confirm your success

## Linux

1. Download md2slides-binary-release.zip from [github.com/caltechlibrary/md2slides/releases/lastest](https://github.com/caltechlibrary/md2slides/releases/lastest)
2. find and unzip **mdslides-binary-release.zip**
3. In the unziped directory and find for *dist/linux-amd64/md2slides*
4. copy the *md2slides* to a "bin" directory (e.g. cp ~/Downloads/md2slides-master/dist/linux-amd64/md2slides ~/bin/)
5. From the shell prompt run `md2slides -h` to confirm your success

## Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Download md2slides-binary-release.zip from [github.com/caltechlibrary/md2slides/releases/lastest](https://github.com/caltechlibrary/md2slides/releases/lastest)
2. find and unzip **mdslides-binary-release.zip**
3. In the unziped directory and find for *dist/raspberrypi-arm7/md2slides*
4. copy the *md2slides* to a "bin" directory (e.g. cp ~/Downloads/md2slides-master/dist/raspberrypi-arm7/md2slides ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
5. From the shell prompt run `md2slides -h` to confirm your success

