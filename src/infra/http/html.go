package http

const HTML_CONTENT = `
<!DOCTYPE HTML>
<html>

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>DT - CHECKUSER</title>

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

        .user {
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

        .details {
            margin-top: 10px;
            display: none;
            flex-direction: column;
            gap: 5px;
            padding: 5px;
            border-radius: 10px;
            border: #2b2c359d solid 2px;
        }

        .detail {
            background: #2c2d30;
            border-radius: 10px;
            color: #FFF;
            font-size: 13px;
            text-align: center;
            padding: 5px;
            font-family: "Roboto", sans-serif;
        }

        .details-not-found {
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
            <h4>TOTAL DE CONEXÕES</h4>
            <div class="container-count">
                <span id="total">00</span>
            </div>
        </div>

        <div class="user">
            <input type="text" class="search" placeholder="Digite um nome...">
            <div class="details">
                <span class="detail"></span>
                <span class="detail"></span>
                <span class="detail"></span>
                <span class="detail"></span>
            </div>
            <div class="details-not-found">
                Usuário não encontrado!
            </div>
        </div>
    </div>
    <script>
        let timeout = null

        const userNotFoundElement = document.querySelector('.details-not-found')
        const [nameElement, expiresAtElement, limitElement, connectionsElement] = document.querySelectorAll('.detail')
        const details = document.querySelector('.details')
        const search = document.querySelector('.search')

        const showUserNotFound = () => {
            userNotFoundElement.style.display = 'block'
            details.style.display = 'none'
        }

        const hideUserNotFound = () => {
            details.style.display = 'flex'
            userNotFoundElement.style.display = 'none'
        }

        const hideUserNotFoundAndDetails = () => {
            userNotFoundElement.style.display = 'none'
            details.style.display = 'none'
        }

        const searchHandler = value => setTimeout(async () => {
            if (!value){
                hideUserNotFoundAndDetails()
                return
            }
            
            const url = "/details/" + value
            const data = await fetch(url).then(r => r.json())

            if (!data?.username) {
                showUserNotFound()
                return
            }

            hideUserNotFound()

            nameElement.innerHTML = data.username ?? ''
            expiresAtElement.innerHTML = data.expires_at ?? ''
            limitElement.innerHTML = data.limit?.toString()?.padStart(2, '0') ?? ''
            connectionsElement.innerHTML = data.connections?.toString()?.padStart(2, '0') ?? ''
        }, 500)

        search.addEventListener('keyup', e => {
            clearTimeout(timeout)
            timeout = searchHandler(e.target.value)
        })

        const main = async () => {
            const data = await fetch('/count').then(e => e.json())
            document.querySelector('#total').innerHTML = data.count.toString().padStart(2, '0')
        }
        main()
    </script>
</body>

</html>
`
