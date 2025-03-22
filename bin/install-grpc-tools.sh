#!/bin/bash

echo 
echo "#:[.'.]:>=============================================="
echo "#:[.'.]:>  __  __  ____  _  __"
echo "#:[.'.]:> |  \\/  |  _ \\| |/ /"
echo "#:[.'.]:> | \\  / | | | | ' / "
echo "#:[.'.]:> | |\\/| | | | |  <  "
echo "#:[.'.]:> | |  | | |_| | . \\ "
echo "#:[.'.]:> |_|  |_|____/|_|\\_\\"
echo "#:[.'.]:>                                   "
echo "#:[.'.]:>  Creador: Marco Antonio - markitos      "
echo "#:[.'.]:>=============================================="
echo "#:[.'.]:>= 🥷 (mArKit0sDevSecOpsKit) 🗡️"
echo "#:[.'.]:>= Markitos DevSecOps Kulture"
echo "#:[.'.]:>=============================================="
echo 

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/../"
set -euo pipefail
IFS=$'\n\t'

#:[.'.]:> Definir funciones de logging
function log_info() {
    echo -e "\033[1;34m[INFO]\033[0m $*"
}

function log_error() {
    echo -e "\033[1;31m[ERROR]\033[0m $*" >&2
}

function log_success() {
    echo -e "\033[1;32m[SUCCESS]\033[0m $*"
}

#:[.'.]:> Mostrar lo que hará el script
echo -e "\033[1;36m🛠️ Este script instalará las siguientes herramientas:\033[0m"
echo -e "  - \033[1;33mprotoc\033[0m (Protocol Buffers Compiler)"
echo -e "  - \033[1;33mPlugins de Go para gRPC\033[0m"
echo
echo -e "\033[1;36m📋 Resumen de acciones:\033[0m"
echo -e "  1. Descargar e instalar protoc."
echo -e "  2. Instalar plugins de Go para gRPC."
echo -e "  3. Actualizar el PATH en ~/.bashrc."
echo
echo -e "\033[1;33m⚠️ Presiona CTRL+C para cancelar o ENTER para continuar...\033[0m"
read -r

#:[.'.]:> Instalar protoc (ultima version marzo 2025)
PROTOC_VERSION=30.1
log_info "Descargando e instalando protoc (versión ${PROTOC_VERSION})..."
cd /tmp
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip
sudo apt install unzip -y
unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d protoc
sudo mv protoc/bin/protoc /usr/local/bin/
sudo mv protoc/include/* /usr/local/include/
rm -rf protoc protoc-${PROTOC_VERSION}-linux-x86_64.zip
log_success "protoc instalado correctamente. Versión: $(protoc --version)"

#:[.'.]:> Instalar plugins de Go para gRPC
log_info "Instalando plugins de Go para gRPC..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
log_success "Plugins de Go para gRPC instalados correctamente."

#:[.'.]:> Actualizar PATH en ~/.bashrc
if ! echo "$PATH" | grep -q "$(go env GOPATH)/bin"; then
    log_info "Actualizando PATH en ~/.bashrc..."
    echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bashrc
    source ~/.bashrc
    log_success "PATH actualizado correctamente."
else
    log_info "El PATH ya incluye $(go env GOPATH)/bin."
fi

#:[.'.]:> Verificar instalación
log_info "Verificando las herramientas instaladas..."
PROTOC_VERSION_INSTALLED=$(protoc --version 2>/dev/null || echo "No instalado")
PROTOC_GEN_GO_VERSION=$(protoc-gen-go --version 2>/dev/null || echo "No instalado")
PROTOC_GEN_GO_GRPC_VERSION=$(protoc-gen-go-grpc --version 2>/dev/null || echo "No instalado")

#:[.'.]:> Mostrar informe final
echo
echo -e "\033[1;36m📋 Informe final:\033[0m"
if [[ "$PROTOC_VERSION_INSTALLED" != "No instalado" ]]; then
    log_success "protoc instalado correctamente. Versión: $PROTOC_VERSION_INSTALLED"
else
    log_error "protoc no se instaló correctamente."
fi

if [[ "$PROTOC_GEN_GO_VERSION" != "No instalado" ]]; then
    log_success "protoc-gen-go instalado correctamente. Versión: $PROTOC_GEN_GO_VERSION"
else
    log_error "protoc-gen-go no se instaló correctamente."
fi

if [[ "$PROTOC_GEN_GO_GRPC_VERSION" != "No instalado" ]]; then
    log_success "protoc-gen-go-grpc instalado correctamente. Versión: $PROTOC_GEN_GO_GRPC_VERSION"
else
    log_error "protoc-gen-go-grpc no se instaló correctamente."
fi

echo
log_success "🎉 Instalación completada. ¡Todo listo para usar!"