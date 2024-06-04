# Registro de asteroides potencialmente peligrosos

La NASA necesita un sistema para registrar y monitorizar asteroides que podrían ser una amenaza para la Tierra.

El sistema debe:

* Almacenar información sobre asteroides: incluyendo nombre, tamaño, composición, etc.
* Registrar la distancia del asteroide a la Tierra en diferentes fechas.
* Permitir la consulta de información sobre asteroides específicos.
* Permitir la actualización de la información de un asteroide.
* Eliminar información de asteroides que ya no son una amenaza.

## Endpoints

### 1. Crear asteroide

* Ruta: `/api/v1/asteroides`
* Descripción: Crea un nuevo registro del asteroide.
* Método: POST
* Entrada: JSON con la información del asteroide:
  * name: string (nombre del asteroide)
  * diameter: float (diámetro en kilómetros)
  * discovery_date: string (fecha de descubrimiento en formato DD-MM-YYYY)
  * observations: string (opcional, información adicional)
  * distances: array (opcional, lista de objetos con date y distance)

```json
{
  "name": "Apophis",
  "diameter": 370,
  "discovery_date": "13-04-2004",
  "observations": "Leve cambio de trayectoria",
  "distances": []
}
```

* Salida: JSON con la información del asteroide creado, incluyendo su ID

```json
{
  "id": "12345",
  "name": "Apophis",
  "diameter": 370,
  "discovery_date": "13-04-2004",
  "observations": "Leve cambio de trayectoria",
  "distances": []
}
```

### 2. Obtener asteroides

* Ruta: `/api/v1/asteroides`
* Método: GET
* Salida: Lista de JSON con la información de todos los asteroides

```json
[
  {
    "id": "12345",
    "name": "Apophis",
    "diameter": 370,
    "discovery_date": "13-04-2004",
    "observations": "Leve cambio de trayectoria",
    "distances": []
  },
  {
    "id": "67890",
    "nombre": "Bennu",
    "diametro": 492,
    "discovery_date": "01-09-1999",
    "observations": "",
    "distances": []
  }
]
```

### 3. Obtener asteroide por ID

* Ruta: `/api/v1/asteroides/{id}`
* Método: GET
* Salida: JSON con la información del asteroide con el ID especificado

```json
{
  "id": "12345",
  "name": "Apophis",
  "diameter": 370,
  "discovery_date": "13-04-2004",
  "observations": "Leve cambio de trayectoria",
  "distances": [
    {
      "date": "2023-01-01",
      "distance": 0.044
    },
    {
      "date": "2024-04-13",
      "distance": 0.031
    }
  ]
}
```
La distancia se mide en UA que es la unidad astronómica (UA). Es una unidad de longitud igual, por definición, a 149 597 870 700 m (149,59 millones de km) que equivale aproximadamente a la distancia media entre la Tierra y el Sol.

### 4. Actualizar asteroide

* Ruta: `/api/v1/asteroides/{id}`
* Método: PATCH
* Entrada: JSON con la información actualizada del asteroide

```json
{
  "name": "Apophis 2.0"
}
```

* Salida: JSON con la información del asteroide actualizado

```json
{
  "id": "12345",
  "name": "Apophis 2.0",
  "diameter": 370,
  "discovery_date": "13-04-2004",
  "observations": "Leve cambio de trayectoria",
  "distances": []
}
```

### 5. Eliminar asteroide

* Ruta: `/api/v1/asteroides/{id}`
* Método: DELETE
* Salida: No hay salida

```json
{
  "message": "Asteroide eliminado exitosamente"
}
```

## Consideraciones adicionales

* Se pueden agregar más endpoints para funcionalidades adicionales paginación, filtros.
* Se puede implementar un sistema de autenticación y autorización para controlar el acceso a la API.
* Se puede implementar un sistema de caché para mejorar el rendimiento del sistema.
* Se puede validar la entrada de datos para evitar errores y datos incorrectos.

Podeis encontrar información sobre asteroides en la base de datos de la NASA [https://ssd.jpl.nasa.gov](https://ssd.jpl.nasa.gov/tools/sbdb_query.html)


## Stack

* Go [https://go.dev](https://go.dev)
* Docker [https://www.docker.com](https://www.docker.com)
* NoSQL
* OpenAPI [https://www.openapis.org](https://www.openapis.org)
* Github
* CI/CD

### Valorable

* Monitoring: uso de herramientas de trazas, logging y monitorización de codigos de estado.
* Frontal: frontal para la visualización de los datos y el trabajo con el CRUD.
* Testing: incluir test unitarios.