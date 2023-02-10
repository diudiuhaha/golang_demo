<html lang="zh-CN">
<head>
    <title></title>
</head>
<body>
<form action="/login" method="post">
    <label>
        用户名:
        <input type="text" name="username">
    </label>
    <label>
        密码:
        <input type="password" name="password">
    </label>
    <label>
        <select name="fruit">
            <option value="apple">apple</option>
            <option value="pear">pear</option>
            <option value="banana">banana</option>
            <option value="other">other</option>
        </select>
    </label>
    <br>
    <label>
        <input type="checkbox" name="interest" value="football">足球
        <input type="checkbox" name="interest" value="basketball">篮球
        <input type="checkbox" name="interest" value="tennis">网球
    </label>
    <br>
    <label>
        <input type="radio" name="gender" value="1">男
        <input type="radio" name="gender" value="2">女

    </label>
    <br>
    <input type="submit" value="登录">
</form>
</body>
</html>
