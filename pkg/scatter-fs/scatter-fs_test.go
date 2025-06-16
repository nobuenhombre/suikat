package scatterfs

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestFileSystem тесты для реальной реализации
func TestFileSystem(t *testing.T) {
	tempDirs := make([]string, 3)
	for i := 0; i < 3; i++ {
		dir, err := os.MkdirTemp("", "testdir_*")
		assert.NoError(t, err)
		tempDirs[i] = dir
		defer os.RemoveAll(dir)
	}

	vfs := New(tempDirs, 10).(*FileSystem)

	t.Run("Create and Open file", func(t *testing.T) {
		path := "testfile.txt"

		file, err := vfs.Create(path)
		assert.NoError(t, err)
		assert.NotNil(t, file)
		file.Close()

		file, err = vfs.Open(path)
		assert.NoError(t, err)
		assert.NotNil(t, file)
		file.Close()
	})

	t.Run("MkdirAll and Remove", func(t *testing.T) {
		path := "dir1/dir2"

		// Создаем директории
		err := vfs.MkdirAll(path, 0755)
		assert.NoError(t, err)

		// Проверяем создание
		found := false
		for _, dir := range tempDirs {
			if _, err := os.Stat(filepath.Join(dir, path)); err == nil {
				found = true
				break
			}
		}
		assert.True(t, found)

		// Удаляем
		err = vfs.Remove(path)
		assert.NoError(t, err)
	})

	t.Run("RemoveAll", func(t *testing.T) {
		path := "dir3/dir4"

		// Создаем структуру директорий
		err := vfs.MkdirAll(path, 0755)
		assert.NoError(t, err)

		// Создаем файл внутри
		filePath := filepath.Join(path, "file.txt")
		file, err := vfs.Create(filePath)
		assert.NoError(t, err)
		file.Close()

		// Удаляем рекурсивно
		err = vfs.RemoveAll(path)
		assert.NoError(t, err)
	})

	t.Run("Stat", func(t *testing.T) {
		path := "statfile.txt"

		file, err := vfs.Create(path)
		assert.NoError(t, err)
		file.Close()

		info, err := vfs.Stat(path)
		assert.NoError(t, err)
		assert.False(t, info.IsDir())
	})

	t.Run("ReadDir", func(t *testing.T) {
		dirPath := "testdir"
		filePath := filepath.Join(dirPath, "file1.txt")

		// Создаем директорию и файл
		err := vfs.MkdirAll(dirPath, 0755)
		assert.NoError(t, err)

		file, err := vfs.Create(filePath)
		assert.NoError(t, err)
		file.Close()

		// Читаем содержимое директории
		entries, err := vfs.ReadDir(dirPath)
		assert.NoError(t, err)
		assert.Len(t, entries, 1)
		assert.Equal(t, "file1.txt", entries[0].Name())
	})

	t.Run("Error cases", func(t *testing.T) {
		// Несуществующий файл
		_, err := vfs.Open("nonexistent.txt")
		assert.Error(t, err)
		assert.True(t, errors.Is(err, fs.ErrNotExist))

		// Удаление несуществующего файла
		err = vfs.Remove("nonexistent.txt")
		assert.Error(t, err)

		// Статистика несуществующего файла
		_, err = vfs.Stat("nonexistent.txt")
		assert.Error(t, err)
	})
}

// MockFileSystem полная mock-реализация IFileSystem
type MockFileSystem struct {
	mock.Mock
}

func (m *MockFileSystem) Create(path string) (*os.File, error) {
	args := m.Called(path)
	return args.Get(0).(*os.File), args.Error(1)
}

func (m *MockFileSystem) Open(path string) (*os.File, error) {
	args := m.Called(path)
	return args.Get(0).(*os.File), args.Error(1)
}

func (m *MockFileSystem) MkdirAll(path string, perm fs.FileMode) error {
	args := m.Called(path, perm)
	return args.Error(0)
}

func (m *MockFileSystem) Remove(path string) error {
	args := m.Called(path)
	return args.Error(0)
}

func (m *MockFileSystem) RemoveAll(path string) error {
	args := m.Called(path)
	return args.Error(0)
}

func (m *MockFileSystem) Stat(path string) (fs.FileInfo, error) {
	args := m.Called(path)
	return args.Get(0).(fs.FileInfo), args.Error(1)
}

func (m *MockFileSystem) ReadDir(path string) ([]fs.DirEntry, error) {
	args := m.Called(path)
	return args.Get(0).([]fs.DirEntry), args.Error(1)
}

// MockFileInfo реализация fs.FileInfo
type MockFileInfo struct {
	mock.Mock
}

func (m *MockFileInfo) Name() string {
	return m.Called().String(0)
}

func (m *MockFileInfo) Size() int64 {
	return m.Called().Get(0).(int64)
}

func (m *MockFileInfo) Mode() fs.FileMode {
	return m.Called().Get(0).(fs.FileMode)
}

func (m *MockFileInfo) ModTime() time.Time {
	return m.Called().Get(0).(time.Time)
}

func (m *MockFileInfo) IsDir() bool {
	return m.Called().Bool(0)
}

func (m *MockFileInfo) Sys() interface{} {
	return m.Called().Get(0)
}

// MockDirEntry упрощенная реализация fs.DirEntry без использования testify/mock
type MockDirEntry struct {
	name  string
	isDir bool
}

func (m *MockDirEntry) Name() string {
	return m.name
}

func (m *MockDirEntry) IsDir() bool {
	return m.isDir
}

func (m *MockDirEntry) Type() fs.FileMode {
	return 0
}

func (m *MockDirEntry) Info() (fs.FileInfo, error) {
	return nil, nil
}

func TestIFileSystem(t *testing.T) {
	t.Run("ReadDir", func(t *testing.T) {
		mockFS := new(MockFileSystem)
		mockEntry := &MockDirEntry{
			name:  "test.txt",
			isDir: false,
		}

		// Настраиваем ожидания для ReadDir
		mockFS.On("ReadDir", "testdir").Return([]fs.DirEntry{mockEntry}, nil)

		// Вызываем тестируемый метод
		entries, err := mockFS.ReadDir("testdir")

		// Проверяем результаты
		assert.NoError(t, err)
		assert.Len(t, entries, 1)
		assert.Equal(t, "test.txt", entries[0].Name())
		assert.False(t, entries[0].IsDir())

		// Проверяем что все ожидания выполнены
		mockFS.AssertExpectations(t)
	})

	t.Run("Stat", func(t *testing.T) {
		mockFS := new(MockFileSystem)
		mockInfo := new(MockFileInfo)

		// Настраиваем ожидания
		mockInfo.On("Name").Return("test.txt")
		mockInfo.On("IsDir").Return(false)
		mockFS.On("Stat", "test.txt").Return(mockInfo, nil)

		// Вызываем тестируемый метод
		info, err := mockFS.Stat("test.txt")

		// Проверяем результаты
		assert.NoError(t, err)
		assert.Equal(t, "test.txt", info.Name())
		assert.False(t, info.IsDir())

		mockFS.AssertExpectations(t)
		mockInfo.AssertExpectations(t)
	})
}
