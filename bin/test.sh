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
# go to root of project
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/../"

set -euo pipefail
IFS=$'\n\t'

SCRIPT_NAME=$(basename "$0")
LOG_FILE="/tmp/${SCRIPT_NAME%.sh}.log"

# Funciones básicas
function log_info() {
    echo "[INFO] $*" | tee -a "$LOG_FILE"
}

function log_error() {
    echo "[ERROR] $*" >&2 | tee -a "$LOG_FILE"
}

#:[.'.]:>-------------------------------------
#:[.'.]:> Tu lógica aquí
#:[.'.]:>-------------------------------------
#:[.'.]:>-------------------------------------
#:[.'.]:> Configuración de variables de entorno para tests
#:[.'.]:>-------------------------------------
#:[.'.]:> Verifica si las variables ya están definidas, si no, usa valores predeterminados
: ${DATABASE_DSN:="host=localhost user=admin password=admin dbname=markitos-svc-boilerplates sslmode=disable"}
: ${HTTP_SERVER_ADDRESS:=":3000"}
: ${GRPC_SERVER_ADDRESS:=":30000"}

#:[.'.]:> Exporta las variables para que estén disponibles para las pruebas
export DATABASE_DSN
export HTTP_SERVER_ADDRESS
export GRPC_SERVER_ADDRESS

#:[.'.]:> Muestra la configuración que vamos a usar para los tests
echo "#:[.'.]:> 🧪 Ejecutando tests con configuración:"
echo "#:[.'.]:> 📊 DATABASE_DSN=$DATABASE_DSN"
echo "#:[.'.]:> 🌐 HTTP_SERVER_ADDRESS=$HTTP_SERVER_ADDRESS"
echo "#:[.'.]:> 📡 GRPC_SERVER_ADDRESS=$GRPC_SERVER_ADDRESS"
echo "#:[.'.]:>-------------------------------------"

go clean -testcache && go test  ./...
#:[.'.]:>-------------------------------------