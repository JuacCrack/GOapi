# Contact API

Esta es una API RESTful para gestionar contactos. La API utiliza Go con el framework Gin y sigue una arquitectura MVC. La documentación y pruebas de la API se realizan con Swagger.

## Características

- CRUD completo para contactos
- Documentación y pruebas con Swagger
- Arquitectura MVC
- Almacenamiento en memoria

## Estructura del Proyecto

ContactAPI/
├── cmd/
│ └── server.go
├── config/
│ └── config.go
├── controllers/
│ └── contact_controller.go
├── docs/
│ └── docs.go
│ └── swagger.json
│ └── swagger.yaml
├── models/
│ └── contact.go
│ └── error_response.go
├── routes/
│ └── routes.go
├── main.go
└── go.mod


## Instalación

1. Clona el repositorio:
    ```sh
    git clone https://github.com/JuacCrack/GOapi.git
    cd GOapi
    ```

2. Instala las dependencias:
    ```sh
    go mod tidy
    ```

3. Genera la documentación de Swagger:
    ```sh
    swag init
    ```

## Uso

1. Ejecuta el servidor:
    ```sh
    go run main.go
    ```

2. Abre tu navegador y ve a `http://localhost:8080/swagger/index.html` para ver la documentación de Swagger y probar los endpoints.

## Endpoints

### Obtener todos los contactos

- **URL**: `/contacts`
- **Método**: `GET`
- **Respuesta**: `200 OK` con un array de contactos

### Obtener un contacto por ID

- **URL**: `/contacts/{id}`
- **Método**: `GET`
- **Parámetros**: `id` (path)
- **Respuesta**:
  - `200 OK` con el contacto solicitado
  - `404 Not Found` si no se encuentra el contacto

### Crear un nuevo contacto

- **URL**: `/contacts`
- **Método**: `POST`
- **Cuerpo**: JSON con los datos del contacto
- **Respuesta**:
  - `201 Created` con el contacto creado
  - `400 Bad Request` si la entrada es inválida

### Actualizar un contacto

- **URL**: `/contacts/{id}`
- **Método**: `PUT`
- **Parámetros**: `id` (path)
- **Cuerpo**: JSON con los datos actualizados del contacto
- **Respuesta**:
  - `200 OK` con el contacto actualizado
  - `400 Bad Request` si la entrada es inválida
  - `404 Not Found` si no se encuentra el contacto

### Eliminar un contacto

- **URL**: `/contacts/{id}`
- **Método**: `DELETE`
- **Parámetros**: `id` (path)
- **Respuesta**:
  - `200 OK` con un mensaje de confirmación
  - `404 Not Found` si no se encuentra el contacto

## Documentación de Swagger

La API está documentada usando Swagger. Puedes acceder a la documentación generada navegando a `http://localhost:8080/swagger/index.html` después de iniciar el servidor.

## Gracias por leer!
