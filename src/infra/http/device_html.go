package http

const DEVICE_HTML_CONTENT = `
<!DOCTYPE HTML>
<html>

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>DT - DEVICES</title>

    <style>
        @import url(https://fonts.googleapis.com/css?family=Roboto:400,300,100,500,700,900);

        html,
        body {
            margin: 0;
            padding: 0;
            font-family: 'Roboto', sans-serif;
        }

        body {
            display: flex;
            justify-content: center;
            align-items: center;
            background: linear-gradient(135deg, #2c2d30 0%, #353535 100%);
            width: 100%;
            height: 100vh;
        }

        .container {
            display: flex;
            flex-direction: column;
            align-items: center;
            margin: 1rem;
            border-radius: 50px;
            border: none;
            background: rgb(39, 39, 39);
            width: 300px;
            padding: 30px;
            box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
        }

        .container h1 {
            color: #FFF;
            font-size: 1.5rem;
            margin-bottom: 5px;
        }

        .container .connectionsArea h4 {
            color: #FFF;
            font-size: 1rem;
            margin-bottom: 20px;
        }

        .connectionsArea {
            display: flex;
            align-items: center;
            gap: 5px;
        }

        .container-count {
            color: #FFF;
            background: #363636;
            padding: 0.3rem;
            border-radius: 50%;
            font-size: 1em;
            display: flex;
            align-items: center;
            justify-content: center;
        }

        .devices {
            display: flex;
            flex-direction: column;
            width: 80%;
        }

        .search {
            padding: 10px;
            background: #2c2d30;
            outline: none;
            border: 0;
            border-radius: 10px;
            color: #FFF;
        }

        .device-id-list {
            margin-top: 10px;
            display: flex;
            flex-direction: column;
            gap: 5px;
            padding: 5px;
            border-radius: 10px;
            border: #2b2c359d solid 2px;
            height: 200px;
            overflow-y: scroll;
        }

        .device-id-list::-webkit-scrollbar {
            display: none;
        }

        .device-id {
            background: #2c2d30;
            border-radius: 10px;
            color: #FFF;
            font-size: 13px;
            text-align: center;
            padding: 5px;
            font-family: "Roboto", sans-serif;
        }

        .device-id-not-found {
            display: none;
            text-align: center;
            margin: 10px 0;
            color: #ff0000;
            background: #2c2d30;
            border-radius: 10px;
            padding: 5px;
            font-family: "Roboto", sans-serif;
        }
    </style>
</head>

<body>
    <div class="container">
        <h1>CHECKUSER - @DuTra01</h1>
        <div class="connectionsArea">
            <h4>TOTAL DE DEVICE ID</h4>
            <div class="container-count">
                <span id="total">00</span>
            </div>
        </div>

        <div class="devices">
            <input type="text" class="search" placeholder="Digite um nome...">
            <div class="device-id-list"></div>
            <div class="device-id-not-found">
                Dispositivos não encontrados!
            </div>
        </div>
    </div>
    <script>
        let timeout = null

        const deviceListElement = document.querySelector('.device-id-list')
        const search = document.querySelector('.search')
        const devicesNotFoundElement = document.querySelector('.device-id-not-found')

        const createDeviceIDElement = deviceID => {
            const element = document.createElement('span')
            element.className = 'device-id'
            element.innerHTML = deviceID
            return element
        }
        const appendDeviceIDElementInList = element => deviceListElement.appendChild(element)

        const cleanDeviceListElement = () => deviceListElement.innerHTML = ''

        const showDevicesNotFound = () => {
            devicesNotFoundElement.style.display = 'block'
            deviceListElement.style.display = 'none'
        }

        const hideDevicesNotFound = () => {
            deviceListElement.style.display = 'flex'
            devicesNotFoundElement.style.display = 'none'
        }

        const searchHandler = value => setTimeout(async () => {
            if (!value) {
                showDevices()
                return
            }

            const url = "/devices/list/" + value
            const data = await fetch(url).then(r => r.json())

            if (!Array.isArray(data) || data.length == 0) {
                showDevicesNotFound()
                return
            }

            hideDevicesNotFound()
            cleanDeviceListElement()
            data.forEach(device => {
                const element = createDeviceIDElement(device)
                appendDeviceIDElementInList(element)
            })
        }, 500)

        search.addEventListener('keyup', e => {
            clearTimeout(timeout)
            timeout = searchHandler(e.target.value)
        })

        const showDevices = async () => {
            const url = '/devices/list'
            const data = await fetch(url).then(e => e.json())

            hideDevicesNotFound()
            cleanDeviceListElement()
            data.forEach(device => {
                const element = createDeviceIDElement(device.username + " - " + device.id)
                appendDeviceIDElementInList(element)
            })
        }
        showDevices()

        const main = async () => {
            const url = '/devices/count'
            const data = await fetch(url).then(e => e.json())
            document.querySelector('#total').innerHTML = data.count.toString().padStart(2, '0')
        }
        main()
    </script>
</body>

</html>
`
