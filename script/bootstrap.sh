#! /usr/bin/env bash
export PSM=${PSM:-tiktok.im.lzz_test}
CURDIR=$(cd $(dirname $0); pwd)

if [ "X$1" != "X" ]; then
    RUNTIME_ROOT=$1
else
    RUNTIME_ROOT=${CURDIR}
fi

if [ "$TCE_HOST_ENV" == "online" ] || [ "$TCE_HOST_ENV" == "ppe" ]; then
    export KITEX_CONF_FILE="kitex.yml"
else
    export KITEX_CONF_FILE="kitex_boe.yml"
fi

export KITEX_RUNTIME_ROOT=$RUNTIME_ROOT
export KITEX_CONF_DIR="$CURDIR/conf"
export KITEX_LOG_DIR="${RUNTIME_LOGDIR:-${RUNTIME_ROOT}/log}"
export MESH_AGENT_SIDECAR_ENABLE_DATA_ABTEST_VM_AGENT=1
export DATA_ABTEST_VM_AGENT_CONFIG_FILE="$CURDIR/conf/vm_agent.conf"

if [ ! -d "$KITEX_LOG_DIR/app" ]; then
    mkdir -p "$KITEX_LOG_DIR/app"
fi

if [ ! -d "$KITEX_LOG_DIR/rpc" ]; then
    mkdir -p "$KITEX_LOG_DIR/rpc"
fi

exec "$CURDIR/bin/tiktok.im.lzz_test"
