#!/bin/bash

#:[.'.]:> ============================================
#:[.'.]:> 🌍 CONFIGURACIÓN CENTRALIZADA DE ENTORNO
#:[.'.]:> ============================================
#:[.'.]:> Este script centraliza todas las configuraciones
#:[.'.]:> predeterminadas del proyecto para evitar duplicación
#:[.'.]:> y mantener consistencia entre los diferentes scripts.
#:[.'.]:> ============================================

#:[.'.]:> Valores predeterminados para conexión a la base de datos
BOILERPLATES_DEFAULT_DATABASE_HOST="localhost"
BOILERPLATES_DEFAULT_DATABASE_USER="admin"
BOILERPLATES_DEFAULT_DATABASE_PASSWORD="admin"
BOILERPLATES_DEFAULT_DATABASE_NAME="markitos-svc-boilerplates"
BOILERPLATES_DEFAULT_DATABASE_SSL_MODE="disable"
BOILERPLATES_DEFAULT_IMAGE_NAME="markitos-svc-boilerplates"
BOILERPLATES_DEFAULT_BOILERPLATES_HTTP_SERVER_ADDRESS=":3003"
BOILERPLATES_DEFAULT_BOILERPLATES_HTTP_SERVER_PORT="3003"
BOILERPLATES_DEFAULT_BOILERPLATES_GRPC_SERVER_ADDRESS=":9000"
BOILERPLATES_DEFAULT_BOILERPLATES_GRPC_SERVER_PORT="9000"

#:[.'.]:> Construir DSN predeterminado
BOILERPLATES_DEFAULT_BOILERPLATES_DATABASE_DSN="host=${BOILERPLATES_DEFAULT_DATABASE_HOST} user=${BOILERPLATES_DEFAULT_DATABASE_USER} password=${BOILERPLATES_DEFAULT_DATABASE_PASSWORD} dbname=${BOILERPLATES_DEFAULT_DATABASE_NAME} sslmode=${BOILERPLATES_DEFAULT_DATABASE_SSL_MODE}"

#:[.'.]:> Otros valores predeterminados
#:[.'.]:> Función para configurar variables de entorno
#:[.'.]:> Esta función establece las variables si no están definidas
#:[.'.]:> y las exporta para que estén disponibles para los procesos hijos
function setup_environment() {
    #:[.'.]:> Establecer variables si no están definidas
    : ${BOILERPLATES_DATABASE_DSN:="${BOILERPLATES_DEFAULT_BOILERPLATES_DATABASE_DSN}"}
    : ${BOILERPLATES_HTTP_SERVER_ADDRESS:="${BOILERPLATES_DEFAULT_BOILERPLATES_HTTP_SERVER_ADDRESS}"}
    : ${BOILERPLATES_GRPC_SERVER_ADDRESS:="${BOILERPLATES_DEFAULT_BOILERPLATES_GRPC_SERVER_ADDRESS}"}
    : ${BOILERPLATES_HTTP_SERVER_PORT:="${BOILERPLATES_DEFAULT_BOILERPLATES_HTTP_SERVER_PORT}"}
    : ${BOILERPLATES_GRPC_SERVER_PORT:="${BOILERPLATES_DEFAULT_BOILERPLATES_GRPC_SERVER_PORT}"}
    : ${BOILERPLATES_IMAGE_NAME:="${BOILERPLATES_DEFAULT_IMAGE_NAME}"}

    #:[.'.]:> Exportar variables
    export BOILERPLATES_DATABASE_DSN
    export BOILERPLATES_HTTP_SERVER_ADDRESS
    export BOILERPLATES_GRPC_SERVER_ADDRESS
    export BOILERPLATES_HTTP_SERVER_PORT
    export BOILERPLATES_GRPC_SERVER_PORT
    export BOILERPLATES_IMAGE_NAME
}

#:[.'.]:> Función para mostrar la configuración actual
#:[.'.]:> Parámetro $1 == "full" mostrará todas las variables
#:[.'.]:> Sin parámetros mostrará solo BOILERPLATES_DATABASE_DSN
function show_config() {
    echo "#:[.'.]:> 🚀 Iniciando con configuración:"
    echo "#:[.'.]:> 📊 BOILERPLATES_DATABASE_DSN=$BOILERPLATES_DATABASE_DSN"
    
    if [[ "${1:-}" == "full" ]]; then
        echo "#:[.'.]:> 🌐 BOILERPLATES_HTTP_SERVER_ADDRESS=$BOILERPLATES_HTTP_SERVER_ADDRESS"
        echo "#:[.'.]:> 📡 BOILERPLATES_GRPC_SERVER_ADDRESS=$BOILERPLATES_GRPC_SERVER_ADDRESS"
    fi
    
    echo "#:[.'.]:>-------------------------------------"
}