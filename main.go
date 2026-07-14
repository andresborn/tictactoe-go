package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	board     [9]string
	turn      int
	whoseTurn string // Player 1, Player 2
	position  int
	winnner   string
	gameOver  bool
}

func checkWhoseTurn(turn int) string {

	if turn%2 == 0 {
		return "Player 2"
	}
	return "Player 1"
}

func checkWin(board [9]string) string {
	// Winning conditions: same chip type in indexes 012, 345, 678, 036, 147, 258, 048, 246
	winner := ""
	switch {
	case board[0] == board[1] && board[1] == board[2]:
		if board[0] == " " {
			winner = ""
		} else {
			winner = board[0]
		}
	case board[3] == board[4] && board[4] == board[5]:
		if board[3] == " " {
			winner = ""
		} else {
			winner = board[3]
		}
	case board[6] == board[7] && board[7] == board[8]:
		if board[6] == " " {
			winner = ""
		} else {
			winner = board[6]
		}
	case board[0] == board[3] && board[3] == board[6]:
		if board[0] == " " {
			winner = ""
		} else {
			winner = board[0]
		}
	case board[1] == board[4] && board[4] == board[7]:
		if board[1] == " " {
			winner = ""
		} else {
			winner = board[1]
		}
	case board[2] == board[5] && board[5] == board[8]:
		if board[2] == " " {
			winner = ""
		} else {
			winner = board[2]
		}
	case board[0] == board[4] && board[4] == board[8]:
		if board[0] == " " {
			winner = ""
		} else {
			winner = board[0]
		}
	case board[2] == board[4] && board[4] == board[6]:
		if board[2] == " " {
			winner = ""
		} else {
			winner = board[2]
		}
	}

	return winner
}

func initialModel() model {
	return model{
		board:     [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "},
		turn:      1,
		whoseTurn: "Player 1",
		position:  0,
		gameOver:  false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.position < 8 {
				m.position += 1
			}
		case "down":
			if m.position > 0 {
				m.position -= 1
			}
		case "enter":
			if m.gameOver {
				return m, nil
			}
			// place chip, check if won, next turn.
			if m.board[m.position] != " " {
				return m, nil
			}
			// Whose chip?
			chip := " "
			if checkWhoseTurn(m.turn) == "Player 1" {
				chip = "X"
			} else {
				chip = "O"
			}
			m.board[m.position] = chip

			// Did I win yet?

			winner := checkWin(m.board)

			if winner != "" {
				if winner == "X" {
					m.winnner = "Player 1"
				} else {
					m.winnner = "Player 2"
				}
				m.gameOver = true
				return m, nil
			}
			// Next turn...
			m.turn++
			m.whoseTurn = checkWhoseTurn(m.turn)
			m.position = 0 // reset position
		}

	}

	return m, nil
}

func (m model) View() tea.View {
	s := "\nWelcome to the game of Tic Tac Toe!!!\n\n"
	s += "Each player will select where they'll place their\n"
	s += "chip on the board using the up and down arrows.\n"

	s += "\nThese are the positions of the board:\n\n"

	s += "  0 | 1 | 2 \n"
	s += " ---|---|---\n"
	s += "  3 | 4 | 5 \n"
	s += " ---|---|---\n"
	s += "  6 | 7 | 8 \n\n"

	s += fmt.Sprintf("Turn: %v\n", m.turn)
	s += fmt.Sprintf("Who's playing: %s\n\n", m.whoseTurn)
	s += fmt.Sprintf("Position to play: %v\n\n", m.position)

	if m.winnner != "" {
		s += fmt.Sprintf("We have a WINNER: %s!!!\n\n", m.winnner)
	}

	s += fmt.Sprintf("  %s | %s | %s  \n", m.board[0], m.board[1], m.board[2])
	s += " ---|---|---\n"
	s += fmt.Sprintf("  %s | %s | %s  \n", m.board[3], m.board[4], m.board[5])
	s += " ---|---|---\n"
	s += fmt.Sprintf("  %s | %s | %s  \n", m.board[6], m.board[7], m.board[8])

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
