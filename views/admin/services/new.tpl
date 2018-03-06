<form class="ui form">
    <h2 id="title" style="margin-top: 15px">
        <i class="user icon"></i>
        {{i18n .Lang "services_table.new_service"}}
    </h2>
    {{template "admin/services/form/body.tpl" .}}
    <button id="save" class="ui primary button" type="submit">{{i18n .Lang "forms.create"}}</button>
    <button id="cancel" class="ui button" type="button">{{i18n .Lang "forms.cancel"}}</button>
</form>

<script type="text/javascript">
    $(document)
            .ready(function() {
                $('.ui.form')
                        .api({
                            url : 'admin/services/new',
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
                                url: "/admin/services",
                                success: function (data) {
                                    $('#main_content').html(data);
                                }
                            });
                        });
            });
</script>
