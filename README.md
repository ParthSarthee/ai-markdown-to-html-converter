# Markdown to HTML Converter

A modern, responsive web application built with Go that converts Markdown text to HTML in real-time. Features a clean, intuitive interface with live preview capabilities.

![Markdown Converter Demo](https://img.shields.io/badge/Go-1.22+-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Status](https://img.shields.io/badge/status-active-success.svg)

## 🚀 Features

- **Real-time Conversion**: Convert Markdown to HTML instantly
- **Live Preview**: View rendered HTML output in a separate preview page
- **Responsive Design**: Optimized for desktop, tablet, and mobile devices
- **Copy to Clipboard**: Easy copying of generated HTML code
- **Syntax Highlighting**: Clear display of HTML output with proper formatting
- **Mobile Auto-convert**: Automatic conversion after typing stops (mobile devices)
- **Keyboard Shortcuts**:
  - `Ctrl/Cmd + Enter`: Convert markdown
  - `Ctrl/Cmd + P`: Print preview (in preview mode)
  - `Escape`: Return to editor (from preview)

## 🛠️ Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML5, CSS3, JavaScript (ES6+)
- **Styling**: Tailwind CSS (editor), Custom CSS (preview)
- **Markdown Parser**: [Blackfriday v2](https://github.com/russross/blackfriday)
- **Server**: Go's built-in HTTP server

## 📋 Prerequisites

- Go 1.22 or higher
- Modern web browser (Chrome, Firefox, Safari, Edge)

## 🔧 Installation & Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/yourusername/md-to-website.git
   cd md-to-website
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Run the application**

   ```bash
   go run main.go
   ```

4. **Access the application**
   Open your browser and navigate to `http://localhost:8080`

## 📁 Project Structure

```
md-to-website/
├── main.go                 # Main server file with handlers
├── go.mod                  # Go module dependencies
├── templates/              # HTML templates
│   ├── editor.html         # Main editor interface
│   └── preview.html        # Preview page template
└── README.md              # Project documentation
```

## 🔌 API Endpoints

- `GET /` - Main editor interface
- `GET /preview` - Preview page with rendered HTML
- `POST /api/convert` - Convert markdown to HTML (JSON API)

### API Usage Example

```javascript
// Convert markdown to HTML
const response = await fetch("/api/convert", {
	method: "POST",
	headers: {
		"Content-Type": "application/json",
	},
	body: JSON.stringify({
		markdown: "# Hello World\nThis is **markdown**!",
	}),
});

const result = await response.json();
console.log(result.html); // "<h1>Hello World</h1>\n<p>This is <strong>markdown</strong>!</p>\n"
```

## 🎨 UI/UX Features

### Responsive Layout

- **Desktop/Laptop**: Side-by-side markdown input and HTML output
- **Mobile/Tablet**: Stacked layout with optimized heights
- **Adaptive Typography**: Scales appropriately across screen sizes

### Smart Button States

- **Convert Button**: Located in markdown input section for easy access
- **Copy & Preview Buttons**: Only enabled when HTML content is available
- **Visual Feedback**: Loading states and success indicators

### Mobile Optimizations

- Auto-conversion after 2 seconds of inactivity
- Touch-friendly button sizes
- Optimized viewport handling
- Orientation change support

## 🔧 Configuration

The application runs on port 8080 by default. To change the port, modify the `main.go` file:

```go
err := http.ListenAndServe(":8080", mux) // Change port here
```

## 🚦 Supported Markdown Features

Thanks to Blackfriday v2, the converter supports:

- Headers (`# ## ###`)
- **Bold** and _italic_ text
- Links `[text](url)`
- Images `![alt](src)`
- Code blocks and `inline code`
- Lists (ordered and unordered)
- Blockquotes
- Tables
- Horizontal rules
- Strikethrough text
- Automatic heading IDs
- External links open in new tabs

## 🔒 Security Features

- Input sanitization through Blackfriday
- Safe HTML rendering
- XSS protection via proper escaping
- Content-Type headers properly set

## 🚀 Performance

- Lightweight Go backend (~10MB memory footprint)
- Fast markdown parsing with Blackfriday
- Minimal JavaScript for optimal loading
- CSS optimizations for smooth animations

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Blackfriday](https://github.com/russross/blackfriday) for excellent Markdown parsing
- [Tailwind CSS](https://tailwindcss.com/) for utility-first CSS framework
- Go community for the robust standard library

---

⭐ **Star this repo if you found it useful!** ⭐
