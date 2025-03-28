# catalog
# catalog
# catalog
# catalog

```
catalog
├─ .air.toml
├─ Dockerfile
├─ Makefile
├─ README.md
├─ cmd
│  └─ main.go
├─ config
│  └─ config.go
├─ docker-compose.yaml
├─ go.mod
├─ go.sum
├─ internal
│  ├─ bootstrap
│  │  └─ bootstrap.go
│  ├─ category
│  │  ├─ category.go
│  │  ├─ category_handler.go
│  │  ├─ category_repository.go
│  │  └─ category_service.go
│  ├─ db
│  │  └─ mongo_db
│  │     ├─ db.go
│  │     └─ mongo_repository
│  │        └─ crud_repository.go
│  ├─ generics
│  │  ├─ handler
│  │  │  └─ crud_handler.go
│  │  └─ service
│  │     └─ crud_service.go
│  ├─ infrastructure
│  │  └─ storage
│  │     └─ minio.go
│  ├─ middleware
│  ├─ product
│  │  ├─ product.go
│  │  ├─ product_handler.go
│  │  ├─ product_repository.go
│  │  └─ product_service.go
│  ├─ responses
│  │  ├─ helper.go
│  │  └─ response.go
│  ├─ routes
│  │  ├─ category_routes.go
│  │  ├─ product_routes.go
│  │  └─ routes.go
│  ├─ services
│  │  └─ upload_service.go
│  ├─ utils
│  │  └─ mql_request_filter
│  │     ├─ filter_from_request.go
│  │     ├─ filter_validator.go
│  │     └─ pagination.go
│  └─ watchers
│     └─ category_watcher.go
├─ minio-data
├─ pkg
├─ scripts
│  └─ gen.go
├─ templates
│  ├─ handler.tmpl
│  ├─ model.tmpl
│  ├─ repository.tmpl
│  └─ service.tmpl
└─ tests
   ├─ handlers
   ├─ repositories
   ├─ routes
   └─ services

```