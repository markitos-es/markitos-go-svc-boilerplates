# ==================================================================================
# 🔥 COMANDOS DISPONIBLES 🔥
# ==================================================================================
# 
# 🧪 PRUEBAS Y DESARROLLO:
# ----------------------
# make test          → Ejecuta tests unitarios
# make testv         → Ejecuta tests con modo verboso
# make postgres      → Levanta PostgreSQL en Docker
# make run           → Ejecuta la aplicación (se pueden usar variables: BOILERPLATES_DATABASE_DSN=... make run)
# make prun          → Ejecuta la aplicación en paralelo
#
# 🔒 SEGURIDAD:
# ----------
# make security      → Analiza código en busca de vulnerabilidades
#
# 🗄️ BASE DE DATOS:
# -------------
# make createdb      → Crea la base de datos
# make dropdb        → Elimina la base de datos
#
# 🛠️ HERRAMIENTAS:
# ------------
# make install-appsec-tools  → Instala herramientas de seguridad
# make install-grpc-tools    → Instala herramientas gRPC
# make certificate           → Genera certificado SSH
# make proto                 → Genera código desde definiciones proto
# make caas                  → Clone As A Service - Crea nuevos servicios usando esta plantilla
#                              (ej: make caas SERVICE_NAME=pepito-svc-mariposas ENTITY_NAME=butterfly)
#
# 🐳 DOCKER:
# -------
# make image         → Construye imagen Docker (ej: make image VERSION=1.2.3)
# make image-run     → Ejecuta imagen Docker (ej: make image-run VERSION=1.2.3)
# make tag           → Crea un tag en Git para la versión especificada
# ==================================================================================

# Definir todos los targets como PHONY para evitar conflictos con archivos del mismo nombre
.PHONY: test testv postgres run prun security createdb dropdb install-appsec-tools install-grpc-tools certificate proto image image-run caas

#:[.'.]:> Ejecuta tests unitarios - ¡Aseguramos que todo funcione como debe!
test:
	bash bin/test.sh

#:[.'.]:> Ejecuta tests en modo verboso - ¡Para cuando queremos todos los detalles!
testv:
	bash bin/testv.sh

#:[.'.]:> Levanta PostgreSQL en Docker - ¡Base de datos lista en segundos!
postgres:
	bash bin/postgres.sh

#:[.'.]:> Ejecuta la aplicación - ¡A darle vida a nuestro servicio!
run:
	bash bin/run.sh

#:[.'.]:> Ejecuta la aplicación en paralelo - ¡Para no bloquear la terminal!
prun:
	bash bin/prun.sh

#:[.'.]:> Analiza seguridad del código - ¡Detectamos vulnerabilidades antes de que sean problema!
security:
	bash bin/security.sh

#:[.'.]:> Crea la base de datos - ¡Preparando el terreno para nuestros datos!
createdb:
	BOILERPLATES_DATABASE_DSN="$(BOILERPLATES_DATABASE_DSN)" bash bin/createdb.sh

#:[.'.]:> Elimina la base de datos - ¡Borrón y cuenta nueva cuando lo necesitemos!
dropdb:
	BOILERPLATES_DATABASE_DSN="$(BOILERPLATES_DATABASE_DSN)" bash bin/dropdb.sh

#:[.'.]:> Instala herramientas de seguridad - ¡El kit completo para estar protegidos!
install-appsec-tools:
	ASK_FOR_SNYK_TOKEN_BYPASS=true SNYK_TOKEN=${SNYK_TOKEN} bash bin/install-appsec-tools.sh

#:[.'.]:> Instala herramientas gRPC - ¡Todo lo necesario para trabajar con protobuf y gRPC!
install-grpc-tools:
	bash bin/install-grpc-tools.sh

#:[.'.]:> Genera certificado SSH para GitHub - ¡Para conectarse fácil y seguro!
certificate:
	bash bin/github-ssh-key.sh $(name) $(email)

#:[.'.]:> Genera código desde definiciones proto - ¡Actualiza las interfaces de comunicación!
proto:
	bash bin/proto.sh

#:[.'.]:> Construye imagen Docker - ¡Empaquetamos la app para distribuirla fácilmente!
image:
	bash bin/image.sh
	@echo "#:[.'.]:> Para probar la imagen ejecuta: make image-run VERSION=$(VERSION)"

#:[.'.]:> Ejecuta imagen Docker - ¡Prueba la imagen antes de desplegarla en producción!
image-run:
	bash bin/image-run.sh

#:[.'.]:> Creacion de un tag para git
tag:
	bash bin/tag.sh

#:[.'.]:> Clone As A Service - ¡Crea un nuevo servicio a partir de esta plantilla!
caas:
	bash bin/caas.sh