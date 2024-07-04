import tkinter as tk
from tkinter import filedialog, messagebox
from PIL import Image, ImageTk

class ImageConverterApp:
    def __init__(self, root):
        self.root = root
        self.root.title("DR64 - Image Converter")

        self.file_path = None
        self.image_label = tk.Label(root)
        self.image_label.pack()

        self.select_button = tk.Button(root, text="Select Image", command=self.select_image)
        self.select_button.pack()

        self.convert_button = tk.Button(root, text="Convert Image", command=self.convert_image)
        self.convert_button.pack()

        self.format_var = tk.StringVar(root)
        self.format_var.set(".png")
        self.format_menu = tk.OptionMenu(root, self.format_var, ".png", ".webp", ".ico", ".gif", ".bmp", ".jpeg", ".tiff")
        self.format_menu.pack()

    def select_image(self):
        self.file_path = filedialog.askopenfilename(filetypes=[("Image files", "*.png;*.webp;*.ico;*.gif;*.bmp;*.jpg;*.jpeg;*.tiff")])
        if self.file_path:
            self.display_image(self.file_path)

    def display_image(self, path):
        image = Image.open(path)
        image.thumbnail((200, 200))
        photo = ImageTk.PhotoImage(image)
        self.image_label.configure(image=photo)
        self.image_label.image = photo

    def convert_image(self):
        if not self.file_path:
            messagebox.showerror("Error", "No image selected")
            return

        save_path = filedialog.asksaveasfilename(defaultextension=self.format_var.get(),
                                                 filetypes=[("Image files", "*.png;*.webp;*.ico;*.gif;*.bmp;*.jpg;*.jpeg;*.tiff")])
        if save_path:
            try:
                image = Image.open(self.file_path)
                image.save(save_path)
                messagebox.showinfo("Success", f"Image successfully saved as {save_path}")
            except Exception as e:
                messagebox.showerror("Error", f"Failed to convert image: {e}")

if __name__ == "__main__":
    root = tk.Tk()
    app = ImageConverterApp(root)
    root.mainloop()
