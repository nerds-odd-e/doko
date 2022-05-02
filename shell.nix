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
  REDIS_TCP_PORT = builtins.getEnv "REDIS_TCP_PORT";
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
    redis
    go_1_18
    go-tools
    go-migrate
    gofumpt
    nodejs-18_x
    buildkit
    docker
    docker-compose_2
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
    set -o err_return
    set -o no_unset
    set -o pipefail

    export NIXPKGS_ALLOW_UNFREE=1
    export GPG_TTY=$(tty)

    export GOROOT="$(readlink -e $(type -p go) | sed  -e 's/\/bin\/go//g')"
    export NODE_HOME=${pkgs.nodejs-18_x}

    export REDIS_BASEDIR=${pkgs.redis}
    export REDIS_TCP_PORT="''${REDIS_TCP_PORT:-6379}"

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

    export PATH=$GOROOT/bin:$GOPATH/bin:$NODE_HOME/bin:$PGSQL_BASEDIR/bin:$REDIS_BASEDIR/bin:$PATH

    export CGO_ENABLED=1
    export MallocNanoZone=0 # to disable OSX automatic memory allocation before ruuning tests
    export DB_NAME=$PGSQL_DOKO_TDD_DB
    export DB_USER=$PGSQL_ROOTUSER
    export DB_PASSWORD=password

    init_psql_db_cluster() {
      mkdir -p $PGSQL_HOME
      mkdir -p $PGSQL_DATADIR

      export PGSQLD_PID=$(ps -ax | grep -v " grep " | grep $PGSQL_BASEDIR/bin/postgres | awk '{ print $1 }')
      if [[ -z "$PGSQLD_PID" ]]; then
        [ ! "$(ls -A ''${PGSQL_DATADIR})" ] && rm -rf $PGSQL_HOME/logfile && $PGSQL_BASEDIR/bin/pg_ctl -U $PGSQL_ROOTUSER -D $PGSQL_DATADIR initdb
        $PGSQL_BASEDIR/bin/pg_ctl -U $PGSQL_ROOTUSER -D $PGSQL_DATADIR -o "-p ''${PGSQL_TCP_PORT}" -l $PGSQL_HOME/logfile start
        export PGSQLD_PID=$!
      fi
    }

    create_doko_db_user() {
      DB_USER_NAME=$1
      if psql -h localhost -p $PGSQL_TCP_PORT $PGSQL_ROOTUSER -t -c '\du' | cut -d \| -f 1 | grep -qw $DB_USER_NAME; then
        echo "PostgreSQL DB User $DB_USER_NAME already exists."
      else
        createuser -e -h localhost -p $PGSQL_TCP_PORT -d -r -S $DB_USER_NAME
      fi
    }

    create_doko_db() {
      DB_NAME=$1
      DB_OWNER_NAME="''${2:-$PGSQL_ROOTUSER}"
      if psql -h localhost -p $PGSQL_TCP_PORT $PGSQL_ROOTUSER -lqt | cut -d \| -f 1 | grep -qw $DB_NAME; then
        echo "PostgreSQL doko DB $DB_NAME already exists. Continuing..."
      else
        createdb -e -h localhost -p $PGSQL_TCP_PORT -O $DB_OWNER_NAME $DB_NAME
      fi
    }

    start_redis_server() {
      EXISTING_REDIS_SERVER_PID=$(lsof -i :6379 -sTCP:LISTEN | awk 'NR > 1 {print $2}' | tail -1)
      if [[ -z $EXISTING_REDIS_SERVER_PID ]]; then
        redis-server &
        export REDIS_PID=$!
        redis-cli config set save "" &
      fi
    }

    generate_dotenv() {
      if [ ! -f ".env" ]; then
        cp env.example .env
        sed -i 's/test-postgres_tdd/localhost/g'
      fi
    }

    cleanup()
    {
      echo -e "\nBYE!!! EXITING doko NIX DEVELOPMENT ENVIRONMENT."
      export PGSQLD_PID=$(ps -ax | grep -v " grep " | grep $PGSQL_BASEDIR/bin/postgres | awk '{ print $1 }')
      export REDIS_PID=$(lsof -i :6379 -sTCP:LISTEN | awk 'NR > 1 {print $2}' | tail -1)

      if [[ ! -z "$PGSQLD_PID" ]]; then
        echo -e "\n\tPostgreSQL Server still running on Port: $PGSQL_TCP_PORT, at PID: $PGSQLD_PID"
        echo -e "\tYou may choose to SHUTDOWN PostgreSQL Server by issuing:"
        echo -e "\t'$PGSQL_BASEDIR/bin/pg_ctl -D $PGSQL_DATADIR stop && rm -f $PGSQL_HOME/logfile'\n"
      fi

      if [[ ! -z "$REDIS_PID" ]]; then
        echo -e "\n\tRedis Server still running on Port: $REDIS_TCP_PORT, at PID: $REDIS_PID"
        echo -e "\tYou may choose to SHUTDOWN Redis Server by issuing:"
        echo -e "\t'$REDIS_BASEDIR/bin/redis-cli shutdown nosave'\n"
      fi
    }

    echo "###################################################################################################################"
    echo "                                                                                "
    echo "##   !! DOKO NIX DEVELOPMENT ENVIRONMENT ;) !!      "
    echo "##   GOROOT: $GOROOT                                "
    echo "##   NODE_HOME: $NODE_HOME                          "
    echo "##   PGSQL_BASEDIR: $PGSQL_BASEDIR                  "
    echo "##   PGSQL_HOME: $PGSQL_HOME                        "
    echo "##   PGSQL_DATADIR: $PGSQL_DATADIR                  "
    echo "##   PGSQL_TCP_PORT: $PGSQL_TCP_PORT                "
    echo "##   REDIS_BASEDIR: $REDIS_BASEDIR                  "
    echo "##   REDIS_TCP_PORT: $REDIS_TCP_PORT                "
    echo "                                                                                "
    echo "###################################################################################################################"

    init_psql_db_cluster

    create_doko_db_user $PGSQL_ROOTUSER
    create_doko_db $PGSQL_TDD_DB
    create_doko_db $PGSQL_TDD_TEST_DB

    start_redis_server

    export GPG_TTY=$(tty)
    export NIX_SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt
    if [[ "$OSTYPE" == "darwin"* ]]; then
       export NIX_SSL_CERT_FILE=/etc/ssl/cert.pem
    fi

    trap cleanup EXIT
  '';
}
