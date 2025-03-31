# Restaurant Management System

Este es un sistema de gestión de restaurantes desarrollado en Go. El objetivo de este proyecto es proporcionar una solución completa para la gestión de restaurantes, incluyendo la administración de menús, reservas y pedidos.

## Estructura del Proyecto

El proyecto está organizado de la siguiente manera:

```
restaurant-management-system
├── cmd
│   └── main.go             # Punto de entrada de la aplicación
├── config
│   └── config.yaml         # Archivo de configuración del sistema
├── internal
│   ├── handlers
│   │   └── handler.go      # Controladores para manejar la lógica de negocio
│   ├── models
│   │   └── model.go        # Definición de entidades y estructuras de datos
│   ├── services
│   │   └── service.go      # Lógica de aplicación y reglas de negocio
│   └── repository
│       └── repository.go   # Gestión del acceso a la base de datos
├── pkg
│   └── utils
│       └── utils.go        # Funciones utilitarias
├── go.mod                  # Módulo de Go para gestionar dependencias
└── README.md               # Documentación del proyecto
```

## Instalación

1. Clona el repositorio:
   ```
   git clone <URL_DEL_REPOSITORIO>
   cd restaurant-management-system
   ```

2. Instala las dependencias:
   ```
   go mod tidy
   ```

## Uso

Para ejecutar la aplicación, utiliza el siguiente comando:

```
go run cmd/main.go
```

Esto iniciará un servidor HTTP en el puerto configurado en `config/config.yaml`.

## Contribuciones

Las contribuciones son bienvenidas. Si deseas contribuir, por favor abre un issue o envía un pull request.

## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.