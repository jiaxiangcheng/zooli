<form class="ui form">
    <h2 id="title">
        <i class="user icon"></i>
        New Service
    </h2>
    {{template "best_practice/services/form/body.tpl" .}}
    <button id="save" class="ui primary button" type="submit">Create</button>
    <button id="cancel" class="ui button" type="button">Cancel</button>
</form>

<script type="text/javascript">
    $(document)
            .ready(function() {
                $('.ui.form')
                        .api({
                            url : '/services/new',
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
                                url: "/services",
                                success: function (data) {
                                    $('#main_content').html(data);
                                }
                            });
                        });
            });
</script>
