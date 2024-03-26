package cli

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(1).Bold(true).MarginBottom(1).MarginTop(1).PaddingLeft(1).PaddingRight(1)
	listTitleStyle    = lipgloss.NewStyle().MarginLeft(0).Bold(true).PaddingLeft(0)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(2)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#FFD700"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(2).PaddingBottom(1).Foreground(lipgloss.Color("#666"))
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)
