# Dragonball API

El proyecto disponibiliza un servicio en Golang con una base de datos SQLite que permite consultar
nombres de personajes de Dragon Ball, los cuales se van a buscar a: https://web.dragonball-api.com/
y se guardan localmente en la DB para ser consultados localmente para futuras consultas.

## Instalación

El proyecto esta dockerizado por lo que puede ser levantado y disponibilizado utilizando el siguiente comando:
```
docker compose up -d 
```

## Ejecución

Puede ser utilizada desde Postman importando y ejecutando el siguiente cURL:
```
curl --location --request POST 'http://localhost:8080/character?name=majin'
```


## test

Teniendo Go 1.23 instalado se puede ejecutar el siguiente comando:
```
go test ./..
```
