package protoeasy

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/pkg/archive"
)

func getAllRelProtoFilePaths(dirPath string) ([]string, error) {
	var relProtoFilePaths []string
	if err := filepath.Walk(
		dirPath,
		func(filePath string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fileInfo.IsDir() {
				return nil
			}
			if filepath.Ext(filePath) != ".proto" {
				return nil
			}
			relFilePath, err := filepath.Rel(dirPath, filePath)
			if err != nil {
				return err
			}
			relProtoFilePaths = append(relProtoFilePaths, relFilePath)
			return nil
		},
	); err != nil {
		return nil, err
	}
	return relProtoFilePaths, nil
}

func filterFilePaths(filePaths []string, excludeFilePatterns []string) ([]string, error) {
	var filteredFilePaths []string
	for _, filePath := range filePaths {
		matched, err := filePathMatches(filePath, excludeFilePatterns)
		if err != nil {
			return nil, err
		}
		if !matched {
			filteredFilePaths = append(filteredFilePaths, filePath)
		}
	}
	return filteredFilePaths, nil
}

func filePathMatches(filePath string, excludeFilePatterns []string) (bool, error) {
	// TODO(pedge): handle this logic correctly
	for _, excludeFilePattern := range excludeFilePatterns {
		if strings.HasPrefix(filePath, excludeFilePattern) {
			return true, nil
		}
		matched, err := filepath.Match(excludeFilePattern, filePath)
		if err != nil {
			return false, err
		}
		if matched {
			return true, nil
		}
	}
	return false, nil
}

func tar(dirPath string, includeFiles []string) (io.ReadCloser, error) {
	return archive.TarWithOptions(
		dirPath,
		&archive.TarOptions{
			IncludeFiles: includeFiles,
			Compression:  archive.Uncompressed,
			NoLchown:     true,
		},
	)
}

func untar(reader io.Reader, dirPath string) error {
	return archive.Untar(
		reader,
		dirPath,
		&archive.TarOptions{
			NoLchown: true,
		},
	)
}