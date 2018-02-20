<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>ZOOLI</title>
    <link rel="shortcut icon" href="/static/img/logo/pear.ico">
    {{ template "common/header.tpl"}}
    <link rel="stylesheet" href="/static/css/menu.css"/>
    <script type="text/javascript" src="/static/js/nav.js"></script>

</head>
<body>
<div class="ui blue huge inverted borderless fixed fluid menu">

    <a class="item" href="/dashboard">
        <i class="large database icon"></i> Zooli
    </a>
    <div class="right menu">
        <span class="item">Hi {{.user.Name}}</span>

        <a class="item" href="/help">
            <i class="help icon"></i> Help
        </a>
        <a class="item" href="/logout">
            <i class="sign out icon"></i> Log out
        </a>
    </div>
</div>

<div class="ui grid">
    <div class="row">
        <div class="column" id="sidebar">
            <div class="ui secondary vertical fluid menu" id="nav_menu">
                <a class="item" href="/dashboard">
                    <span><i class="dashboard icon"></i>Dashboard</span>
                </a>

                {{ if eq .user.Role.Name .roleAdmin}}

                <a class="item" href="/users">
                    <span><i class="users icon"></i>Users</span>
                </a>
                <a class="item" href="/companies">
                    <span><i class="world icon"></i>Companies</span>
                </a>
                <a class="item" href="/services">
                    <span><i class="cubes icon"></i>Services</span>
                </a>
                <a class="item" href="/stores">
                    <span><i class="shopping bag icon"></i>Stores</span>
                </a>
                <div class="ui hidden divider"></div>

                <a class="item" href="/analytics">
                    <span><i class="bar chart icon"></i>Analytics</span>
                </a>
                {{end}}

                {{ if eq .user.Role.Name .roleManager}}
                        <a class="item" href="/stores/{{.user.StoreID}}">
                            <span><i class="world icon"></i>My Store</span>
                        </a>

                        <a class="item" href="/orders">
                            <span><i class="payment icon"></i>Orders</span>
                        </a>

                        <a class="item" href="/products">
                            <span><i class="cubes icon"></i>Products</span>
                        </a>
                {{end}}
            </div>
        </div>

        <div class="column" id="main_content">
        {{.LayoutContent}}
        </div>
    </div>
</div>
<script type="text/javascript" src="/static/js/canvas-nest.min.js"></script>
</body>

</html>
