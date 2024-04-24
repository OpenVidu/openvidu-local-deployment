package templates

type App502Data struct {
	Title   string
	Message string
}

const App502Template = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>502 - Application Not Found</title>
    <!-- Bootstrap CSS CDN -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
    <!-- Custom styles -->
    <style>
        body {
            padding-top: 50px;
            background-color: #f7f7f7;
        }
        .container {
            padding: 40px;
            background: #fff;
            border-radius: 4px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            max-width: 600px;
            margin: auto;
            text-align: center;
        }
        .error-code {
            font-size: 45px;
            font-weight: bold;
        }
    </style>
</head>
<body>
    <div class="container">
        <div>
            <div class="error-code">502 - Bad Gateway</div>
            <h1 class="display-5">{{.Title}}</h1>
			<hr class="my-4">
            <p>{{ .Message }}</p>
        </div>
    </div>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
`