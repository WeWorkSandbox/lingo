package generator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func WriteFile(path string, data io.Reader, perm os.FileMode) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return fmt.Errorf("unable to open file for write: %w", err)
	}
	if _, err = io.Copy(f, data); err != nil {
		return fmt.Errorf("unable to copy file data: %w", err)
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("unable to close file: %w", err)
	}
	return nil
}

func writeSchema(root, schemaName string, contents io.Reader) error {
	dirPath := buildSchemaDir(root, schemaName)
	schemaFile := filepath.Join(dirPath, "schema.go")

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}
	return WriteFile(schemaFile, contents, os.ModePerm)
}

func writeTable(root, schemaName, name string, contents io.Reader) error {
	dirPath := buildTableDir(root, schemaName, name)
	path := filepath.Join(dirPath, "table.go")

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}
	return WriteFile(path, contents, os.ModePerm)
}

func writePackageMembers(root, schemaName, name string, contents io.Reader) error {
	dirPath := buildTableDir(root, schemaName, name)
	path := filepath.Join(dirPath, "exported.go")

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}
	return WriteFile(path, contents, os.ModePerm)
}

func buildSchemaDir(root, schemaName string) string {
	pkg := strings.ToLower(LittleQ(ToNonExported(schemaName)))
	dirPath := filepath.Join(root, pkg)
	return dirPath
}

func buildTableDir(root string, schemaName string, name string) string {
	tablePkgName := strings.ToLower(LittleQ(ToNonExported(name)))
	dirPath := filepath.Join(buildSchemaDir(root, schemaName), tablePkgName)
	return dirPath
}
