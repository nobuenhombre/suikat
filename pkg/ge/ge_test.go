package ge

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIdentityErrorError(t *testing.T) {
	mainErr := errors.New("main error")

	middleLevelError := Pin(mainErr, Params{"Middle": "Middle"})

	hiLevelError := Pin(middleLevelError, Params{"Hi": 123})

	actualMsg := hiLevelError.Error()

	// Проверяем наличие обязательных элементов верхнего уровня
	require.Contains(t, actualMsg, "CreatedAt:")
	require.Contains(t, actualMsg, "Way: TestIdentityErrorError() github.com/nobuenhombre/suikat/pkg/ge ge_test.go line")
	require.Contains(t, actualMsg, "Params: (Hi = 123)")
	require.Contains(t, actualMsg, "ParentError:")

	// Проверяем наличие элементов вложенной ошибки (с отступами)
	require.Contains(t, actualMsg, "\tCreatedAt:")
	require.Contains(t, actualMsg, "\tWay: TestIdentityErrorError() github.com/nobuenhombre/suikat/pkg/ge ge_test.go line")
	require.Contains(t, actualMsg, "\tParams: (Middle = Middle)")
	require.Contains(t, actualMsg, "\tParentError:")
	require.Contains(t, actualMsg, "\t\tmain error")
}

func TestIdentityErrorRootError(t *testing.T) {
	mainErr := errors.New("main error")

	middleLevelError := Pin(mainErr, Params{"Middle": "Middle"})

	hiLevelError := Pin(middleLevelError, Params{"Hi": 123})

	identityErr, ok := errors.AsType[*IdentityError](hiLevelError)
	require.True(t, ok)

	rootErr := identityErr.RootError()
	require.Equal(t, mainErr, rootErr)
}

func TestIdentityErrorRootError2(t *testing.T) {
	mainErr := New("main error")

	middleLevelError := Pin(mainErr, Params{"Middle": "Middle"})

	hiLevelError := Pin(middleLevelError, Params{"Hi": 123})

	identityErr, ok := errors.AsType[*IdentityError](hiLevelError)
	require.True(t, ok)

	rootErr := identityErr.RootError()
	require.Equal(t, mainErr, rootErr)
}
