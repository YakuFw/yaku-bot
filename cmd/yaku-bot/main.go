package main

import (
	"log"

	"github.com/YakuFw/yaku-bot/internal/bot"
)

func main() {
	log.Println("Iniciando Yaku Bot SSH VPN Manager...")

	// Iniciar servidor del bot (bloqueante)
	bot.StartBot()
}
