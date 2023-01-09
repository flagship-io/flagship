#!/usr/bin/env bash

echo "Installing Flagship CLI..."
echo "------------------------"

# Determining the Linux distribution and architecture
distro=$(lsb_release -i -s)
arch=$(uname -m)

echo "Distribution: $distro"
echo "Architecture: $arch"

# Flagship CLI version
version="0.3.1"

echo "Version: v$version"
echo "------------------------"

# URL for downloading the archive based on the distribution and architecture
url=""

case "$distro" in
  "Darwin")
    case "$arch" in
      "x86_64")
        url="https://github.com/flagship-io/flagship/releases/download/${version}/flagship_${version}-darwin-amd64.tar.gz"
        ;;
      "arm64")
        url="https://github.com/flagship-io/flagship/releases/download/${version}/flagship_${version}-darwin-arm64.tar.gz"
        ;;
      *)
        echo "Unsupported architecture"
        exit 1
        ;;
    esac
    ;;
  "Ubuntu"|"Debian"|"Raspbian")
  echo "Downloading Flagship CLI..."
    case "$arch" in
      "i686")
        url="https://github.com/flagship-io/flagship/releases/download/${version}/flagship_${version}-linux-386.tar.gz"
        ;;
      "x86_64")
        url="https://github.com/flagship-io/flagship/releases/download/${version}/flagship_${version}-linux-amd64.tar.gz"
        echo $url
        ;;
      "arm64")
        url="https://github.com/flagship-io/flagship/releases/download/${version}/flagship_${version}-linux-arm64.tar.gz"
        ;;
      *)
        echo "Unsupported architecture"
        exit 1
        ;;
    esac
    ;;
  *)
    echo "Unsupported distribution"
    exit 1
    ;;
esac

# Downloading the archive to home directory (and check if url is not 404)
echo "Downloading Flagship CLI..."
wget -q --spider $url
if [ $? -eq 0 ]; then
  wget -O ~/flagship.tar.gz $url -q --show-progress
else
  echo "------------------------"
  echo "Flagship CLI archive not found"
  echo "------------------------"
  exit 1
fi

# Extracting the archive (if it exists)
echo "Extracting Flagship CLI..."
if [ -f ~/flagship.tar.gz ]; then
  tar -xzf ~/flagship.tar.gz -C ~/
else
  echo "Flagship CLI archive not found"
  exit 1
fi

# Removing the archive
echo "Removing archive..."
rm ~/flagship.tar.gz

# Moving the binary to /usr/local/bin
echo "Moving Flagship CLI to /usr/local/bin..."
sudo mv ~/flagship /usr/local/bin/

# Making the binary executable
echo "Making Flagship CLI executable..."
sudo chmod +x /usr/local/bin/flagship

# Sending a message to the user
echo "-----------------------------------------"
echo "Flagship CLI successfully installed"
echo "-----------------------------------------"