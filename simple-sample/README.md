Para Docker:

Crear Dockerfile
Crear docker-compose.yml
Correr:
- go mod init nombre_proyecto
- go mod tidy
- go.sum (en caso de existir modulos)

Pasos para la levantar la aplicación:
- docker-compose build
- docker-compose up -d
- docker-compose down


Fuentes:

- https://devjaime.medium.com/gu%C3%ADa-definitiva-para-la-organizaci%C3%B3n-de-proyectos-en-golang-65d666e06ab4