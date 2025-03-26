# ==================================================================================
# 🔥 COMANDOS DISPONIBLES 🔥
# ==================================================================================
# 
# 🧪 PRUEBAS Y DESARROLLO:
# ----------------------
# make test          → Ejecuta tests unitarios
# make testv         → Ejecuta tests con modo verboso
# make postgres      → Levanta PostgreSQL en Docker
# make run           → Ejecuta la aplicación (se pueden usar variables: DATABASE_DSN=... make run)
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
#
# 🐳 DOCKER:
# -------
# make image         → Construye imagen Docker (ej: make image VERSION=1.2.3)
# make image-run     → Ejecuta imagen Docker (ej: make image-run VERSION=1.2.3)
# ==================================================================================

# Variables de entorno por defecto (pueden ser sobreescritas en la línea de comandos)
DATABASE_DSN ?= host=localhost user=admin password=admin dbname=markitos-svc-boilerplates sslmode=disable
HTTP_SERVER_ADDRESS ?= :3000
GRPC_SERVER_ADDRESS ?= :30000
VERSION ?= 1.0.0

# Definir todos los targets como PHONY para evitar conflictos con archivos del mismo nombre
.PHONY: test testv postgres run prun security createdb dropdb install-appsec-tools install-grpc-tools certificate proto image image-run

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
	DATABASE_DSN="$(DATABASE_DSN)" HTTP_SERVER_ADDRESS="$(HTTP_SERVER_ADDRESS)" GRPC_SERVER_ADDRESS="$(GRPC_SERVER_ADDRESS)" bash bin/run.sh

#:[.'.]:> Ejecuta la aplicación en paralelo - ¡Para no bloquear la terminal!
prun:
	DATABASE_DSN="$(DATABASE_DSN)" HTTP_SERVER_ADDRESS="$(HTTP_SERVER_ADDRESS)" GRPC_SERVER_ADDRESS="$(GRPC_SERVER_ADDRESS)" GIN_MODE=release bash bin/run.sh &

#:[.'.]:> Analiza seguridad del código - ¡Detectamos vulnerabilidades antes de que sean problema!
security:
	@echo "#:[.'.]:> Ejecutando análisis de seguridad en el código Go..."
	@echo "#:[.'.]:> Ejecutando análisis Snyk..."
	@SNYK_TOKEN=${SNYK_TOKEN} snyk code test
	@SNYK_TOKEN=${SNYK_TOKEN} snyk test --all-projects --detection-depth=10
	@echo "#:[.'.]:> Ejecutando Gitleaks para detectar secrets..."
	@gitleaks detect --source . --verbose

#:[.'.]:> Crea la base de datos - ¡Preparando el terreno para nuestros datos!
createdb:
	DATABASE_DSN="$(DATABASE_DSN)" bash bin/createdb.sh

#:[.'.]:> Elimina la base de datos - ¡Borrón y cuenta nueva cuando lo necesitemos!
dropdb:
	DATABASE_DSN="$(DATABASE_DSN)" bash bin/dropdb.sh

#:[.'.]:> Instala herramientas de seguridad - ¡El kit completo para estar protegidos!
install-appsec-tools:
	bash bin/install-appsec-tools.sh

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
	@echo "#:[.'.]:> Construyendo imagen Docker versión: $(VERSION)"; \
	docker build -t markitos-svc-boilerplates:$(VERSION) -t markitos-svc-boilerplates:latest .; \
	echo "#:[.'.]:> ¡Imagen markitos-svc-boilerplates:$(VERSION) creada con éxito! 🚀"

#:[.'.]:> Ejecuta imagen Docker - ¡Prueba la imagen antes de desplegarla en producción!
image-run:
	@echo "#:[.'.]:> Ejecutando imagen Docker versión: $(VERSION)"; \
	docker run -p 3000:3000 -p 30000:30000 \
	-e DATABASE_DSN="$(DATABASE_DSN)" \
	-e HTTP_SERVER_ADDRESS="$(HTTP_SERVER_ADDRESS)" \
	-e GRPC_SERVER_ADDRESS="$(GRPC_SERVER_ADDRESS)" \
	markitos-svc-boilerplates:$(VERSION)