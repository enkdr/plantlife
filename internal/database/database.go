package database


func main() {
	var err error

	if len(os.Args) >= 1 {
		err = godotenv.Load(".env.dev")
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}
	}

	connStr := fmt.Sprintf("port=%s host=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	fmt.Println(connStr)

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


