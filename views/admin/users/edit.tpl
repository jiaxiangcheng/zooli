<form class="ui form">
    <h2 id="title">
        <i class="user icon"></i>
        User Information
    </h2>

    {{template "admin/users/form/body.tpl" .}}
    <button id="save" class="ui primary button" type="submit">{{i18n .Lang "forms.save"}}</button>
    <button id="cancel" class="ui button" type="button">{{i18n .Lang "forms.cancel"}}</button>
</form>

<script type="text/javascript">
    $(document)
            .ready(function() {
                $('.ui.form')
                        .api({
                            url : "/admin/users/{{.userForm.ID}}",
                            method : 'POST',
                            serializeForm : true,
                            onSuccess    : function(response) {
                                $('#main_content').html(response);
                            },
                            onFailure    : function(response) {
                                $('#main_content').html(response);
                            }
                        });
                $('#cancel')
                        .on('click', function () {
                            $.ajax({
                                async: false,
                                type: "get",
                                url: "/admin/users",
                                success: function (data) {
                                    $('#main_content').html(data);
                                }
                            });
                        });
            });
</script>
