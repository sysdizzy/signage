package main

import (
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FileSignature struct {
	Path string `xml:"path,attr"`
	Hash string `xml:"hash,attr"`
}

type SignedFiles struct {
	XMLName xml.Name        `xml:"signed_files"`
	Files   []FileSignature `xml:"file"`
}

func sign() error {
	signedFiles := SignedFiles{}
	pathList := strings.Split(os.Getenv("PATH"), ":")

	for _, dir := range pathList {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && (info.Mode()&0111) != 0 {
				hash, err := generateHash(path)
				if err != nil {
					return err
				}

				signedFiles.Files = append(signedFiles.Files, FileSignature{
					Path: path,
					Hash: hash,
				})
			}
			return nil
		})

		if err != nil {
			return err
		}
	}

	xmlData, err := xml.MarshalIndent(signedFiles, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("~/signage/signed.xml", xmlData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func generateHash(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash), nil
}
func verify() error {
	xmlData, err := ioutil.ReadFile("~/signage/signed.xml")
	if err != nil {
		return err
	}

	signedFiles := SignedFiles{}
	err = xml.Unmarshal(xmlData, &signedFiles)
	if err != nil {
		return err
	}

	diffLogFile, err := os.Create("~/signage/diff.log")
	if err != nil {
		return err
	}
	defer diffLogFile.Close()

	for _, file := range signedFiles.Files {
		currentHash, err := generateHash(file.Path)
		if err != nil {
			return err
		}

		if file.Hash != currentHash {
			diffLog := fmt.Sprintf("File: %s\nPrevious Hash: %s\nCurrent Hash: %s\n\n",
				file.Path, file.Hash, currentHash)
			_, err = diffLogFile.WriteString(diffLog)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: signage <command>")
		fmt.Println("Commands: sign, verify")
		return
	}

	command := os.Args[1]

	switch command {
	case "sign":
		err := sign()
		if err != nil {
			fmt.Println("Error while signing:", err)
			os.Exit(1)
		}
		fmt.Println("Signage complete.")

	case "verify":
		err := verify()
		if err != nil {
			fmt.Println("Error while verifying:", err)
			os.Exit(1)
		}
		fmt.Println("Verification complete.")

	default:
		fmt.Println("Invalid command:", command)
		fmt.Println("Usage: signage <command>")
		fmt.Println("Commands: sign, verify")
	}
}
