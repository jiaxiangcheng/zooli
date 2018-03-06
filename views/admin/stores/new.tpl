<form class="ui form" enctype="multipart/form-data">
    <h2 id="title" style="margin-top: 15px">
        <i class="shopping bag icon"></i>
        {{i18n .Lang "stores_table.new_store"}}
    </h2>
    {{template "admin/stores/form/body.tpl" .}}
    <button id="save" class="ui primary button" type="submit">{{i18n .Lang "forms.create"}}</button>
    <button id="cancel" class="ui button" type="button">{{i18n .Lang "forms.cancel"}}</button>
</form>

<script type="text/javascript">
    $(document)
            .ready(function() {
                $('.ui.form')
                        .api({
                            url : 'admin/stores/new',
                            method : 'POST',
                            cache: false,
                            processData: false,
                            contentType: false,
                            beforeSend: (settings)=>{
                            settings.data = new FormData($(".ui.form")[0]);
                            return settings;
                            },
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
                                url: "/admin/stores",
                                success: function (data) {
                                    $('#main_content').html(data);
                                }
                            });
                        });
            });
</script>
