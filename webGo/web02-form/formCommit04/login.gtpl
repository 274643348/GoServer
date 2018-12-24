
<html>
<head>
    <meta charset="UTF-8">
    <title>login</title>
</head>
<body>
<form action ="/login" method="post">
    用户名:<input type="text" name="username">
    秘密:<input type="password" name="password">
    <input type="hidden" name="token" value={{.}}>
    <button type="summit">登录</button>

</form>
</body>
</html>