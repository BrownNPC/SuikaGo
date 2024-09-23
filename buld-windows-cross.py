import os
import shutil

# Set environment variables for cross-compilation
os.environ['CGO_ENABLED'] = '1'
os.environ['CC'] = 'x86_64-w64-mingw32-gcc'
os.environ['GOOS'] = 'windows'
os.environ['GOARCH'] = 'amd64'

# Define the distribution directory variable
distDir = "dist-windows"
zipFile = "game-windows"#.zip

print("Building...")
# Compile the Go project for Windows
os.system('go build -o game.exe -tags static -ldflags "-s -w -H windowsgui" src/*')

print("Packaging...")

# Clean up old distDir if it exists
if os.path.exists(distDir):
    shutil.rmtree(distDir)  # This will remove the directory and its contents

# Create a fresh distDir
os.mkdir(distDir)

# List of files and folders to copy
files = [
    'game.exe',
    'Config.json',
    'Magic.json',   
]
folders = [
    'assets'
]

# Copy files to distDir
for f in files:
    shutil.copy2(f, distDir)

# Copy directories to distDir
for f in folders:
    shutil.copytree(f, os.path.join(distDir, f))

# Create a zip archive of the distDir directory
shutil.make_archive( zipFile, 'zip', distDir)
shutil.move(zipFile + ".zip", 'dist-windows/')

print("Done!")
