package main

import "sync"

// Тестовая структура с мьютексами
type TestStruct struct {
	sync.Mutex            // должна быть ошибка
	mu         sync.Mutex // должна быть ошибка
}
