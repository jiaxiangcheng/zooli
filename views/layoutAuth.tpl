<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>ZOOLI</title>
    <link rel="shortcut icon" href="/static/img/logo/pear.ico">
    {{ template "common/header.tpl"}}
    <link rel="stylesheet" href="/static/css/main.css"/>
    <script type="text/javascript" src="/static/js/main.js"></script>

</head>

<body>
<div class="ui sidebar vertical left menu overlay borderless visible sidemenu inverted grey" id="nav_sidebar" style="-webkit-transition-duration: 0.1s; transition-duration: 0.1s;" data-color="grey">
    <a class="item logo" href="/dashboard">
        <img class="ui mini image spaced" src="/static/img/logo/z.png" alt="Zooli" /><span style="color: black;">   Zooli</span>
    </a>

    <a class="item" href="/dashboard">
        <i class="dashboard icon"></i><span>Dashboard</span>
    </a>

    {{ if eq .user.Role.Name .roleAdmin}}

        <a class="item" href="/admin/users">
            <i class="users icon"></i>
            <span>Users</span>
        </a>
        <a class="item" href="/admin/companies">
            <i class="world icon"></i>
            <span>Companies</span>
        </a>
        <a class="item" href="/admin/services">
            <i class="cubes icon"></i>
            <span>Services</span>
        </a>
        <a class="item" href="/admin/stores">
            <i class="shopping bag icon"></i>
            <span>Stores</span>
        </a>
        <div class="ui divider"></div>
        <a class="item" href="/admin/analytics">
            <i class="bar chart icon"></i>
            <span>Analytics</span>
        </a>
        <div class="ui divider"></div>
        <a class="item">
            <div class="ui inverted progress tiny yellow" id="sidebar_progress1">
                <div class="bar">

                </div>
                <div class="label colhidden" style="margin-top: 10px"><span class="colhidden">Monthly Bandwidth Transfer</span></div>
            </div>

        </a>

        <a class="item">
            <div class="ui inverted progress tiny teal" id="sidebar_progress2">
                <div class="bar">

                </div>
                <div class="label colhidden" style="margin-top: 10px"><span class="colhidden">Disk Space Usage</span></div>
            </div>

        </a>
        <a class="item">
            <div class="ui inverted progress tiny blue" id="sidebar_progress3">
                <div class="bar">

                </div>
                <div class="label colhidden" style="margin-top: 10px"><span class="colhidden">Earn money</span></div>
            </div>

        </a>
        <div class="ui divider"></div>
    {{end}}

    {{ if eq .user.Role.Name .roleManager}}
    {{if .user.StoreID}}
        <a class="item" href="/public/store">
            <span><i class="world icon"></i>My Store</span>
        </a>

        <a class="item" href="/public/orders">
            <span><i class="payment icon"></i>Orders</span>
        </a>

        <a class="item" href="/public/products">
            <span><i class="cubes icon"></i>Products</span>
        </a>
    {{end}}
    {{end}}

</div>

<div class="pusher">
    <!--navbar-->
    <div class="navslide">
        <div class="ui menu icon borderless grid blue">
            <a class="item labeled" id="open_btn">
                <i class="bars big icon"></i>
            </a>
            <a class="item labeled expandit" id="full_screen_toggle">
                <i class="expand arrows alternate big icon"></i>
            </a>
            <div class="item ui colhidden">
                <div class="ui icon input">
                    <input type="text" placeholder="Search...">
                    <i class="search icon"></i>
                </div>
            </div>
            <div class="right menu colhidden">

            {{ if eq .user.Role.Name .roleAdmin}}
                    {{template "random.tpl"}}
            {{end}}

                <div class="ui dropdown item" id="select_language">
                {{.CurLang}} <i class="dropdown icon"></i>
                    <div class="menu">
                    {{range .RestLangs}}
                        <a class="item" data-value="{{.Lang}}">{{i18n $.Lang .Name}}</a>
                    {{end}}
                    </div>
                </div>
                <a class="item" href="/help">
                    <i class="help icon"></i> {{i18n $.Lang "help"}}
                </a>
                <a class="item" href="/logout">
                    <i class="sign out icon"></i> {{i18n $.Lang "log out"}}
                </a>

            </div>
        </div>
    </div>
    <!--navbar-->
    <!--maincontent-->
    <div class="mainWrap navslide">
        <div class="ui equal width left aligned padded grid stackable" id="main_content">
            {{.LayoutContent}}
        </div>
    </div>
    <!--maincontent-->
</div>

</body>
