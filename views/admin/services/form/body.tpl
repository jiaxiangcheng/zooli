<div class="ui five wid raised segment">
    <div class="ui error message"></div>
    {{template "common/flash.tpl" .}}
    <div class="field">
        <label>{{i18n .Lang "forms.name"}}</label>
        <input name="name" value="{{.serviceForm.Name}}" type="text" placeholder="{{i18n .Lang "forms.name"}}"/>
    </div>
</div>

<script>
    $(document)
            .ready(function() {
                $('.ui.form')
                        .form({
                            fields: {
                                servicename: {
                                    identifier  : 'name',
                                    rules: [
                                        {
                                            type   : 'empty',
                                            prompt : 'Please enter your service name'
                                        }
                                    ]
                                }
                            }
                        });
            });
</script>
