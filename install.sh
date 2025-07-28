#!/bin/bash

get_arch() {
    case "$(uname -m)" in
        x86_64 | x64 | amd64 ) echo 'amd64' ;;
        armv8 | arm64 | aarch64 ) echo 'arm64' ;;
        * ) echo 'unsupported' ;;
    esac
}

check_url_access() {
    local test_url=$1
    echo -e "\n🔍 Testando acesso externo a: $test_url"
    
    if curl -s --max-time 5 "$test_url" >/dev/null; then
        echo -e "\e[1;32m✅ A URL está acessível externamente.\e[0m"
        return
    fi

    echo -e "\e[1;31m❌ Não foi possível acessar a URL externamente.\e[0m"
    echo -ne "\e[1;33mDeseja abrir a porta no iptables automaticamente? [s/N]: \e[0m"
    read answer

    if [[ "$answer" =~ ^[Ss]$ ]]; then
        local port=$(echo "$test_url" | grep -oE ':[0-9]+' | tr -d ':')
        sudo iptables -I INPUT -p tcp --dport "$port" -j ACCEPT
        sudo iptables-save > /etc/iptables.rules
        echo -e "\e[1;32m✔ Porta $port liberada no iptables.\e[0m"
        return
    fi

    echo -e "\e[1;33m⚠ Porta não foi aberta. Faça isso manualmente se necessário.\e[0m"
}

install_checkuser() {
    local latest_release=$(curl -s https://api.github.com/repos/DTunnel0/CheckUser-Go/releases/latest | grep "tag_name" | cut -d'"' -f4)
    local arch=$(get_arch)

    if [ "$arch" = "unsupported" ]; then
        echo -e "\e[1;31mArquitetura de CPU não suportada!\e[0m"
        exit 1
    fi

    local name="checkuser-linux-$arch"
    echo "⬇️  Baixando $name..."
    wget -q "https://github.com/DTunnel0/CheckUser-Go/releases/download/$latest_release/$name" -O /usr/local/bin/checkuser
    chmod +x /usr/local/bin/checkuser

    local addr=$(curl -s https://ipv4.icanhazip.com)
    local domain_json=$(curl -s https://dns.dtunnel.com.br/api/v1/dns/create -X POST --data '{"content": "'"$addr"'", "proxied": true}')
    local url=$(echo "$domain_json" | grep -o '"domain": *"[^"]*"' | grep -o '"[^"]*"$' | tr -d '"')

    local port="2052"
    local sslEnabled=""
    local final_url="http://$addr:$port"

    if [[ -n "$url" ]]; then
        port="2053"
        sslEnabled="--ssl"
        final_url="https://$url:$port"
    fi

    if systemctl is-active --quiet checkuser; then
        echo "🛑 Parando serviço checkuser existente..."
        systemctl stop checkuser
        systemctl disable checkuser
        rm -f /etc/systemd/system/checkuser.service
        systemctl daemon-reload
    fi

    cat << EOF | tee /etc/systemd/system/checkuser.service > /dev/null
[Unit]
Description=CheckUser Service
After=network.target nss-lookup.target

[Service]
User=root
CapabilityBoundingSet=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
NoNewPrivileges=true
ExecStart=/usr/local/bin/checkuser --start --port $port $sslEnabled
Restart=always

[Install]
WantedBy=multi-user.target
EOF

    systemctl daemon-reload &>/dev/null
    systemctl start checkuser &>/dev/null
    systemctl enable checkuser &>/dev/null

    echo -e "\n\e[1;32m✅ CheckUser instalado com sucesso!\e[0m"
    echo -e "\e[1;34m🌐 URL: \e[1;36m$final_url\e[0m"

    check_url_access "$final_url"

    echo -e "\nPressione Enter para continuar..."
    read
}

reinstall_checkuser() {
    echo "♻️  Reinstalando CheckUser..."
    systemctl stop checkuser &>/dev/null
    systemctl disable checkuser &>/dev/null
    rm -f /usr/local/bin/checkuser /etc/systemd/system/checkuser.service
    systemctl daemon-reload &>/dev/null
    install_checkuser
}

uninstall_checkuser() {
    echo "🧹 Desinstalando CheckUser..."
    systemctl stop checkuser &>/dev/null
    systemctl disable checkuser &>/dev/null
    rm -f /usr/local/bin/checkuser /etc/systemd/system/checkuser.service
    systemctl daemon-reload &>/dev/null
    echo -e "\e[1;31m✔ CheckUser removido.\e[0m"
    echo -e "\nPressione Enter para continuar..."
    read
}

main() {
    clear
    echo '---------------------------------'
    echo -ne '     \e[1;33mCHECKUSER\e[0m'
    if [[ -x /usr/local/bin/checkuser ]]; then
        echo -e ' \e[1;32mv'$(/usr/local/bin/checkuser --version | cut -d' ' -f2)'\e[0m'
    fi

    if [[ ! -x /usr/local/bin/checkuser ]]; then
        echo -e ' \e[1;31m[DESINSTALADO]\e[0m'
    fi
    echo '---------------------------------'

    echo -e '\e[1;32m[01] - \e[1;31mINSTALAR CHECKUSER\e[0m'
    echo -e '\e[1;32m[02] - \e[1;31mREINSTALAR CHECKUSER\e[0m'
    echo -e '\e[1;32m[03] - \e[1;31mDESINSTALAR CHECKUSER\e[0m'
    echo -e '\e[1;32m[00] - \e[1;31mSAIR\e[0m'
    echo '---------------------------------'
    echo -ne '\e[1;32mEscolha uma opção: \e[0m'; 
    read option

    case $option in
        1) install_checkuser; main ;;
        2) reinstall_checkuser; main ;;
        3) uninstall_checkuser; main ;;
        0) echo "Saindo." ;;
        *) echo -e "\e[1;31mOpção inválida. Tente novamente.\e[0m"; read; main ;;
    esac
}

main
