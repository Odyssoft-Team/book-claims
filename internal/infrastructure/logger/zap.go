package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewZapLogger(env string, logFilePath string) *zap.Logger {
	var encoderCfg zapcore.EncoderConfig
	var level zapcore.LevelEnabler
	var consoleEncoder zapcore.Encoder
	var fileEncoder zapcore.Encoder

	// 1. Configuración del Encoder y Nivel de Logeo según el entorno
	if env == "production" {
		// --- Configuración para Producción ---
		encoderCfg = zap.NewProductionEncoderConfig()
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder   // Formato de tiempo estándar para JSON
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder // Ej. "INFO", "WARN"
		encoderCfg.CallerKey = "caller"                      // Incluye el archivo:línea de la llamada
		encoderCfg.StacktraceKey = "stacktrace"              // Clave para el stacktrace
		encoderCfg.MessageKey = "message"                    // Clave para el mensaje principal del log
		level = zap.InfoLevel                                // Nivel mínimo para producción (Info, Warn, Error, Fatal, Panic)

		consoleEncoder = zapcore.NewJSONEncoder(encoderCfg) // Consola también en JSON en producción
		fileEncoder = zapcore.NewJSONEncoder(encoderCfg)    // Archivo en JSON
	} else {
		// --- Configuración para Desarrollo (env = "development" o cualquier otra cosa) ---
		encoderCfg = zap.NewDevelopmentEncoderConfig()            // Configuración más legible para desarrollo
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder // Niveles con colores en la consola
		encoderCfg.CallerKey = "caller"                           // Incluye el archivo:línea
		encoderCfg.StacktraceKey = "stacktrace"                   // Clave para el stacktrace
		encoderCfg.MessageKey = "message"                         // Clave para el mensaje principal del log
		level = zap.DebugLevel                                    // Nivel mínimo para desarrollo (incluye Debug)

		consoleEncoder = zapcore.NewConsoleEncoder(encoderCfg) // Consola con formato legible (no JSON)
		fileEncoder = zapcore.NewConsoleEncoder(encoderCfg)    // Archivo en formato legible (puedes cambiarlo a JSON si lo prefieres)
	}

	// 2. Configuración de Lumberjack para la rotación de logs en archivo
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    50,
		Compress:   true,
	}

	// 3. Definir los "sink" de salida (donde irán los logs físicamente)
	fileWriter := zapcore.AddSync(lumberjackLogger)
	consoleWriter := zapcore.AddSync(os.Stderr)

	// 4. Crear los Cores de Zap para cada destino
	fileCore := zapcore.NewCore(
		fileEncoder,
		fileWriter,
		level,
	)

	consoleCore := zapcore.NewCore(
		consoleEncoder,
		consoleWriter,
		level,
	)

	// 5. Combinar los Cores con zapcore.NewTee para enviar logs a múltiples destinos
	core := zapcore.NewTee(fileCore, consoleCore)

	// 6. Construir el Logger de Zap
	var opts []zap.Option
	opts = append(opts, zap.AddCaller())
	// opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))
	if env != "production" {
		opts = append(opts, zap.Development())
	}

	l := zap.New(core, opts...)

	return l
}


