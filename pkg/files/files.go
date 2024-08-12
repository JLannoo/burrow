package files

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

type FileManager struct {
	// Path to the directory where the files are stored
	Path string
	// Path to the file where the secret key is stored
	SecretKeyFileName string
	// Path to the file where the master password is temporarily stored
	MasterPasswordFileName string
}

// Creates a new FileManager with the given path
func NewFileManager(dirPath string) *FileManager {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, 0755)
		fmt.Println("Created directory", dirPath)
	}

	return &FileManager{
		Path:                   dirPath,
		SecretKeyFileName:      ".key",
		MasterPasswordFileName: ".master",
	}
}

// Writes the given data to the file
func (fm *FileManager) WriteToFile(data []byte, filename string) error {
	finalPath := path.Join(fm.Path, filename)

	// If the file is a file in a subdirectory, create the subdirectory
	if strings.Contains(filename, "/") {
		subdir := path.Dir(finalPath)
		if _, err := os.Stat(subdir); os.IsNotExist(err) {
			os.MkdirAll(subdir, 0755)
		}
	}

	file, err := os.Create(finalPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// Reads the data from the file
func (fm *FileManager) ReadFromFile(filename string) ([]byte, error) {
	finalPath := path.Join(fm.Path, filename)
	file, err := os.Open(finalPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (fm *FileManager) GetFileUpdateTime(filename string) (int64, error) {
	finalPath := path.Join(fm.Path, filename)
	file, err := os.Open(finalPath)
	if err != nil {
		return 0, err
	}

	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return stat.ModTime().UnixMilli(), nil
}

func (fm *FileManager) WriteToSecretKeyFile(data []byte) error {
	return fm.WriteToFile(data, fm.SecretKeyFileName)
}

func (fm *FileManager) ReadFromSecretKeyFile() ([]byte, error) {
	return fm.ReadFromFile(fm.SecretKeyFileName)
}

func (fm *FileManager) WriteToMasterPasswordFile(data []byte) error {
	return fm.WriteToFile(data, fm.MasterPasswordFileName)
}

func (fm *FileManager) ReadFromMasterPasswordFile() ([]byte, error) {
	masterPassword, err := fm.ReadFromFile(fm.MasterPasswordFileName)
	if err != nil {
		return nil, errors.New("master password file not found")
	}

	return masterPassword, nil
}

var dir, _ = os.UserHomeDir()
var Manager = NewFileManager(path.Join(dir, ".burrow"))
