<form class="ui form" enctype="multipart/form-data">
    <h2 id="title" style="margin-top: 15px">
        <i class="world icon"></i>
        {{i18n .Lang "products_table.new_product"}}
    </h2>

    {{template "public/products/form/body.tpl" .}}
    <button id="save" class="ui primary button" type="submit">{{i18n .Lang "forms.create"}}</button>
    <button id="cancel" class="ui button" type="button">{{i18n .Lang "forms.cancel"}}</button>
</form>

<script type="text/javascript">
    $(document)
            .ready(function() {
                $('.ui.form')
                        .api({
                            url : 'public/products/new',
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
                                url: "public/products",
                                success: function (data) {
                                    $('#main_content').html(data);
                                }
                            });
                        });
            });
</script>
