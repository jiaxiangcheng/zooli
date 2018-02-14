
<form class="ui form">
    <h2 class="title"><i class="user icon"></i>User Information</h2>
        {{template "best_practice/users/form/body.tpl" .}}
    <button id="save" class="ui primary button" type="button">Save</button>
    <button id="cancel" class="ui button" type="button">Cancel</button>
</form>

<script type="text/javascript">
    $(document)
            .ready(function() {
                $('#save')
                        .on('click', function() {
                            $.ajax({
                                //async: false,
                                type: "POST",
                                url: "/users/{{.userForm.ID}}",
                                data: $("form").serialize(),
                                success: function (data) {
                                    $('#main_content').html(data);
                                }
                            });
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