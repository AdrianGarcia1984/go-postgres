# INTRODUCCION

 Esta API se encarga de llevar el control de los medidores de energia de la empresa enerBit, donde se presenta el reto de registrar y hacer seguimiento de los medidores de energía que se han instalado en los predios de nuestros clientes, para efectos de pruebas en modo local se manejara la url localhost:3000

 Esta diseñada bajo Go, con los framework gorilla-mux para manejar la conexion, gorm para el ORM, y postgres como base de datos.

## preparar la instalacion

inicialmente se debe tener una imagen en docker o un servidor de postgres, necesita de estos elementos como variable de entorno para la conexion, en el archivo env.example, encontrara lo siguiente

~~~
HOST = ""
USER = "" 
PASSWORD = "" 
DBNAME = "" 
PORT = ""
~~~

se debe clonar este repositorio

~~~
$ git clone https://https://github.com/AdrianGarcia1984/go-postgres
~~~

tener instalado la ultima version de go, ademas de la ejecucion la base de datos
## USAGE

despues de descargado el repositorio, instalar los archivos necesarios de go, con el siguiente comando en la terminal de vscode

~~~ 
go mod init
~~~

crear las variables de entorno como se indico anteriormente


## inicar la ejecucion

se debe iniciar el proyecto con 

~~~
go run .
~~~

como se inidico inicialmente se utilizara para este caso el localhost, se documento el proyecto, conforme al requerimiento con SWAGGO, con lo cual se debe iniciar en un navegador la siguiente url
<a>http://localhost:3000/swagger/index.html</a>