{ pkgs ? import <nixpkgs> { } }:
with pkgs;
let
  inherit (pkgs) stdenv;
  apple_sdk = darwin.apple_sdk.frameworks;
in mkShell {
  name = "doko";
  PGSQL_HOME = builtins.getEnv "PGSQL_HOME";
  PGSQL_DATADIR = builtins.getEnv "PGSQL_DATADIR";
  PGSQL_TCP_PORT = builtins.getEnv "PGSQL_TCP_PORT";
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
    postgresql_14
    go_1_18
    go-tools
    go-migrate
    gofumpt
    nodejs-18_x
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
    bash --posix $(whoami)
    export DOKO_ROOT=$PWD
    . $DOKO_ROOT/scripts/durable-storage-utils.sh

    export NIXPKGS_ALLOW_UNFREE=1
    export GPG_TTY=$(tty)

    export GOROOT="$(readlink -e $(type -p go) | sed  -e 's/\/bin\/go//g')"
    export NODE_HOME=${pkgs.nodejs-18_x}

    export PGSQL_BASEDIR=${pkgs.postgresql_14}
    export PGSQL_HOME="''${PGSQL_HOME:-$PWD/pgsql}"
    export PGSQL_DATADIR="''${PGSQL_DATADIR:-$PGSQL_HOME/data}"
    export PGSQL_PID_FILE=$PGSQL_HOME/pgsql.pid
    export PGSQL_TCP_PORT="''${PGSQL_TCP_PORT:-5432}"
    export PGSQL_ROOTUSER=postgres
    export PGSQL_DBUSER=doko
    export PGSQL_DOKO_DEV_DB=doko_development
    export PGSQL_DOKO_TEST_DB=doko_test
    export PGSQL_DOKO_E2E_TEST_DB=doko_e2e_test
    export PGSQL_TDD_DB=tdd
    export PGSQL_TDD_TEST_DB=tdd_test

    export PATH=$GOROOT/bin:$GOPATH/bin:$NODE_HOME/bin:$PGSQL_BASEDIR/bin:$PATH

    export CGO_ENABLED=1
    export MallocNanoZone=0 # to disable OSX automatic memory allocation before ruuning tests
    export DB_NAME=$PGSQL_DOKO_TDD_DB
    export DB_USER=$PGSQL_ROOTUSER
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
      export PGSQLD_PID=$(ps -ax | grep -v " grep " | grep $PGSQL_BASEDIR/bin/postgres | awk '{ print $1 }')

      if [[ ! -z "$PGSQLD_PID" ]]; then
        echo -e "\n\tPostgreSQL Server still running on Port: $PGSQL_TCP_PORT, at PID: $PGSQLD_PID"
        echo -e "\tYou may choose to SHUTDOWN PostgreSQL Server by issuing:"
        echo -e "\t'$PGSQL_BASEDIR/bin/pg_ctl -D $PGSQL_DATADIR stop && rm -f $PGSQL_HOME/logfile'\n"
      fi
    }

    echo "###################################################################################################################"
    echo "                                                                                "
    echo "##   !! DOKO NIX DEVELOPMENT ENVIRONMENT ;) !!      "
    echo "##   DOKO_ROOT: $DOKO_ROOT                          "
    echo "##   GOROOT: $GOROOT                                "
    echo "##   NODE_HOME: $NODE_HOME                          "
    echo "##   PGSQL_BASEDIR: $PGSQL_BASEDIR                  "
    echo "##   PGSQL_HOME: $PGSQL_HOME                        "
    echo "##   PGSQL_DATADIR: $PGSQL_DATADIR                  "
    echo "##   PGSQL_TCP_PORT: $PGSQL_TCP_PORT                "
    echo "                                                                                "
    echo "###################################################################################################################"

    echo -e "Generating .env file from env.example if required..." && generate_dotenv

    start_psql_db_cluster

    echo -e "Create DB user/s..." && create_doko_db_user $PGSQL_ROOTUSER

    echo -e "Create databases..."
    create_doko_db $PGSQL_TDD_DB
    create_doko_db $PGSQL_TDD_TEST_DB

    export GPG_TTY=$(tty)
    export NIX_SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt
    if [[ "$OSTYPE" == "darwin"* ]]; then
       export NIX_SSL_CERT_FILE=/etc/ssl/cert.pem
    fi

    trap cleanup EXIT
  '';
}
