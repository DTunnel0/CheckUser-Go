getArch(){
    case "$(uname -m)" in
        x86_64 | x64 | amd64 ) echo 'amd64' ;;
        armv8 | arm64 | aarch64 ) echo 'arm64' ;;
        * ) red "Unsupported CPU architecture! " && rm -f install.sh && exit 1 ;;
    esac
}

latest_release=$(curl -s https://api.github.com/repos/DTunnel0/CheckUser-Go/releases/latest | grep "tag_name" | cut -d'"' -f4)
arch=$(getArch)
name="checkuser-linux-$arch"
wget https://github.com/DTunnel0/CheckUser-Go/releases/download/$latest_release/$name -O checkuser
chmod +x checkuser
