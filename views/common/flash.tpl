<script type="text/javascript" src="/static/js/message.js"></script>
{{ if .flash.error }}
<div class="ui negative message">
    <i class="close icon"></i>
    <div class="header">
        Error
    </div>
    <p>{{str2html .flash.error}}</p>
</div>
{{ end }}

{{ if .flash.success }}
<div class="ui positive message">
    <i class="close icon"></i>
    <div class="header">
        Success
    </div>
    <p>{{str2html .flash.success}}</p>
</div>
{{ end }}