#!/bin/bash

get_arch() {
    case "$(uname -m)" in
        x86_64 | x64 | amd64 ) echo 'amd64' ;;
        armv8 | arm64 | aarch64 ) echo 'arm64' ;;
        * ) echo 'unsupported' ;;
    esac
}

install_checkuser() {
    local latest_release=$(curl -s https://api.github.com/repos/DTunnel0/CheckUser-Go/releases/latest | grep "tag_name" | cut -d'"' -f4)
    local arch=$(get_arch)

    if [ "$arch" = "unsupported" ]; then
        echo "Arquitetura de CPU não suportada!"
        exit 1
    fi

    local name="checkuser-linux-$arch"
    echo "Baixando $name..."
    wget -q "https://github.com/DTunnel0/CheckUser-Go/releases/download/$latest_release/$name" -O /usr/local/bin/checkuser
    chmod +x /usr/local/bin/checkuser

    read -p "Porta: " -ei 8000 port

    if [ -z "$port" ]; then
        echo "Porta não fornecida. Saindo."
        exit 1
    fi

    if systemctl status checkuser >/dev/null 2>&1; then
        echo "Parando o serviço checkuser existente..."
        sudo systemctl stop checkuser
        sudo systemctl disable checkuser
        sudo rm /etc/systemd/system/checkuser.service
        sudo systemctl daemon-reload
        echo "Serviço checkuser existente foi parado e removido."
    fi

    cat << EOF | sudo tee /etc/systemd/system/checkuser.service > /dev/null
[Unit]
Description=CheckUser Service
After=network.target nss-lookup.target

[Service]
User=root
CapabilityBoundingSet=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
NoNewPrivileges=true
ExecStart=/usr/local/bin/checkuser --start --port $port
Restart=always

[Install]
WantedBy=multi-user.target
EOF

    sudo systemctl daemon-reload
    sudo systemctl start checkuser
    sudo systemctl enable checkuser

    local addr=$(curl -s icanhazip.com)
    echo "URL: http://$addr:$port"
    echo "O serviço CheckUser foi instalado e iniciado."
    read
}

reinstall_checkuser() {
    sudo systemctl stop checkuser
    sudo systemctl disable checkuser
    sudo rm /usr/local/bin/checkuser
    sudo rm /etc/systemd/system/checkuser.service
    sudo systemctl daemon-reload
    echo "Serviço checkuser removido."

    install_checkuser
}

uninstall_checkuser() {
    sudo systemctl stop checkuser
    sudo systemctl disable checkuser
    sudo rm /usr/local/bin/checkuser
    sudo rm /etc/systemd/system/checkuser.service
    sudo systemctl daemon-reload
    echo "Serviço checkuser removido."
    read
}

main() {
    clear

    echo -n 'CHECKUSER MENU '
    if [[ -e /usr/local/bin/checkuser ]]; then
        echo -e '\e[32m[INSTALADO]\e[0m - Versao:' $(/usr/local/bin/checkuser --version | cut -d' ' -f2)
    else
        echo -e '\e[31m[DESINSTALADO]\e[0m'
    fi

    echo
    echo '[01] - INSTALAR CHECKUSER'
    echo '[02] - REINSTALAR CHECKUSER'
    echo '[03] - DESINSTALAR CHECKUSER'
    echo '[00] - SAIR'
    echo
    read -p 'Escolha uma opção: ' option

    case $option in
        1) install_checkuser; main ;;
        2) reinstall_checkuser; main ;;
        3) uninstall_checkuser; main ;;
        4) echo "Saindo." && break ;;
        *) echo "Opção inválida. Tente novamente.";read; main ;;
    esac
}

main