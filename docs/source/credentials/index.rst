CREDENTIALS - Файлы через Environments
======================================

В GKE (Google Kubernetes Engine) микро сервисы запаковываются через Docker.
Есть всякие файлы от google - например у нас есть некий файл credentials.json (содержащий например токены авторизации).
Не хочется хранить такие файлы в контейнере Docker.
Хочется передавать их как нормальные параметры через Environments или попросту через секреты GKE.

Сказано, сделано.

Шаг 1 - Превращаем файл в строку
--------------------------------

Берем любой файл - текстовый, бинарный, любой.
Конвертируем его в Base64 строку.

Например см. пример тут :ref:`colorog-example`


Шаг 2 - Пишем эту строку в Environments
---------------------------------------

это просто


Шаг 3 - Читаем данные из Environments
-------------------------------------

.. code-block:: go
  :linenos:

    import (
      "fmt"

      "golang.org/x/oauth2/google"
      "github.com/nobuenhombre/suikat/pkg/credentials"
    )

    ...

    type CredentialsConfig struct {
      GDriveClientSecret credentials.EnvCred
    }

    func New() *CredentialsConfig {
      return &CredentialsConfig{
        GDriveClientSecret: credentials.New("GOOGLE_DRIVE_CLIENT_SECRET"),
      }
    }

    func main() {
      conf := New()

      clientSecret, err := conf.GDriveClientSecret.GetBytes()
      if err != nil {
        panic(fmt.Sprintf("Unable to get client secret [%v]", err))
      }

      drvConf, err := google.ConfigFromJSON(clientSecret, drive.DriveScope)
      if err != nil {
        panic("Unable to parse client secret file to drvConf")
      }

      ...
    }

Доступные методы чтения
-----------------------

.. code-block:: go
  :linenos:

    // читает из environments, декодирует base64, отдает []byte
    func (d *Data) GetBytes() ([]byte, error)

    // читает из environments, декодирует base64, отдает string
    func (d *Data) GetString() (string, error)

    // читает из environments, декодирует base64,
    // создает временный файл, сохраняет в него содержимое
    // отдает string - полное имя временного файла, не забудьте потом удалить этот файл
    func (d *Data) GetFile() (string, error)
