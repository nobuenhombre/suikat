package fifo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFifoCreate(t *testing.T) {
	fifoName := "/tmp/fifo-test"

	fifo := New(fifoName)

	err := fifo.Create()
	require.NoError(t, err)

	exist, err := fifo.IsExists()
	require.NoError(t, err)
	require.True(t, exist)

	err = fifo.Delete()
	require.NoError(t, err)
}

func TestFifoWriteAndRead(t *testing.T) {
	fifoName := "/tmp/fifo-test"

	fifo := New(fifoName)

	err := fifo.Create()
	require.NoError(t, err)

	exist, err := fifo.IsExists()
	require.NoError(t, err)
	require.True(t, exist)

	fifoWriter := New(fifoName)
	err = fifoWriter.OpenToWrite()
	require.NoError(t, err)

	fifoReader := New(fifoName)
	err = fifoReader.OpenToRead()
	require.NoError(t, err)

	go func() {
		for i := 1; i <= 15; i++ {
			err = fifoWriter.Write(fmt.Sprintf("%d", i))
			require.NoError(t, err)
		}

		err = fifoWriter.Close()
		require.NoError(t, err)
	}()

	rcv := func(msg string) error {
		t.Logf("%s", msg)
		return nil
	}

	err = fifoReader.Read(rcv)
	require.NoError(t, err)

	err = fifoReader.Close()
	require.NoError(t, err)

	err = fifo.Delete()
	require.NoError(t, err)
}
