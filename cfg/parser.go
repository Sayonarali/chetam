package cfg

import (
	"fmt"

	env "github.com/caarlos0/env/v11"
)

// Parse парсит переданный конфиг на основе тегов вида `env:`
// Если cfg - nil, либо не является указателем на структуру - функция возвращает ошибку
func Parse(cfg any) error {
	if cfg == nil {
		return fmt.Errorf("error: cfg must be a non-nil value")
	}
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	return nil
}
