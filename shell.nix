{ pkgs ? import <nixpkgs> { } }:
with pkgs;
let
  inherit (pkgs) stdenv;
  apple_sdk = darwin.apple_sdk.frameworks;
in mkShell {
  name = "doko";
  buildInputs = [
    buildkit
    docker
    docker-compose_2
    go_1_18
    go-tools
    go-migrate
    nodejs-17_x
    libiconv
    git
    git-secret
    gitAndTools.delta
    binutils-unwrapped
    hostname
    inetutils
    openssh
    pkg-config
    rsync
    fasd
    fzf
    gnupg
    jq
    less
    lesspipe
    lsof
    lzma
    jetbrains-mono
    yamllint
  ] ++ lib.optionals stdenv.isDarwin [
    darwin.apple_sdk.libs.utmp
    apple_sdk.ApplicationServices
    apple_sdk.CoreServices
    apple_sdk.OpenGL
    apple_sdk.QTKit
    apple_sdk.Security
    apple_sdk.SystemConfiguration
    xcodebuild
  ] ++ lib.optionals (!stdenv.isDarwin) [
    dbeaver
    ungoogled-chromium
    psmisc
    x11vnc
    xclip
    xvfb-run
  ];
}
