<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>ZOOLI</title>

    {{ template "best_practice/common/header.tpl"}}

    </head>
    <body>
    <div class="ui huge blue inverted borderless fixed fluid menu">

        <a class="item" href="/dashboard">
            <i class="large database icon"></i> Zooli
        </a>
        <div class="right menu">
            <a class="item" href="/help">
                <i class="help icon"></i> Help
            </a>
        </div>
    </div>


            {{.LayoutContent}}

    <script type="text/javascript" src="/static/js/canvas-nest.min.js"></script>
    </body>
</html>