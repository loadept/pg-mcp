# pg-mcp

**Versión:** 0.1.0  
**Lenguaje:** Go 1.25+  
**Tipo:** Ejecutable binario

## Descripción
Servidor MCP para consultas PostgreSQL de solo lectura.

## Ejecución
```bash
# Linux/macOS
./pg-mcp

# Windows
pg-mcp.exe
```

## Variables de Entorno
```bash
POSTGRES_URI="postgres://usuario:contraseña@localhost:5432/basedatos?sslmode=disable"
```

## Herramientas
- `execute_query`: Ejecuta consultas SELECT (máx. 50 filas)
- `get_table_info`: Información de estructura de tabla
- `list_tables`: Lista tablas por esquema (paginado)

## Compilación
```bash
go build -o pg-mcp cmd/mcp/main.go
```
