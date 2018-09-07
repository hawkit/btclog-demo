package btclog

import (
	"testing"
	"bytes"
	"fmt"
	"strings"
)

var testBuffer = bytes.NewBuffer(make([]byte, 0 , 120))

var testLogger = &slog{lvl: LevelInfo, tag: "subsys", b: NewBackend(testBuffer)}

func TestSlog_Trace(t *testing.T) {
	helperLog(LevelTrace, t)
}

func TestSlog_Debug(t *testing.T) {
	helperLog(LevelDebug, t)
}

func TestSlog_Info(t *testing.T) {
	helperLog(LevelInfo, t)
}

func TestSlog_Warn(t *testing.T) {
	helperLog(LevelWarn, t)
}

func TestSlog_Error(t *testing.T) {
	helperLog(LevelError, t)
}

func TestSlog_Critical(t *testing.T) {
	helperLog(LevelCritical, t)
}

func logAll(logger Logger) {

	logger.Trace("Trace")
	logger.Tracef("Tracef-%s", "test")
	logger.Debug("Debug")
	logger.Debugf("Debugf-%s", "test")
	logger.Info("Info")
	logger.Infof("Infof-%s", "test")
	logger.Warn("Warn")
	logger.Warnf("Warnf-%s", "test")
	logger.Error("Error")
	logger.Errorf("Errorf-%s", "test")
	logger.Critical("Critical")
	logger.Criticalf("Criticalf-%s", "test")
}

func helperLog(lvl Level, t *testing.T)  {
	defer testBuffer.Reset()
	testLogger.SetLevel(lvl)
	logAll(testLogger)

	str := fmt.Sprintf("%s", testBuffer)

	var lvlStrs []string
	lvlStrs = append(lvlStrs, "test")
	for i := int(lvl); i < int(LevelOff); i++ {
		lvlStrs = append(lvlStrs, levelStrs[i])
	}
	for _, levelStr := range lvlStrs {
		if !strings.Contains(str, levelStr) {
			t.Errorf("can't find '%s', expected log with '%s'", levelStr, levelStr)
		}
	}

	fmt.Println(str)
}