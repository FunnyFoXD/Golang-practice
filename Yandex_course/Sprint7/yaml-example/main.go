package main

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v3" // импортируем пакет для работы с YAML
)

// Artist содержит данные об артисте
type Artist struct {
    ID    int      `yaml:"id"`
    Name  string   `yaml:"name"`
    Genre string   `yaml:"genre"`
    Songs []string `yaml:"songs"`
}

func main() {
    yamlFile, err := os.ReadFile("artist.yaml")
    if err != nil {
        fmt.Printf("ошибка при чтении файла: %s", err.Error())
        return
    }

    fmt.Println(yamlFile)
    fmt.Println(string(yamlFile))

    var artist Artist

    err = yaml.Unmarshal(yamlFile, &artist)
    if err != nil {
        fmt.Printf("ошибка при десериализации: %s", err.Error())
        return
    }

    fmt.Println("ID:", artist.ID)
    fmt.Println("Name:", artist.Name)
    fmt.Println("Genre:", artist.Genre)
    fmt.Printf("Songs: %v\n", artist.Songs)
}