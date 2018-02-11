{{ if .flash.error }}
<p class="bg-danger" style="padding: 10px; border-radius: 5px">
{{str2html .flash.error}}
</p>
{{ end }}

{{ if .flash.success }}
<p class="bg-success" style="padding: 10px; border-radius: 5px">
{{str2html .flash.success}}
</p>
{{ end }}