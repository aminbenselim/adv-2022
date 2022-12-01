const fs = require('fs');

// Define the function
function readFileLines(filePath) {
  // Read the contents of the file
  const fileContents = fs.readFileSync(filePath, 'utf8');

  // Split the file contents into an array of lines
  const lines = fileContents.split('\n');

  // Return the array of lines
  return lines;
}

