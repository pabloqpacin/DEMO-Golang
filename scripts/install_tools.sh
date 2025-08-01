#!/usr/bin/env bash


install_errcheck() {
    if ! command -v errcheck &> /dev/null; then
        go install github.com/kisielk/errcheck@latest
    fi
}

# install_staticcheck() {
#     if ! command -v staticcheck &> /dev/null; then
#         go install honnef.co/go/tools/cmd/staticcheck@latest
#     fi
# }


if true; then
    install_errcheck
    # install_staticcheck
fi

