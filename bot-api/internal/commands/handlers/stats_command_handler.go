package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/constants/messages"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	"github.com/davidPardoC/budbot/internal/users/usecases"

	userModels "github.com/davidPardoC/budbot/internal/users/models"
)

type StatsCommandHandler struct {
	telegramService services.ITelegramService
	userUseCases    usecases.IUserUseCases
}

func NewStatsCommandHandler(telegramServices services.ITelegramService, userUseCases usecases.IUserUseCases) *StatsCommandHandler {
	return &StatsCommandHandler{
		telegramService: telegramServices,
		userUseCases:    userUseCases,
	}
}

func (h *StatsCommandHandler) HandleCommand(chatID int64, args []string) {
	userStats, err := h.userUseCases.GetCurrentMothStats(chatID)

	if err != nil {
		payload := builders.NewTelegramMessageBuilder(chatID).SetText(messages.ErrorGettingStats).Build()
		h.telegramService.SendMessage(payload)
		return
	}

	message := h.CreateStatsMessage(*userStats)
	payload := builders.NewTelegramMessageBuilder(chatID).SetText(message).SetParseMode("Markdown").Build()
	h.telegramService.SendMessage(payload)
}

func (h *StatsCommandHandler) ValidateArgs(args []string) bool {
	return true
}

func (h *StatsCommandHandler) CreateStatsMessage(stats userModels.UserStats) string {
	currentMoth := time.Now().Month()

	message := fmt.Sprintf(`This are your stats for the current month of %s:

	Total income: $%.1f
	Spent: $%.1f
	`,
		currentMoth.String(),
		stats.TotalIncome,
		stats.Spent,
	)

	if stats.Budget == 0 {
		message = message + fmt.Sprintf(`%s
		`, messages.NoBudgetMessage)
	} else {
		progressBar := generateProgressBar(stats.Spent, stats.Budget)

		message = message + fmt.Sprintf(`Budget: $%.2f
		
		Spent Percentage: *%.1f %%*
		%s`,
			stats.Budget,
			stats.SpentPercentage,
			progressBar)
	}

	return message
}

func generateProgressBar(spent, total float64) (progressBar string) {

	green := "🟩"
	white := "⬜️"
	yellow := "🟨"
	red := "🟥"
	upperTreshold := 0.8
	lowerTreshold := 0.6

	progressIcon := green

	progress := spent / total
	barLength := 10
	completed := int(progress * float64(barLength))

	if progress > lowerTreshold {
		progressIcon = yellow
	}
	if progress > upperTreshold {
		progressIcon = red
	}

	if completed > barLength {
		completed = barLength
	}

	progressBar = strings.Repeat(progressIcon, completed) + strings.Repeat(white, barLength-completed)

	if progress > 1 {
		progressBar = progressBar + "💀"
	}

	return
}
