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

        .container h4 {
            color: #FFF;
            font-size: 1rem;
            margin-bottom: 20px;
        }

        .container .container-count {
            width: 80px;
            height: 80px;
            color: #FFF;
            background: #363636;
            padding: 1rem;
            border-radius: 50%;
            font-size: 3em;
            display: flex;
            align-items: center;
            justify-content: center;
        }
    </style>
</head>

<body>
    <div class="container">
        <h1>CHECKUSER - @DuTra01</h1>
        <h4>TOTAL DE CONEXÕES</h4>
        <div class="container-count">
            <span id="total">00</span>
        </div>
    </div>
    <script>
        const main = async() => {
            const data = await fetch('/count').then(e => e.json())
            document.querySelector('#total').innerHTML = String(data.data).padStart(2, '0')
        }
        main()
    </script>
</body>

</html>
`
