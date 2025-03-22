echo 
echo "# ============================================="
echo "#  __  __  ____  _  __"
echo "# |  \\/  |  _ \\| |/ /"
echo "# | \\  / | | | | ' / "
echo "# | |\\/| | | | |  <  "
echo "# | |  | | |_| | . \\ "
echo "# |_|  |_|____/|_|\\_\\"
echo "#                                   "
echo "#  Creador: Marco Antonio - markitos      "
echo "# ============================================="
echo "#  🥷 (mArKit0sDevSecOpsKit) 🗡️"
echo "#  Markitos DevSecOps Kulture"
echo "# ============================================="
echo 
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/../"
set -euo pipefail
IFS=$'\n\t'
SCRIPT_NAME=$(basename "$0")
function log_info() {
    echo "[INFO] $*"
}
function log_error() {
    echo "[ERROR] $*" >&2
}

#:[.'.]:>-------------------------------------
#:[.'.]:> Tu lógica aquí
#:[.'.]:>-------------------------------------
log_info "Argumentos proporcionados: $*"
protoc \
    --proto_path=infrastructure/proto \
    --go_out=infrastructure/gapi \
    --go-grpc_out=infrastructure/gapi \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative \
    infrastructure/proto/*.proto
#:[.'.]:>-------------------------------------