package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Введи имя домена:")
	var domain_name string
	fmt.Scanln(&domain_name)
	env_name := "DEFAULT_EMAIL"
	email, exists := os.LookupEnv(env_name)

	if !exists {
		fmt.Printf("Назначь переменную среды %s", env_name)
		os.Exit(0)
	}
	// Путь к файлу
	filePath := "/etc/nginx/nginx.conf"

	// Открываем файл для записи
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer file.Close()

	// Запись текста в файл
	_, err = file.WriteString(fmt.Sprintf(config, domain_name))
	if err != nil {
		log.Fatalf("Ошибка записи в файл: %v", err)
	}

	cmd := exec.Command(
		"certbot",
		"-d", domain_name,          // Укажите домен
		"--nginx",                   // Использование nginx
		"--non-interactive",         // Без взаимодействия с пользователем
		"--agree-tos",               // Согласие с условиями использования
		"-m", email, // Email администратора
	)
	cmd.Stderr = os.Stderr // перенаправил ошибки
	cmd.Run()

	cmd = exec.Command("nginx", "-s", "reload")
	cmd.Stderr = os.Stderr
	cmd.Run()
}
