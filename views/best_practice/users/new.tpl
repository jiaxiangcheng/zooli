<form class="ui form">
    {{template "best_practice/users/form/header.tpl" .}}
    {{template "best_practice/users/form/body.tpl" .}}
    <button id="save" class="ui primary button" type="submit">Create</button>
    <button id="cancel" class="ui button" type="button">Cancel</button>
</form>

<script type="text/javascript">
    $(document)
            .ready(function() {
                $('.ui.form')
                        .api({
                            url : '/users/new',
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
                                url: "/users",
                                success: function (data) {
                                    $('#main_content').html(data);
                                }
                            });
                        });
            });
</script>