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

# USO:
#:[.'.]:> En desarrollo:
#:[.'.]:>   ./bin/run.sh                           → Usa valores por defecto
#:[.'.]:>   DATABASE_DSN=... HTTP_SERVER_ADDRESS=... ./bin/run.sh → Personalizado
#:[.'.]:> 
#:[.'.]:> Con make:
#:[.'.]:>   make run                               → Usa valores por defecto
#:[.'.]:>   DATABASE_DSN=... make run              → Personalizado

# go to root of project
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/../"

set -euo pipefail
IFS=$'\n\t'

#:[.'.]:>-------------------------------------
#:[.'.]:> Configuración de variables de entorno
#:[.'.]:>-------------------------------------
#:[.'.]:> Verifica si las variables ya están definidas, si no, usa valores predeterminados
: ${DATABASE_DSN:="host=localhost user=admin password=admin dbname=markitos-svc-boilerplates sslmode=disable"}
: ${HTTP_SERVER_ADDRESS:=":3000"}
: ${GRPC_SERVER_ADDRESS:=":30000"}

#:[.'.]:> Exporta las variables para que estén disponibles para el proceso hijo
export DATABASE_DSN
export HTTP_SERVER_ADDRESS
export GRPC_SERVER_ADDRESS

#:[.'.]:> Muestra la configuración que vamos a usar
echo "#:[.'.]:> 🚀 Iniciando con configuración:"
echo "#:[.'.]:> 📊 DATABASE_DSN=$DATABASE_DSN"
echo "#:[.'.]:> 🌐 HTTP_SERVER_ADDRESS=$HTTP_SERVER_ADDRESS"
echo "#:[.'.]:> 📡 GRPC_SERVER_ADDRESS=$GRPC_SERVER_ADDRESS"
echo "#:[.'.]:>-------------------------------------"

#:[.'.]:> Ejecuta la aplicación con la configuración establecida
go run cmd/main.go
#:[.'.]:>-------------------------------------