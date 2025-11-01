package messages

const (
	MsgErrInvalidCommand   = "Неправильная команда (h для помощи)"
	MsgErrInvalidArgsCount = "Неверное количество аргументов"
	MsgErrInvalidArg       = "Неверный аргумент"

	MsgErrInternal    = "Внутренняя ошибка"
	MsgErrEmptyResult = "Пустой результат"

	MsgFileEmpty     = "Пустой файл"
	MsgFileInfoErr   = "Ошибка получения информации о файле '%s': %s"
	MsgFileUploadErr = "Ошибка загрузки файла '%s': %s"
	MsgFileUploaded  = "Файл '%s' загружен"
	MsgFileDeleteErr = "Ошибка удаления файла '%s': %s"
	MsgFileTotal     = "Итого: %d"
	MsgFileDeleted   = "Удалено: %d"
	MsgFileRenamed   = "Переименовано: %d"
	MsgFileRenameErr = "Ошибка переименования файла '%s': %s"

	MsgOK = "OK"
)
