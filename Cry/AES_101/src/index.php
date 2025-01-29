<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Encrypt Data</title>

    <style>
    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }

    body {
        font-family: 'Arial', sans-serif;
        background: #f3f4f6;
        display: flex;
        justify-content: center;
        align-items: center;
        min-height: 100vh;
    }

    .container {
        background: #ffffff;
        padding: 40px;
        border-radius: 10px;
        box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
        max-width: 400px;
        width: 100%;
    }

    h1 {
        font-size: 24px;
        color: #333;
        margin-bottom: 20px;
        text-align: center;
    }

    form {
        display: flex;
        flex-direction: column;
    }

    textarea {
        padding: 12px;
        border: 1px solid #ddd;
        border-radius: 8px;
        resize: none;
        margin-bottom: 15px;
        min-height: 100px;
        font-size: 16px;
    }

    button {
        padding: 12px;
        border: none;
        background-color: #4a90e2;
        color: white;
        font-size: 16px;
        border-radius: 8px;
        cursor: pointer;
    }

    button:hover {
        background-color: #357ab8;
    }

    .result {
        background: #f0f0f0;
        padding: 10px;
        margin-top: 15px;
        border-radius: 8px;
        word-break: break-word;
    }
    </style>
    </head>
    <body>

    <div class="container">
    <h1>Data Encryption Form</h1>
    <form action="chall.php" method="POST">
        <label for="encrypt">Enter Text to Encrypt:</label>
        <textarea id="encrypt" name="encrypt" placeholder="Enter your text here..."></textarea>
        <button type="submit">Encrypt Text</button>
    </form>
    </div>

</body>
</html>
