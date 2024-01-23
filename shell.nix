{ pkgs ? import <nixpkgs> { } }:
with pkgs;
let
  inherit (pkgs) stdenv;
  apple_sdk = darwin.apple_sdk.frameworks;
in mkShell {
  name = "doko";
  PG_HOME = builtins.getEnv "PG_HOME";
  PGDATA = builtins.getEnv "PGDATA";
  PG_TCP_PORT = builtins.getEnv "PG_TCP_PORT";
  buildInputs = [
    bash
    zsh
    less
    lesspipe
    autoconf
    binutils-unwrapped
    pkg-config
    gnupg
    lsof
    git
    git-secret
    gitAndTools.delta
    postgresql_16
    go_1_20
    go-tools
    go-migrate
    gofumpt
    nodejs_20
    hostname
    inetutils
    openssh
    rsync
    bat
    fasd
    fzf
    platinum-searcher
    zoxide
    jq
    yq
    yamllint
    lzma
    jetbrains-mono
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
  shellHook = ''
    #!/usr/bin/env bash

    export DOKO_ROOT=$PWD
    . $DOKO_ROOT/scripts/durable-storage-utils.sh

    export NIXPKGS_ALLOW_UNFREE=1
    export GPG_TTY=$(tty)

    export GOROOT="$(readlink -e $(type -p go) | sed  -e 's/\/bin\/go//g')"
    export NODE_HOME=${pkgs.nodejs_20}

    export PG_BASEDIR=${pkgs.postgresql_16}
    export PG_HOME="''${PG_HOME:-$PWD/pgsql}"
    export PGDATA="''${PGDATA:-$PG_HOME/data}"
    export PG_PID_FILE=$PG_HOME/pgsql.pid
    export PG_TCP_PORT="''${PG_TCP_PORT:-5432}"
    export PG_ROOTUSER=postgres
    export PG_DBUSER=doko
    export PG_DOKO_DEV_DB=doko_development
    export PG_DOKO_TEST_DB=doko_test
    export PG_DOKO_E2E_TEST_DB=doko_e2e_test
    export PG_TDD_DB=tdd
    export PG_TDD_TEST_DB=tdd_test

    export PATH=$GOROOT/bin:$NODE_HOME/bin:$PG_BASEDIR/bin:$PATH

    export CGO_ENABLED=1
    export MallocNanoZone=0 # to disable OSX automatic memory allocation before ruuning tests
    export DB_NAME=$PG_TDD_DB
    export DB_USER=$PG_ROOTUSER
    export DB_PASSWORD=password

    generate_dotenv() {
      if [[ ! -f ".env" ]]; then
        cp ./env.example ./.env
        sed -i 's/test-postgres_tdd/localhost/g'
        . ./.env
      fi
    }

    cleanup()
    {
      echo -e "\nBYE!!! EXITING doko NIX DEVELOPMENT ENVIRONMENT."
      export PG_PID=$(ps -ax | grep -v " grep " | grep $PG_BASEDIR/bin/postgres | awk '{ print $1 }')

      if [[ ! -z "$PG_PID" ]]; then
        echo -e "\n\tPostgreSQL Server still running on Port: $PG_TCP_PORT, at PID: $PG_PID"
        echo -e "\tYou may choose to SHUTDOWN PostgreSQL Server by issuing:"
        echo -e "\t'$PG_BASEDIR/bin/pg_ctl -D $PGDATA stop && rm -f $PG_HOME/logfile'\n"
      fi
    }

    echo "###################################################################################################################"
    echo "                                                                                "
    echo "##   !! DOKO NIX DEVELOPMENT ENVIRONMENT ;) !!      "
    echo "##   DOKO_ROOT: $DOKO_ROOT                          "
    echo "##   GOROOT: $GOROOT                                "
    echo "##   NODE_HOME: $NODE_HOME                          "
    echo "##   PG_BASEDIR: $PG_BASEDIR                        "
    echo "##   PG_HOME: $PG_HOME                           "
    echo "##   PGDATA: $PGDATA                                "
    echo "##   PG_TCP_PORT: $PG_TCP_PORT                      "
    echo "                                                                                "
    echo "###################################################################################################################"

    echo -e "Generating .env file from env.example if required..." && generate_dotenv

    start_psql_db_cluster

    echo -e "Create DB user/s..." && create_doko_db_user $PG_ROOTUSER

    echo -e "Create databases..."
    create_doko_db $PG_TDD_DB
    create_doko_db $PG_TDD_TEST_DB

    export GPG_TTY=$(tty)
    export NIX_SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt
    if [[ "$OSTYPE" == "darwin"* ]]; then
       export NIX_SSL_CERT_FILE=/etc/ssl/cert.pem
    fi

    trap cleanup EXIT
  '';
}
