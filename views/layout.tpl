<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>ZOOLI</title>
<link rel="shortcut icon" href="/static/img/logo/pear.ico">
{{ template "common/header.tpl"}}

<body>
    <div class="ui huge blue inverted borderless fixed fluid menu">

        <a class="item" href="/dashboard">
            <i class="large database icon"></i> Zooli
        </a>
        <div class="right menu">
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
        </div>
    </div>

    {{.LayoutContent}}

    <script src="//cdn.bootcss.com/canvas-nest.js/1.0.1/canvas-nest.min.js"></script>
</body>
