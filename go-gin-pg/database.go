package main

func PostgresConnect() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "ptpit",
		Addr:     "localhost:5432",
		Database: "coti",
	})
}
