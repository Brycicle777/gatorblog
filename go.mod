module github.com/Brycicle777/gatorblog

go 1.23.6

replace github.com/Brycicle777/gatorblog => ./gatorblog

replace internal/config v0.0.0 => ./internal/config

require internal/config v0.0.0

replace internal/database v0.0.0 => ./internal/database

require (
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
	internal/database v0.0.0
)
