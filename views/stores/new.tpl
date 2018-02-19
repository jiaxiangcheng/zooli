<form class="ui form" enctype="multipart/form-data">
    <h2 id="title">
        <i class="shopping bag icon"></i>
        New Store
    </h2>
    {{template "stores/form/body.tpl" .}}
    <button id="save" class="ui primary button" type="submit">Create</button>
    <button id="cancel" class="ui button" type="button">Cancel</button>
</form>

<script type="text/javascript">
    $(document)
            .ready(function() {
                $('.ui.form')
                        .api({
                            url : '/stores/new',
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
                                url: "/stores",
                                success: function (data) {
                                    $('#main_content').html(data);
                                }
                            });
                        });
            });
</script>

