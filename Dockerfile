# Usa una imagen base oficial de Go para la construcción
FROM golang:1.21 as builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente al directorio de trabajo
COPY . .

# Compila la aplicación Go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Usa una imagen base ligera para la imagen final
FROM alpine:latest

# Instala ca-certificates en caso de que tu aplicación haga llamadas HTTPS
RUN apk --no-cache add ca-certificates

# Copia el ejecutable desde el contenedor de compilación
COPY --from=builder /app/myapp /myapp

# Copia el módulo Go y sus dependencias
COPY --from=builder /app/go.mod ./go.mod

# Copia los directorios 'view', 'static' y 'db' al contenedor
COPY --from=builder /app/views /views
COPY --from=builder /app/static /static
COPY --from=builder /app/db 	/db

# Expone el puerto (debe coincidir con la variable de entorno PORT)
EXPOSE 8080

# Ejecuta la aplicación
CMD ["/myapp"]
