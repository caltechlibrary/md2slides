#!/bin/bash

function softwareCheck() {
    for CMD in mkpage mkslides; do
        APP=$(which $CMD)
        if [ "$APP" = "" ]; then
            echo "Skipping, missing $CMD"
            exit 1;
        fi 
    done
}

function mkPage () {
    nav="$1"
    content="$2"
    html="$3"

    echo "Rendering $html"
    mkpage \
        "nav=$nav" \
        "content=$content" \
        page.tmpl > $html
}

function mkExample () {
    cd example
    mkslides presentation.md
    cd ..
}

echo "Checking for required software"
softwareCheck
echo "Generating website with mkpage"
mkPage nav.md index.md index.html
mkPage nav.md README.md readme.html
mkPage nav.md INSTALL.md install.html
mkPage nav.md "markdown:$(cat LICENSE)" license.html
echo "Generating example presentation.md"
mkExample

