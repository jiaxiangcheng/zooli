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