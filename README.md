# FlexiDoc
Simple tool to processes a DOCX file with placeholders and a JSON file

## Features
- Easy download of the latest project release.
- Generate a DOCX file using a template containing placeholder tags.
- Command-line interface (CLI) for flexible usage.

## Getting Started

Download the latest release of the project from the [Releases](https://github.com/jam3on/FlexiDoc/releases/) page.

## Usage

1. **JSON File**: Create a JSON file containing the placeholder data. Example:
   ```json
    {
        "title": "Report Title",
        "date": "2025-01-01",
        "author": "John Doe"
    }
    ```
2. **DOCX Template**: Prepare a DOCX file with placeholders wrapped in curly braces. Example:
    ```docx
    Title: {title}
    Date: {date}
    Author: {author}
    ```

3. **Run the Tool** : Use the following CLI arguments to run the tool:

-j or --json: Path to the JSON file containing the placeholders.

-t or --template: Path to the DOCX template file.

-o or --output (optional): Path to save the generated DOCX file. Defaults to the current directory.

**Example Command**:

```bash
.\FlexiDoc.exe -j data.json -t template.docx -o output.docx
```

## Disclaimer

This tool is provided "as is" without warranty of any kind. Use it at your own risk. The authors are not responsible for any loss of funds or data due to improper usage or external threats.

## Contributing

Contributions are welcome! Feel free to fork the repository, open issues, or submit pull requests to improve this tool.

## License

This project is licensed under the [MIT License](LICENSE).