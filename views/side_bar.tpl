<div class="ui sidebar vertical left menu overlay borderless visible sidemenu inverted grey" id="nav_sidebar" style="-webkit-transition-duration: 0.1s; transition-duration: 0.1s;" data-color="grey">
    <a class="item logo" href="/dashboard">
        <img class="ui mini image spaced" src="/static/img/logo/z.png" alt="Zooli" /><span style="color: black;">   Zooli</span>
    </a>

    <a class="item" href="/dashboard">
        <i class="dashboard icon"></i><span>{{i18n .Lang "nav_items.dashboard"}}</span>
    </a>

{{ if eq .user.Role.Name .roleAdmin}}

    <a class="item" href="/admin/users">
        <i class="users icon"></i>
        <span>{{i18n .Lang "nav_items.users"}}</span>
    </a>
    <a class="item" href="/admin/companies">
        <i class="world icon"></i>
        <span>{{i18n .Lang "nav_items.companies"}}</span>
    </a>
    <a class="item" href="/admin/services">
        <i class="cubes icon"></i>
        <span>{{i18n .Lang "nav_items.services"}}</span>
    </a>
    <a class="item" href="/admin/stores">
        <i class="shopping bag icon"></i>
        <span>{{i18n .Lang "nav_items.stores"}}</span>
    </a>
    <div class="ui divider"></div>
    <a class="item" href="/admin/analytics">
        <i class="bar chart icon"></i>
        <span>{{i18n .Lang "nav_items.analytics"}}</span>
    </a>
    <div class="ui divider"></div>
    <a class="item">
        <div class="ui inverted progress tiny yellow" id="sidebar_progress1">
            <div class="bar">

            </div>
            <div class="label colhidden" style="margin-top: 10px"><span class="colhidden">{{i18n .Lang "nav_items.month_bandwidth"}}</span></div>
        </div>

    </a>

    <a class="item">
        <div class="ui inverted progress tiny teal" id="sidebar_progress2">
            <div class="bar">

            </div>
            <div class="label colhidden" style="margin-top: 10px"><span class="colhidden">{{i18n .Lang "nav_items.disk_usage"}}</span></div>
        </div>

    </a>
    <a class="item">
        <div class="ui inverted progress tiny blue" id="sidebar_progress3">
            <div class="bar">

            </div>
            <div class="label colhidden" style="margin-top: 10px"><span class="colhidden">{{i18n .Lang "nav_items.earn_money"}}</span></div>
        </div>

    </a>
    <div class="ui divider"></div>
{{end}}

{{ if eq .user.Role.Name .roleManager}}
{{if .user.StoreID}}
    <a class="item" href="/public/store">
        <i class="world icon"></i>
        <span>{{i18n .Lang "nav_items.my_store"}}</span>
    </a>

    <a class="item" href="/public/orders">
        <i class="payment icon"></i>
        <span>{{i18n .Lang "nav_items.orders"}}</span>
    </a>

    <a class="item" href="/public/products">
        <i class="cubes icon"></i>
        <span>{{i18n .Lang "nav_items.products"}}</span>
    </a>
{{end}}
{{end}}
</div>