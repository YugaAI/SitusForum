export MYSQL_Url='mysql://root:situsForum@tcp(localhost:3306)/SitusForum'

migrate-create:
	@migrate create -ext sql -dir script/migrations -seq $(name)

migrate-up:
	@migrate -database $(MYSQL_Url)  -path script/migrations up

migrate-down:
	@migrate -database $(MYSQL_Url)  -path script/migrations down

migrate-force:
	@migrate -database $(MYSQL_Url)  -path script/migrations force $(version)

migrate-version:
	@migrate -database $(MYSQL_Url)  -path script/migrations version