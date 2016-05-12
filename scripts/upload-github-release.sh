#!/bin/bash

VERSION=$(< ./VERSION)
FILES=./releases/*
RELEASE_TYPE=$1

function github_release() {
    $GOPATH/bin/github-release "$@"
}

echo "--- Creating GitHub release v$VERSION"

github_release release \
    --user "itmh" \
    --repo "commit-msg" \
    --tag "$VERSION" \
    --name "$VERSION" \
    --description "README.md" \
    $RELEASE_TYPE

if [ $? -eq 0 ]; then
    echo "Done."
fi

echo "--- Uploading files for release v$VERSION"

for fullfile in $FILES
do
    filename=$(basename "$fullfile")
    echo $filename

    github_release upload \
        --user "itmh" \
        --repo "commit-msg" \
        --tag "$VERSION" \
        --name "$filename" \
        --file "$fullfile"

    if [ $? -eq 0 ]; then
        echo "Done."
    fi
done
