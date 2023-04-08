Se requiere:
Realizar un CRUD (create, read, update, delete) en una base de datos Postgres:
1. Usando algún ORM (gorm,bun, etc..).
2. Definiendo el modelo de datos como estructuras de Go.
3. Implementar un sistema de migraciones que permita crear la tabla en la
base de datos
4. Las conexiones a la base de datos deben definirse como una variable de
entorno

Restricciones de negocio:
* No puede eliminar medidores que estén instalados en este momento. (check)
* No puede crear dos medidores con el mismo serial-marca (check)

* Solo puede existir un medidor por predio (En un mismo periodo de tiempo)

* No puede cambiar: la fecha de instalación, serial, marca (check)
Realizar una API en HTTP o preferiblemente en gRPC para interactuar con la
aplicación:
1. En el caso de ser gRPC debe crear el archivo .proto y adjuntarlo al
repositorio
2. En el caso de ser http agregar documentación al API con Swagger
(Swaggo)

Realice un endpoint que retorne los medidores instalados actualmente que
tengan cortado o inactivo el servicio de energía.
Realice un endpoint que a partir de un serial y marca retorne la instalación más
reciente.

Publicar todos los eventos de creación de registros de la base de datos en un
stream de Redis Stream de forma asíncrona. (Si hizo la implementación de gRPC
usar streams para escuchar los eventos de redis).

Agregar casos de pruebas para las restricciones.