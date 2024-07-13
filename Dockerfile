FROM golang:1.22-alpine3.19 AS builder

# Copiar os arquivos
COPY ${PWD} /app
WORKDIR /app

# Realizar o build da aplicação
RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o /app/appbin cmd/api/main.go

# Imagem para expor a aplicação
FROM alpine:3.19
LABEL MAINTAINER Samir Mamude <mamude@gmail.com>

# Instalar certificados
RUN apk --update add ca-certificates && rm -rf /var/cache/apk/*

# Adicionar novo usuário
RUN adduser -D appuser
USER appuser

# Configurar diretório da aplicação
COPY --from=builder --chown=appuser:appuser /app /home/appuser/app
RUN chmod -R 755 /home/appuser/app
WORKDIR /home/appuser/app

# Setar ambiente
ENV TZ=America/Sao_Paulo
ENV GIN_MODE=release
ENV PORT=8080
ENV TEMPLATE=/home/appuser/app/cmd/web/templates/*
ENV ASSETS=/home/appuser/app/cmd/web/assets
ENV TEMP_FILES=/home/appuser/app/temp

# Configurar ports
EXPOSE 8080

# Inicializar container
CMD [ "./appbin" ]