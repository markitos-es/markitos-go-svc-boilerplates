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

#:[.'.]:> Preguntar por el SNYK_TOKEN
echo -e "\033[1;36m🔑 ¿Tienes un SNYK_TOKEN para configurar? (Déjalo en blanco para usar 'replace_me')\033[0m"
read -p "Introduce tu SNYK_TOKEN: " SNYK_TOKEN
SNYK_TOKEN=${SNYK_TOKEN:-replace_me}

#:[.'.]:> Mostrar lo que hará el script
echo -e "\033[1;36m🛠️ Este script instalará las siguientes herramientas en ~/.local/bin:\033[0m"
echo -e "  - \033[1;33mSnyk CLI\033[0m (https://snyk.io)"
echo -e "  - \033[1;33mGitleaks\033[0m (https://github.com/gitleaks/gitleaks)"
echo
echo -e "\033[1;36m📋 Resumen de acciones:\033[0m"
echo -e "  1. Descargar los binarios."
echo -e "  2. Moverlos a ~/.local/bin."
echo -e "  3. Hacerlos ejecutables."
echo -e "  4. Actualizar el PATH y configurar SNYK_TOKEN."
echo -e "  5. Verificar las versiones instaladas."
echo
echo -e "\033[1;33m⚠️ Presiona CTRL+C para cancelar o ENTER para continuar...\033[0m"
read -r

#:[.'.]:> Crear directorio ~/.local/bin si no existe
mkdir -p ~/.local/bin

#:[.'.]:> Instalar Snyk CLI
log_info "Descargando e instalando Snyk CLI..."
cd /tmp
curl -s https://static.snyk.io/cli/latest/snyk-linux -o snyk
chmod +x ./snyk
mv ./snyk ~/.local/bin/snyk
chmod u+x ~/.local/bin/snyk

#:[.'.]:> Instalar Gitleaks
log_info "Descargando e instalando Gitleaks..."
cd /tmp
wget -q https://github.com/gitleaks/gitleaks/releases/download/v8.24.0/gitleaks_8.24.0_linux_x64.tar.gz
tar xfz gitleaks_8.24.0_linux_x64.tar.gz
mv ./gitleaks ~/.local/bin/gitleaks
chmod u+x ~/.local/bin/gitleaks

#:[.'.]:> Actualizar PATH y configurar SNYK_TOKEN
if ! echo "$PATH" | grep -q "$HOME/.local/bin"; then
    log_info "Actualizando PATH en ~/.bashrc..."
    echo 'export PATH=${PATH}:${HOME}/.local/bin' >> ~/.bashrc
fi
log_info "Configurando SNYK_TOKEN en ~/.bashrc..."
echo "export SNYK_TOKEN=${SNYK_TOKEN}" >> ~/.bashrc
source ~/.bashrc

#:[.'.]:> Verificar las herramientas instaladas
log_info "Verificando las herramientas instaladas..."
SNYK_VERSION=$(~/.local/bin/snyk --version 2>/dev/null || echo "No instalado")
GITLEAKS_VERSION=$(~/.local/bin/gitleaks version 2>/dev/null || echo "No instalado")

#:[.'.]:> Mostrar informe final
echo
echo -e "\033[1;36m📋 Informe final:\033[0m"
if [[ "$SNYK_VERSION" != "No instalado" ]]; then
    log_success "Snyk CLI instalado correctamente. Versión: $SNYK_VERSION"
else
    log_error "Snyk CLI no se instaló correctamente."
fi

if [[ "$GITLEAKS_VERSION" != "No instalado" ]]; then
    log_success "Gitleaks instalado correctamente. Versión: $GITLEAKS_VERSION"
else
    log_error "Gitleaks no se instaló correctamente."
fi

#:[.'.]:> Limpiar archivos temporales
log_info "Limpiando archivos temporales..."
rm -f /tmp/gitleaks_8.24.0_linux_x64.tar.gz
rm -f /tmp/README.md /tmp/LICENSE

echo
log_success "🎉 Instalación completada. ¡Todo listo para usar!"