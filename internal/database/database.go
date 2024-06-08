package database






func main() {
	var err error

	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln("Failed to retrieve configs:", err)
	}

	connStr := fmt.Sprintf("port=%s host=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DB_PORT, config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	} else {
		log.Println("Successfully connected to the database")
	}

}


