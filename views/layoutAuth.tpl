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
<!-- side bar -->
{{template "side_bar.tpl" .}}

<div class="pusher">
    <!-- top menu-->
    <div class="navslide">
        {{template "top_menu.tpl" .}}
    </div>
    <!-- main content -->
    <div class="mainWrap navslide">
        <div class="ui equal width left aligned padded grid stackable" id="main_content">
            {{.LayoutContent}}
        </div>
    </div>
</div>

</body>
