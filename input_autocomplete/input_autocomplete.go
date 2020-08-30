package input_autocomplete

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func keyboardListener(input *Input) error {
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			return err
		}

		switch key {
		case keyboard.KeyEnter:
			fmt.Println("")
			return nil
		case keyboard.KeyArrowLeft:
			input.MoveCursorLeft()
		case keyboard.KeyArrowRight:
			input.MoveCursorRight()
		case keyboard.KeyBackspace:
			input.RemoveChar()
			input.Print()
			input.MoveCursorLeftTo(len(input.currentText) - input.cursor.GetPosition())
		case keyboard.KeyBackspace2:
			input.RemoveChar()
			input.Print()
			input.MoveCursorLeftTo(len(input.currentText) - input.cursor.GetPosition())
		default:
			c := string(char)
			input.AddChar(c)
			fmt.Print(c)
		}
	}
}

func Read(text string) (string, error) {
	if err := keyboard.Open(); err != nil {
		return "", err
	}

	defer keyboard.Close()

	input := NewInput(text)

	input.Print()

	if err := keyboardListener(input); err != nil {
		return "", err
	}

	return input.GetCurrentText(), nil
}
