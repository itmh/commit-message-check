#!/bin/bash

VERSION=$(< ./VERSION)
FILES=./releases/*
RELEASE_TYPE=$1

function github_release() {
    $GOPATH/bin/github-release "$@"
}

echo "--- Creating GitHub release v$VERSION"

github_release release \
    --user "leominov" \
    --repo "commit-message-check" \
    --tag "$VERSION" \
    --name "$VERSION" \
    --description "" \
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
        --user "leominov" \
        --repo "commit-message-check" \
        --tag "$VERSION" \
        --name "$filename" \
        --file "$fullfile"

    if [ $? -eq 0 ]; then
        echo "Done."
    fi
done
