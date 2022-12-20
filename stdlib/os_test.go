package stdlib_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/snple/slim"
	"github.com/snple/slim/require"
)

func TestReadFile(t *testing.T) {
	content := []byte("the quick brown fox jumps over the lazy dog")
	tf, err := ioutil.TempFile("", "test")
	require.NoError(t, err)
	defer func() { _ = os.Remove(tf.Name()) }()

	_, err = tf.Write(content)
	require.NoError(t, err)
	_ = tf.Close()

	module(t, "os").call("read_file", tf.Name()).
		expect(&slim.Bytes{Value: content})
}

func TestReadFileArgs(t *testing.T) {
	module(t, "os").call("read_file").expectError()
}
func TestFileStatArgs(t *testing.T) {
	module(t, "os").call("stat").expectError()
}

func TestFileStatFile(t *testing.T) {
	content := []byte("the quick brown fox jumps over the lazy dog")
	tf, err := ioutil.TempFile("", "test")
	require.NoError(t, err)
	defer func() { _ = os.Remove(tf.Name()) }()

	_, err = tf.Write(content)
	require.NoError(t, err)
	_ = tf.Close()

	stat, err := os.Stat(tf.Name())
	if err != nil {
		t.Logf("could not get tmp file stat: %s", err)
		return
	}

	module(t, "os").call("stat", tf.Name()).expect(&slim.ImmutableMap{
		Value: map[string]slim.Object{
			"name":      &slim.String{Value: stat.Name()},
			"mtime":     &slim.Time{Value: stat.ModTime()},
			"size":      &slim.Int{Value: stat.Size()},
			"mode":      &slim.Int{Value: int64(stat.Mode())},
			"directory": slim.FalseValue,
		},
	})
}

func TestFileStatDir(t *testing.T) {
	td, err := ioutil.TempDir("", "test")
	require.NoError(t, err)
	defer func() { _ = os.RemoveAll(td) }()

	stat, err := os.Stat(td)
	require.NoError(t, err)

	module(t, "os").call("stat", td).expect(&slim.ImmutableMap{
		Value: map[string]slim.Object{
			"name":      &slim.String{Value: stat.Name()},
			"mtime":     &slim.Time{Value: stat.ModTime()},
			"size":      &slim.Int{Value: stat.Size()},
			"mode":      &slim.Int{Value: int64(stat.Mode())},
			"directory": slim.TrueValue,
		},
	})
}

func TestOSExpandEnv(t *testing.T) {
	curMaxStringLen := slim.MaxStringLen
	defer func() { slim.MaxStringLen = curMaxStringLen }()
	slim.MaxStringLen = 12

	_ = os.Setenv("slim", "FOO BAR")
	module(t, "os").call("expand_env", "$slim").expect("FOO BAR")

	_ = os.Setenv("slim", "FOO")
	module(t, "os").call("expand_env", "$slim $slim").expect("FOO FOO")

	_ = os.Setenv("slim", "123456789012")
	module(t, "os").call("expand_env", "$slim").expect("123456789012")

	_ = os.Setenv("slim", "1234567890123")
	module(t, "os").call("expand_env", "$slim").expectError()

	_ = os.Setenv("slim", "123456")
	module(t, "os").call("expand_env", "$slim$slim").expect("123456123456")

	_ = os.Setenv("slim", "123456")
	module(t, "os").call("expand_env", "${slim}${slim}").
		expect("123456123456")

	_ = os.Setenv("slim", "123456")
	module(t, "os").call("expand_env", "$slim $slim").expectError()

	_ = os.Setenv("slim", "123456")
	module(t, "os").call("expand_env", "${slim} ${slim}").expectError()
}
