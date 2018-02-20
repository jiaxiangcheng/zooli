<form class="ui form">
    <h2 id="title">
        <i class="cubes icon"></i>
        Product Information
    </h2>

    {{template "products/form/body.tpl" .}}
    <button id="save" class="ui primary button" type="submit">Save</button>
    <button id="cancel" class="ui button" type="button">Cancel</button>
</form>

<script type="text/javascript">
    $(document)
            .ready(function() {
                $('.ui.form')
                        .api({
                            url : "/products/{{.productForm.ID}}",
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
                                url: "/products",
                                success: function (data) {
                                    $('#main_content').html(data);
                                }
                            });
                        });
            });
</script>
