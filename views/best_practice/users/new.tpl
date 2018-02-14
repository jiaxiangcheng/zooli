
<form action="javascript:void(0);" class="ui form">
    <h2 class="title"><i class="user icon"></i>User Information</h2>
        {{template "best_practice/users/form/body.tpl" .}}
    <button id="save" class="ui primary button" type="submit">Create</button>
    <button id="cancel" class="ui button" type="submit">Cancel</button>
</form>
{{template "best_practice/common/flash.tpl" .}}

<script type="text/javascript">
    $(document)
            .ready(function() {
                $('#save')
                        .on('click', function() {
                            $.ajax({
                                async: false,
                                type: "post",
                                dataType: "json",
                                url: "/users/new",
                                data: $("form").serialize(),
                                success: function (data) {
                                    $('main_content').html(data);
                                }
                            });
                        });
                $('#cancel')
                        .on('click', function () {
                            window.location.href = "/users";
                        });
            });
</script>