package log

import (
	"os"
	"strings"

	"github.com/op/go-logging"

	"github.com/dabao-zhao/xgin/config"
)

var (
	log *logging.Logger

	NameToLevel = map[string]logging.Level{
		"DEBUG":    logging.DEBUG,
		"INFO":     logging.INFO,
		"NOTICE":   logging.NOTICE,
		"WARNING":  logging.WARNING,
		"ERROR":    logging.ERROR,
		"CRITICAL": logging.CRITICAL,
	}

	LevelToName = map[logging.Level]string{
		logging.DEBUG:    "DEBUG",
		logging.INFO:     "INFO",
		logging.NOTICE:   "NOTICE",
		logging.WARNING:  "WARNING",
		logging.ERROR:    "ERROR",
		logging.CRITICAL: "CRITICAL",
	}
)

func NewLogger(logConfig *config.LogConfig) (logger *logging.Logger, err error) {
	if log != nil {
		return log, nil
	}

	var (
		ok                  bool
		module              = "xgin"
		level               logging.Level
		backend             []logging.Backend
		consoleBackend      logging.Backend
		fileBackend         logging.Backend
		fileHandler         *Writer
		consoleLevelBackend logging.LeveledBackend
		fileLevelBackend    logging.LeveledBackend
		leveledBackend      logging.LeveledBackend
	)

	log = logging.MustGetLogger(module)
	// output log to stdout
	if logConfig.Console.Enable {
		consoleBackend = logging.NewBackendFormatter(
			logging.NewLogBackend(os.Stdout, "", 0),
			logging.MustStringFormatter(logConfig.Console.Format),
		)
		consoleLevelBackend = logging.AddModuleLevel(consoleBackend)
		if level, ok = NameToLevel[strings.ToUpper(logConfig.Console.Level)]; !ok {
			level = logging.DEBUG
		}
		consoleLevelBackend.SetLevel(level, module)
		backend = append(backend, consoleLevelBackend)
	}

	// file output handler
	if logConfig.File.Enable {
		fileHandler, err = NewWriter(logConfig.File.Path)
		if err != nil {
			return
		}
		fileBackend = logging.NewBackendFormatter(
			logging.NewLogBackend(fileHandler, "", 0),
			logging.MustStringFormatter(logConfig.File.Format),
		)
		fileLevelBackend = logging.AddModuleLevel(fileBackend)
		if level, ok = NameToLevel[strings.ToUpper(logConfig.File.Level)]; !ok {
			level = logging.WARNING
		}
		fileLevelBackend.SetLevel(level, module)
		backend = append(backend, fileLevelBackend)
	}

	leveledBackend = logging.MultiLogger(backend...)
	log.SetBackend(leveledBackend)

	return log, err
}
