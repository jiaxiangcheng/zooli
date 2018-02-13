<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>ZOOLI Dashboard</title>

{{ template "best_practice/common/header.tpl"}}

</head>
<body>
<div class="ui blue inverted menu">

    <a class="item" href="/dashboard">
        <i class="large database icon"></i> Zooli
    </a>
    <div class="right menu">
        <a class="item" href="/logout">
            <i class="external icon"></i> Log out
        </a>
    </div>
</div>


{{.LayoutContent}}

</body>

</html>