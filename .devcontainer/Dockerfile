# Usar la imagen base de Go
FROM golang:1.21.1

# Evitar preguntas de la interfaz de usuario durante la instalación
ARG DEBIAN_FRONTEND=noninteractive

# Actualizar los paquetes e instalar fish
RUN apt-get update && apt-get install -y fish

# Cambiar la shell predeterminada a fish para el usuario root
RUN chsh -s /usr/bin/fish root

# Instalar herramientas de desarrollo y análisis de Go
RUN go install golang.org/x/tools/gopls@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && go install github.com/securego/gosec/v2/cmd/gosec@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest
