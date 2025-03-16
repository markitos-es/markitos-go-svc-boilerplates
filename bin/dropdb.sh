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
echo "#  🥷 (Cultura DevSecOps) 🗡️"
echo "#  Markitos DevSecOps Kulture. "
echo "# ============================================="
echo 
# go to root of project
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/.."
set -euo pipefail
IFS=$'\n\t'

# Funciones básicas
function log_info() {
    echo "[INFO] $*"
}

function log_error() {
    echo "[ERROR] $*" >&2
}

# Cargar variables de entorno desde app.env
if [ -f app.env ]; then
    set -o allexport
    source app.env
    set +o allexport
else
    log_error "El archivo app.env no existe"
    exit 1
fi
# Verificar que DATABASE_DSN esté definido
if [ -z "${DATABASE_DSN:-}" ]; then
    log_error "DATABASE_DSN no está definido en app.env"
    exit 1
fi

# Extraer datos de conexión de DATABASE_DSN
DB_NAME=$(echo $DATABASE_DSN | awk -F'[ =]' '{for(i=1;i<=NF;i++){if($i=="dbname"){print $(i+1)}}}')
DB_USER=$(echo $DATABASE_DSN | awk -F'[ =]' '{for(i=1;i<=NF;i++){if($i=="user"){print $(i+1)}}}')

# Verificar si la base de datos existe
function database_exists() {
    docker exec markitos-common-postgres psql -U admin -d postgres -tAc "SELECT 1 FROM pg_database WHERE datname='$DB_NAME'" | grep -q 1
}

# Verificar si el usuario existe
function user_exists() {
    docker exec markitos-common-postgres psql -U admin -d postgres -tAc "SELECT 1 FROM pg_roles WHERE rolname='$DB_USER'" | grep -q 1
}

# Terminar conexiones activas del usuario
function terminate_connections() {
    docker exec markitos-common-postgres psql -U admin -d postgres -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE usename='$DB_USER';" || true
}

#:[.'.]:>-------------------------------------
#:[.'.]:> Tu lógica aquí
#:[.'.]:>-------------------------------------
log_info "Eliminando base de datos $DB_NAME"

if database_exists "$DB_NAME"; then
    docker exec markitos-common-postgres dropdb --username=admin "$DB_NAME"
    log_info "Base de datos $DB_NAME eliminada"
else
    log_info "La base de datos $DB_NAME no existe"
fi

if user_exists "$DB_USER"; then
    log_info "Terminando conexiones activas del usuario $DB_USER"
    terminate_connections || true
    docker exec markitos-common-postgres psql -U admin -d postgres -c "DROP USER \"$DB_USER\";" || true
    log_info "Usuario $DB_USER eliminado"
else
    log_info "El usuario $DB_USER no existe"
fi
#:[.'.]:>-------------------------------------