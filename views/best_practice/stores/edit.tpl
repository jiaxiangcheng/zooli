<form class="ui form">
        {{template "best_practice/stores/form/header.tpl" .}}
        {{template "best_practice/stores/form/body.tpl" .}}
        <button id="save" class="ui primary button" type="submit">Save</button>
        <button id="cancel" class="ui button" type="button">Cancel</button>
    </form>
    
    <script type="text/javascript">
        $(document)
                .ready(function() {
                    $('.ui.form')
                            .api({
                                url : "/stores/{{.storeForm.ID}}",
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
                                    url: "/stores",
                                    success: function (data) {
                                        $('#main_content').html(data);
                                    }
                                });
                            });
                });
    </script>