// Import required modules
const fs = require("fs");
const path = require("path");
const sharp = require("sharp");

// Function to convert an image to WebP
async function convertToWebP(filePath, outputFolder) {
    const fileName = path.basename(filePath, path.extname(filePath)) + ".webp";
    const outputPath = path.join(outputFolder, fileName);

    try {
        await sharp(filePath)
            .webp({ quality: 80 }) // Adjust quality as needed
            .toFile(outputPath);
        console.log(`Converted: ${filePath} -> ${outputPath}`);
    } catch (err) {
        console.error(`Failed to convert ${filePath}:`, err.message);
    }
}

// Function to process all images in a folder
async function processFolder(inputFolder, outputFolder) {
    if (!fs.existsSync(outputFolder)) {
        fs.mkdirSync(outputFolder, { recursive: true });
    }

    fs.readdir(inputFolder, (err, files) => {
        if (err) {
            console.error("Error reading folder:", err.message);
            return;
        }

        // Filter for image files (JPEG, JPG, PNG)
        const imageFiles = files.filter((file) =>
            /\.(jpe?g|png)$/i.test(file)
        );

        if (imageFiles.length === 0) {
            console.log("No image files found in the folder.");
            return;
        }

        // Convert each image to WebP
        imageFiles.forEach((file) => {
            const filePath = path.join(inputFolder, file);
            convertToWebP(filePath, outputFolder);
        });
    });
}

// Main script
const inputFolder = process.argv[2]; // First argument: Input folder
const outputFolder = process.argv[3]; // Second argument: Output folder

if (!inputFolder || !outputFolder) {
    console.error("Usage: node convert-to-webp.js <input-folder> <output-folder>");
    process.exit(1);
}

processFolder(inputFolder, outputFolder);
