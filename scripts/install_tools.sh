#!/usr/bin/env bash


install_pkgsite() {
    if ! command -v pkgsite &> /dev/null; then
        go install golang.org/x/pkgsite/cmd/pkgsite@latest
    fi

    # pkgsite -open .
    # open http://localhost:8080
    # open http://localhost:8080/testing # testing package
}


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
    install_pkgsite
    install_errcheck
    # install_staticcheck
fi

