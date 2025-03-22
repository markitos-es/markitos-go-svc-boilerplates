# Seguridad en el Desarrollo 🛡️

Este proyecto incluye herramientas y scripts para garantizar la seguridad en el desarrollo siguiendo las mejores prácticas de DevSecOps. A continuación, se describen las herramientas disponibles, cómo utilizarlas y cómo integrarlas en tu flujo de trabajo.

---

## Índice 📚

1. [Herramientas de Seguridad Incluidas 🔒](#herramientas-de-seguridad-incluidas-)
  - [Snyk CLI](#1-snyk-cli)
  - [Gitleaks](#2-gitleaks)
  - [Generación de Claves SSH](#3-generación-de-claves-ssh)
2. [Ejecución de Seguridad con Makefile ⚙️](#ejecución-de-seguridad-con-makefile-️)
  - [Análisis de Seguridad](#1-análisis-de-seguridad)
  - [Generación de Certificados SSH](#2-generación-de-certificados-ssh)
3. [Ejemplo de Flujo de Seguridad 🛠️](#ejemplo-de-flujo-de-seguridad-️)
4. [Configuración Adicional](#configuración-adicional)

---

## Herramientas de Seguridad Incluidas 🔒

### 1. **Snyk CLI**
Snyk es una herramienta que analiza el código fuente y las dependencias en busca de vulnerabilidades de seguridad.

- **Instalación**: Usa el script `install-appsec-tools.sh` para instalar Snyk CLI.
- **Uso**:
  ```bash
  make security
  ```
  Esto ejecutará un análisis de seguridad del código fuente utilizando Snyk.

### 2. **Gitleaks**
Gitleaks es una herramienta para detectar secretos y credenciales expuestos en el repositorio.

- **Instalación**: Usa el script `install-appsec-tools.sh` para instalar Gitleaks.
- **Uso**:
  ```bash
  make security
  ```
  Esto ejecutará un análisis de seguridad en busca de secretos expuestos utilizando Gitleaks.

### 3. **Generación de Claves SSH**
Puedes generar claves SSH para autenticarte con GitHub u otros servicios de forma segura.

- **Uso**:
  ```bash
  make certificate name=<nombre-de-la-clave> email=<tu-email>
  ```
  Ejemplo:
  ```bash
  make certificate name=github-key email=tuemail@example.com
  ```
  Esto generará una clave SSH en el directorio `~/.ssh` con el nombre especificado y mostrará la clave pública para que puedas añadirla a tu cuenta de GitHub.

---

## Ejecución de Seguridad con Makefile ⚙️

El archivo Makefile incluye comandos para ejecutar las herramientas de seguridad y generar certificados SSH. A continuación, se describen los comandos disponibles:

### 1. **Análisis de Seguridad**
Ejecuta un análisis de seguridad completo utilizando Snyk y Gitleaks.

```bash
make security
```

### 2. **Generación de Certificados SSH**
Genera una clave SSH para autenticarte con servicios como GitHub.

```bash
make certificate name=<nombre-de-la-clave> email=<tu-email>
```

Ejemplo:
```bash
make certificate name=github-key email=tuemail@example.com
```

Esto generará una clave SSH en el directorio `~/.ssh` con el nombre especificado y mostrará la clave pública.

---

## Ejemplo de Flujo de Seguridad 🛠️

1. **Instalar herramientas de seguridad**:
  ```bash
  make install-appsec-tools
  ```

2. **Ejecutar análisis de seguridad**:
  ```bash
  make security
  ```

3. **Generar una clave SSH para GitHub**:
  ```bash
  make certificate name=github-key email=tuemail@example.com
  ```

4. **Añadir la clave pública a tu cuenta de GitHub**:
  Copia la clave pública generada (mostrada en la terminal) y añádela a tu cuenta de GitHub en la sección **SSH and GPG keys**.

---

## Configuración Adicional

### Variables de Entorno
El script `install-appsec-tools.sh` configura automáticamente las siguientes variables de entorno en tu archivo `~/.bashrc`:

- **`PATH`**: Incluye el directorio `~/.local/bin` donde se instalan las herramientas.
- **`SNYK_TOKEN`**: Token de autenticación para Snyk CLI. Si no se proporciona, se configura con el valor `replace_me`.

Para asegurarte de que las variables están configuradas correctamente, ejecuta:
```bash
source ~/.bashrc
```

---

Con estas herramientas y configuraciones, puedes garantizar un desarrollo seguro y seguir las mejores prácticas de DevSecOps. ¡Asegúrate de integrarlas en tu flujo de trabajo! 🚀
