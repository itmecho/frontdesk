.PHONY: build

build: build-migrations
	go build -o frontdesk ./cmd/frontdesk

test:
	go test $(FLAGS) ./...

clean:
	rm -f frontdesk

local: build
	docker-compose up -d
	./frontdesk --auth-secret 9ucypsd58P9M3GqxewSuAA8H26auDVUk

migration:
	@test -n "$(name)" || { echo "'name' variable must be set"; exit 1; }
	@test -d pkg/store/postgres/migrations || mkdir -p pkg/store/postgres/migrations
	@touch pkg/store/postgres/migrations/$(shell date +%s)_$(name).{up,down}.sql

build-migrations:
	go install github.com/go-bindata/go-bindata/...
	cd pkg/store/postgres/migrations; \
	$(GOPATH)/bin/go-bindata -pkg migrations -o migrations.go .
