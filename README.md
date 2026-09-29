# Chat Cliente-Servidor (Proyecto1 de Modelado y Programación)

Sistema de comunicación en tiempo real en consola desarrollado con una arquitectura distribuida cliente-servidor.

## Requisitos del Sistema

Para compilar y ejecutar el proyecto se requiere contar con las siguientes herramientas instaladas:

* **Servidor:** Go (Golang) v1.18 o superior.
* **Cliente:** SDK de .NET (C# 10+) y el sistema de construcción **Meson** (junto con **Ninja**).

## Instrucciones de Compilación y Ejecución

### 1. Servidor (Golang)

Navega a la carpeta del servidor e inícialo directamente con el entorno de Go:

# Para compilar:
     go run . 1234 --Número de puerto

### 2. Cliente (C#)

Navega a la carpeta del cliente e utiliza Meson para compilar el proyecto:

# Para compilar:
     meson setup build	
     meson compile -C build
    ./build/cliente.exe -i localhost -p 1234 --Dirección ip y puerto

