package fifo

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/nobuenhombre/suikat/pkg/ge"
	"io"
	"os"
	"syscall"
)

type conn struct {
	Name string
	f    *os.File
}

type MessageReceiver func(string) error

type Service interface {
	Create() error
	IsExists() (bool, error)
	Delete() error
	OpenToRead() error
	OpenToWrite() error
	Close() error
	Write(msg string) error
	Read(rcv MessageReceiver) error
}

func New(name string) Service {
	return &conn{
		Name: name,
	}
}

func (c *conn) Create() error {
	err := syscall.Mkfifo(c.Name, 0666)
	if err != nil {
		return ge.Pin(err, ge.Params{"name": c.Name})
	}

	return nil
}

func (c *conn) IsExists() (bool, error) {
	_, err := os.Stat(c.Name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, ge.Pin(err, ge.Params{"name": c.Name})
	}

	return true, nil
}

func (c *conn) Delete() error {
	if c.f != nil {
		err := c.Close()
		if err != nil {
			return ge.Pin(err, ge.Params{"name": c.Name})
		}
	}

	isExists, err := c.IsExists()
	if err != nil {
		return ge.Pin(err, ge.Params{"name": c.Name})
	}

	if !isExists {
		return nil
	}

	err = os.Remove(c.Name)
	if err != nil {
		return ge.Pin(err)
	}

	return nil
}

func (c *conn) OpenToRead() error {
	file, err := os.OpenFile(c.Name, os.O_CREATE, 0666)
	if err != nil {
		return ge.Pin(err, ge.Params{"name": c.Name})
	}

	c.f = file

	return nil
}

func (c *conn) OpenToWrite() error {
	file, err := os.OpenFile(c.Name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return ge.Pin(err, ge.Params{"name": c.Name})
	}

	c.f = file

	return nil
}

func (c *conn) Close() error {
	if c.f == nil {
		return ge.New("fifo file is not opened")
	}

	err := c.f.Close()
	if err != nil {
		return ge.Pin(err, ge.Params{"name": c.Name})
	}

	c.f = nil

	return nil
}

func (c *conn) Write(msg string) error {
	if c.f == nil {
		return ge.New("fifo file is not opened")
	}

	_, err := c.f.WriteString(fmt.Sprintf("%s\n", msg))
	if err != nil {
		return ge.Pin(err, ge.Params{"name": c.Name, "msg": msg})
	}

	return nil
}

func (c *conn) Read(rcv MessageReceiver) error {
	if c.f == nil {
		return ge.New("fifo file is not opened")
	}

	reader := bufio.NewReader(c.f)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return ge.Pin(err, ge.Params{"name": c.Name})
		}

		err = rcv(string(line))
		if err != nil {
			return ge.Pin(err, ge.Params{"name": c.Name, "line": string(line)})
		}
	}

	return nil
}
