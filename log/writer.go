//  Copyright 2019 The bigfile Authors. All rights reserved.
//  Use of this source code is governed by a MIT-style
//  license that can be found in the LICENSE file.

// Package log plan to provide a log collect component that can
// output application log to console and file meanwhile. In addition,
// log that is exported to file can be rotated automatically by Mode
// and file size.
package log

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Writer implement io.Writer interface, this writer will
// write content to different file by configuration.
type Writer struct {
	mu sync.Mutex

	// dir represent the directory that include log file
	dir string

	// basename represent log file basename, example: bigfile
	basename string

	// ext log file ext
	ext string

	// handler represent current file descriptor, that's used
	// to write log content
	handler *os.File
}

func (w *Writer) GetDir() string {
	return w.dir
}

func (w *Writer) GetBasename() string {
	return w.basename
}

func (w *Writer) GetExt() string {
	return w.ext
}

func (w *Writer) GetHandler() *os.File {
	return w.handler
}

// NewWriter is used return a writer, that will change to another writer
// when the size of current file will be up to maxBytes.
func NewWriter(file string) (*Writer, error) {
	var (
		dir              = filepath.Dir(file)
		err              error
		stat             os.FileInfo
		completeFileName string
		basename         string
		ext              string
		handler          *os.File
	)
	stat, err = os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	} else if err == nil && !stat.IsDir() {
		return nil, errors.New("the directory of log file is illegal")
	} else if err != nil {
		return nil, err
	}

	dir = strings.TrimSuffix(dir, "/")
	ext = filepath.Ext(file)
	basename = strings.TrimSuffix(filepath.Base(file), ext)
	ext = strings.TrimPrefix(ext, ".")

	completeFileName = fmt.Sprintf("%s/%s.%s", dir, basename, ext)
	if stat, err = os.Stat(completeFileName); err != nil {
		if os.IsNotExist(err) {
			if handler, err = os.OpenFile(completeFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &Writer{
		dir:      dir,
		basename: basename,
		ext:      ext,
		handler:  handler,
	}, nil
}

// Write is used to implement io.Writer, because of mutex before write,
// this will lead to a bad performance. I will optimize this in the future.
func (w *Writer) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	n, err = w.handler.Write(p)
	return n, err
}

// Close is used to implement io.Close
func (w *Writer) Close() error {
	return w.handler.Close()
}
