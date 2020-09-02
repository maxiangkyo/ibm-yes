version=v4.27.0
curl -sSL https://github.com/v2ray/v2ray-core/releases/download/$version/v2ray-linux-64.zip > v2ray-linux-64.zip
unzip -o v2ray-linux-64.zip && rm -f v2ray-linux-64.zip
chmod +x v2ctl v2ray