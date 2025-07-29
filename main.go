package main

import "github.com/joho/godotenv"

func main() {
    godotenv.Load() // Carrega .env local
    connectToDB()
    ...
}
