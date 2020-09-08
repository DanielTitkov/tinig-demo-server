.PHONY: run
run: 
	go run cmd/tinig/main.go

.PHONY: db
db:
	cd deployments/dev && docker-compose up -d

.PHONY: entgen
entgen:
	cd internal/repository/entgo && go generate ./ent