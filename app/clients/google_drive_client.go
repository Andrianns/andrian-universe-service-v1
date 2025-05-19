package clients

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type GoogleDrive interface {
	EnsureFolder(name string) (string, error)
	UploadFile(file *multipart.FileHeader, folderID string) (string, error)
	ShareFile(fileID string) (string, error)
	GetFileURLByName(fileName, folderName string) (string, error)
}

type googleDriveClient struct {
	srv *drive.Service
}

func NewGoogleDriveClient() (GoogleDrive, error) {
	ctx := context.Background()
	creds, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("GOOGLE_CREDENTIALS")), drive.DriveScope)
	if err != nil {
		return nil, err
	}

	srv, err := drive.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		return nil, err
	}

	return &googleDriveClient{srv: srv}, nil
}

// Implementation

func (g *googleDriveClient) EnsureFolder(name string) (string, error) {
	q := fmt.Sprintf("mimeType='application/vnd.google-apps.folder' and name='%s' and trashed=false", name)
	files, err := g.srv.Files.List().Q(q).Fields("files(id)").Do()
	if err != nil {
		return "", err
	}
	if len(files.Files) > 0 {
		return files.Files[0].Id, nil
	}

	folder := &drive.File{Name: name, MimeType: "application/vnd.google-apps.folder"}
	created, err := g.srv.Files.Create(folder).Fields("id").Do()
	if err != nil {
		return "", err
	}
	return created.Id, nil
}

func (g *googleDriveClient) UploadFile(file *multipart.FileHeader, folderID string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		io.Copy(pw, src) // ✅ now using correct writer
	}()

	fileMeta := &drive.File{
		Name:    file.Filename,
		Parents: []string{folderID},
	}

	upload, err := g.srv.Files.Create(fileMeta).Media(pr).Fields("id").Do()
	if err != nil {
		return "", err
	}

	return upload.Id, nil
}

func (g *googleDriveClient) ShareFile(fileID string) (string, error) {
	perm := &drive.Permission{
		Type: "anyone",
		Role: "reader",
	}
	_, err := g.srv.Permissions.Create(fileID, perm).Do()
	if err != nil {
		return "", err
	}
	return "https://drive.google.com/uc?id=" + fileID, nil
}

func (g *googleDriveClient) GetFileURLByName(fileName, folderName string) (string, error) {
	folderID, err := g.EnsureFolder(folderName)
	if err != nil {
		return "", err
	}
	query := fmt.Sprintf("name='%s' and '%s' in parents and trashed=false", fileName, folderID)
	res, err := g.srv.Files.List().Q(query).Fields("files(id)").Do()
	if err != nil {
		return "", err
	}
	if len(res.Files) == 0 {
		return "", fmt.Errorf("file not found")
	}
	return "https://drive.google.com/uc?id=" + res.Files[0].Id, nil
}
