package util

import (
	"bufio"
	"errors"
	"os"
)

func FormatFile(inFilePath string, outFilePath string) (string, error) {
	formatter, err := NewFormatter(inFilePath, outFilePath)
	if err != nil {
		return "", err
	}
	defer formatter.Close()
	formatter.StartFormat()
	return "Done!", nil
}

type Formatter struct {
	inFilePath         string
	outFilePath        string
	inFile             *os.File
	outFile            *os.File
	currentIndentation int
	fileReader         *bufio.Reader
}

func NewFormatter(inFilePath string, outFilePath string) (*Formatter, error) {
	inFile, err := os.Open(inFilePath)
	if err != nil {
		return nil, err
	}
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(inFile)

	return &Formatter{
		inFilePath:  inFilePath,
		outFilePath: outFilePath,
		inFile:      inFile,
		outFile:     outFile,
		fileReader:  reader,
	}, nil
}

func (f *Formatter) Close() {
	f.inFile.Close()
	f.outFile.Close()
	f.fileReader = nil
}

func (f *Formatter) Next() (rune, error) {
	data, _, err := f.fileReader.ReadRune()
	if err != nil {
		return 0, err
	}
	return data, nil
}

func (f *Formatter) StartFormat() error {
	err := f.processElement()
	if err != nil {
		return err
	}
	return nil
}

func (f *Formatter) processElements() error {
	f.newLine()
	f.indent()

	err := f.processElement()
	if err != nil {
		return err
	}

	data, err := f.Next()
	if err != nil {
		return err
	}
	if data == ',' {
		_, err = f.outFile.Write([]byte(","))
		if err != nil {
			return err
		}
		return f.processElements()
	}

	if data == ']' {
		err = f.fileReader.UnreadRune()
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Formatter) processElement() error {
	err := f.consumeWhiteSpace()
	if err != nil {
		return err
	}
	err = f.processValue()
	if err != nil {
		return err
	}
	err = f.consumeWhiteSpace()
	if err != nil {
		return err
	}

	return err
}

func (f *Formatter) processValue() error {
	data, err := f.Next()
	if err != nil {
		return err
	}

	switch data {
	case '{':
		err = f.fileReader.UnreadRune()
		if err == nil {
			err = f.processObject()
		}
	case '[':
		f.fileReader.UnreadRune()
		err = f.processArray()
	case '"':
		f.fileReader.UnreadRune()
		err = f.processString()
	case 't':
		err = f.processAndMatch("rue", "true")
	case 'f':
		err = f.processAndMatch("alse", "false")
	case 'n':
		err = f.processAndMatch("ull", "null")
	default:
		f.fileReader.UnreadRune()
		err = f.processNumber()
	}

	return err
}

func (f *Formatter) processArray() error {
	data, err := f.Next()
	if err != nil {
		return err
	}
	if data != '[' {
		return errors.New("invalid character. expecting '['")
	}
	f.outFile.Write([]byte("["))

	err = f.consumeWhiteSpace()
	if err != nil {
		return err
	}
	data, _ = f.Next()
	if data == ']' {
		// empty array
		f.outFile.Write([]byte("]"))
		return nil
	}
	f.fileReader.UnreadRune()
	err = f.updateIndentation(1)
	if err != nil {
		return err
	}
	err = f.processElements()
	if err != nil {
		return err
	}

	data, err = f.Next()
	if err != nil {
		return err
	}
	if data != ']' {
		return errors.New("invalid character. expecting ']'")
	}
	f.newLine()
	err = f.updateIndentation(-1)
	if err != nil {
		return err
	}
	f.indent()
	_, err = f.outFile.Write([]byte("]"))
	if err != nil {
		return err
	}
	return err
}
func (f *Formatter) processObject() error {
	data, err := f.Next()
	if err != nil {
		return err
	}
	if data != '{' {
		return errors.New("invalid character. expecting '{'")
	}

	f.outFile.Write([]byte("{"))
	err = f.consumeWhiteSpace()
	if err != nil {
		return err
	}
	data, _ = f.Next()
	if data == '}' {
		// empty object
		f.outFile.Write([]byte("}"))
		return nil
	}
	f.fileReader.UnreadRune()
	err = f.updateIndentation(1)
	if err != nil {
		return err
	}
	err = f.processMembers()
	if err != nil {
		return err
	}
	data, err = f.Next()
	if err != nil {
		return err
	}
	if data != '}' {
		return errors.New("invalid character. expecting '}'")
	}
	f.updateIndentation(-1)
	f.newLine()
	f.indent()
	_, err = f.outFile.Write([]byte("}"))
	return err
}

func (f *Formatter) processMembers() error {
	err := f.newLine()
	if err != nil {
		return err
	}

	err = f.processMember()
	if err != nil {
		return err
	}

	data, err := f.Next()
	if err != nil {
		return err
	}
	if data == ',' {
		_, err = f.outFile.Write([]byte(","))
		if err != nil {
			return err
		}
		return f.processMembers()
	}

	if data == '}' {
		err = f.fileReader.UnreadRune()
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Formatter) processMember() error {
	err := f.consumeWhiteSpace()
	if err != nil {
		return err
	}
	f.indent()
	err = f.processString()
	if err != nil {
		return err
	}
	err = f.consumeWhiteSpace()
	if err != nil {
		return err
	}
	data, err := f.Next()
	if err != nil {
		return err
	}
	if data != ':' {
		return errors.New("invalid character. expecting ':'")
	}
	_, err = f.outFile.Write([]byte(": "))
	if err != nil {
		return err
	}
	return f.processElement()
}

func (f *Formatter) processString() error {
	data, err := f.Next()
	if err != nil {
		return err
	}

	if data != '"' {
		return errors.New("invalid character. expecting '\"'")
	}

	_, err = f.outFile.Write([]byte("\""))
	if err != nil {
		return err
	}

	lastRune := rune(0)
	// TODO: Not upto standard
	for {
		data, err = f.Next()
		if err != nil {
			return err
		}
		_, err = f.outFile.Write([]byte(string(data)))
		if err != nil {
			return nil
		}
		if lastRune != '\\' && data == '"' {
			break
		}
		lastRune = data
	}
	return nil
}

func (f *Formatter) indent() {
	for i := 0; i < f.currentIndentation; i++ {
		f.outFile.Write([]byte("  "))
	}
}

func (f *Formatter) newLine() error {
	_, err := f.outFile.Write([]byte("\n"))
	return err
}

func (f *Formatter) updateIndentation(delta int) error {
	f.currentIndentation += delta
	if f.currentIndentation < 0 {
		return errors.New("invalid indentation")
	}
	return nil
}

func (f *Formatter) processNumber() error {
	allowed := []rune{'-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 'e', 'E'}
	// TODO: This validation is not enough
	for {
		data, err := f.Next()
		if err != nil {
			return err
		}
		if IndexOf(allowed, data) == -1 {
			f.fileReader.UnreadRune()
			return nil
		}
		_, err = f.outFile.Write([]byte(string(data)))
		if err != nil {
			return err
		}
	}
}

func (f *Formatter) processAndMatch(match string, writeString string) error {
	for i := 0; i < len(match); i++ {
		data, err := f.Next()
		if err != nil {
			return err
		}
		if data != rune(match[i]) {
			return errors.New("invalid character")
		}
	}

	_, err := f.outFile.Write([]byte(writeString))
	if err != nil {
		return err
	}

	return nil
}

func (f *Formatter) consumeWhiteSpace() error {
	for {
		data, err := f.Next()
		if err != nil {
			return err
		}
		if data != '\u0020' && data != '\u000A' && data != '\u000D' && data != '\u0009' {
			return f.fileReader.UnreadRune()
		}
	}
}
