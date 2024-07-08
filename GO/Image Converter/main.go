package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	_ "image/gif"
	_ "image/jpeg"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/webp"
)

const windowTitle = "DR64 - Image Converter"

func main() {
	gtk.Init(nil)

	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetTitle(windowTitle)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	grid, _ := gtk.GridNew()
	win.Add(grid)

	imageWidget, _ := gtk.ImageNew()
	grid.Attach(imageWidget, 0, 0, 1, 1)

	selectButton, _ := gtk.ButtonNewWithLabel("Select Image")
	grid.Attach(selectButton, 0, 1, 1, 1)

	convertButton, _ := gtk.ButtonNewWithLabel("Convert Image")
	grid.Attach(convertButton, 0, 2, 1, 1)

	formatCombo, _ := gtk.ComboBoxTextNew()
	formats := []string{".png", ".webp", ".ico", ".gif", ".bmp", ".jpeg", ".tiff"}
	for _, format := range formats {
		formatCombo.AppendText(format)
	}
	formatCombo.SetActive(0)
	grid.Attach(formatCombo, 0, 3, 1, 1)

	var filePath string

	selectButton.Connect("clicked", func() {
		fileChooserDialog, _ := gtk.FileChooserDialogNewWith2Buttons(
			"Select Image", win, gtk.FILE_CHOOSER_ACTION_OPEN,
			"Cancel", gtk.RESPONSE_CANCEL, "Open", gtk.RESPONSE_ACCEPT)

		fileFilter, _ := gtk.FileFilterNew()
		fileFilter.AddPattern("*.png")
		fileFilter.AddPattern("*.webp")
		fileFilter.AddPattern("*.ico")
		fileFilter.AddPattern("*.gif")
		fileFilter.AddPattern("*.bmp")
		fileFilter.AddPattern("*.jpg")
		fileFilter.AddPattern("*.jpeg")
		fileFilter.AddPattern("*.tiff")
		fileChooserDialog.AddFilter(fileFilter)

		if fileChooserDialog.Run() == gtk.RESPONSE_ACCEPT {
			filePath = fileChooserDialog.GetFilename()
			pixbuf, err := gdk.PixbufNewFromFileAtScale(filePath, 200, 200, true)
			if err != nil {
				fmt.Println("Error loading image:", err)
				return
			}
			imageWidget.SetFromPixbuf(pixbuf)
		}
		fileChooserDialog.Destroy()
	})

	convertButton.Connect("clicked", func() {
		if filePath == "" {
			showErrorDialog(win, "No image selected")
			return
		}

		fileChooserDialog, _ := gtk.FileChooserDialogNewWith2Buttons(
			"Save Image As", win, gtk.FILE_CHOOSER_ACTION_SAVE,
			"Cancel", gtk.RESPONSE_CANCEL, "Save", gtk.RESPONSE_ACCEPT)

		fileChooserDialog.SetDoOverwriteConfirmation(true)

		saveFormat := formatCombo.GetActiveText()
		fileChooserDialog.SetCurrentName("untitled" + saveFormat)

		if fileChooserDialog.Run() == gtk.RESPONSE_ACCEPT {
			savePath := fileChooserDialog.GetFilename()
			err := convertImage(filePath, savePath)
			if err != nil {
				showErrorDialog(win, fmt.Sprintf("Failed to convert image: %v", err))
			} else {
				showInfoDialog(win, fmt.Sprintf("Image successfully saved as %s", savePath))
			}
		}
		fileChooserDialog.Destroy()
	})

	win.SetDefaultSize(400, 300)
	win.ShowAll()

	gtk.Main()
}

func convertImage(inputPath, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	switch outputPath[len(outputPath)-4:] {
	case ".png":
		err = png.Encode(outputFile, img)
	// Add cases for other formats as needed
	default:
		return fmt.Errorf("unsupported format")
	}

	return err
}

func showErrorDialog(parent *gtk.Window, message string) {
	dialog := gtk.MessageDialogNew(parent, gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_ERROR, gtk.BUTTONS_CLOSE, message)
	dialog.Run()
	dialog.Destroy()
}

func showInfoDialog(parent *gtk.Window, message string) {
	dialog := gtk.MessageDialogNew(parent, gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_INFO, gtk.BUTTONS_CLOSE, message)
	dialog.Run()
	dialog.Destroy()
}
